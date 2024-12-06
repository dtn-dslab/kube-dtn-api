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

type DPUBindingInterface interface {
	List(ctx context.Context, opts metav1.ListOptions) (*kubedtnv1.DPUBindingList, error)
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*kubedtnv1.DPUBinding, error)
	Create(ctx context.Context, dpubinding *kubedtnv1.DPUBinding) (*kubedtnv1.DPUBinding, error)
	Update(ctx context.Context, dpubinding *kubedtnv1.DPUBinding, opts metav1.UpdateOptions) (*kubedtnv1.DPUBinding, error)
	UpdateStatus(ctx context.Context, dpubinding *kubedtnv1.DPUBinding, opts metav1.UpdateOptions) (*kubedtnv1.DPUBinding, error)

	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)

	Unstructured(ctx context.Context, name string, opts metav1.GetOptions) (*unstructured.Unstructured, error)

	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*kubedtnv1.DPUBinding, error)
	PatchStatus(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions) (*kubedtnv1.DPUBinding, error)
}

const dpubindingResourceName = "dpubindings"

type dpubindingClient struct {
	dInterface dynamic.NamespaceableResourceInterface
	restClient rest.Interface
	ns         string
}

func (c *Clientset) DPUBinding(namespace string) DPUBindingInterface {
	dInterface := c.dynamicClient.Resource(schema.GroupVersionResource{
		Group:    kubedtnv1.GroupVersion.Group,
		Version:  kubedtnv1.GroupVersion.Version,
		Resource: dpubindingResourceName,
	})
	return &dpubindingClient{
		dInterface: dInterface,
		restClient: c.restClient,
		ns:         namespace,
	}
}

func (t *dpubindingClient) List(ctx context.Context, opts metav1.ListOptions) (*kubedtnv1.DPUBindingList, error) {
	result := kubedtnv1.DPUBindingList{}
	err := t.restClient.
		Get().
		Namespace(t.ns).
		Resource(dpubindingResourceName).
		VersionedParams(&opts, scheme.ParameterCodec).
		Do(ctx).
		Into(&result)

	return &result, err
}

func (t *dpubindingClient) Get(ctx context.Context, name string, opts metav1.GetOptions) (*kubedtnv1.DPUBinding, error) {
	result := kubedtnv1.DPUBinding{}
	err := t.restClient.
		Get().
		Namespace(t.ns).
		Resource(dpubindingResourceName).
		Name(name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Do(ctx).
		Into(&result)

	return &result, err
}

func (t *dpubindingClient) Create(ctx context.Context, dpubinding *kubedtnv1.DPUBinding) (*kubedtnv1.DPUBinding, error) {
	result := kubedtnv1.DPUBinding{}
	err := t.restClient.
		Post().
		Namespace(t.ns).
		Resource(dpubindingResourceName).
		Body(dpubinding).
		Do(ctx).
		Into(&result)

	return &result, err
}

func (t *dpubindingClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return t.restClient.
		Get().
		Namespace(t.ns).
		Resource(dpubindingResourceName).
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch(ctx)
}

func (t *dpubindingClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return t.restClient.
		Delete().
		Namespace(t.ns).
		Resource(dpubindingResourceName).
		VersionedParams(&opts, scheme.ParameterCodec).
		Name(name).
		Do(ctx).
		Error()
}

func (t *dpubindingClient) Update(ctx context.Context, dpubinding *kubedtnv1.DPUBinding, opts metav1.UpdateOptions) (*kubedtnv1.DPUBinding, error) {
	result := kubedtnv1.DPUBinding{}
	err := t.restClient.
		Put().
		Namespace(t.ns).
		Resource(dpubindingResourceName).
		Name(dpubinding.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(dpubinding).
		Do(ctx).
		Into(&result)
	return &result, err
}

func (t *dpubindingClient) UpdateStatus(ctx context.Context, dpubinding *kubedtnv1.DPUBinding, opts metav1.UpdateOptions) (*kubedtnv1.DPUBinding, error) {
	result := kubedtnv1.DPUBinding{}
	err := t.restClient.
		Put().
		Namespace(t.ns).
		Resource(dpubindingResourceName).
		Name(dpubinding.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(dpubinding).
		Do(ctx).
		Into(&result)
	return &result, err
}

func (t *dpubindingClient) Unstructured(ctx context.Context, name string, opts metav1.GetOptions) (*unstructured.Unstructured, error) {
	return t.dInterface.Namespace(t.ns).Get(ctx, name, opts)
}

func (t *dpubindingClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*kubedtnv1.DPUBinding, error) {
	result := kubedtnv1.DPUBinding{}
	err := t.restClient.
		Patch(pt).
		Namespace(t.ns).
		Resource(dpubindingResourceName).
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(&result)
	return &result, err
}

func (t *dpubindingClient) PatchStatus(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions) (*kubedtnv1.DPUBinding, error) {
	result := kubedtnv1.DPUBinding{}
	err := t.restClient.
		Patch(pt).
		Namespace(t.ns).
		Resource(dpubindingResourceName).
		Name(name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(&result)
	return &result, err
}
