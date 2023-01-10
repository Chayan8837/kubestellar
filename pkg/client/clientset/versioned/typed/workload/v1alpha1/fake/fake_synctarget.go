/*
Copyright The KCP Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"

	v1alpha1 "github.com/kcp-dev/kcp/pkg/apis/workload/v1alpha1"
)

// FakeSyncTargets implements SyncTargetInterface
type FakeSyncTargets struct {
	Fake *FakeWorkloadV1alpha1
}

var synctargetsResource = schema.GroupVersionResource{Group: "workload.kcp.dev", Version: "v1alpha1", Resource: "synctargets"}

var synctargetsKind = schema.GroupVersionKind{Group: "workload.kcp.dev", Version: "v1alpha1", Kind: "SyncTarget"}

// Get takes name of the syncTarget, and returns the corresponding syncTarget object, and an error if there is any.
func (c *FakeSyncTargets) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.SyncTarget, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(synctargetsResource, name), &v1alpha1.SyncTarget{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SyncTarget), err
}

// List takes label and field selectors, and returns the list of SyncTargets that match those selectors.
func (c *FakeSyncTargets) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.SyncTargetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(synctargetsResource, synctargetsKind, opts), &v1alpha1.SyncTargetList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.SyncTargetList{ListMeta: obj.(*v1alpha1.SyncTargetList).ListMeta}
	for _, item := range obj.(*v1alpha1.SyncTargetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested syncTargets.
func (c *FakeSyncTargets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(synctargetsResource, opts))
}

// Create takes the representation of a syncTarget and creates it.  Returns the server's representation of the syncTarget, and an error, if there is any.
func (c *FakeSyncTargets) Create(ctx context.Context, syncTarget *v1alpha1.SyncTarget, opts v1.CreateOptions) (result *v1alpha1.SyncTarget, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(synctargetsResource, syncTarget), &v1alpha1.SyncTarget{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SyncTarget), err
}

// Update takes the representation of a syncTarget and updates it. Returns the server's representation of the syncTarget, and an error, if there is any.
func (c *FakeSyncTargets) Update(ctx context.Context, syncTarget *v1alpha1.SyncTarget, opts v1.UpdateOptions) (result *v1alpha1.SyncTarget, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(synctargetsResource, syncTarget), &v1alpha1.SyncTarget{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SyncTarget), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeSyncTargets) UpdateStatus(ctx context.Context, syncTarget *v1alpha1.SyncTarget, opts v1.UpdateOptions) (*v1alpha1.SyncTarget, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(synctargetsResource, "status", syncTarget), &v1alpha1.SyncTarget{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SyncTarget), err
}

// Delete takes name of the syncTarget and deletes it. Returns an error if one occurs.
func (c *FakeSyncTargets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(synctargetsResource, name, opts), &v1alpha1.SyncTarget{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSyncTargets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(synctargetsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.SyncTargetList{})
	return err
}

// Patch applies the patch and returns the patched syncTarget.
func (c *FakeSyncTargets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.SyncTarget, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(synctargetsResource, name, pt, data, subresources...), &v1alpha1.SyncTarget{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.SyncTarget), err
}
