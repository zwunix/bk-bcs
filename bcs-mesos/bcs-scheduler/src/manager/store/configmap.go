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

package store

import (
	"bk-bcs/bcs-common/common/blog"
	commtypes "bk-bcs/bcs-common/common/types"
	"encoding/json"
)

func getConfigMapRootPath() string {
	return "/" + bcsRootNode + "/" + configMapNode
}

func (store *managerStore) SaveConfigMap(configmap *commtypes.BcsConfigMap) error {

	data, err := json.Marshal(configmap)
	if err != nil {
		return err
	}

	path := getConfigMapRootPath() + "/" + configmap.ObjectMeta.NameSpace + "/" + configmap.ObjectMeta.Name

	return store.Db.Insert(path, string(data))
}

func (store *managerStore) FetchConfigMap(ns, name string) (*commtypes.BcsConfigMap, error) {

	path := getConfigMapRootPath() + "/" + ns + "/" + name

	data, err := store.Db.Fetch(path)
	if err != nil {
		return nil, err
	}

	configmap := &commtypes.BcsConfigMap{}
	if err := json.Unmarshal(data, configmap); err != nil {
		blog.Error("fail to unmarshal configmap(%s). err:%s", string(data), err.Error())
		return nil, err
	}

	return configmap, nil
}

func (store *managerStore) DeleteConfigMap(ns, name string) error {

	path := getConfigMapRootPath() + "/" + ns + "/" + name
	if err := store.Db.Delete(path); err != nil {
		blog.Error("fail to delete configmap(%s) err:%s", path, err.Error())
		return err
	}

	return nil
}