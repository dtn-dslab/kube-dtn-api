package v1

import (
	"context"

	kubedtnv1 "dslab.sjtu/kube-dtn/api/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
)

// NetworkNodeInterface has methods to work with NetworkNode resources.
type NetworkNodeInterface interface {
	List(ctx context.Context, opts metav1.ListOptions) (*kubedtnv1.NetworkNodeList, error)
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*kubedtnv1.NetworkNode, error)
	Create(ctx context.Context, networkNode *kubedtnv1.NetworkNode) (*kubedtnv1.NetworkNode, error)
	Update(ctx context.Context, networkNode *kubedtnv1.NetworkNode, opts metav1.UpdateOptions) (*kubedtnv1.NetworkNode, error)
	UpdateStatus(ctx context.Context, networkNode *kubedtnv1.NetworkNode, opts metav1.UpdateOptions) (*kubedtnv1.NetworkNode, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Unstructured(ctx context.Context, name string, opts metav1.GetOptions, subresources ...string) (*unstructured.Unstructured, error)
}

// Interface is the clientset for networknode API group
type Interface interface {
	NetworkNode(namespace string) NetworkNodeInterface
}

// Clientset is a clientset for networknode API group
type Clientset struct {
	dInterface dynamic.NamespaceableResourceInterface
	restClient rest.Interface
}

const resourceName = "networknodes"

var gvr = schema.GroupVersionResource{
	Group:    kubedtnv1.GroupVersion.Group,
	Version:  kubedtnv1.GroupVersion.Version,
	Resource: resourceName,
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
	dInterface := dClient.Resource(gvr)
	rClient, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &Clientset{
		dInterface: dInterface,
		restClient: rClient,
	}, nil
}

func (c *Clientset) NetworkNode(namespace string) NetworkNodeInterface {
	return &networkNodeClient{
		dInterface: c.dInterface,
		restClient: c.restClient,
		ns:         namespace,
	}
}

type networkNodeClient struct {
	dInterface dynamic.NamespaceableResourceInterface
	restClient rest.Interface
	ns         string
}

func (t *networkNodeClient) List(ctx context.Context, opts metav1.ListOptions) (*kubedtnv1.NetworkNodeList, error) {
	result := kubedtnv1.NetworkNodeList{}
	err := t.restClient.
		Get().
		Namespace(t.ns).
		Resource(resourceName).
		VersionedParams(&opts, scheme.ParameterCodec).
		Do(ctx).
		Into(&result)

	return &result, err
}

func (t *networkNodeClient) Get(ctx context.Context, name string, opts metav1.GetOptions) (*kubedtnv1.NetworkNode, error) {
	result := kubedtnv1.NetworkNode{}
	err := t.restClient.
		Get().
		Namespace(t.ns).
		Resource(resourceName).
		Name(name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Do(ctx).
		Into(&result)

	return &result, err
}

func (t *networkNodeClient) Create(ctx context.Context, networkNode *kubedtnv1.NetworkNode) (*kubedtnv1.NetworkNode, error) {
	result := kubedtnv1.NetworkNode{}
	err := t.restClient.
		Post().
		Namespace(t.ns).
		Resource(resourceName).
		Body(networkNode).
		Do(ctx).
		Into(&result)

	return &result, err
}

func (t *networkNodeClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return t.restClient.
		Get().
		Namespace(t.ns).
		Resource(resourceName).
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch(ctx)
}

func (t *networkNodeClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return t.restClient.
		Delete().
		Namespace(t.ns).
		Resource(resourceName).
		VersionedParams(&opts, scheme.ParameterCodec).
		Name(name).
		Do(ctx).
		Error()
}

func (t *networkNodeClient) Update(ctx context.Context, networkNode *kubedtnv1.NetworkNode, opts metav1.UpdateOptions) (*kubedtnv1.NetworkNode, error) {
	result := kubedtnv1.NetworkNode{}
	err := t.restClient.
		Put().
		Namespace(t.ns).
		Resource(resourceName).
		Name(networkNode.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(networkNode).
		Do(ctx).
		Into(&result)
	return &result, err
}

func (t *networkNodeClient) UpdateStatus(ctx context.Context, networkNode *kubedtnv1.NetworkNode, opts metav1.UpdateOptions) (*kubedtnv1.NetworkNode, error) {
	result := kubedtnv1.NetworkNode{}
	err := t.restClient.
		Put().
		Namespace(t.ns).
		Resource(resourceName).
		Name(networkNode.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(networkNode).
		Do(ctx).
		Into(&result)
	return &result, err
}

func (t *networkNodeClient) Unstructured(ctx context.Context, name string, opts metav1.GetOptions, subresources ...string) (*unstructured.Unstructured, error) {
	return t.dInterface.Namespace(t.ns).Get(ctx, name, opts, subresources...)
}

func init() {
	kubedtnv1.AddToScheme(scheme.Scheme)
}
