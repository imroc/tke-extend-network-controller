package manager

import (
	"crypto/tls"

	"sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/client"

	corev1 "k8s.io/api/core/v1"

	kubelib "github.com/imroc/tke-extend-network-controller/pkg/kube"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/metrics/filters"
	metricsserver "sigs.k8s.io/controller-runtime/pkg/metrics/server"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

func GetOptions(scheme *runtime.Scheme, metricsAddr, probeAddr string, enableLeaderElection bool) manager.Options {
	tlsOpts := []func(*tls.Config){}

	webhookServer := webhook.NewServer(webhook.Options{
		TLSOpts: tlsOpts,
	})

	return manager.Options{
		Cache: cache.Options{
			ByObject: map[client.Object]cache.ByObject{
				&corev1.Pod{}: {
					Transform: kubelib.StripPodUnusedFields,
				},
			},
		},
		Scheme: scheme,
		// Metrics endpoint is enabled in 'config/default/kustomization.yaml'. The Metrics options configure the server.
		// More info:
		// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.18.4/pkg/metrics/server
		// - https://book.kubebuilder.io/reference/metrics.html
		Metrics: metricsserver.Options{
			BindAddress:   metricsAddr,
			SecureServing: false,
			// TODO(user): TLSOpts is used to allow configuring the TLS config used for the server. If certificates are
			// not provided, self-signed certificates will be generated by default. This option is not recommended for
			// production environments as self-signed certificates do not offer the same level of trust and security
			// as certificates issued by a trusted Certificate Authority (CA). The primary risk is potentially allowing
			// unauthorized access to sensitive metrics data. Consider replacing with CertDir, CertName, and KeyName
			// to provide certificates, ensuring the server communicates using trusted and secure certificates.
			TLSOpts: tlsOpts,
			// FilterProvider is used to protect the metrics endpoint with authn/authz.
			// These configurations ensure that only authorized users and service accounts
			// can access the metrics endpoint. The RBAC are configured in 'config/rbac/kustomization.yaml'. More info:
			// https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.18.4/pkg/metrics/filters#WithAuthenticationAndAuthorization
			FilterProvider: filters.WithAuthenticationAndAuthorization,
		},
		WebhookServer:          webhookServer,
		HealthProbeBindAddress: probeAddr,
		LeaderElection:         enableLeaderElection,
		LeaderElectionID:       "0ce7cb44.cloud.tencent.com",
		// LeaderElectionReleaseOnCancel defines if the leader should step down voluntarily
		// when the Manager ends. This requires the binary to immediately end when the
		// Manager is stopped, otherwise, this setting is unsafe. Setting this significantly
		// speeds up voluntary leader transitions as the new leader don't have to wait
		// LeaseDuration time first.
		//
		// In the default scaffold provided, the program ends immediately after
		// the manager stops, so would be fine to enable this option. However,
		// if you are doing or is intended to do any operation such as perform cleanups
		// after the manager stops then its usage might be unsafe.
		// LeaderElectionReleaseOnCancel: true,
	}
}
