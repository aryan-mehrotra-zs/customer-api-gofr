package customer

// func initialiseTest(t *testing.T) (sqlmock.Sqlmock, *gofr.Context, stores.Customer) {
// 	db, mockDB, err := sqlmock.New()
// 	if err != nil {
// 		t.Errorf("error in creating mockDB :%v", err)
// 	}
//
// 	app := gofr.Gofr{DataStore: datastore.DataStore{ORM: db}}
// 	ctx := gofr.NewContext(nil, nil, &app)
//
// 	store := New()
//
// 	return mockDB, ctx, store
// }
//
// func TestStore_Create(t *testing.T) {
// 	mock, ctx, store := initialiseTest(t)
//
// 	row := sqlmock.NewRows([]string{"id"}).AddRow(1)
//
// 	mock.ExpectQuery(create).WithArgs("Aryan").WillReturnRows(row)
//
// 	id, err := store.Create(ctx, models.Customer{Name: "Aryan"})
//
// 	if err != nil {
// 		t.Errorf("error")
// 	}
//
// 	if id != 1 {
// 		t.Errorf("id")
// 	}
// }
//
// func TestStore_Get(t *testing.T) {
// 	mock, ctx, store := initialiseTest(t)
//
// 	row := sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "Aryan")
// 	mock.ExpectQuery(get).WithArgs("1").WillReturnRows(row)
// 	customer, err := store.Get(ctx, "1")
// 	if err != nil {
// 		t.Errorf("nil")
// 	}
//
// 	if customer != (models.Customer{ID: 1, Name: "Aryan"}) {
// 		t.Errorf("name")
// 	}
// }
