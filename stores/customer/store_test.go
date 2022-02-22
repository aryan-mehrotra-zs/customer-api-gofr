package customer

import (
	"context"
	"regexp"
	"testing"

	"developer.zopsmart.com/go/gofr/pkg/datastore"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/DATA-DOG/go-sqlmock"

	"example.com/customer-api/models"
	"example.com/customer-api/stores"
)

func initialiseTest(t *testing.T) (sqlmock.Sqlmock, *gofr.Context, stores.Customer) {
	db, mockDB, err := sqlmock.New()
	if err != nil {
		t.Errorf("error in creating mockDB :%v", err)
	}

	app := gofr.Gofr{DataStore: datastore.DataStore{ORM: db}}
	ctx := gofr.NewContext(nil, nil, &app)

	ctx.Context = context.TODO()

	store := New()

	return mockDB, ctx, store
}

func TestStore_Create(t *testing.T) {
	mock, ctx, store := initialiseTest(t)

	row := sqlmock.NewRows([]string{"id"}).AddRow(1)

	mock.ExpectQuery(regexp.QuoteMeta(create)).WithArgs("Aryan").WillReturnRows(row)

	id, err := store.Create(ctx, models.Customer{Name: "Aryan"})

	if err != nil {
		t.Errorf("[TEST], failed. DESC success\n GOT: %v \nWANT: %v", err, nil)
	}

	if id != 1 {
		t.Errorf("[TEST], failed. DESC success\n GOT: %v \nWANT: %v", id, 1)
	}
}

func TestStore_Get(t *testing.T) {
	mock, ctx, store := initialiseTest(t)

	row := sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "Aryan")
	mock.ExpectQuery(regexp.QuoteMeta(get)).WithArgs("1").WillReturnRows(row)

	customer, err := store.Get(ctx, "1")
	if err != nil {
		t.Errorf("[TEST], failed. DESC success\n GOT: %v \nWANT: %v", err, nil)
	}

	if customer != (models.Customer{ID: 1, Name: "Aryan"}) {
		t.Errorf("[TEST], failed. DESC success\n GOT: %v \nWANT: %v", customer, models.Customer{ID: 1, Name: "Aryan"})
	}
}

func TestStore_Update(t *testing.T) {
	mock, ctx, store := initialiseTest(t)

	mock.ExpectExec(regexp.QuoteMeta(update)).WithArgs("Aryan", 1).WillReturnResult(sqlmock.NewResult(1, 1)).WillReturnError(nil)

	err := store.Update(ctx, models.Customer{ID: 1, Name: "Aryan"})
	if err != nil {
		t.Errorf("[TEST], failed. DESC success\n GOT: %v \nWANT: %v", err, nil)
	}
}

func TestStore_Delete(t *testing.T) {
	mock, ctx, store := initialiseTest(t)

	mock.ExpectExec(regexp.QuoteMeta(delete)).WithArgs("1").WillReturnResult(sqlmock.NewResult(1, 1)).WillReturnError(nil)

	err := store.Delete(ctx, "1")
	if err != nil {
		t.Errorf("[TEST], failed. DESC success\n GOT: %v \nWANT: %v", err, nil)
	}
}
