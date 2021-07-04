// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// GetBitableViewList 根据 app_token 和 table_id，获取数据表的所有视图
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table-view/list
func (r *BitableService) GetBitableViewList(ctx context.Context, request *GetBitableViewListReq, options ...MethodOptionFunc) (*GetBitableViewListResp, *Response, error) {
	if r.cli.mock.mockBitableGetBitableViewList != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Bitable#GetBitableViewList mock enable")
		return r.cli.mock.mockBitableGetBitableViewList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Bitable",
		API:                   "GetBitableViewList",
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/bitable/v1/apps/:app_token/tables/:table_id/views",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(getBitableViewListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockBitableGetBitableViewList(f func(ctx context.Context, request *GetBitableViewListReq, options ...MethodOptionFunc) (*GetBitableViewListResp, *Response, error)) {
	r.mockBitableGetBitableViewList = f
}

func (r *Mock) UnMockBitableGetBitableViewList() {
	r.mockBitableGetBitableViewList = nil
}

type GetBitableViewListReq struct {
	PageSize  *int64  `query:"page_size" json:"-"`  // 分页大小, 示例值：10, 最大值：`100`
	PageToken *string `query:"page_token" json:"-"` // 分页标记，第一次请求不填，表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token，下次遍历可采用该 page_token 获取查询结果, 示例值："vewTpR1urY"
	AppToken  string  `path:"app_token" json:"-"`   // bitable app token, 示例值："appbcbWCzen6D8dezhoCH2RpMAh"
	TableID   string  `path:"table_id" json:"-"`    // table id, 示例值："tblsRc9GRRXKqhvW"
}

type getBitableViewListResp struct {
	Code int64                   `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                  `json:"msg,omitempty"`  // 错误描述
	Data *GetBitableViewListResp `json:"data,omitempty"`
}

type GetBitableViewListResp struct {
	Items     []*GetBitableViewListRespItem `json:"items,omitempty"`      // 视图信息
	PageToken string                        `json:"page_token,omitempty"` // 分页标记，当 has_more 为 true 时，会同时返回新的 page_token，否则不返回 page_token
	HasMore   bool                          `json:"has_more,omitempty"`   // 是否还有更多项
}

type GetBitableViewListRespItem struct {
	ViewID   string `json:"view_id,omitempty"`   // 视图Id
	ViewName string `json:"view_name,omitempty"` // 视图名字
	ViewType string `json:"view_type,omitempty"` // 视图类型
}