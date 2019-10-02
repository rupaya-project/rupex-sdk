// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import bson "github.com/globalsign/mgo/bson"
import common "github.com/ethereum/go-ethereum/common"

import mock "github.com/stretchr/testify/mock"
import types "github.com/tomochain/tomox-sdk/types"

// WalletDao is an autogenerated mock type for the WalletDao type
type WalletDao struct {
    mock.Mock
}

// Create provides a mock function with given fields: wallet
func (_m *WalletDao) Create(wallet *types.Wallet) error {
    ret := _m.Called(wallet)

    var r0 error
    if rf, ok := ret.Get(0).(func(*types.Wallet) error); ok {
        r0 = rf(wallet)
    } else {
        r0 = ret.Error(0)
    }

    return r0
}

// GetAll provides a mock function with given fields:
func (_m *WalletDao) GetAll() ([]types.Wallet, error) {
    ret := _m.Called()

    var r0 []types.Wallet
    if rf, ok := ret.Get(0).(func() []types.Wallet); ok {
        r0 = rf()
    } else {
        if ret.Get(0) != nil {
            r0 = ret.Get(0).([]types.Wallet)
        }
    }

    var r1 error
    if rf, ok := ret.Get(1).(func() error); ok {
        r1 = rf()
    } else {
        r1 = ret.Error(1)
    }

    return r0, r1
}

// GetByAddress provides a mock function with given fields: addr
func (_m *WalletDao) GetByAddress(addr common.Address) (*types.Wallet, error) {
    ret := _m.Called(addr)

    var r0 *types.Wallet
    if rf, ok := ret.Get(0).(func(common.Address) *types.Wallet); ok {
        r0 = rf(addr)
    } else {
        if ret.Get(0) != nil {
            r0 = ret.Get(0).(*types.Wallet)
        }
    }

    var r1 error
    if rf, ok := ret.Get(1).(func(common.Address) error); ok {
        r1 = rf(addr)
    } else {
        r1 = ret.Error(1)
    }

    return r0, r1
}

// GetByID provides a mock function with given fields: id
func (_m *WalletDao) GetByID(id bson.ObjectId) (*types.Wallet, error) {
    ret := _m.Called(id)

    var r0 *types.Wallet
    if rf, ok := ret.Get(0).(func(bson.ObjectId) *types.Wallet); ok {
        r0 = rf(id)
    } else {
        if ret.Get(0) != nil {
            r0 = ret.Get(0).(*types.Wallet)
        }
    }

    var r1 error
    if rf, ok := ret.Get(1).(func(bson.ObjectId) error); ok {
        r1 = rf(id)
    } else {
        r1 = ret.Error(1)
    }

    return r0, r1
}

// GetDefaultAdminWallet provides a mock function with given fields:
func (_m *WalletDao) GetDefaultAdminWallet() (*types.Wallet, error) {
    ret := _m.Called()

    var r0 *types.Wallet
    if rf, ok := ret.Get(0).(func() *types.Wallet); ok {
        r0 = rf()
    } else {
        if ret.Get(0) != nil {
            r0 = ret.Get(0).(*types.Wallet)
        }
    }

    var r1 error
    if rf, ok := ret.Get(1).(func() error); ok {
        r1 = rf()
    } else {
        r1 = ret.Error(1)
    }

    return r0, r1
}

// GetOperatorWallets provides a mock function with given fields:
func (_m *WalletDao) GetOperatorWallets() ([]*types.Wallet, error) {
    ret := _m.Called()

    var r0 []*types.Wallet
    if rf, ok := ret.Get(0).(func() []*types.Wallet); ok {
        r0 = rf()
    } else {
        if ret.Get(0) != nil {
            r0 = ret.Get(0).([]*types.Wallet)
        }
    }

    var r1 error
    if rf, ok := ret.Get(1).(func() error); ok {
        r1 = rf()
    } else {
        r1 = ret.Error(1)
    }

    return r0, r1
}
