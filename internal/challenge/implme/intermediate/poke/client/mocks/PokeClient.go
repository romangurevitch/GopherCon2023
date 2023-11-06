// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	structs "github.com/mtslzr/pokeapi-go/structs"
	mock "github.com/stretchr/testify/mock"
)

// PokeClient is an autogenerated mock type for the PokeClient type
type PokeClient struct {
	mock.Mock
}

type PokeClient_Expecter struct {
	mock *mock.Mock
}

func (_m *PokeClient) EXPECT() *PokeClient_Expecter {
	return &PokeClient_Expecter{mock: &_m.Mock}
}

// FetchPokemon provides a mock function with given fields: ID
func (_m *PokeClient) FetchPokemon(ID string) (*structs.Pokemon, error) {
	ret := _m.Called(ID)

	var r0 *structs.Pokemon
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*structs.Pokemon, error)); ok {
		return rf(ID)
	}
	if rf, ok := ret.Get(0).(func(string) *structs.Pokemon); ok {
		r0 = rf(ID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*structs.Pokemon)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PokeClient_FetchPokemon_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FetchPokemon'
type PokeClient_FetchPokemon_Call struct {
	*mock.Call
}

// FetchPokemon is a helper method to define mock.On call
//   - ID string
func (_e *PokeClient_Expecter) FetchPokemon(ID interface{}) *PokeClient_FetchPokemon_Call {
	return &PokeClient_FetchPokemon_Call{Call: _e.mock.On("FetchPokemon", ID)}
}

func (_c *PokeClient_FetchPokemon_Call) Run(run func(ID string)) *PokeClient_FetchPokemon_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *PokeClient_FetchPokemon_Call) Return(_a0 *structs.Pokemon, _a1 error) *PokeClient_FetchPokemon_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *PokeClient_FetchPokemon_Call) RunAndReturn(run func(string) (*structs.Pokemon, error)) *PokeClient_FetchPokemon_Call {
	_c.Call.Return(run)
	return _c
}

type mockConstructorTestingTNewPokeClient interface {
	mock.TestingT
	Cleanup(func())
}

// NewPokeClient creates a new instance of PokeClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPokeClient(t mockConstructorTestingTNewPokeClient) *PokeClient {
	mock := &PokeClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}