// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/julysNICK/stock_system/db/sqlc (interfaces: StoreDB)

// Package mockdb is a generated GoMock package.
package mockdb

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	db "github.com/julysNICK/stock_system/db/sqlc"
)

// MockStoreDB is a mock of StoreDB interface.
type MockStoreDB struct {
	ctrl     *gomock.Controller
	recorder *MockStoreDBMockRecorder
}

// MockStoreDBMockRecorder is the mock recorder for MockStoreDB.
type MockStoreDBMockRecorder struct {
	mock *MockStoreDB
}

// NewMockStoreDB creates a new mock instance.
func NewMockStoreDB(ctrl *gomock.Controller) *MockStoreDB {
	mock := &MockStoreDB{ctrl: ctrl}
	mock.recorder = &MockStoreDBMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStoreDB) EXPECT() *MockStoreDBMockRecorder {
	return m.recorder
}

// CreateProduct mocks base method.
func (m *MockStoreDB) CreateProduct(arg0 context.Context, arg1 db.CreateProductParams) (db.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProduct", arg0, arg1)
	ret0, _ := ret[0].(db.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProduct indicates an expected call of CreateProduct.
func (mr *MockStoreDBMockRecorder) CreateProduct(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProduct", reflect.TypeOf((*MockStoreDB)(nil).CreateProduct), arg0, arg1)
}

// CreateSale mocks base method.
func (m *MockStoreDB) CreateSale(arg0 context.Context, arg1 db.CreateSaleParams) (db.Sale, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSale", arg0, arg1)
	ret0, _ := ret[0].(db.Sale)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSale indicates an expected call of CreateSale.
func (mr *MockStoreDBMockRecorder) CreateSale(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSale", reflect.TypeOf((*MockStoreDB)(nil).CreateSale), arg0, arg1)
}

// CreateStockAlert mocks base method.
func (m *MockStoreDB) CreateStockAlert(arg0 context.Context, arg1 db.CreateStockAlertParams) (db.StockAlert, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateStockAlert", arg0, arg1)
	ret0, _ := ret[0].(db.StockAlert)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateStockAlert indicates an expected call of CreateStockAlert.
func (mr *MockStoreDBMockRecorder) CreateStockAlert(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateStockAlert", reflect.TypeOf((*MockStoreDB)(nil).CreateStockAlert), arg0, arg1)
}

// CreateStore mocks base method.
func (m *MockStoreDB) CreateStore(arg0 context.Context, arg1 db.CreateStoreParams) (db.Store, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateStore", arg0, arg1)
	ret0, _ := ret[0].(db.Store)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateStore indicates an expected call of CreateStore.
func (mr *MockStoreDBMockRecorder) CreateStore(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateStore", reflect.TypeOf((*MockStoreDB)(nil).CreateStore), arg0, arg1)
}

// CreateSupplier mocks base method.
func (m *MockStoreDB) CreateSupplier(arg0 context.Context, arg1 db.CreateSupplierParams) (db.Supplier, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateSupplier", arg0, arg1)
	ret0, _ := ret[0].(db.Supplier)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateSupplier indicates an expected call of CreateSupplier.
func (mr *MockStoreDBMockRecorder) CreateSupplier(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateSupplier", reflect.TypeOf((*MockStoreDB)(nil).CreateSupplier), arg0, arg1)
}

// DeleteSale mocks base method.
func (m *MockStoreDB) DeleteSale(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSale", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSale indicates an expected call of DeleteSale.
func (mr *MockStoreDBMockRecorder) DeleteSale(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSale", reflect.TypeOf((*MockStoreDB)(nil).DeleteSale), arg0, arg1)
}

// DeleteStockAlert mocks base method.
func (m *MockStoreDB) DeleteStockAlert(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteStockAlert", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteStockAlert indicates an expected call of DeleteStockAlert.
func (mr *MockStoreDBMockRecorder) DeleteStockAlert(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteStockAlert", reflect.TypeOf((*MockStoreDB)(nil).DeleteStockAlert), arg0, arg1)
}

// GetProduct mocks base method.
func (m *MockStoreDB) GetProduct(arg0 context.Context, arg1 int64) (db.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProduct", arg0, arg1)
	ret0, _ := ret[0].(db.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProduct indicates an expected call of GetProduct.
func (mr *MockStoreDBMockRecorder) GetProduct(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProduct", reflect.TypeOf((*MockStoreDB)(nil).GetProduct), arg0, arg1)
}

// GetProductForUpdate mocks base method.
func (m *MockStoreDB) GetProductForUpdate(arg0 context.Context, arg1 int64) (db.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProductForUpdate", arg0, arg1)
	ret0, _ := ret[0].(db.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProductForUpdate indicates an expected call of GetProductForUpdate.
func (mr *MockStoreDBMockRecorder) GetProductForUpdate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProductForUpdate", reflect.TypeOf((*MockStoreDB)(nil).GetProductForUpdate), arg0, arg1)
}

// GetSale mocks base method.
func (m *MockStoreDB) GetSale(arg0 context.Context, arg1 int64) (db.Sale, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSale", arg0, arg1)
	ret0, _ := ret[0].(db.Sale)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSale indicates an expected call of GetSale.
func (mr *MockStoreDBMockRecorder) GetSale(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSale", reflect.TypeOf((*MockStoreDB)(nil).GetSale), arg0, arg1)
}

// GetStockAlert mocks base method.
func (m *MockStoreDB) GetStockAlert(arg0 context.Context, arg1 int64) (db.StockAlert, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStockAlert", arg0, arg1)
	ret0, _ := ret[0].(db.StockAlert)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStockAlert indicates an expected call of GetStockAlert.
func (mr *MockStoreDBMockRecorder) GetStockAlert(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStockAlert", reflect.TypeOf((*MockStoreDB)(nil).GetStockAlert), arg0, arg1)
}

// GetStore mocks base method.
func (m *MockStoreDB) GetStore(arg0 context.Context, arg1 int64) (db.Store, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStore", arg0, arg1)
	ret0, _ := ret[0].(db.Store)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStore indicates an expected call of GetStore.
func (mr *MockStoreDBMockRecorder) GetStore(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStore", reflect.TypeOf((*MockStoreDB)(nil).GetStore), arg0, arg1)
}

// GetStoreForUpdate mocks base method.
func (m *MockStoreDB) GetStoreForUpdate(arg0 context.Context, arg1 int64) (db.Store, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStoreForUpdate", arg0, arg1)
	ret0, _ := ret[0].(db.Store)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStoreForUpdate indicates an expected call of GetStoreForUpdate.
func (mr *MockStoreDBMockRecorder) GetStoreForUpdate(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStoreForUpdate", reflect.TypeOf((*MockStoreDB)(nil).GetStoreForUpdate), arg0, arg1)
}

// GetSupplier mocks base method.
func (m *MockStoreDB) GetSupplier(arg0 context.Context, arg1 int64) (db.Supplier, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSupplier", arg0, arg1)
	ret0, _ := ret[0].(db.Supplier)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSupplier indicates an expected call of GetSupplier.
func (mr *MockStoreDBMockRecorder) GetSupplier(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSupplier", reflect.TypeOf((*MockStoreDB)(nil).GetSupplier), arg0, arg1)
}

// ListProducts mocks base method.
func (m *MockStoreDB) ListProducts(arg0 context.Context, arg1 db.ListProductsParams) ([]db.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListProducts", arg0, arg1)
	ret0, _ := ret[0].([]db.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProducts indicates an expected call of ListProducts.
func (mr *MockStoreDBMockRecorder) ListProducts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProducts", reflect.TypeOf((*MockStoreDB)(nil).ListProducts), arg0, arg1)
}

// ListSales mocks base method.
func (m *MockStoreDB) ListSales(arg0 context.Context, arg1 db.ListSalesParams) ([]db.Sale, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSales", arg0, arg1)
	ret0, _ := ret[0].([]db.Sale)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSales indicates an expected call of ListSales.
func (mr *MockStoreDBMockRecorder) ListSales(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSales", reflect.TypeOf((*MockStoreDB)(nil).ListSales), arg0, arg1)
}

// ListStores mocks base method.
func (m *MockStoreDB) ListStores(arg0 context.Context, arg1 db.ListStoresParams) ([]db.Store, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListStores", arg0, arg1)
	ret0, _ := ret[0].([]db.Store)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListStores indicates an expected call of ListStores.
func (mr *MockStoreDBMockRecorder) ListStores(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListStores", reflect.TypeOf((*MockStoreDB)(nil).ListStores), arg0, arg1)
}

// ProductTx mocks base method.
func (m *MockStoreDB) ProductTx(arg0 context.Context, arg1 db.ProductTxParams) (db.ProductTxResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProductTx", arg0, arg1)
	ret0, _ := ret[0].(db.ProductTxResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProductTx indicates an expected call of ProductTx.
func (mr *MockStoreDBMockRecorder) ProductTx(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProductTx", reflect.TypeOf((*MockStoreDB)(nil).ProductTx), arg0, arg1)
}

// SaleTx mocks base method.
func (m *MockStoreDB) SaleTx(arg0 context.Context, arg1 db.SaleTxParams) (db.SaleTxResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaleTx", arg0, arg1)
	ret0, _ := ret[0].(db.SaleTxResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SaleTx indicates an expected call of SaleTx.
func (mr *MockStoreDBMockRecorder) SaleTx(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaleTx", reflect.TypeOf((*MockStoreDB)(nil).SaleTx), arg0, arg1)
}

// StockAlertTx mocks base method.
func (m *MockStoreDB) StockAlertTx(arg0 context.Context, arg1 db.StockAlertTxParams) (db.StockAlertTxResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StockAlertTx", arg0, arg1)
	ret0, _ := ret[0].(db.StockAlertTxResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StockAlertTx indicates an expected call of StockAlertTx.
func (mr *MockStoreDBMockRecorder) StockAlertTx(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StockAlertTx", reflect.TypeOf((*MockStoreDB)(nil).StockAlertTx), arg0, arg1)
}

// UpdateProduct mocks base method.
func (m *MockStoreDB) UpdateProduct(arg0 context.Context, arg1 db.UpdateProductParams) (db.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProduct", arg0, arg1)
	ret0, _ := ret[0].(db.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateProduct indicates an expected call of UpdateProduct.
func (mr *MockStoreDBMockRecorder) UpdateProduct(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProduct", reflect.TypeOf((*MockStoreDB)(nil).UpdateProduct), arg0, arg1)
}

// UpdateStockAlert mocks base method.
func (m *MockStoreDB) UpdateStockAlert(arg0 context.Context, arg1 db.UpdateStockAlertParams) (db.StockAlert, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStockAlert", arg0, arg1)
	ret0, _ := ret[0].(db.StockAlert)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateStockAlert indicates an expected call of UpdateStockAlert.
func (mr *MockStoreDBMockRecorder) UpdateStockAlert(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStockAlert", reflect.TypeOf((*MockStoreDB)(nil).UpdateStockAlert), arg0, arg1)
}

// UpdateStore mocks base method.
func (m *MockStoreDB) UpdateStore(arg0 context.Context, arg1 db.UpdateStoreParams) (db.Store, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateStore", arg0, arg1)
	ret0, _ := ret[0].(db.Store)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateStore indicates an expected call of UpdateStore.
func (mr *MockStoreDBMockRecorder) UpdateStore(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateStore", reflect.TypeOf((*MockStoreDB)(nil).UpdateStore), arg0, arg1)
}

// UpdateSupplier mocks base method.
func (m *MockStoreDB) UpdateSupplier(arg0 context.Context, arg1 db.UpdateSupplierParams) (db.Supplier, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateSupplier", arg0, arg1)
	ret0, _ := ret[0].(db.Supplier)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateSupplier indicates an expected call of UpdateSupplier.
func (mr *MockStoreDBMockRecorder) UpdateSupplier(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateSupplier", reflect.TypeOf((*MockStoreDB)(nil).UpdateSupplier), arg0, arg1)
}
