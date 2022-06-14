// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

//nolint:revive
package dividend

import "github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib/wasmtypes"

type ImmutableInitParams struct {
	proxy wasmtypes.Proxy
}

// optional owner of contract, defaults to contract creator
func (s ImmutableInitParams) Owner() wasmtypes.ScImmutableAgentID {
	return wasmtypes.NewScImmutableAgentID(s.proxy.Root(ParamOwner))
}

type MutableInitParams struct {
	proxy wasmtypes.Proxy
}

// optional owner of contract, defaults to contract creator
func (s MutableInitParams) Owner() wasmtypes.ScMutableAgentID {
	return wasmtypes.NewScMutableAgentID(s.proxy.Root(ParamOwner))
}

type ImmutableMemberParams struct {
	proxy wasmtypes.Proxy
}

// address of dividend recipient
func (s ImmutableMemberParams) Address() wasmtypes.ScImmutableAddress {
	return wasmtypes.NewScImmutableAddress(s.proxy.Root(ParamAddress))
}

// relative division factor
func (s ImmutableMemberParams) Factor() wasmtypes.ScImmutableUint64 {
	return wasmtypes.NewScImmutableUint64(s.proxy.Root(ParamFactor))
}

type MutableMemberParams struct {
	proxy wasmtypes.Proxy
}

// address of dividend recipient
func (s MutableMemberParams) Address() wasmtypes.ScMutableAddress {
	return wasmtypes.NewScMutableAddress(s.proxy.Root(ParamAddress))
}

// relative division factor
func (s MutableMemberParams) Factor() wasmtypes.ScMutableUint64 {
	return wasmtypes.NewScMutableUint64(s.proxy.Root(ParamFactor))
}

type ImmutableSetOwnerParams struct {
	proxy wasmtypes.Proxy
}

// new owner of smart contract
func (s ImmutableSetOwnerParams) Owner() wasmtypes.ScImmutableAgentID {
	return wasmtypes.NewScImmutableAgentID(s.proxy.Root(ParamOwner))
}

type MutableSetOwnerParams struct {
	proxy wasmtypes.Proxy
}

// new owner of smart contract
func (s MutableSetOwnerParams) Owner() wasmtypes.ScMutableAgentID {
	return wasmtypes.NewScMutableAgentID(s.proxy.Root(ParamOwner))
}

type ImmutableGetFactorParams struct {
	proxy wasmtypes.Proxy
}

// address of dividend recipient
func (s ImmutableGetFactorParams) Address() wasmtypes.ScImmutableAddress {
	return wasmtypes.NewScImmutableAddress(s.proxy.Root(ParamAddress))
}

type MutableGetFactorParams struct {
	proxy wasmtypes.Proxy
}

// address of dividend recipient
func (s MutableGetFactorParams) Address() wasmtypes.ScMutableAddress {
	return wasmtypes.NewScMutableAddress(s.proxy.Root(ParamAddress))
}
