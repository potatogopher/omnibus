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

// ShowUserNotFound test setup
func ShowUserNotFound(t *testing.T, ctrl app.UserController, userID int) {
	ShowUserNotFoundCtx(t, context.Background(), ctrl, userID)
}

// ShowUserNotFoundCtx test setup
func ShowUserNotFoundCtx(t *testing.T, ctx context.Context, ctrl app.UserController, userID int) {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/users/%v", userID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "UserTest"), rw, req, nil)
	showCtx, err := app.NewShowUserContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	err = ctrl.Show(showCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	if rw.Code != 404 {
		t.Errorf("invalid response status code: got %+v, expected 404", rw.Code)
	}

}

// ShowUserOK test setup
func ShowUserOK(t *testing.T, ctrl app.UserController, userID int) *app.User {
	return ShowUserOKCtx(t, context.Background(), ctrl, userID)
}

// ShowUserOKCtx test setup
func ShowUserOKCtx(t *testing.T, ctx context.Context, ctrl app.UserController, userID int) *app.User {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/users/%v", userID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "UserTest"), rw, req, nil)
	showCtx, err := app.NewShowUserContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	err = ctrl.Show(showCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	a, ok := resp.(*app.User)
	if !ok {
		t.Errorf("invalid response media: got %+v, expected instance of app.User", resp)
	}

	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}

	err = a.Validate()
	if err != nil {
		t.Errorf("invalid response payload: got %v", err)
	}
	return a

}

// ShowUserOKTiny test setup
func ShowUserOKTiny(t *testing.T, ctrl app.UserController, userID int) *app.UserTiny {
	return ShowUserOKTinyCtx(t, context.Background(), ctrl, userID)
}

// ShowUserOKTinyCtx test setup
func ShowUserOKTinyCtx(t *testing.T, ctx context.Context, ctrl app.UserController, userID int) *app.UserTiny {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/users/%v", userID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "UserTest"), rw, req, nil)
	showCtx, err := app.NewShowUserContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	err = ctrl.Show(showCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	a, ok := resp.(*app.UserTiny)
	if !ok {
		t.Errorf("invalid response media: got %+v, expected instance of app.UserTiny", resp)
	}

	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}

	err = a.Validate()
	if err != nil {
		t.Errorf("invalid response payload: got %v", err)
	}
	return a

}

// ShowUserOKLink test setup
func ShowUserOKLink(t *testing.T, ctrl app.UserController, userID int) *app.UserLink {
	return ShowUserOKLinkCtx(t, context.Background(), ctrl, userID)
}

// ShowUserOKLinkCtx test setup
func ShowUserOKLinkCtx(t *testing.T, ctx context.Context, ctrl app.UserController, userID int) *app.UserLink {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/users/%v", userID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "UserTest"), rw, req, nil)
	showCtx, err := app.NewShowUserContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	err = ctrl.Show(showCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	a, ok := resp.(*app.UserLink)
	if !ok {
		t.Errorf("invalid response media: got %+v, expected instance of app.UserLink", resp)
	}

	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}

	err = a.Validate()
	if err != nil {
		t.Errorf("invalid response payload: got %v", err)
	}
	return a

}

// CreateUserCreated test setup
func CreateUserCreated(t *testing.T, ctrl app.UserController, payload *app.CreateUserPayload) {
	CreateUserCreatedCtx(t, context.Background(), ctrl, payload)
}

// CreateUserCreatedCtx test setup
func CreateUserCreatedCtx(t *testing.T, ctx context.Context, ctrl app.UserController, payload *app.CreateUserPayload) {
	err := payload.Validate()
	if err != nil {
		panic(err)
	}
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("POST", fmt.Sprintf("/users"), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "UserTest"), rw, req, nil)
	createCtx, err := app.NewCreateUserContext(goaCtx, service)
	createCtx.Payload = payload

	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	err = ctrl.Create(createCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	if rw.Code != 201 {
		t.Errorf("invalid response status code: got %+v, expected 201", rw.Code)
	}

}

// UpdateUserNoContent test setup
func UpdateUserNoContent(t *testing.T, ctrl app.UserController, userID int, payload *app.UpdateUserPayload) {
	UpdateUserNoContentCtx(t, context.Background(), ctrl, userID, payload)
}

// UpdateUserNoContentCtx test setup
func UpdateUserNoContentCtx(t *testing.T, ctx context.Context, ctrl app.UserController, userID int, payload *app.UpdateUserPayload) {
	err := payload.Validate()
	if err != nil {
		panic(err)
	}
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("PUT", fmt.Sprintf("/users/%v", userID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "UserTest"), rw, req, nil)
	updateCtx, err := app.NewUpdateUserContext(goaCtx, service)
	updateCtx.Payload = payload

	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	err = ctrl.Update(updateCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	if rw.Code != 204 {
		t.Errorf("invalid response status code: got %+v, expected 204", rw.Code)
	}

}

// UpdateUserNotFound test setup
func UpdateUserNotFound(t *testing.T, ctrl app.UserController, userID int, payload *app.UpdateUserPayload) {
	UpdateUserNotFoundCtx(t, context.Background(), ctrl, userID, payload)
}

// UpdateUserNotFoundCtx test setup
func UpdateUserNotFoundCtx(t *testing.T, ctx context.Context, ctrl app.UserController, userID int, payload *app.UpdateUserPayload) {
	err := payload.Validate()
	if err != nil {
		panic(err)
	}
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("PUT", fmt.Sprintf("/users/%v", userID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "UserTest"), rw, req, nil)
	updateCtx, err := app.NewUpdateUserContext(goaCtx, service)
	updateCtx.Payload = payload

	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	err = ctrl.Update(updateCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	if rw.Code != 404 {
		t.Errorf("invalid response status code: got %+v, expected 404", rw.Code)
	}

}

// DeleteUserNoContent test setup
func DeleteUserNoContent(t *testing.T, ctrl app.UserController, userID int) {
	DeleteUserNoContentCtx(t, context.Background(), ctrl, userID)
}

// DeleteUserNoContentCtx test setup
func DeleteUserNoContentCtx(t *testing.T, ctx context.Context, ctrl app.UserController, userID int) {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("DELETE", fmt.Sprintf("/users/%v", userID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "UserTest"), rw, req, nil)
	deleteCtx, err := app.NewDeleteUserContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	err = ctrl.Delete(deleteCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	if rw.Code != 204 {
		t.Errorf("invalid response status code: got %+v, expected 204", rw.Code)
	}

}

// DeleteUserNotFound test setup
func DeleteUserNotFound(t *testing.T, ctrl app.UserController, userID int) {
	DeleteUserNotFoundCtx(t, context.Background(), ctrl, userID)
}

// DeleteUserNotFoundCtx test setup
func DeleteUserNotFoundCtx(t *testing.T, ctx context.Context, ctrl app.UserController, userID int) {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("DELETE", fmt.Sprintf("/users/%v", userID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "UserTest"), rw, req, nil)
	deleteCtx, err := app.NewDeleteUserContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	err = ctrl.Delete(deleteCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	if rw.Code != 404 {
		t.Errorf("invalid response status code: got %+v, expected 404", rw.Code)
	}

}
