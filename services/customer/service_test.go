package customer

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"strconv"
	"testing"

	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"developer.zopsmart.com/go/gofr/pkg/gofr/request"
	"developer.zopsmart.com/go/gofr/pkg/gofr/responder"
	"github.com/golang/mock/gomock"

	"example.com/customer-api/models"
	"example.com/customer-api/stores"
)

func initialiseTest(t *testing.T) (*stores.MockCustomer, service, *gofr.Context) {
	app := gofr.New()

	ctrl := gomock.NewController(t)
	mockStore := stores.NewMockCustomer(ctrl)
	s := New(mockStore)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("", "http://cars", http.NoBody)

	req := request.NewHTTPRequest(r)
	res := responder.NewContextualResponder(w, r)

	ctx := gofr.NewContext(res, req, app)

	return mockStore, s, ctx
}

func TestService_Create(t *testing.T) {
	mockStore, s, ctx := initialiseTest(t)

	mockStore.EXPECT().Create(ctx, models.Customer{Name: "Aryan"}).Return(int64(1), nil)

	output, err := s.Create(ctx, models.Customer{Name: "Aryan"})

	if !reflect.DeepEqual(err, nil) {
		t.Errorf("[TEST], failed. DESC success\n GOT: %v \nWANT: %v", err, nil)
	}

	if output != 1 {
		t.Errorf("[TEST], failed. DESC success\n GOT: %v \nWANT: %v", output, 1)
	}
}

func TestService_CreateInvalidName(t *testing.T) {
	_, s, ctx := initialiseTest(t)

	output, err := s.Create(ctx, models.Customer{Name: ""})

	if !reflect.DeepEqual(err, errors.InvalidParam{Param: []string{"name"}}) {
		t.Errorf("[TEST], failed. DESC invalid name\n GOT: %v \nWANT: %v", err, errors.InvalidParam{Param: []string{"name"}})
	}

	if output != 0 {
		t.Errorf("[TEST], failed. DESC invalid name\n GOT: %v \nWANT: %v", output, 0)
	}
}

func TestService_Get(t *testing.T) {
	mockStore, s, ctx := initialiseTest(t)

	expOutput := models.Customer{ID: 1, Name: "Aryan"}

	mockStore.EXPECT().Get(ctx, "1").Return(expOutput, nil)

	output, err := s.Get(ctx, "1")

	if !reflect.DeepEqual(err, nil) {
		t.Errorf("[TEST], failed. DESC success\n GOT: %v \nWANT: %v", err, nil)
	}

	if output != expOutput {
		t.Errorf("[TEST], failed. DESC success\n GOT: %v \nWANT: %v", output, 1)
	}
}

func TestService_GetInvalidID(t *testing.T) {
	_, s, ctx := initialiseTest(t)

	expOutput := models.Customer{}

	output, err := s.Get(ctx, "-1")

	if !reflect.DeepEqual(err, errors.InvalidParam{Param: []string{"id"}}) {
		t.Errorf("[TEST], failed. DESC invalid id\n GOT: %v \nWANT: %v", err, errors.InvalidParam{Param: []string{"id"}})
	}

	if output != expOutput {
		t.Errorf("[TEST], failed. DESC invalid id\n GOT: %v \nWANT: %v", output, 0)
	}
}

func TestService_Update(t *testing.T) {
	mockStore, s, ctx := initialiseTest(t)

	customer := models.Customer{ID: 1, Name: "Aryan"}

	mockStore.EXPECT().Update(ctx, customer).Return(nil)
	mockStore.EXPECT().Get(ctx, strconv.Itoa(customer.ID)).Return(customer, nil)

	output, err := s.Update(ctx, customer)

	if !reflect.DeepEqual(err, nil) {
		t.Errorf("[TEST], failed. DESC success\n GOT: %v \nWANT: %v", err, nil)
	}

	if output != customer {
		t.Errorf("[TEST], failed. DESC success\n GOT: %v \nWANT: %v", output, customer)
	}
}

func TestService_UpdateInvalidCustomer(t *testing.T) {
	_, s, ctx := initialiseTest(t)

	customer := models.Customer{ID: 0, Name: ""}
	expOutput := models.Customer{}

	output, err := s.Update(ctx, customer)

	if !reflect.DeepEqual(err, errors.InvalidParam{Param: []string{"id", "name"}}) {
		t.Errorf("[TEST], failed. DESC invalid id\n GOT: %v \nWANT: %v", err, errors.InvalidParam{Param: []string{"id", "name"}})
	}

	if output != expOutput {
		t.Errorf("[TEST], failed. DESC invalid id\n GOT: %v \nWANT: %v", output, expOutput)
	}
}

func TestService_Delete(t *testing.T) {
	mockStore, s, ctx := initialiseTest(t)

	mockStore.EXPECT().Delete(ctx, "1").Return(nil)

	err := s.Delete(ctx, "1")

	if !reflect.DeepEqual(err, nil) {
		t.Errorf("[TEST], failed. DESC success\n GOT: %v \nWANT: %v", err, nil)
	}
}

func TestService_DeleteInvalidID(t *testing.T) {
	_, s, ctx := initialiseTest(t)

	err := s.Delete(ctx, "0")

	if !reflect.DeepEqual(err, errors.InvalidParam{Param: []string{"id"}}) {
		t.Errorf("[TEST], failed. DESC success\n GOT: %v \nWANT: %v", err, errors.InvalidParam{Param: []string{"id"}})
	}
}
