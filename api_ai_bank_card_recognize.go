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

// RecognizeAiBankCard 银行卡识别接口, 支持JPG/JPEG/PNG/BMP四种文件类型的一次性的识别。
//
// 单租户限流: 10QPS, 同租户下的应用没有限流, 共享本租户的 10QPS 限流
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/ai/document_ai-v1/bank_card/recognize
func (r *AIService) RecognizeAiBankCard(ctx context.Context, request *RecognizeAiBankCardReq, options ...MethodOptionFunc) (*RecognizeAiBankCardResp, *Response, error) {
	if r.cli.mock.mockAIRecognizeAiBankCard != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] AI#RecognizeAiBankCard mock enable")
		return r.cli.mock.mockAIRecognizeAiBankCard(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "AI",
		API:                   "RecognizeAiBankCard",
		Method:                "POST",
		URL:                   r.cli.openBaseURL + "/open-apis/document_ai/v1/bank_card/recognize",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(recognizeAiBankCardResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

// MockAIRecognizeAiBankCard mock AIRecognizeAiBankCard method
func (r *Mock) MockAIRecognizeAiBankCard(f func(ctx context.Context, request *RecognizeAiBankCardReq, options ...MethodOptionFunc) (*RecognizeAiBankCardResp, *Response, error)) {
	r.mockAIRecognizeAiBankCard = f
}

// UnMockAIRecognizeAiBankCard un-mock AIRecognizeAiBankCard method
func (r *Mock) UnMockAIRecognizeAiBankCard() {
	r.mockAIRecognizeAiBankCard = nil
}

// RecognizeAiBankCardReq ...
type RecognizeAiBankCardReq struct {
	File *RecognizeAiBankCardReqFile `json:"file,omitempty"` // 识别的银行卡源文件, 示例值: file binary
}

// RecognizeAiBankCardResp ...
type RecognizeAiBankCardResp struct {
	BankCard *RecognizeAiBankCardRespBankCard `json:"bank_card,omitempty"` // 银行卡信息
}

// RecognizeAiBankCardRespBankCard ...
type RecognizeAiBankCardRespBankCard struct {
	Entities []*RecognizeAiBankCardRespBankCardEntitie `json:"entities,omitempty"` // 识别出的实体类型
}

// RecognizeAiBankCardRespBankCardEntitie ...
type RecognizeAiBankCardRespBankCardEntitie struct {
	Type  string `json:"type,omitempty"`  // 识别的字段种类, 可选值有: card_number: 银行卡卡号, date_of_expiry: 有效日期
	Value string `json:"value,omitempty"` // 识别出字段的文本信息
}

// recognizeAiBankCardResp ...
type recognizeAiBankCardResp struct {
	Code int64                    `json:"code,omitempty"` // 错误码, 非 0 表示失败
	Msg  string                   `json:"msg,omitempty"`  // 错误描述
	Data *RecognizeAiBankCardResp `json:"data,omitempty"`
}