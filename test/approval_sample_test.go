// Code generated by lark_sdk_gen. DO NOT EDIT.

package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/chyroc/lark"
)

func Test_Approval_Sample_Failed(t *testing.T) {
	as := assert.New(t)

	t.Run("request failed", func(t *testing.T) {
		cli := AppALLPermission.Ins()
		cli.Mock().MockGetTenantAccessToken(mockGetTenantAccessTokenFailed)
		cli.Mock().MockGetAppAccessToken(mockGetTenantAccessTokenFailed)
		moduleCli := cli.Approval

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetApproval(ctx, &lark.GetApprovalReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetApprovalInstanceList(ctx, &lark.GetApprovalInstanceListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetApprovalInstance(ctx, &lark.GetApprovalInstanceReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CreateApprovalInstance(ctx, &lark.CreateApprovalInstanceReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.ApproveApprovalInstance(ctx, &lark.ApproveApprovalInstanceReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.RejectApprovalInstance(ctx, &lark.RejectApprovalInstanceReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.TransferApprovalInstance(ctx, &lark.TransferApprovalInstanceReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CancelApprovalInstance(ctx, &lark.CancelApprovalInstanceReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.UploadApprovalFile(ctx, &lark.UploadApprovalFileReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.SearchApprovalInstance(ctx, &lark.SearchApprovalInstanceReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CreateApprovalCarbonCopy(ctx, &lark.CreateApprovalCarbonCopyReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})
	})

	t.Run("request mock failed", func(t *testing.T) {
		cli := AppALLPermission.Ins()
		moduleCli := cli.Approval

		t.Run("", func(t *testing.T) {
			cli.Mock().MockApprovalGetApproval(func(ctx context.Context, request *lark.GetApprovalReq, options ...lark.MethodOptionFunc) (*lark.GetApprovalResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockApprovalGetApproval()

			_, _, err := moduleCli.GetApproval(ctx, &lark.GetApprovalReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockApprovalGetApprovalInstanceList(func(ctx context.Context, request *lark.GetApprovalInstanceListReq, options ...lark.MethodOptionFunc) (*lark.GetApprovalInstanceListResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockApprovalGetApprovalInstanceList()

			_, _, err := moduleCli.GetApprovalInstanceList(ctx, &lark.GetApprovalInstanceListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockApprovalGetApprovalInstance(func(ctx context.Context, request *lark.GetApprovalInstanceReq, options ...lark.MethodOptionFunc) (*lark.GetApprovalInstanceResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockApprovalGetApprovalInstance()

			_, _, err := moduleCli.GetApprovalInstance(ctx, &lark.GetApprovalInstanceReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockApprovalCreateApprovalInstance(func(ctx context.Context, request *lark.CreateApprovalInstanceReq, options ...lark.MethodOptionFunc) (*lark.CreateApprovalInstanceResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockApprovalCreateApprovalInstance()

			_, _, err := moduleCli.CreateApprovalInstance(ctx, &lark.CreateApprovalInstanceReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockApprovalApproveApprovalInstance(func(ctx context.Context, request *lark.ApproveApprovalInstanceReq, options ...lark.MethodOptionFunc) (*lark.ApproveApprovalInstanceResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockApprovalApproveApprovalInstance()

			_, _, err := moduleCli.ApproveApprovalInstance(ctx, &lark.ApproveApprovalInstanceReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockApprovalRejectApprovalInstance(func(ctx context.Context, request *lark.RejectApprovalInstanceReq, options ...lark.MethodOptionFunc) (*lark.RejectApprovalInstanceResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockApprovalRejectApprovalInstance()

			_, _, err := moduleCli.RejectApprovalInstance(ctx, &lark.RejectApprovalInstanceReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockApprovalTransferApprovalInstance(func(ctx context.Context, request *lark.TransferApprovalInstanceReq, options ...lark.MethodOptionFunc) (*lark.TransferApprovalInstanceResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockApprovalTransferApprovalInstance()

			_, _, err := moduleCli.TransferApprovalInstance(ctx, &lark.TransferApprovalInstanceReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockApprovalCancelApprovalInstance(func(ctx context.Context, request *lark.CancelApprovalInstanceReq, options ...lark.MethodOptionFunc) (*lark.CancelApprovalInstanceResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockApprovalCancelApprovalInstance()

			_, _, err := moduleCli.CancelApprovalInstance(ctx, &lark.CancelApprovalInstanceReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockApprovalUploadApprovalFile(func(ctx context.Context, request *lark.UploadApprovalFileReq, options ...lark.MethodOptionFunc) (*lark.UploadApprovalFileResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockApprovalUploadApprovalFile()

			_, _, err := moduleCli.UploadApprovalFile(ctx, &lark.UploadApprovalFileReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockApprovalSearchApprovalInstance(func(ctx context.Context, request *lark.SearchApprovalInstanceReq, options ...lark.MethodOptionFunc) (*lark.SearchApprovalInstanceResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockApprovalSearchApprovalInstance()

			_, _, err := moduleCli.SearchApprovalInstance(ctx, &lark.SearchApprovalInstanceReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockApprovalCreateApprovalCarbonCopy(func(ctx context.Context, request *lark.CreateApprovalCarbonCopyReq, options ...lark.MethodOptionFunc) (*lark.CreateApprovalCarbonCopyResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockApprovalCreateApprovalCarbonCopy()

			_, _, err := moduleCli.CreateApprovalCarbonCopy(ctx, &lark.CreateApprovalCarbonCopyReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})
	})

	t.Run("response is failed", func(t *testing.T) {
		cli := AppNoPermission.Ins()
		moduleCli := cli.Approval

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetApproval(ctx, &lark.GetApprovalReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetApprovalInstanceList(ctx, &lark.GetApprovalInstanceListReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetApprovalInstance(ctx, &lark.GetApprovalInstanceReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CreateApprovalInstance(ctx, &lark.CreateApprovalInstanceReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.ApproveApprovalInstance(ctx, &lark.ApproveApprovalInstanceReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.RejectApprovalInstance(ctx, &lark.RejectApprovalInstanceReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.TransferApprovalInstance(ctx, &lark.TransferApprovalInstanceReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CancelApprovalInstance(ctx, &lark.CancelApprovalInstanceReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.UploadApprovalFile(ctx, &lark.UploadApprovalFileReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.SearchApprovalInstance(ctx, &lark.SearchApprovalInstanceReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CreateApprovalCarbonCopy(ctx, &lark.CreateApprovalCarbonCopyReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})
	})
}
