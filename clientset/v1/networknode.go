package v1

import (
	"context"

	kubedtnv1 "dslab.sjtu/kube-dtn/api/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
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

	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*kubedtnv1.NetworkNode, error)

	PatchStatus(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*kubedtnv1.NetworkNode, error)
}

const networkNodeResourceName = "networknodes"

type networkNodeClient struct {
	dInterface dynamic.NamespaceableResourceInterface
	restClient rest.Interface
	ns         string
}

func (c *Clientset) NetworkNode(namespace string) NetworkNodeInterface {
	dInterface := c.dynamicClient.Resource(schema.GroupVersionResource{
		Group:    kubedtnv1.GroupVersion.Group,
		Version:  kubedtnv1.GroupVersion.Version,
		Resource: networkNodeResourceName,
	})
	return &networkNodeClient{
		dInterface: dInterface,
		restClient: c.restClient,
		ns:         namespace,
	}
}

func (t *networkNodeClient) List(ctx context.Context, opts metav1.ListOptions) (*kubedtnv1.NetworkNodeList, error) {
	result := kubedtnv1.NetworkNodeList{}
	err := t.restClient.
		Get().
		Namespace(t.ns).
		Resource(networkNodeResourceName).
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
		Resource(networkNodeResourceName).
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
		Resource(networkNodeResourceName).
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
		Resource(networkNodeResourceName).
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch(ctx)
}

func (t *networkNodeClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return t.restClient.
		Delete().
		Namespace(t.ns).
		Resource(networkNodeResourceName).
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
		Resource(networkNodeResourceName).
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
		Resource(networkNodeResourceName).
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

func (t *networkNodeClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*kubedtnv1.NetworkNode, error) {
	result := kubedtnv1.NetworkNode{}
	err := t.restClient.
		Patch(pt).
		Namespace(t.ns).
		Resource(networkNodeResourceName).
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(&result)
	return &result, err
}

func (t *networkNodeClient) PatchStatus(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*kubedtnv1.NetworkNode, error) {
	result := kubedtnv1.NetworkNode{}
	err := t.restClient.
		Patch(pt).
		Namespace(t.ns).
		Resource(networkNodeResourceName).
		Name(name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(&result)
	return &result, err
}
