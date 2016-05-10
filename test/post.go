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

// ShowPostOK test setup
func ShowPostOK(t *testing.T, ctrl app.PostController, postID int) *app.Post {
	return ShowPostOKCtx(t, context.Background(), ctrl, postID)
}

// ShowPostOKCtx test setup
func ShowPostOKCtx(t *testing.T, ctx context.Context, ctrl app.PostController, postID int) *app.Post {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/posts/%v", postID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "PostTest"), rw, req, nil)
	showCtx, err := app.NewShowPostContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	err = ctrl.Show(showCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	a, ok := resp.(*app.Post)
	if !ok {
		t.Errorf("invalid response media: got %+v, expected instance of app.Post", resp)
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

// ShowPostOKLink test setup
func ShowPostOKLink(t *testing.T, ctrl app.PostController, postID int) *app.PostLink {
	return ShowPostOKLinkCtx(t, context.Background(), ctrl, postID)
}

// ShowPostOKLinkCtx test setup
func ShowPostOKLinkCtx(t *testing.T, ctx context.Context, ctrl app.PostController, postID int) *app.PostLink {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/posts/%v", postID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "PostTest"), rw, req, nil)
	showCtx, err := app.NewShowPostContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	err = ctrl.Show(showCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	a, ok := resp.(*app.PostLink)
	if !ok {
		t.Errorf("invalid response media: got %+v, expected instance of app.PostLink", resp)
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

// ShowPostNotFound test setup
func ShowPostNotFound(t *testing.T, ctrl app.PostController, postID int) {
	ShowPostNotFoundCtx(t, context.Background(), ctrl, postID)
}

// ShowPostNotFoundCtx test setup
func ShowPostNotFoundCtx(t *testing.T, ctx context.Context, ctrl app.PostController, postID int) {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/posts/%v", postID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "PostTest"), rw, req, nil)
	showCtx, err := app.NewShowPostContext(goaCtx, service)
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

// CreatePostCreated test setup
func CreatePostCreated(t *testing.T, ctrl app.PostController, payload *app.CreatePostPayload) {
	CreatePostCreatedCtx(t, context.Background(), ctrl, payload)
}

// CreatePostCreatedCtx test setup
func CreatePostCreatedCtx(t *testing.T, ctx context.Context, ctrl app.PostController, payload *app.CreatePostPayload) {
	err := payload.Validate()
	if err != nil {
		panic(err)
	}
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("POST", fmt.Sprintf("/posts"), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "PostTest"), rw, req, nil)
	createCtx, err := app.NewCreatePostContext(goaCtx, service)
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

// UpdatePostNoContent test setup
func UpdatePostNoContent(t *testing.T, ctrl app.PostController, postID int, payload *app.UpdatePostPayload) {
	UpdatePostNoContentCtx(t, context.Background(), ctrl, postID, payload)
}

// UpdatePostNoContentCtx test setup
func UpdatePostNoContentCtx(t *testing.T, ctx context.Context, ctrl app.PostController, postID int, payload *app.UpdatePostPayload) {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("PUT", fmt.Sprintf("/posts/%v", postID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "PostTest"), rw, req, nil)
	updateCtx, err := app.NewUpdatePostContext(goaCtx, service)
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

// UpdatePostNotFound test setup
func UpdatePostNotFound(t *testing.T, ctrl app.PostController, postID int, payload *app.UpdatePostPayload) {
	UpdatePostNotFoundCtx(t, context.Background(), ctrl, postID, payload)
}

// UpdatePostNotFoundCtx test setup
func UpdatePostNotFoundCtx(t *testing.T, ctx context.Context, ctrl app.PostController, postID int, payload *app.UpdatePostPayload) {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("PUT", fmt.Sprintf("/posts/%v", postID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "PostTest"), rw, req, nil)
	updateCtx, err := app.NewUpdatePostContext(goaCtx, service)
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

// DeletePostNoContent test setup
func DeletePostNoContent(t *testing.T, ctrl app.PostController, postID int) {
	DeletePostNoContentCtx(t, context.Background(), ctrl, postID)
}

// DeletePostNoContentCtx test setup
func DeletePostNoContentCtx(t *testing.T, ctx context.Context, ctrl app.PostController, postID int) {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("DELETE", fmt.Sprintf("/posts/%v", postID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "PostTest"), rw, req, nil)
	deleteCtx, err := app.NewDeletePostContext(goaCtx, service)
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

// DeletePostNotFound test setup
func DeletePostNotFound(t *testing.T, ctrl app.PostController, postID int) {
	DeletePostNotFoundCtx(t, context.Background(), ctrl, postID)
}

// DeletePostNotFoundCtx test setup
func DeletePostNotFoundCtx(t *testing.T, ctx context.Context, ctrl app.PostController, postID int) {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("DELETE", fmt.Sprintf("/posts/%v", postID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "PostTest"), rw, req, nil)
	deleteCtx, err := app.NewDeletePostContext(goaCtx, service)
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
