// Code generated by lark_sdk_gen. DO NOT EDIT.
/**
 * Copyright 2022 chyroc
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package lark

import (
	"context"
)

// CreateTaskCustomFieldOption 为单选或多选字段添加一个自定义选项。一个单选/多选字段最大支持100个选项。
//
// 新添加的选项如果不隐藏, 其名字不能和已存在的不隐藏选项的名字重复。
// 需要对自定义字段的编辑权限。
//
// doc: https://open.larkoffice.com/document/uAjLw4CM/ukTMukTMukTM/task-v2/custom_field-option/create
func (r *TaskService) CreateTaskCustomFieldOption(ctx context.Context, request *CreateTaskCustomFieldOptionReq, options ...MethodOptionFunc) (*CreateTaskCustomFieldOptionResp, *Response, error) {
	if r.cli.mock.mockTaskCreateTaskCustomFieldOption != nil {
		r.cli.Log(ctx, LogLevelDebug, "[lark] Task#CreateTaskCustomFieldOption mock enable")
		return r.cli.mock.mockTaskCreateTaskCustomFieldOption(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Task",
		API:                   "CreateTaskCustomFieldOption",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/task/v2/custom_fields/:custom_field_guid/options",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(createTaskCustomFieldOptionResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockTaskCreateTaskCustomFieldOption mock TaskCreateTaskCustomFieldOption method
func (r *Mock) MockTaskCreateTaskCustomFieldOption(f func(ctx context.Context, request *CreateTaskCustomFieldOptionReq, options ...MethodOptionFunc) (*CreateTaskCustomFieldOptionResp, *Response, error)) {
	r.mockTaskCreateTaskCustomFieldOption = f
}

// UnMockTaskCreateTaskCustomFieldOption un-mock TaskCreateTaskCustomFieldOption method
func (r *Mock) UnMockTaskCreateTaskCustomFieldOption() {
	r.mockTaskCreateTaskCustomFieldOption = nil
}

// CreateTaskCustomFieldOptionReq ...
type CreateTaskCustomFieldOptionReq struct {
	CustomFieldGuid string  `path:"custom_field_guid" json:"-"` // 要添加选项的自定义字段GUID, 该字段必须是, 示例值: "b13adf3c-cad6-4e02-8929-550c112b5633"
	Name            *string `json:"name,omitempty"`             // 选项名称, 最大50个字符, 示例值: "高优"
	ColorIndex      *int64  `json:"color_index,omitempty"`      // 颜色索引值, 支持0～54中的一个数字。如果不填写, 则会随机选一个, 示例值: 10, 取值范围: `0` ～ `54`
	InsertBefore    *string `json:"insert_before,omitempty"`    // 要放到某个option之前的option_guid, 示例值: "2bd905f8-ef38-408b-aa1f-2b2ad33b2913"
	InsertAfter     *string `json:"insert_after,omitempty"`     // 要放到某个option之后的option_guid, 示例值: "b13adf3c-cad6-4e02-8929-550c112b5633"
	IsHidden        *bool   `json:"is_hidden,omitempty"`        // 是否隐藏, 示例值: false, 默认值: `false`
}

// CreateTaskCustomFieldOptionResp ...
type CreateTaskCustomFieldOptionResp struct {
	Option *CreateTaskCustomFieldOptionRespOption `json:"option,omitempty"` // 创建的选项
}

// CreateTaskCustomFieldOptionRespOption ...
type CreateTaskCustomFieldOptionRespOption struct {
	Guid       string `json:"guid,omitempty"`        // 选项的GUID
	Name       string `json:"name,omitempty"`        // 选项名称, 不能为空, 最大50个字符
	ColorIndex int64  `json:"color_index,omitempty"` // 选项的颜色索引值, 可以是0～54中的一个数字。如果不填写则会随机选一个。
	IsHidden   bool   `json:"is_hidden,omitempty"`   // 选项是否隐藏。隐藏后的选项在界面不可见, 也不可以再通过openapi将字段值设为该选项。
}

// createTaskCustomFieldOptionResp ...
type createTaskCustomFieldOptionResp struct {
	Code  int64                            `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg   string                           `json:"msg,omitempty"`  // 错误描述
	Data  *CreateTaskCustomFieldOptionResp `json:"data,omitempty"`
	Error *ErrorDetail                     `json:"error,omitempty"`
}
