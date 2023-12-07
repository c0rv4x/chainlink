// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	big "math/big"

	bind "github.com/ethereum/go-ethereum/accounts/abi/bind"

	generated "github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated"

	i_keeper_registry_master_wrapper_2_1 "github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/i_keeper_registry_master_wrapper_2_1"

	mock "github.com/stretchr/testify/mock"

	types "github.com/ethereum/go-ethereum/core/types"
)

// Registry is an autogenerated mock type for the Registry type
type Registry struct {
	mock.Mock
}

// CheckCallback provides a mock function with given fields: opts, id, values, extraData
func (_m *Registry) CheckCallback(opts *bind.CallOpts, id *big.Int, values [][]byte, extraData []byte) (i_keeper_registry_master_wrapper_2_1.CheckCallback, error) {
	ret := _m.Called(opts, id, values, extraData)

	if len(ret) == 0 {
		panic("no return value specified for CheckCallback")
	}

	var r0 i_keeper_registry_master_wrapper_2_1.CheckCallback
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.CallOpts, *big.Int, [][]byte, []byte) (i_keeper_registry_master_wrapper_2_1.CheckCallback, error)); ok {
		return rf(opts, id, values, extraData)
	}
	if rf, ok := ret.Get(0).(func(*bind.CallOpts, *big.Int, [][]byte, []byte) i_keeper_registry_master_wrapper_2_1.CheckCallback); ok {
		r0 = rf(opts, id, values, extraData)
	} else {
		r0 = ret.Get(0).(i_keeper_registry_master_wrapper_2_1.CheckCallback)
	}

	if rf, ok := ret.Get(1).(func(*bind.CallOpts, *big.Int, [][]byte, []byte) error); ok {
		r1 = rf(opts, id, values, extraData)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetActiveUpkeepIDs provides a mock function with given fields: opts, startIndex, maxCount
func (_m *Registry) GetActiveUpkeepIDs(opts *bind.CallOpts, startIndex *big.Int, maxCount *big.Int) ([]*big.Int, error) {
	ret := _m.Called(opts, startIndex, maxCount)

	if len(ret) == 0 {
		panic("no return value specified for GetActiveUpkeepIDs")
	}

	var r0 []*big.Int
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.CallOpts, *big.Int, *big.Int) ([]*big.Int, error)); ok {
		return rf(opts, startIndex, maxCount)
	}
	if rf, ok := ret.Get(0).(func(*bind.CallOpts, *big.Int, *big.Int) []*big.Int); ok {
		r0 = rf(opts, startIndex, maxCount)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*big.Int)
		}
	}

	if rf, ok := ret.Get(1).(func(*bind.CallOpts, *big.Int, *big.Int) error); ok {
		r1 = rf(opts, startIndex, maxCount)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetState provides a mock function with given fields: opts
func (_m *Registry) GetState(opts *bind.CallOpts) (i_keeper_registry_master_wrapper_2_1.GetState, error) {
	ret := _m.Called(opts)

	if len(ret) == 0 {
		panic("no return value specified for GetState")
	}

	var r0 i_keeper_registry_master_wrapper_2_1.GetState
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) (i_keeper_registry_master_wrapper_2_1.GetState, error)); ok {
		return rf(opts)
	}
	if rf, ok := ret.Get(0).(func(*bind.CallOpts) i_keeper_registry_master_wrapper_2_1.GetState); ok {
		r0 = rf(opts)
	} else {
		r0 = ret.Get(0).(i_keeper_registry_master_wrapper_2_1.GetState)
	}

	if rf, ok := ret.Get(1).(func(*bind.CallOpts) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUpkeep provides a mock function with given fields: opts, id
func (_m *Registry) GetUpkeep(opts *bind.CallOpts, id *big.Int) (i_keeper_registry_master_wrapper_2_1.KeeperRegistryBase21UpkeepInfo, error) {
	ret := _m.Called(opts, id)

	if len(ret) == 0 {
		panic("no return value specified for GetUpkeep")
	}

	var r0 i_keeper_registry_master_wrapper_2_1.KeeperRegistryBase21UpkeepInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.CallOpts, *big.Int) (i_keeper_registry_master_wrapper_2_1.KeeperRegistryBase21UpkeepInfo, error)); ok {
		return rf(opts, id)
	}
	if rf, ok := ret.Get(0).(func(*bind.CallOpts, *big.Int) i_keeper_registry_master_wrapper_2_1.KeeperRegistryBase21UpkeepInfo); ok {
		r0 = rf(opts, id)
	} else {
		r0 = ret.Get(0).(i_keeper_registry_master_wrapper_2_1.KeeperRegistryBase21UpkeepInfo)
	}

	if rf, ok := ret.Get(1).(func(*bind.CallOpts, *big.Int) error); ok {
		r1 = rf(opts, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUpkeepPrivilegeConfig provides a mock function with given fields: opts, upkeepId
func (_m *Registry) GetUpkeepPrivilegeConfig(opts *bind.CallOpts, upkeepId *big.Int) ([]byte, error) {
	ret := _m.Called(opts, upkeepId)

	if len(ret) == 0 {
		panic("no return value specified for GetUpkeepPrivilegeConfig")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.CallOpts, *big.Int) ([]byte, error)); ok {
		return rf(opts, upkeepId)
	}
	if rf, ok := ret.Get(0).(func(*bind.CallOpts, *big.Int) []byte); ok {
		r0 = rf(opts, upkeepId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(*bind.CallOpts, *big.Int) error); ok {
		r1 = rf(opts, upkeepId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUpkeepTriggerConfig provides a mock function with given fields: opts, upkeepId
func (_m *Registry) GetUpkeepTriggerConfig(opts *bind.CallOpts, upkeepId *big.Int) ([]byte, error) {
	ret := _m.Called(opts, upkeepId)

	if len(ret) == 0 {
		panic("no return value specified for GetUpkeepTriggerConfig")
	}

	var r0 []byte
	var r1 error
	if rf, ok := ret.Get(0).(func(*bind.CallOpts, *big.Int) ([]byte, error)); ok {
		return rf(opts, upkeepId)
	}
	if rf, ok := ret.Get(0).(func(*bind.CallOpts, *big.Int) []byte); ok {
		r0 = rf(opts, upkeepId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	if rf, ok := ret.Get(1).(func(*bind.CallOpts, *big.Int) error); ok {
		r1 = rf(opts, upkeepId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ParseLog provides a mock function with given fields: log
func (_m *Registry) ParseLog(log types.Log) (generated.AbigenLog, error) {
	ret := _m.Called(log)

	if len(ret) == 0 {
		panic("no return value specified for ParseLog")
	}

	var r0 generated.AbigenLog
	var r1 error
	if rf, ok := ret.Get(0).(func(types.Log) (generated.AbigenLog, error)); ok {
		return rf(log)
	}
	if rf, ok := ret.Get(0).(func(types.Log) generated.AbigenLog); ok {
		r0 = rf(log)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(generated.AbigenLog)
		}
	}

	if rf, ok := ret.Get(1).(func(types.Log) error); ok {
		r1 = rf(log)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewRegistry creates a new instance of Registry. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewRegistry(t interface {
	mock.TestingT
	Cleanup(func())
}) *Registry {
	mock := &Registry{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
