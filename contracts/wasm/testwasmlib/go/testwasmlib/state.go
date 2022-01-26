// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package testwasmlib

import "github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib/wasmtypes"

type MapStringToImmutableStringArray struct {
	proxy wasmtypes.Proxy
}

func (m MapStringToImmutableStringArray) GetStringArray(key string) ImmutableStringArray {
	return ImmutableStringArray{proxy: m.proxy.Key(wasmtypes.BytesFromString(key))}
}

type MapStringToImmutableStringMap struct {
	proxy wasmtypes.Proxy
}

func (m MapStringToImmutableStringMap) GetStringMap(key string) ImmutableStringMap {
	return ImmutableStringMap{proxy: m.proxy.Key(wasmtypes.BytesFromString(key))}
}

type ImmutableTestWasmLibState struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableTestWasmLibState) Arrays() MapStringToImmutableStringArray {
	return MapStringToImmutableStringArray{proxy: s.proxy.Root(StateArrays)}
}

func (s ImmutableTestWasmLibState) Maps() MapStringToImmutableStringMap {
	return MapStringToImmutableStringMap{proxy: s.proxy.Root(StateMaps)}
}

func (s ImmutableTestWasmLibState) Random() wasmtypes.ScImmutableUint64 {
	return wasmtypes.NewScImmutableUint64(s.proxy.Root(StateRandom))
}

type MapStringToMutableStringArray struct {
	proxy wasmtypes.Proxy
}

func (m MapStringToMutableStringArray) Clear() {
	m.proxy.ClearMap()
}

func (m MapStringToMutableStringArray) GetStringArray(key string) MutableStringArray {
	return MutableStringArray{proxy: m.proxy.Key(wasmtypes.BytesFromString(key))}
}

type MapStringToMutableStringMap struct {
	proxy wasmtypes.Proxy
}

func (m MapStringToMutableStringMap) Clear() {
	m.proxy.ClearMap()
}

func (m MapStringToMutableStringMap) GetStringMap(key string) MutableStringMap {
	return MutableStringMap{proxy: m.proxy.Key(wasmtypes.BytesFromString(key))}
}

type MutableTestWasmLibState struct {
	proxy wasmtypes.Proxy
}

func (s MutableTestWasmLibState) AsImmutable() ImmutableTestWasmLibState {
	return ImmutableTestWasmLibState(s)
}

func (s MutableTestWasmLibState) Arrays() MapStringToMutableStringArray {
	return MapStringToMutableStringArray{proxy: s.proxy.Root(StateArrays)}
}

func (s MutableTestWasmLibState) Maps() MapStringToMutableStringMap {
	return MapStringToMutableStringMap{proxy: s.proxy.Root(StateMaps)}
}

func (s MutableTestWasmLibState) Random() wasmtypes.ScMutableUint64 {
	return wasmtypes.NewScMutableUint64(s.proxy.Root(StateRandom))
}
