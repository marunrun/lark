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

// GetPassportSession 该接口用于查询用户的登录信息。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/passport-v1/session/query
func (r *PassportService) GetPassportSession(ctx context.Context, request *GetPassportSessionReq, options ...MethodOptionFunc) (*GetPassportSessionResp, *Response, error) {
	if r.cli.mock.mockPassportGetPassportSession != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Passport#GetPassportSession mock enable")
		return r.cli.mock.mockPassportGetPassportSession(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Passport",
		API:                   "GetPassportSession",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/passport/v1/sessions/query",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getPassportSessionResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockPassportGetPassportSession mock PassportGetPassportSession method
func (r *Mock) MockPassportGetPassportSession(f func(ctx context.Context, request *GetPassportSessionReq, options ...MethodOptionFunc) (*GetPassportSessionResp, *Response, error)) {
	r.mockPassportGetPassportSession = f
}

// UnMockPassportGetPassportSession un-mock PassportGetPassportSession method
func (r *Mock) UnMockPassportGetPassportSession() {
	r.mockPassportGetPassportSession = nil
}

// GetPassportSessionReq ...
type GetPassportSessionReq struct {
	UserIDType *IDType  `query:"user_id_type" json:"-"` // 用户 ID 类型, 示例值: "open_id", 可选值有: open_id: 标识一个用户在某个应用中的身份。同一个用户在不同应用中的 Open ID 不同。[了解更多: 如何获取 Open ID](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-openid), union_id: 标识一个用户在某个应用开发商下的身份。同一用户在同一开发商下的应用中的 Union ID 是相同的, 在不同开发商下的应用中的 Union ID 是不同的。通过 Union ID, 应用开发商可以把同个用户在多个应用中的身份关联起来。[了解更多: 如何获取 Union ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-union-id), user_id: 标识一个用户在某个租户内的身份。同一个用户在租户 A 和租户 B 内的 User ID 是不同的。在同一个租户内, 一个用户的 User ID 在所有应用（包括商店应用）中都保持一致。User ID 主要用于在不同的应用间打通用户数据。[了解更多: 如何获取 User ID？](https://open.feishu.cn/document/uAjLw4CM/ugTN1YjL4UTN24CO1UjN/trouble-shooting/how-to-obtain-user-id), 默认值: `open_id`, 当值为 `user_id`, 字段权限要求: 获取用户 user ID
	UserIDs    []string `json:"user_ids,omitempty"`     // 用户 ID, 示例值: ["47f621ff"], 最大长度: `100`
}

// GetPassportSessionResp ...
type GetPassportSessionResp struct {
	MaskSessions []*GetPassportSessionRespMaskSession `json:"mask_sessions,omitempty"` // 用户登录信息
}

// GetPassportSessionRespMaskSession ...
type GetPassportSessionRespMaskSession struct {
	CreateTime   string `json:"create_time,omitempty"`   // 创建时间
	TerminalType int64  `json:"terminal_type,omitempty"` // 客户端类型, 可选值有: 0: 未知, 1: 个人电脑, 2: 浏览器, 3: 安卓手机, 4: Apple手机, 5: 服务端
	UserID       string `json:"user_id,omitempty"`       // 用户ID
}

// getPassportSessionResp ...
type getPassportSessionResp struct {
	Code int64                   `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                  `json:"msg,omitempty"`  // 错误描述
	Data *GetPassportSessionResp `json:"data,omitempty"`
}
