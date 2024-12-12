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

// CNIExclusionInterface has methods to work with CNIExclusion resources.
type CNIExclusionInterface interface {
	List(ctx context.Context, opts metav1.ListOptions) (*kubedtnv1.CNIExclusionList, error)
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*kubedtnv1.CNIExclusion, error)
	Create(ctx context.Context, cniExclusion *kubedtnv1.CNIExclusion) (*kubedtnv1.CNIExclusion, error)
	Update(ctx context.Context, cniExclusion *kubedtnv1.CNIExclusion, opts metav1.UpdateOptions) (*kubedtnv1.CNIExclusion, error)
	UpdateStatus(ctx context.Context, cniExclusion *kubedtnv1.CNIExclusion, opts metav1.UpdateOptions) (*kubedtnv1.CNIExclusion, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Unstructured(ctx context.Context, name string, opts metav1.GetOptions) (*unstructured.Unstructured, error)

	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*kubedtnv1.CNIExclusion, error)
	PatchStatus(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*kubedtnv1.CNIExclusion, error)
}

const cniExclusionResourceName = "cniexclusions"

type cniExclusionClient struct {
	dInterface dynamic.NamespaceableResourceInterface
	restClient rest.Interface
	ns         string
}

func (c *Clientset) CNIExclusion(namespace string) CNIExclusionInterface {
	dInterface := c.dynamicClient.Resource(schema.GroupVersionResource{
		Group:    kubedtnv1.GroupVersion.Group,
		Version:  kubedtnv1.GroupVersion.Version,
		Resource: cniExclusionResourceName,
	})
	return &cniExclusionClient{
		dInterface: dInterface,
		restClient: c.restClient,
		ns:         namespace,
	}
}

func (t *cniExclusionClient) List(ctx context.Context, opts metav1.ListOptions) (*kubedtnv1.CNIExclusionList, error) {
	result := kubedtnv1.CNIExclusionList{}
	err := t.restClient.
		Get().
		Namespace(t.ns).
		Resource(cniExclusionResourceName).
		VersionedParams(&opts, scheme.ParameterCodec).
		Do(ctx).
		Into(&result)

	return &result, err
}

func (t *cniExclusionClient) Get(ctx context.Context, name string, opts metav1.GetOptions) (*kubedtnv1.CNIExclusion, error) {
	result := kubedtnv1.CNIExclusion{}
	err := t.restClient.
		Get().
		Namespace(t.ns).
		Resource(cniExclusionResourceName).
		Name(name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Do(ctx).
		Into(&result)

	return &result, err
}

func (t *cniExclusionClient) Create(ctx context.Context, cniExclusion *kubedtnv1.CNIExclusion) (*kubedtnv1.CNIExclusion, error) {
	result := kubedtnv1.CNIExclusion{}
	err := t.restClient.
		Post().
		Namespace(t.ns).
		Resource(cniExclusionResourceName).
		Body(cniExclusion).
		Do(ctx).
		Into(&result)

	return &result, err
}

func (t *cniExclusionClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return t.restClient.
		Get().
		Namespace(t.ns).
		Resource(cniExclusionResourceName).
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch(ctx)
}

func (t *cniExclusionClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return t.restClient.
		Delete().
		Namespace(t.ns).
		Resource(cniExclusionResourceName).
		VersionedParams(&opts, scheme.ParameterCodec).
		Name(name).
		Do(ctx).
		Error()
}

func (t *cniExclusionClient) Update(ctx context.Context, cniExclusion *kubedtnv1.CNIExclusion, opts metav1.UpdateOptions) (*kubedtnv1.CNIExclusion, error) {
	result := kubedtnv1.CNIExclusion{}
	err := t.restClient.
		Put().
		Namespace(t.ns).
		Resource(cniExclusionResourceName).
		Name(cniExclusion.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(cniExclusion).
		Do(ctx).
		Into(&result)
	return &result, err
}

func (t *cniExclusionClient) UpdateStatus(ctx context.Context, cniExclusion *kubedtnv1.CNIExclusion, opts metav1.UpdateOptions) (*kubedtnv1.CNIExclusion, error) {
	result := kubedtnv1.CNIExclusion{}
	err := t.restClient.
		Put().
		Namespace(t.ns).
		Resource(cniExclusionResourceName).
		Name(cniExclusion.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(cniExclusion).
		Do(ctx).
		Into(&result)
	return &result, err
}

func (t *cniExclusionClient) Unstructured(ctx context.Context, name string, opts metav1.GetOptions) (*unstructured.Unstructured, error) {
	return t.dInterface.Namespace(t.ns).Get(ctx, name, opts)
}

func (t *cniExclusionClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*kubedtnv1.CNIExclusion, error) {
	result := kubedtnv1.CNIExclusion{}
	err := t.restClient.
		Patch(pt).
		Namespace(t.ns).
		Resource(cniExclusionResourceName).
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(&result)
	return &result, err
}

func (t *cniExclusionClient) PatchStatus(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*kubedtnv1.CNIExclusion, error) {
	result := kubedtnv1.CNIExclusion{}
	err := t.restClient.
		Patch(pt).
		Namespace(t.ns).
		Resource(cniExclusionResourceName).
		Name(name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(&result)
	return &result, err
}
