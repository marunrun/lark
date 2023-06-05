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

// GetCoreHrSecurityGroupList 通过部门或工作地点, 查询对应的 HRBP/属地 BP。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/security_group/query
func (r *CoreHrService) GetCoreHrSecurityGroupList(ctx context.Context, request *GetCoreHrSecurityGroupListReq, options ...MethodOptionFunc) (*GetCoreHrSecurityGroupListResp, *Response, error) {
	if r.cli.mock.mockCoreHrGetCoreHrSecurityGroupList != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] CoreHr#GetCoreHrSecurityGroupList mock enable")
		return r.cli.mock.mockCoreHrGetCoreHrSecurityGroupList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "CoreHr",
		API:                   "GetCoreHrSecurityGroupList",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/corehr/v1/security_groups/query",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getCoreHrSecurityGroupListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockCoreHrGetCoreHrSecurityGroupList mock CoreHrGetCoreHrSecurityGroupList method
func (r *Mock) MockCoreHrGetCoreHrSecurityGroupList(f func(ctx context.Context, request *GetCoreHrSecurityGroupListReq, options ...MethodOptionFunc) (*GetCoreHrSecurityGroupListResp, *Response, error)) {
	r.mockCoreHrGetCoreHrSecurityGroupList = f
}

// UnMockCoreHrGetCoreHrSecurityGroupList un-mock CoreHrGetCoreHrSecurityGroupList method
func (r *Mock) UnMockCoreHrGetCoreHrSecurityGroupList() {
	r.mockCoreHrGetCoreHrSecurityGroupList = nil
}

// GetCoreHrSecurityGroupListReq ...
type GetCoreHrSecurityGroupListReq struct {
	DepartmentIDType *DepartmentIDType                    `query:"department_id_type" json:"-"` // 此次调用中使用的部门 ID 类型, 示例值: people_corehr_department_id, 可选值有: open_department_id: 以 open_department_id 来标识部门, department_id: 以 department_id 来标识部门, people_corehr_department_id: 以 people_corehr_department_id 来标识部门, 默认值: `people_corehr_department_id`
	ItemList         []*GetCoreHrSecurityGroupListReqItem `json:"item_list,omitempty"`          // 角色列表, 一次最多支持查询 50 个
}

// GetCoreHrSecurityGroupListReqItem ...
type GetCoreHrSecurityGroupListReqItem struct {
	RoleKey        string  `json:"role_key,omitempty"`         // 角色类型的唯一标识, HRBP: 与部门有关, role_key 固定为 「hrbp」, 属地 BP: 与部门、工作地点有关, role_key 固定为 「location_bp」, 示例值: "location_bp"
	DepartmentID   string  `json:"department_id,omitempty"`    // 部门 ID, 查询 HRBP 需输入部门 ID, 示例值: "7063072995761456670"
	WorkLocationID *string `json:"work_location_id,omitempty"` // 工作地点 ID, 查询属地 BP 需要输入部门 ID 与 工作地点 ID, 示例值: "6892687221355185677"
}

// GetCoreHrSecurityGroupListResp ...
type GetCoreHrSecurityGroupListResp struct {
	HrbpList []*GetCoreHrSecurityGroupListRespHrbp `json:"hrbp_list,omitempty"` // HRBP/属地 BP 信息
}

// GetCoreHrSecurityGroupListRespHrbp ...
type GetCoreHrSecurityGroupListRespHrbp struct {
	EmploymentIDList []string `json:"employment_id_list,omitempty"` // HRBP/属地 BP 的雇员ID : 对于 HRBP 而言, 若入参的部门没有找到对应的 HRBP, 将向上找寻, 即向其上级部门取对应的 HRBP, 且同一部门可能有多个 HRBP；, 对于 属地 BP 而言, 若入参的部门和地点没有找到对应的属地 BP, 将优先拿地点向上找寻, 即向其上级地点取对应的属地 BP
	DepartmentID     string   `json:"department_id,omitempty"`      // 部门 ID
	WorkLocationID   string   `json:"work_location_id,omitempty"`   // 工作地点 ID
}

// getCoreHrSecurityGroupListResp ...
type getCoreHrSecurityGroupListResp struct {
	Code int64                           `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                          `json:"msg,omitempty"`  // 错误描述
	Data *GetCoreHrSecurityGroupListResp `json:"data,omitempty"`
}