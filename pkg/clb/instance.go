package clb

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/imroc/tke-extend-network-controller/pkg/util"
	clb "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/clb/v20180317"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

func GetClbExternalAddress(ctx context.Context, lbId, region string) (address string, err error) {
	lb, err := GetClb(ctx, lbId, region)
	if err != nil {
		return
	}
	if lb.LoadBalancerDomain != nil && *lb.LoadBalancerDomain != "" {
		address = *lb.LoadBalancerDomain
		return
	}
	if len(lb.LoadBalancerVips) > 0 {
		address = *lb.LoadBalancerVips[0]
		return
	}
	err = fmt.Errorf("no external address found for clb %s", lbId)
	return
}

func GetClb(ctx context.Context, lbId, region string) (instance *clb.LoadBalancer, err error) {
	client := GetClient(region)
	req := clb.NewDescribeLoadBalancersRequest()
	req.LoadBalancerIds = []*string{&lbId}

	resp, err := client.DescribeLoadBalancersWithContext(ctx, req)
	if err != nil {
		return
	}
	if *resp.Response.TotalCount == 0 {
		err = fmt.Errorf("%s is not exists in %s", lbId, region)
		return
	} else if *resp.Response.TotalCount > 1 {
		err = fmt.Errorf("%s found %d instances in %s", lbId, *resp.Response.TotalCount, region)
		return
	}
	instance = resp.Response.LoadBalancerSet[0]
	return
}

// TODO: 支持部分成功
func Create(ctx context.Context, region, vpcId, extensiveParameters string, num int) (ids []string, err error) {
	if vpcId == "" {
		vpcId = defaultVpcId
	}
	req := clb.NewCreateLoadBalancerRequest()
	req.LoadBalancerType = common.StringPtr("OPEN")
	req.VpcId = &vpcId
	req.Number = common.Uint64Ptr(uint64(num))
	req.Tags = append(req.Tags,
		&clb.TagInfo{
			TagKey:   common.StringPtr("tke-clusterId"), // 与集群关联
			TagValue: common.StringPtr(clusterId),
		},
		&clb.TagInfo{
			TagKey:   common.StringPtr("tke-createdBy-flag"), // 用于删除集群时自动清理集群关联的自动创建的 CLB
			TagValue: common.StringPtr("yes"),
		},
	)
	if extensiveParameters != "" {
		err = json.Unmarshal([]byte(extensiveParameters), req)
		if err != nil {
			return
		}
	}
	client := GetClient(region)
	resp, err := client.CreateLoadBalancerWithContext(ctx, req)
	if err != nil {
		return
	}
	ids = util.ConvertStringPointSlice(resp.Response.LoadBalancerIds)
	if len(ids) == 0 {
		ids, err = Wait(ctx, region, *resp.Response.RequestId)
		if err != nil {
			return nil, err
		}
	}
	if len(ids) == 0 {
		return nil, fmt.Errorf("no loadbalancer created")
	}
	for _, lbId := range ids {
		for {
			lb, err := GetClb(ctx, lbId, region)
			if err != nil {
				return nil, err
			}
			if *lb.Status == 0 { // 创建中，等待一下
				log.FromContext(ctx).Info("lb is still creating", "lbId", lbId)
				time.Sleep(time.Second * 3)
				continue
			}
			break
		}
	}
	return
}

func Delete(ctx context.Context, region string, lbIds ...string) error {
	req := clb.NewDeleteLoadBalancerRequest()
	for _, lbId := range lbIds {
		req.LoadBalancerIds = append(req.LoadBalancerIds, &lbId)
	}
	client := GetClient(region)
	resp, err := client.DeleteLoadBalancerWithContext(ctx, req)
	if err != nil {
		if IsLbIdNotFoundError(err) {
			if len(lbIds) == 1 { // lb 已全部删除，忽略
				return nil
			} else { // lb 可能全部删除，也可能部分删除，挨个尝试一下
				for _, lbId := range lbIds {
					if err := Delete(ctx, region, lbId); err != nil {
						return err
					}
				}
				return nil
			}
		}
		return err
	}
	_, err = Wait(ctx, region, *resp.Response.RequestId)
	return err
}
