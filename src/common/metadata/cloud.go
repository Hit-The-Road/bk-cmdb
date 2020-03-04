/*
 * Tencent is pleased to support the open source community by making 蓝鲸 available.,
 * Copyright (C) 2017,-2018 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the ",License",); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an ",AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */
package metadata

import (
	"time"

	"configcenter/src/common/mapstr"
)

type CloudAccount struct {
	AccountName string      `json:"bk_account_name" bson:"bk_account_name"`
	CloudVendor AccountType `json:"bk_cloud_vendor" bson:"bk_cloud_vendor"`
	AccountID   int64       `json:"bk_account_id" bson:"bk_account_id"`
	SecreteID   string      `json:"bk_secret_id" bson:"bk_secret_id"`
	SecreteKey  string      `json:"bk_secret_key" bson:"bk_secret_key"`
	Description string      `json:"bk_description" bson:"bk_description"`
	OwnerID     string      `json:"bk_supplier_account" bson:"bk_supplier_account"`
	Creator     string      `json:"bk_creator" bson:"bk_creator"`
	LastEditor  string      `json:"bk_last_editor" bson:"bk_last_editor"`
	CreateTime  time.Time   `json:"create_time" bson:"create_time"`
	LastTime    time.Time   `json:"last_time" bson:"last_time"`
}

type AccountType string

const (
	AWS          AccountType = "aws"
	TencentCloud AccountType = "tencent_cloud"
)

var SupportCloudVendors = []string{"aws", "tencent_cloud"}

type SearchCloudOption struct {
	Condition mapstr.MapStr `json:"condition" bson:"condition" field:"condition"`
	Page      BasePage      `json:"page" bson:"page" field:"page"`
	Fields    []string      `json:"fields,omitempty" bson:"fields,omitempty"`
	// 对于condition里的属性值是否精确匹配，默认为false，即使用模糊匹配和忽略大小写
	Exact     bool          `json:"exact" bson:"exact"`
}

type MultipleCloudAccount struct {
	Count int64          `json:"count"`
	Info  []CloudAccount `json:"info"`
}

type CloudAccountVerify struct {
	SecretID    string      `json:"bk_secret_id"`
	SecretKey   string      `json:"bk_secret_key"`
	CloudVendor AccountType `json:"bk_cloud_vendor"`
}

type VpcInfo struct {
	VpcName string `json:"bk_vpc_name"`
	VpcID   string `json:"bk_vpc_id"`
	Region  string `json:"bk_region"`
}

type CloudSyncTask struct {
	TaskID            int64       `json:"bk_task_id" bson:"bk_task_id"`
	TaskName          string      `json:"bk_task_name" bson:"bk_task_name"`
	ResourceType      string      `json:"bk_resource_type" bson:"bk_resource_type"`
	AccountID         int64       `json:"bk_account_id" bson:"bk_account_id"`
	CloudVendor       AccountType `json:"bk_cloud_vendor" bson:"bk_cloud_vendor"`
	SyncStatus        int         `json:"bk_sync_status" bson:"bk_sync_status"`
	OwnerID           string      `json:"bk_supplier_account" bson:"bk_supplier_account"`
	StatusDescription string      `json:"bk_status_description" bson:"bk_status_description"`
	LastSyncTime      time.Time   `json:"bk_last_sync_time" bson:"bk_last_sync_time"`
	SyncAll           bool        `json:"bk_sync_all" bson:"bk_sync_all"`
	SyncAllDir        int64       `json:"bk_sync_all_dir" bson:"bk_sync_all_dir"`
	SyncVpcs          TargetVpcs  `json:"bk_sync_vpcs" bson:"bk_sync_vpcs"`
	Creator           string      `json:"bk_creator" bson:"bk_creator"`
	CreateTime        time.Time   `json:"create_time" bson:"create_time"`
	LastEditor        string      `json:"bk_last_editor" bson:"bk_last_editor"`
	LastTime          time.Time   `json:"last_time" bson:"last_time"`
}

type TargetVpcs []VpcSyncInfo

type VpcSyncInfo struct {
	VpcID        string `json:"bk_vpc_id" bson:"bk_vpc_id"`
	VpcName      string `json:"bk_vpc_name" bson:"bk_vpc_name"`
	VpcRegion    string `json:"bk_vpc_region" bson:"bk_vpc_region"`
	VpcHostCount int64  `json:"bk_host_count" bson:"bk_host_count"`
	SyncDir      int64  `json:"bk_sync_dir" bson:"bk_sync_dir"`
}

type MultipleCloudSyncTask struct {
	Count int64           `json:"count"`
	Info  []CloudSyncTask `json:"info"`
}
