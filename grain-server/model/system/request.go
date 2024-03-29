// Copyright © 2023 Grain. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package model

// PageReq 分页查询使用的结构体
type PageReq struct {
	// 这个字段对前端隐藏,只服务于后端
	UID string ` form:"-" json:"-"`
	// 前端传过来用于查询数据的ID
	ID string ` form:"id" json:"id"`
	// 所查询的数据总量
	Total int64 `json:"total,omitempty"  form:"total"`
	// 页
	Page int `json:"page,omitempty"  form:"page"`
	// 页大小
	PageSize int `json:"pageSize,omitempty"  form:"pageSize"`
	// 查询关键词
	Keyword string `json:"keyword"  form:"keyword"`
	// 查询类型
	Type int `json:"type"  form:"type"`
}
