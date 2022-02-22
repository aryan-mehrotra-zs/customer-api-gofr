package customer

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"developer.zopsmart.com/go/gofr/pkg/gofr/request"
	"developer.zopsmart.com/go/gofr/pkg/gofr/responder"
	"github.com/golang/mock/gomock"

	"example.com/customer-api/models"
	"example.com/customer-api/services"
)

func initialiseTest(t *testing.T, method string, body io.Reader, pathParam map[string]string) (*gofr.Context, handler, *services.MockCustomer) {
	app := gofr.New()

	ctrl := gomock.NewController(t)
	mockService := services.NewMockCustomer(ctrl)

	h := New(mockService)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(method, "http://cars", body)

	req := request.NewHTTPRequest(r)
	res := responder.NewContextualResponder(w, r)

	ctx := gofr.NewContext(res, req, app)

	if pathParam != nil {
		ctx.SetPathParams(pathParam)
	}

	return ctx, h, mockService
}

func TestHandler_Create(t *testing.T) {
	cases := []struct {
		desc   string
		body   []byte
		input  models.Customer
		output int64
		err    error
	}{
		{"success", []byte(`{"name":"Aryan"}`), models.Customer{Name: "Aryan"}, int64(1),
			nil},
		{"failed", []byte(`{"name":""}`), models.Customer{Name: ""}, int64(0),
			errors.InvalidParam{Param: []string{"name"}},
		},
	}

	for i, tc := range cases {
		ctx, h, mockService := initialiseTest(t, http.MethodPost, bytes.NewReader(tc.body), nil)

		mockService.EXPECT().Create(ctx, tc.input).Return(tc.output, tc.err)

		output, err := h.Create(ctx)

		if !reflect.DeepEqual(err, tc.err) {
			t.Errorf("[TEST %d], failed. DESC %v\n GOT: %v \nWANT: %v", i, tc.desc, err, tc.err)
		}

		if output.(int64) != tc.output {
			t.Errorf("[TEST %v], failed. DESC %v\n GOT: %v \nWANT: %v", i, tc.desc, output, tc.output)
		}
	}
}

func TestHandler_CreateBindError(t *testing.T) {
	ctx, h, _ := initialiseTest(t, http.MethodPost, bytes.NewReader([]byte(`{"name":""`)), nil)

	output, err := h.Create(ctx)

	if !reflect.DeepEqual(err, errors.InvalidParam{Param: []string{"body"}}) {
		t.Errorf("[TEST], failed. DESC bind error\n GOT: %v \nWANT: %v", err, errors.InvalidParam{Param: []string{"body"}})
	}

	if output.(int64) != 0 {
		t.Errorf("[TEST], failed. DESC bind error\n GOT: %v \nWANT: %v", output, 0)
	}
}

func TestHandler_Get(t *testing.T) {
	cases := []struct {
		desc   string
		input  string
		output models.Customer
		err    error
	}{
		{"success", "1", models.Customer{ID: 1, Name: "Aryan"}, nil},
		{"failure", "0", models.Customer{}, errors.InvalidParam{Param: []string{"id"}}},
	}

	for i, tc := range cases {
		ctx, h, mockService := initialiseTest(t, http.MethodGet, http.NoBody, map[string]string{"id": tc.input})

		mockService.EXPECT().Get(ctx, tc.input).Return(tc.output, tc.err)

		resp, err := h.Get(ctx)

		if !reflect.DeepEqual(err, tc.err) {
			t.Errorf("[TEST %d], failed. DESC %v\n GOT: %v \nWANT: %v", i, tc.desc, err, tc.err)
		}

		if resp.(models.Customer) != tc.output {
			t.Errorf("[TEST], failed. DESC bind error\n GOT: %v \nWANT: %v", resp, tc.output)
		}
	}
}

func TestHandler_Update(t *testing.T) {
	cases := []struct {
		desc   string
		body   []byte
		input  models.Customer
		output models.Customer
		err    error
	}{
		{"success", []byte(`{"id":1,"name":"Umang"}`), models.Customer{ID: 1, Name: "Umang"}, models.Customer{ID: 1, Name: "Umang"}, nil},
		{"failure", []byte(`{"name":"Umang"}`), models.Customer{Name: "Umang"}, models.Customer{}, errors.InvalidParam{Param: []string{"id"}}},
	}

	for i, tc := range cases {
		ctx, h, mockService := initialiseTest(t, http.MethodPut, bytes.NewReader(tc.body), nil)

		mockService.EXPECT().Update(ctx, tc.input).Return(tc.output, tc.err)

		output, err := h.Update(ctx)

		if !reflect.DeepEqual(err, tc.err) {
			t.Errorf("[TEST %d], failed. DESC %v\n GOT: %v \nWANT: %v", i, tc.desc, err, tc.err)
		}

		if output.(models.Customer) != tc.output {
			t.Errorf("[TEST %d], failed. DESC %v\n GOT: %v \nWANT: %v", i, tc.desc, output, tc.output)
		}
	}
}

func TestHandler_UpdateBindError(t *testing.T) {
	ctx, h, _ := initialiseTest(t, http.MethodPut, bytes.NewReader([]byte(`{"name":""`)), nil)

	output, err := h.Update(ctx)

	if !reflect.DeepEqual(err, errors.InvalidParam{Param: []string{"body"}}) {
		t.Errorf("[TEST], failed. DESC bind error\n GOT: %v \nWANT: %v", err, errors.InvalidParam{Param: []string{"body"}})
	}

	if output.(models.Customer) != (models.Customer{}) {
		t.Errorf("[TEST], failed. DESC bind error\n GOT: %v \nWANT: %v", output, 0)
	}
}

func TestHandler_Delete(t *testing.T) {
	cases := []struct {
		desc  string
		input string
		err   error
	}{
		{"success", "1", nil},
		{"failure", "0", errors.InvalidParam{Param: []string{"id"}}},
	}

	for i, tc := range cases {
		ctx, h, mockService := initialiseTest(t, http.MethodDelete, nil, map[string]string{"id": tc.input})

		mockService.EXPECT().Delete(ctx, tc.input).Return(tc.err)

		_, err := h.Delete(ctx)

		if !reflect.DeepEqual(err, tc.err) {
			t.Errorf("[TEST %v], failed. DESC %v\n GOT: %v \nWANT: %v", i, tc.desc, err, tc.err)
		}
	}
}
