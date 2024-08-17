package handler

import (
	"io"
	"net/http/httptest"

	"github.com/labstack/echo/v4"
)

type (
	mockEchoContext struct {
		echo.Context

		mockParam      func() string
		mockBind       func(i interface{}) error
		mockAttachment func() error
	}
)

func newMockEchoContext(m *mockEchoContext) echo.Context {
	if m == nil {
		m = &mockEchoContext{}
	}

	m.Context = echo.New().NewContext(
		httptest.NewRequest(echo.GET, "/", nil),
		httptest.NewRecorder(),
	)

	return m
}

func (m *mockEchoContext) getResponseBody() []byte {
	httpResp := m.Response().Writer.(*httptest.ResponseRecorder)
	body, _ := io.ReadAll(httpResp.Body)
	return body
}

func (m *mockEchoContext) Param(name string) string {
	if m.mockParam != nil {
		return m.mockParam()
	}
	return ""
}
func (m *mockEchoContext) Bind(i interface{}) error {
	if m.mockBind != nil {
		return m.mockBind(i)
	}
	return nil
}
func (m *mockEchoContext) Attachment(file string, name string) error {
	if m.mockAttachment != nil {
		return m.mockAttachment()
	}
	return nil
}
