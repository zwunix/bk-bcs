/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v2 "bk-bcs/bcs-mesos/pkg/apis/bkbcs/v2"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeTaskGroups implements TaskGroupInterface
type FakeTaskGroups struct {
	Fake *FakeBkbcsV2
	ns   string
}

var taskgroupsResource = schema.GroupVersionResource{Group: "bkbcs.tencent.com", Version: "v2", Resource: "taskgroups"}

var taskgroupsKind = schema.GroupVersionKind{Group: "bkbcs.tencent.com", Version: "v2", Kind: "TaskGroup"}

// Get takes name of the taskGroup, and returns the corresponding taskGroup object, and an error if there is any.
func (c *FakeTaskGroups) Get(name string, options v1.GetOptions) (result *v2.TaskGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(taskgroupsResource, c.ns, name), &v2.TaskGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2.TaskGroup), err
}

// List takes label and field selectors, and returns the list of TaskGroups that match those selectors.
func (c *FakeTaskGroups) List(opts v1.ListOptions) (result *v2.TaskGroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(taskgroupsResource, taskgroupsKind, c.ns, opts), &v2.TaskGroupList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v2.TaskGroupList{ListMeta: obj.(*v2.TaskGroupList).ListMeta}
	for _, item := range obj.(*v2.TaskGroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested taskGroups.
func (c *FakeTaskGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(taskgroupsResource, c.ns, opts))

}

// Create takes the representation of a taskGroup and creates it.  Returns the server's representation of the taskGroup, and an error, if there is any.
func (c *FakeTaskGroups) Create(taskGroup *v2.TaskGroup) (result *v2.TaskGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(taskgroupsResource, c.ns, taskGroup), &v2.TaskGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2.TaskGroup), err
}

// Update takes the representation of a taskGroup and updates it. Returns the server's representation of the taskGroup, and an error, if there is any.
func (c *FakeTaskGroups) Update(taskGroup *v2.TaskGroup) (result *v2.TaskGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(taskgroupsResource, c.ns, taskGroup), &v2.TaskGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2.TaskGroup), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeTaskGroups) UpdateStatus(taskGroup *v2.TaskGroup) (*v2.TaskGroup, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(taskgroupsResource, "status", c.ns, taskGroup), &v2.TaskGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2.TaskGroup), err
}

// Delete takes name of the taskGroup and deletes it. Returns an error if one occurs.
func (c *FakeTaskGroups) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(taskgroupsResource, c.ns, name), &v2.TaskGroup{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeTaskGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(taskgroupsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v2.TaskGroupList{})
	return err
}

// Patch applies the patch and returns the patched taskGroup.
func (c *FakeTaskGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v2.TaskGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(taskgroupsResource, c.ns, name, data, subresources...), &v2.TaskGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v2.TaskGroup), err
}
