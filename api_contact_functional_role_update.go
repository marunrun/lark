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

// UpdateContactFunctionalRole 通过本接口可以修改角色名称
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/functional_role/update
func (r *ContactService) UpdateContactFunctionalRole(ctx context.Context, request *UpdateContactFunctionalRoleReq, options ...MethodOptionFunc) (*UpdateContactFunctionalRoleResp, *Response, error) {
	if r.cli.mock.mockContactUpdateContactFunctionalRole != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Contact#UpdateContactFunctionalRole mock enable")
		return r.cli.mock.mockContactUpdateContactFunctionalRole(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Contact",
		API:                   "UpdateContactFunctionalRole",
		Method:                "PUT",
		URL:                   r.cli.openBaseURL + "/open-apis/contact/v3/functional_roles/:role_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(updateContactFunctionalRoleResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockContactUpdateContactFunctionalRole mock ContactUpdateContactFunctionalRole method
func (r *Mock) MockContactUpdateContactFunctionalRole(f func(ctx context.Context, request *UpdateContactFunctionalRoleReq, options ...MethodOptionFunc) (*UpdateContactFunctionalRoleResp, *Response, error)) {
	r.mockContactUpdateContactFunctionalRole = f
}

// UnMockContactUpdateContactFunctionalRole un-mock ContactUpdateContactFunctionalRole method
func (r *Mock) UnMockContactUpdateContactFunctionalRole() {
	r.mockContactUpdateContactFunctionalRole = nil
}

// UpdateContactFunctionalRoleReq ...
type UpdateContactFunctionalRoleReq struct {
	RoleID   string `path:"role_id" json:"-"`    // 角色的唯一标识, 单租户下唯一, 示例值: "7vrj3vk70xk7v5r"
	RoleName string `json:"role_name,omitempty"` // 修改的角色名称, 在单租户下唯一, 示例值: "考勤管理员", 长度范围: `1` ～ `50` 字符
}

// UpdateContactFunctionalRoleResp ...
type UpdateContactFunctionalRoleResp struct {
}

// updateContactFunctionalRoleResp ...
type updateContactFunctionalRoleResp struct {
	Code int64                            `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                           `json:"msg,omitempty"`  // 错误描述
	Data *UpdateContactFunctionalRoleResp `json:"data,omitempty"`
}
