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

type KubeDTNInterface interface {
	List(ctx context.Context, opts metav1.ListOptions) (*kubedtnv1.KubeDTNList, error)
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*kubedtnv1.KubeDTN, error)
	Create(ctx context.Context, kubeDTN *kubedtnv1.KubeDTN) (*kubedtnv1.KubeDTN, error)
	Update(ctx context.Context, kubeDTN *kubedtnv1.KubeDTN, opts metav1.UpdateOptions) (*kubedtnv1.KubeDTN, error)
	UpdateStatus(ctx context.Context, kubeDTN *kubedtnv1.KubeDTN, opts metav1.UpdateOptions) (*kubedtnv1.KubeDTN, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error

	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Unstructured(ctx context.Context, name string, opts metav1.GetOptions) (*unstructured.Unstructured, error)

	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*kubedtnv1.KubeDTN, error)

	PatchStatus(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*kubedtnv1.KubeDTN, error)
}

const kubeDTNResourceName = "kubedtns"

type kubeDTNClient struct {
	dInterface dynamic.NamespaceableResourceInterface
	restClient rest.Interface
	ns         string
}

func (c *Clientset) KubeDTN(namespace string) KubeDTNInterface {
	dInterface := c.dynamicClient.Resource(schema.GroupVersionResource{
		Group:    kubedtnv1.GroupVersion.Group,
		Version:  kubedtnv1.GroupVersion.Version,
		Resource: kubeDTNResourceName,
	})
	return &kubeDTNClient{
		dInterface: dInterface,
		restClient: c.restClient,
		ns:         namespace,
	}
}

func (t *kubeDTNClient) List(ctx context.Context, opts metav1.ListOptions) (*kubedtnv1.KubeDTNList, error) {
	result := kubedtnv1.KubeDTNList{}
	err := t.restClient.
		Get().
		Namespace(t.ns).
		Resource(kubeDTNResourceName).
		VersionedParams(&opts, scheme.ParameterCodec).
		Do(ctx).
		Into(&result)
	return &result, err
}

func (t *kubeDTNClient) Get(ctx context.Context, name string, opts metav1.GetOptions) (*kubedtnv1.KubeDTN, error) {
	result := kubedtnv1.KubeDTN{}
	err := t.restClient.
		Get().
		Namespace(t.ns).
		Resource(kubeDTNResourceName).
		Name(name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Do(ctx).
		Into(&result)
	return &result, err
}

func (t *kubeDTNClient) Create(ctx context.Context, kubeDTN *kubedtnv1.KubeDTN) (*kubedtnv1.KubeDTN, error) {
	result := kubedtnv1.KubeDTN{}
	err := t.restClient.
		Post().
		Namespace(t.ns).
		Resource(kubeDTNResourceName).
		Body(kubeDTN).
		Do(ctx).
		Into(&result)
	return &result, err
}

func (t *kubeDTNClient) Update(ctx context.Context, kubeDTN *kubedtnv1.KubeDTN, opts metav1.UpdateOptions) (*kubedtnv1.KubeDTN, error) {
	result := kubedtnv1.KubeDTN{}
	err := t.restClient.
		Put().
		Namespace(t.ns).
		Resource(kubeDTNResourceName).
		Name(kubeDTN.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(kubeDTN).
		Do(ctx).
		Into(&result)
	return &result, err
}

func (t *kubeDTNClient) UpdateStatus(ctx context.Context, kubeDTN *kubedtnv1.KubeDTN, opts metav1.UpdateOptions) (*kubedtnv1.KubeDTN, error) {
	result := kubedtnv1.KubeDTN{}
	err := t.restClient.
		Put().
		Namespace(t.ns).
		Resource(kubeDTNResourceName).
		Name(kubeDTN.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(kubeDTN).
		Do(ctx).
		Into(&result)
	return &result, err
}

func (t *kubeDTNClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return t.restClient.
		Delete().
		Namespace(t.ns).
		Resource(kubeDTNResourceName).
		VersionedParams(&opts, scheme.ParameterCodec).
		Name(name).
		Do(ctx).
		Error()
}

func (t *kubeDTNClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return t.restClient.
		Get().
		Namespace(t.ns).
		Resource(kubeDTNResourceName).
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch(ctx)
}

func (t *kubeDTNClient) Unstructured(ctx context.Context, name string, opts metav1.GetOptions) (*unstructured.Unstructured, error) {
	result := unstructured.Unstructured{}
	err := t.restClient.
		Get().
		Namespace(t.ns).
		Resource(kubeDTNResourceName).
		Name(name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Do(ctx).
		Into(&result)
	return &result, err
}

func (t *kubeDTNClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*kubedtnv1.KubeDTN, error) {
	result := kubedtnv1.KubeDTN{}
	err := t.restClient.
		Patch(pt).
		Namespace(t.ns).
		Resource(kubeDTNResourceName).
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(&result)
	return &result, err
}

func (t *kubeDTNClient) PatchStatus(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*kubedtnv1.KubeDTN, error) {
	result := kubedtnv1.KubeDTN{}
	err := t.restClient.
		Patch(pt).
		Namespace(t.ns).
		Resource(kubeDTNResourceName).
		Name(name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(&result)
	return &result, err
}
