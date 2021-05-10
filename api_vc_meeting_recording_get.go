package lark

import (
	"context"
)

// GetMeetingRecording
//
// > 获取一个会议的录制文件
// 会议结束后并且收到了"录制完成"的事件方可获取录制文件；只有会议owner（通过开放平台预约的会议即为预约人）有权限获取；录制时间太短(<5s)有可能无法生成录制文件
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/vc-v1/meeting.recording/get
func (r *VCAPI) GetMeetingRecording(ctx context.Context, request *GetMeetingRecordingReq) (*GetMeetingRecordingResp, *Response, error) {
	req := &RawRequestReq{
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/vc/v1/meetings/:meeting_id/recording",
		Body:                  request,
		NeedTenantAccessToken: false,
		NeedAppAccessToken:    false,
		NeedHelpdeskAuth:      false,
		IsFile:                false,
	}
	resp := new(getMeetingRecordingResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	if err != nil {
		return nil, response, err
	} else if resp.Code != 0 {
		return nil, response, NewError("VC", "GetMeetingRecording", resp.Code, resp.Msg)
	}

	return resp.Data, response, nil
}

type GetMeetingRecordingReq struct {
	MeetingID string `path:"meeting_id" json:"-"` // 会议id，示例值："6911188411932033028"
}

type getMeetingRecordingResp struct {
	Code int                      `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                   `json:"msg,omitempty"`  // 错误描述
	Data *GetMeetingRecordingResp `json:"data,omitempty"` // -
}

type GetMeetingRecordingResp struct {
	Recording *GetMeetingRecordingRespRecording `json:"recording,omitempty"` // 录制文件数据
}

type GetMeetingRecordingRespRecording struct {
	URL      string `json:"url,omitempty"`      // 录制文件URL
	Duration string `json:"duration,omitempty"` // 录制总时长（单位msec）
}