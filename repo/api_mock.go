// Code generated by MockGen. DO NOT EDIT.
// Source: api.go
//
// Generated by this command:
//
//	mockgen -source api.go -destination api_mock.go -package repo
//

// Package repo is a generated GoMock package.
package repo

import (
	reflect "reflect"

	postsql "github.com/atomicals-go/repo/postsql"
	v3 "github.com/bits-and-blooms/bloom/v3"
	gomock "go.uber.org/mock/gomock"
)

// MockDB is a mock of DB interface.
type MockDB struct {
	ctrl     *gomock.Controller
	recorder *MockDBMockRecorder
}

// MockDBMockRecorder is the mock recorder for MockDB.
type MockDBMockRecorder struct {
	mock *MockDB
}

// NewMockDB creates a new mock instance.
func NewMockDB(ctrl *gomock.Controller) *MockDB {
	mock := &MockDB{ctrl: ctrl}
	mock.recorder = &MockDBMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDB) EXPECT() *MockDBMockRecorder {
	return m.recorder
}

// BloomFilter mocks base method.
func (m *MockDB) BloomFilter() (map[string]*bloomFilterInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BloomFilter")
	ret0, _ := ret[0].(map[string]*bloomFilterInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BloomFilter indicates an expected call of BloomFilter.
func (mr *MockDBMockRecorder) BloomFilter() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BloomFilter", reflect.TypeOf((*MockDB)(nil).BloomFilter))
}

// BtcTx mocks base method.
func (m *MockDB) BtcTx(txID string) (*postsql.BtcTx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BtcTx", txID)
	ret0, _ := ret[0].(*postsql.BtcTx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BtcTx indicates an expected call of BtcTx.
func (mr *MockDBMockRecorder) BtcTx(txID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BtcTx", reflect.TypeOf((*MockDB)(nil).BtcTx), txID)
}

// BtcTxHeight mocks base method.
func (m *MockDB) BtcTxHeight(txID string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BtcTxHeight", txID)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BtcTxHeight indicates an expected call of BtcTxHeight.
func (mr *MockDBMockRecorder) BtcTxHeight(txID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BtcTxHeight", reflect.TypeOf((*MockDB)(nil).BtcTxHeight), txID)
}

// ContainerItemByNameHasExist mocks base method.
func (m *MockDB) ContainerItemByNameHasExist(container, item string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContainerItemByNameHasExist", container, item)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ContainerItemByNameHasExist indicates an expected call of ContainerItemByNameHasExist.
func (mr *MockDBMockRecorder) ContainerItemByNameHasExist(container, item any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainerItemByNameHasExist", reflect.TypeOf((*MockDB)(nil).ContainerItemByNameHasExist), container, item)
}

// DeleteBtcTxUntil mocks base method.
func (m *MockDB) DeleteBtcTxUntil(blockHeight int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBtcTxUntil", blockHeight)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteBtcTxUntil indicates an expected call of DeleteBtcTxUntil.
func (mr *MockDBMockRecorder) DeleteBtcTxUntil(blockHeight any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBtcTxUntil", reflect.TypeOf((*MockDB)(nil).DeleteBtcTxUntil), blockHeight)
}

// DeleteFtUTXO mocks base method.
func (m *MockDB) DeleteFtUTXO(locationID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFtUTXO", locationID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFtUTXO indicates an expected call of DeleteFtUTXO.
func (mr *MockDBMockRecorder) DeleteFtUTXO(locationID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFtUTXO", reflect.TypeOf((*MockDB)(nil).DeleteFtUTXO), locationID)
}

// DistributedFtByName mocks base method.
func (m *MockDB) DistributedFtByName(tickerName string) (*postsql.GlobalDistributedFt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DistributedFtByName", tickerName)
	ret0, _ := ret[0].(*postsql.GlobalDistributedFt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DistributedFtByName indicates an expected call of DistributedFtByName.
func (mr *MockDBMockRecorder) DistributedFtByName(tickerName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DistributedFtByName", reflect.TypeOf((*MockDB)(nil).DistributedFtByName), tickerName)
}

// ExecAllSql mocks base method.
func (m *MockDB) ExecAllSql(blockHeight, txIndex int64, txID, operation string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExecAllSql", blockHeight, txIndex, txID, operation)
	ret0, _ := ret[0].(error)
	return ret0
}

// ExecAllSql indicates an expected call of ExecAllSql.
func (mr *MockDBMockRecorder) ExecAllSql(blockHeight, txIndex, txID, operation any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExecAllSql", reflect.TypeOf((*MockDB)(nil).ExecAllSql), blockHeight, txIndex, txID, operation)
}

// FtUTXOsByLocationID mocks base method.
func (m *MockDB) FtUTXOsByLocationID(locationID string) ([]*postsql.UTXOFtInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FtUTXOsByLocationID", locationID)
	ret0, _ := ret[0].([]*postsql.UTXOFtInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FtUTXOsByLocationID indicates an expected call of FtUTXOsByLocationID.
func (mr *MockDBMockRecorder) FtUTXOsByLocationID(locationID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FtUTXOsByLocationID", reflect.TypeOf((*MockDB)(nil).FtUTXOsByLocationID), locationID)
}

// FtUTXOsByUserPK mocks base method.
func (m *MockDB) FtUTXOsByUserPK(UserPK string) ([]*postsql.UTXOFtInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FtUTXOsByUserPK", UserPK)
	ret0, _ := ret[0].([]*postsql.UTXOFtInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FtUTXOsByUserPK indicates an expected call of FtUTXOsByUserPK.
func (mr *MockDBMockRecorder) FtUTXOsByUserPK(UserPK any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FtUTXOsByUserPK", reflect.TypeOf((*MockDB)(nil).FtUTXOsByUserPK), UserPK)
}

// InsertBloomFilter mocks base method.
func (m *MockDB) InsertBloomFilter(name string, filter *v3.BloomFilter) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertBloomFilter", name, filter)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertBloomFilter indicates an expected call of InsertBloomFilter.
func (mr *MockDBMockRecorder) InsertBloomFilter(name, filter any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertBloomFilter", reflect.TypeOf((*MockDB)(nil).InsertBloomFilter), name, filter)
}

// InsertBtcTx mocks base method.
func (m *MockDB) InsertBtcTx(btcTx *postsql.BtcTx) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertBtcTx", btcTx)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertBtcTx indicates an expected call of InsertBtcTx.
func (mr *MockDBMockRecorder) InsertBtcTx(btcTx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertBtcTx", reflect.TypeOf((*MockDB)(nil).InsertBtcTx), btcTx)
}

// InsertDirectFtUTXO mocks base method.
func (m *MockDB) InsertDirectFtUTXO(entity *postsql.GlobalDirectFt) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertDirectFtUTXO", entity)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertDirectFtUTXO indicates an expected call of InsertDirectFtUTXO.
func (mr *MockDBMockRecorder) InsertDirectFtUTXO(entity any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertDirectFtUTXO", reflect.TypeOf((*MockDB)(nil).InsertDirectFtUTXO), entity)
}

// InsertDistributedFt mocks base method.
func (m *MockDB) InsertDistributedFt(ft *postsql.GlobalDistributedFt) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertDistributedFt", ft)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertDistributedFt indicates an expected call of InsertDistributedFt.
func (mr *MockDBMockRecorder) InsertDistributedFt(ft any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertDistributedFt", reflect.TypeOf((*MockDB)(nil).InsertDistributedFt), ft)
}

// InsertFtUTXO mocks base method.
func (m *MockDB) InsertFtUTXO(UTXO *postsql.UTXOFtInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertFtUTXO", UTXO)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertFtUTXO indicates an expected call of InsertFtUTXO.
func (mr *MockDBMockRecorder) InsertFtUTXO(UTXO any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertFtUTXO", reflect.TypeOf((*MockDB)(nil).InsertFtUTXO), UTXO)
}

// InsertMod mocks base method.
func (m *MockDB) InsertMod(mod *postsql.ModInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertMod", mod)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertMod indicates an expected call of InsertMod.
func (mr *MockDBMockRecorder) InsertMod(mod any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertMod", reflect.TypeOf((*MockDB)(nil).InsertMod), mod)
}

// InsertNftUTXO mocks base method.
func (m *MockDB) InsertNftUTXO(UTXO *postsql.UTXONftInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertNftUTXO", UTXO)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertNftUTXO indicates an expected call of InsertNftUTXO.
func (mr *MockDBMockRecorder) InsertNftUTXO(UTXO any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertNftUTXO", reflect.TypeOf((*MockDB)(nil).InsertNftUTXO), UTXO)
}

// Location mocks base method.
func (m *MockDB) Location() (*postsql.Location, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Location")
	ret0, _ := ret[0].(*postsql.Location)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Location indicates an expected call of Location.
func (mr *MockDBMockRecorder) Location() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Location", reflect.TypeOf((*MockDB)(nil).Location))
}

// Mod mocks base method.
func (m *MockDB) Mod(atomicalsID string) (*postsql.ModInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Mod", atomicalsID)
	ret0, _ := ret[0].(*postsql.ModInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Mod indicates an expected call of Mod.
func (mr *MockDBMockRecorder) Mod(atomicalsID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Mod", reflect.TypeOf((*MockDB)(nil).Mod), atomicalsID)
}

// NftContainerByNameHasExist mocks base method.
func (m *MockDB) NftContainerByNameHasExist(containerName string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NftContainerByNameHasExist", containerName)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NftContainerByNameHasExist indicates an expected call of NftContainerByNameHasExist.
func (mr *MockDBMockRecorder) NftContainerByNameHasExist(containerName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NftContainerByNameHasExist", reflect.TypeOf((*MockDB)(nil).NftContainerByNameHasExist), containerName)
}

// NftRealmByNameHasExist mocks base method.
func (m *MockDB) NftRealmByNameHasExist(realmName string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NftRealmByNameHasExist", realmName)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NftRealmByNameHasExist indicates an expected call of NftRealmByNameHasExist.
func (mr *MockDBMockRecorder) NftRealmByNameHasExist(realmName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NftRealmByNameHasExist", reflect.TypeOf((*MockDB)(nil).NftRealmByNameHasExist), realmName)
}

// NftSubRealmByNameHasExist mocks base method.
func (m *MockDB) NftSubRealmByNameHasExist(parentRealmAtomicalsID, subRealm string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NftSubRealmByNameHasExist", parentRealmAtomicalsID, subRealm)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NftSubRealmByNameHasExist indicates an expected call of NftSubRealmByNameHasExist.
func (mr *MockDBMockRecorder) NftSubRealmByNameHasExist(parentRealmAtomicalsID, subRealm any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NftSubRealmByNameHasExist", reflect.TypeOf((*MockDB)(nil).NftSubRealmByNameHasExist), parentRealmAtomicalsID, subRealm)
}

// NftUTXOsByLocationID mocks base method.
func (m *MockDB) NftUTXOsByLocationID(locationID string) ([]*postsql.UTXONftInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NftUTXOsByLocationID", locationID)
	ret0, _ := ret[0].([]*postsql.UTXONftInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NftUTXOsByLocationID indicates an expected call of NftUTXOsByLocationID.
func (mr *MockDBMockRecorder) NftUTXOsByLocationID(locationID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NftUTXOsByLocationID", reflect.TypeOf((*MockDB)(nil).NftUTXOsByLocationID), locationID)
}

// NftUTXOsByUserPK mocks base method.
func (m *MockDB) NftUTXOsByUserPK(UserPK string) ([]*postsql.UTXONftInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NftUTXOsByUserPK", UserPK)
	ret0, _ := ret[0].([]*postsql.UTXONftInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NftUTXOsByUserPK indicates an expected call of NftUTXOsByUserPK.
func (mr *MockDBMockRecorder) NftUTXOsByUserPK(UserPK any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NftUTXOsByUserPK", reflect.TypeOf((*MockDB)(nil).NftUTXOsByUserPK), UserPK)
}

// ParentContainerHasExist mocks base method.
func (m *MockDB) ParentContainerHasExist(parentContainerAtomicalsID string) (*postsql.UTXONftInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParentContainerHasExist", parentContainerAtomicalsID)
	ret0, _ := ret[0].(*postsql.UTXONftInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParentContainerHasExist indicates an expected call of ParentContainerHasExist.
func (mr *MockDBMockRecorder) ParentContainerHasExist(parentContainerAtomicalsID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParentContainerHasExist", reflect.TypeOf((*MockDB)(nil).ParentContainerHasExist), parentContainerAtomicalsID)
}

// ParentRealmHasExist mocks base method.
func (m *MockDB) ParentRealmHasExist(parentRealmAtomicalsID string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParentRealmHasExist", parentRealmAtomicalsID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ParentRealmHasExist indicates an expected call of ParentRealmHasExist.
func (mr *MockDBMockRecorder) ParentRealmHasExist(parentRealmAtomicalsID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParentRealmHasExist", reflect.TypeOf((*MockDB)(nil).ParentRealmHasExist), parentRealmAtomicalsID)
}

// UpdateBloomFilter mocks base method.
func (m *MockDB) UpdateBloomFilter(name string, filter *v3.BloomFilter) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBloomFilter", name, filter)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateBloomFilter indicates an expected call of UpdateBloomFilter.
func (mr *MockDBMockRecorder) UpdateBloomFilter(name, filter any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBloomFilter", reflect.TypeOf((*MockDB)(nil).UpdateBloomFilter), name, filter)
}

// UpdateDistributedFt mocks base method.
func (m *MockDB) UpdateDistributedFt(ft *postsql.GlobalDistributedFt) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateDistributedFt", ft)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateDistributedFt indicates an expected call of UpdateDistributedFt.
func (mr *MockDBMockRecorder) UpdateDistributedFt(ft any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDistributedFt", reflect.TypeOf((*MockDB)(nil).UpdateDistributedFt), ft)
}

// UpdateNftUTXO mocks base method.
func (m *MockDB) UpdateNftUTXO(UTXO *postsql.UTXONftInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateNftUTXO", UTXO)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateNftUTXO indicates an expected call of UpdateNftUTXO.
func (mr *MockDBMockRecorder) UpdateNftUTXO(UTXO any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateNftUTXO", reflect.TypeOf((*MockDB)(nil).UpdateNftUTXO), UTXO)
}
