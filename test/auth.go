package test

import (
	"bytes"
	"fmt"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/goatest"
	"goa-blog/app"
	"golang.org/x/net/context"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TokenAuthOK test setup
func TokenAuthOK(t *testing.T, ctrl app.AuthController, payload *app.TokenAuthPayload) *app.Authorize {
	return TokenAuthOKCtx(t, context.Background(), ctrl, payload)
}

// TokenAuthOKCtx test setup
func TokenAuthOKCtx(t *testing.T, ctx context.Context, ctrl app.AuthController, payload *app.TokenAuthPayload) *app.Authorize {
	err := payload.Validate()
	if err != nil {
		panic(err)
	}
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("POST", fmt.Sprintf("/auth/token"), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "AuthTest"), rw, req, nil)
	tokenCtx, err := app.NewTokenAuthContext(goaCtx, service)
	tokenCtx.Payload = payload

	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	err = ctrl.Token(tokenCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	a, ok := resp.(*app.Authorize)
	if !ok {
		t.Errorf("invalid response media: got %+v, expected instance of app.Authorize", resp)
	}

	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	return a

}
