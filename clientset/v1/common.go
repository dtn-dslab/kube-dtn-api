package v1

import (
	kubedtnv1 "dslab.sjtu/kube-dtn/api/v1"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
)

// Clientset is a clientset for API group
type Clientset struct {
	// dInterface dynamic.NamespaceableResourceInterface
	dynamicClient dynamic.Interface // To create each dynamicInterface, reuse the same dynamic client for all resources in this API Group is OK?
	restClient    rest.Interface
}

// Clientset is the implementation of the Interface with these public methods.
type Interface interface {
	NetworkNode(namespace string) NetworkNodeInterface
	PhysicalInterface(namespace string) PhysicalInterfaceInterface
}

// NewForConfig returns a new clientset for the given config
func NewForConfig(c *rest.Config) (*Clientset, error) {
	config := *c
	config.ContentConfig.GroupVersion = &kubedtnv1.GroupVersion
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()
	config.UserAgent = rest.DefaultKubernetesUserAgent()
	dClient, err := dynamic.NewForConfig(c)
	if err != nil {
		return nil, err
	}

	rClient, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &Clientset{
		dynamicClient: dClient,
		restClient:    rClient,
	}, nil
}

func init() {
	kubedtnv1.AddToScheme(scheme.Scheme)
}
