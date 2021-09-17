// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// GetContactGroupList 通过该接口可查询企业的用户组列表，如果应用的通讯录权限范围是“全部员工”，则可获取企业全部用户组列表。如果应用的通讯录权限范围不是“全部员工”，则仅可获取通讯录权限范围内的用户组。[点击了解通讯录权限范围](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/scope/overview)。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/contact-v3/group/simplelist
func (r *ContactService) GetContactGroupList(ctx context.Context, request *GetContactGroupListReq, options ...MethodOptionFunc) (*GetContactGroupListResp, *Response, error) {
	if r.cli.mock.mockContactGetContactGroupList != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Contact#GetContactGroupList mock enable")
		return r.cli.mock.mockContactGetContactGroupList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Contact",
		API:                   "GetContactGroupList",
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/contact/v3/group/simplelist",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getContactGroupListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockContactGetContactGroupList(f func(ctx context.Context, request *GetContactGroupListReq, options ...MethodOptionFunc) (*GetContactGroupListResp, *Response, error)) {
	r.mockContactGetContactGroupList = f
}

func (r *Mock) UnMockContactGetContactGroupList() {
	r.mockContactGetContactGroupList = nil
}

type GetContactGroupListReq struct {
	PageSize  *int64  `query:"page_size" json:"-"`  // 分页大小, 示例值：50, 最大值：`100`
	PageToken *string `query:"page_token" json:"-"` // 分页标记，第一次请求不填，表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token，下次遍历可采用该 page_token 获取查询结果, 示例值："AQD9/Rn9eij9Pm39ED40/dk53s4Ebp882DYfFaPFbz00L4CMZJrqGdzNyc8BcZtDbwVUvRmQTvyMYicnGWrde9X56TgdBuS+JKiSIkdexPw="
	Type      *int64  `query:"type" json:"-"`       // 用户组类型, 示例值：1, 可选值有: `1`：普通用户组, 默认值: `1`
}

type getContactGroupListResp struct {
	Code int64                    `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                   `json:"msg,omitempty"`  // 错误描述
	Data *GetContactGroupListResp `json:"data,omitempty"`
}

type GetContactGroupListResp struct {
	Grouplist []*GetContactGroupListRespGroup `json:"grouplist,omitempty"`  // 用户组列表
	PageToken string                          `json:"page_token,omitempty"` // 分页标记，当 has_more 为 true 时，会同时返回新的 page_token，否则不返回 page_token
	HasMore   bool                            `json:"has_more,omitempty"`   // 是否还有更多项
}

type GetContactGroupListRespGroup struct {
	ID                    string `json:"id,omitempty"`                      // 用户组ID
	Name                  string `json:"name,omitempty"`                    // 用户组名字
	Description           string `json:"description,omitempty"`             // 用户组描述
	MemberUserCount       int64  `json:"member_user_count,omitempty"`       // 用户组成员中用户的数量
	MemberDepartmentCount int64  `json:"member_department_count,omitempty"` // 用户组成员中部门的数量
}