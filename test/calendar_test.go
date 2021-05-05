package test

import (
	"testing"

	"github.com/chyroc/go-ptr"
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"

	"github.com/chyroc/lark"
)

func Test_Calendar_Failed(t *testing.T) {
	as := assert.New(t)

	t.Run("request failed", func(t *testing.T) {
		cli := AppALLPermission.Ins()
		cli.Mock().MockGetTenantAccessToken(mockGetTenantAccessTokenFailed)
		helpdeskCli := cli.Calendar()

		t.Run("", func(t *testing.T) {
			_, _, err := helpdeskCli.CreateCalendar(ctx, &lark.CreateCalendarReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := helpdeskCli.DeleteCalendar(ctx, &lark.DeleteCalendarReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := helpdeskCli.GetCalendar(ctx, &lark.GetCalendarReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := helpdeskCli.GetCalendarList(ctx, &lark.GetCalendarListReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := helpdeskCli.UpdateCalendar(ctx, &lark.UpdateCalendarReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := helpdeskCli.SearchCalendar(ctx, &lark.SearchCalendarReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})
		t.Run("", func(t *testing.T) {
			_, _, err := helpdeskCli.SubscribeCalendar(ctx, &lark.SubscribeCalendarReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := helpdeskCli.UnsubscribeCalendar(ctx, &lark.UnsubscribeCalendarReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})
	})

	t.Run("response failed", func(t *testing.T) {
		cli := AppNoPermission.Ins()
		moduleCli := cli.Calendar()

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CreateCalendar(ctx, &lark.CreateCalendarReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0)
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.DeleteCalendar(ctx, &lark.DeleteCalendarReq{
				CalendarID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0)
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetCalendar(ctx, &lark.GetCalendarReq{
				CalendarID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0)
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetCalendarList(ctx, &lark.GetCalendarListReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0)
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.UpdateCalendar(ctx, &lark.UpdateCalendarReq{
				CalendarID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0)
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.SearchCalendar(ctx, &lark.SearchCalendarReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0)
		})
		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.SubscribeCalendar(ctx, &lark.SubscribeCalendarReq{
				CalendarID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0)
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.UnsubscribeCalendar(ctx, &lark.UnsubscribeCalendarReq{
				CalendarID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0)
		})
	})
}

func Test_Calendar(t *testing.T) {
	as := assert.New(t)
	moduleCli := AppALLPermission.Ins().Calendar()

	t.Run("", func(t *testing.T) {
		calendarID := ""
		summary := "summary-test"

		{
			resp, _, err := moduleCli.CreateCalendar(ctx, &lark.CreateCalendarReq{
				Summary:      ptr.String("summary-test"),
				Description:  ptr.String("desc-test"),
				Permissions:  nil,
				Color:        nil,
				SummaryAlias: nil,
			})
			spew.Dump(resp, err)
			as.Nil(err)
			calendarID = resp.Calendar.CalendarID
		}

		{
			resp, _, err := moduleCli.GetCalendar(ctx, &lark.GetCalendarReq{
				CalendarID: calendarID,
			})
			spew.Dump(resp, err)
			as.Nil(err)
			as.Equal(summary, resp.Summary)
		}

		{
			resp, _, err := moduleCli.GetCalendarList(ctx, &lark.GetCalendarListReq{})
			spew.Dump(resp, err)
			as.Nil(err)
		}

		{
			resp, _, err := moduleCli.UpdateCalendar(ctx, &lark.UpdateCalendarReq{
				CalendarID: calendarID,
				Summary:    ptr.String(summary + "-update"),
			})
			spew.Dump(resp, err)
			as.Nil(err)
		}

		{
			resp, _, err := moduleCli.SubscribeCalendar(ctx, &lark.SubscribeCalendarReq{
				CalendarID: calendarID,
			})
			spew.Dump(resp, err)
			as.Nil(err)
		}

		{
			resp, _, err := moduleCli.UnsubscribeCalendar(ctx, &lark.UnsubscribeCalendarReq{
				CalendarID: calendarID,
			})
			spew.Dump(resp, err)
			as.Nil(err)
		}

		{
			resp, _, err := moduleCli.DeleteCalendar(ctx, &lark.DeleteCalendarReq{
				CalendarID: calendarID,
			})
			spew.Dump(resp, err)
			as.Nil(err)
		}
	})
}