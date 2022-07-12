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

// SortChatTab 会话标签页排序
//
// 注意事项:
// - 当前消息标签页固定为第一顺位, 不参与排序, 但是请求体中必须包含。
// - 请求体必须包含全部的TabID
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/im-v1/chat-tab/sort_tabs
func (r *ChatService) SortChatTab(ctx context.Context, request *SortChatTabReq, options ...MethodOptionFunc) (*SortChatTabResp, *Response, error) {
	if r.cli.mock.mockChatSortChatTab != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Chat#SortChatTab mock enable")
		return r.cli.mock.mockChatSortChatTab(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Chat",
		API:                   "SortChatTab",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/im/v1/chats/:chat_id/chat_tabs/sort_tabs",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(sortChatTabResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockChatSortChatTab mock ChatSortChatTab method
func (r *Mock) MockChatSortChatTab(f func(ctx context.Context, request *SortChatTabReq, options ...MethodOptionFunc) (*SortChatTabResp, *Response, error)) {
	r.mockChatSortChatTab = f
}

// UnMockChatSortChatTab un-mock ChatSortChatTab method
func (r *Mock) UnMockChatSortChatTab() {
	r.mockChatSortChatTab = nil
}

// SortChatTabReq ...
type SortChatTabReq struct {
	ChatID string   `path:"chat_id" json:"-"`  // 群ID, 示例值: "oc_a0553eda9014c201e6969b478895c230"
	TabIDs []string `json:"tab_ids,omitempty"` // 会话标签页ID列表, 示例值: ["7101214603622940671", "7101214603622940672"]
}

// SortChatTabResp ...
type SortChatTabResp struct {
	ChatTabs []*SortChatTabRespChatTab `json:"chat_tabs,omitempty"` // 会话标签页
}

// SortChatTabRespChatTab ...
type SortChatTabRespChatTab struct {
	TabID      string                            `json:"tab_id,omitempty"`      // TabID
	TabName    string                            `json:"tab_name,omitempty"`    // Tab名称
	TabType    string                            `json:"tab_type,omitempty"`    // Tab类型, 可选值有: message: 消息类型, doc_list: 云文档列表, doc: 文档, pin: Pin, meeting_minute: 会议纪要, chat_announcement: 群公告, url: URL, file: 文件
	TabContent *SortChatTabRespChatTabTabContent `json:"tab_content,omitempty"` // Tab内容
}

// SortChatTabRespChatTabTabContent ...
type SortChatTabRespChatTabTabContent struct {
	URL           string `json:"url,omitempty"`            // URL类型
	Doc           string `json:"doc,omitempty"`            // Doc链接
	MeetingMinute string `json:"meeting_minute,omitempty"` // 会议纪要
}

// sortChatTabResp ...
type sortChatTabResp struct {
	Code int64            `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string           `json:"msg,omitempty"`  // 错误描述
	Data *SortChatTabResp `json:"data,omitempty"`
}