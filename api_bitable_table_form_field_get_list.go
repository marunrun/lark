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

// GetBitableTableFormFieldList 列出表单的所有问题项
//
// 该接口支持调用频率上限为 20 QPS
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-form-field/list
func (r *BitableService) GetBitableTableFormFieldList(ctx context.Context, request *GetBitableTableFormFieldListReq, options ...MethodOptionFunc) (*GetBitableTableFormFieldListResp, *Response, error) {
	if r.cli.mock.mockBitableGetBitableTableFormFieldList != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Bitable#GetBitableTableFormFieldList mock enable")
		return r.cli.mock.mockBitableGetBitableTableFormFieldList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Bitable",
		API:                   "GetBitableTableFormFieldList",
		Method:                "GET",
		URL:                   r.cli.openBaseURL + "/open-apis/bitable/v1/apps/:app_token/tables/:table_id/forms/:form_id/fields",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(getBitableTableFormFieldListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockBitableGetBitableTableFormFieldList mock BitableGetBitableTableFormFieldList method
func (r *Mock) MockBitableGetBitableTableFormFieldList(f func(ctx context.Context, request *GetBitableTableFormFieldListReq, options ...MethodOptionFunc) (*GetBitableTableFormFieldListResp, *Response, error)) {
	r.mockBitableGetBitableTableFormFieldList = f
}

// UnMockBitableGetBitableTableFormFieldList un-mock BitableGetBitableTableFormFieldList method
func (r *Mock) UnMockBitableGetBitableTableFormFieldList() {
	r.mockBitableGetBitableTableFormFieldList = nil
}

// GetBitableTableFormFieldListReq ...
type GetBitableTableFormFieldListReq struct {
	AppToken  string  `path:"app_token" json:"-"`   // 多维表格文档 Token, 示例值: "bascnCMII2ORej2RItqpZZUNMIe"
	TableID   string  `path:"table_id" json:"-"`    // 表格 ID, 示例值: "tblxI2tWaxP5dG7p"
	FormID    string  `path:"form_id" json:"-"`     // 表单 ID, 示例值: "vewTpR1urY"
	PageSize  *int64  `query:"page_size" json:"-"`  // 分页大小, 示例值: 10, 最大值: `100`
	PageToken *string `query:"page_token" json:"-"` // 分页标记, 第一次请求不填, 表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token, 下次遍历可采用该 page_token 获取查询结果, 示例值: "vewTpR1urY"
}

// GetBitableTableFormFieldListResp ...
type GetBitableTableFormFieldListResp struct {
	Items     []*GetBitableTableFormFieldListRespItem `json:"items,omitempty"`      // 表单问题信息
	PageToken string                                  `json:"page_token,omitempty"` // 分页标记, 当 has_more 为 true 时, 会同时返回新的 page_token, 否则不返回 page_token
	HasMore   bool                                    `json:"has_more,omitempty"`   // 是否还有更多项
	Total     int64                                   `json:"total,omitempty"`      // 总数
}

// GetBitableTableFormFieldListRespItem ...
type GetBitableTableFormFieldListRespItem struct {
	FieldID     string `json:"field_id,omitempty"`    // 表单问题 ID
	Title       string `json:"title,omitempty"`       // 表单问题
	Description string `json:"description,omitempty"` // 问题描述
	Required    bool   `json:"required,omitempty"`    // 是否必填
	Visible     bool   `json:"visible,omitempty"`     // 是否可见
}

// getBitableTableFormFieldListResp ...
type getBitableTableFormFieldListResp struct {
	Code int64                             `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                            `json:"msg,omitempty"`  // 错误描述
	Data *GetBitableTableFormFieldListResp `json:"data,omitempty"`
}
