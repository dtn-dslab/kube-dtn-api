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

type PhysicalInterfaceInterface interface {
	List(ctx context.Context, opts metav1.ListOptions) (*kubedtnv1.PhysicalInterfaceList, error)
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*kubedtnv1.PhysicalInterface, error)
	Create(ctx context.Context, physicalInterface *kubedtnv1.PhysicalInterface) (*kubedtnv1.PhysicalInterface, error)
	Update(ctx context.Context, physicalInterface *kubedtnv1.PhysicalInterface, opts metav1.UpdateOptions) (*kubedtnv1.PhysicalInterface, error)

	UpdateStatus(ctx context.Context, physicalInterface *kubedtnv1.PhysicalInterface, opts metav1.UpdateOptions) (*kubedtnv1.PhysicalInterface, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)

	Unstructured(ctx context.Context, name string, opts metav1.GetOptions, subresources ...string) (*unstructured.Unstructured, error)
}

const physicalInterfaceResourceName = "physicalinterfaces"

type physicalInterfaceClient struct {
	dInterface dynamic.NamespaceableResourceInterface
	restClient rest.Interface
	ns         string
}

func (c *Clientset) PhysicalInterface(namespace string) PhysicalInterfaceInterface {
	dInterface := c.dynamicClient.Resource(schema.GroupVersionResource{
		Group:    kubedtnv1.GroupVersion.Group,
		Version:  kubedtnv1.GroupVersion.Version,
		Resource: physicalInterfaceResourceName,
	})
	return &physicalInterfaceClient{
		dInterface: dInterface,
		restClient: c.restClient,
		ns:         namespace,
	}
}

func (t *physicalInterfaceClient) List(ctx context.Context, opts metav1.ListOptions) (*kubedtnv1.PhysicalInterfaceList, error) {
	result := kubedtnv1.PhysicalInterfaceList{}
	err := t.restClient.
		Get().
		Namespace(t.ns).
		Resource(physicalInterfaceResourceName).
		VersionedParams(&opts, scheme.ParameterCodec).
		Do(ctx).
		Into(&result)
	return &result, err
}

func (t *physicalInterfaceClient) Get(ctx context.Context, name string, opts metav1.GetOptions) (*kubedtnv1.PhysicalInterface, error) {
	result := kubedtnv1.PhysicalInterface{}
	err := t.restClient.
		Get().
		Namespace(t.ns).
		Resource(physicalInterfaceResourceName).
		Name(name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Do(ctx).
		Into(&result)
	return &result, err
}

func (t *physicalInterfaceClient) Create(ctx context.Context, physicalInterface *kubedtnv1.PhysicalInterface) (*kubedtnv1.PhysicalInterface, error) {
	result := kubedtnv1.PhysicalInterface{}
	err := t.restClient.
		Post().
		Namespace(t.ns).
		Resource(physicalInterfaceResourceName).
		Body(physicalInterface).
		Do(ctx).
		Into(&result)
	return &result, err
}

func (t *physicalInterfaceClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return t.restClient.
		Get().
		Namespace(t.ns).
		Resource(physicalInterfaceResourceName).
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch(ctx)
}

func (t *physicalInterfaceClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return t.restClient.
		Delete().
		Namespace(t.ns).
		Resource(physicalInterfaceResourceName).
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

func (t *physicalInterfaceClient) Update(ctx context.Context, physicalInterface *kubedtnv1.PhysicalInterface, opts metav1.UpdateOptions) (*kubedtnv1.PhysicalInterface, error) {
	result := kubedtnv1.PhysicalInterface{}
	err := t.restClient.
		Put().
		Namespace(t.ns).
		Resource(physicalInterfaceResourceName).
		Name(physicalInterface.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(physicalInterface).
		Do(ctx).
		Into(&result)
	return &result, err
}

func (t *physicalInterfaceClient) UpdateStatus(ctx context.Context, physicalInterface *kubedtnv1.PhysicalInterface, opts metav1.UpdateOptions) (*kubedtnv1.PhysicalInterface, error) {
	result := kubedtnv1.PhysicalInterface{}
	err := t.restClient.
		Put().
		Namespace(t.ns).
		Resource(physicalInterfaceResourceName).
		Name(physicalInterface.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(physicalInterface).
		Do(ctx).
		Into(&result)
	return &result, err
}

func (t *physicalInterfaceClient) Unstructured(ctx context.Context, name string, opts metav1.GetOptions, subresources ...string) (*unstructured.Unstructured, error) {
	return t.dInterface.Namespace(t.ns).Get(ctx, name, opts, subresources...)
}
