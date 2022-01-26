// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package tokenregistry

import "github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib/wasmtypes"

type ArrayOfImmutableColor struct {
	proxy wasmtypes.Proxy
}

func (a ArrayOfImmutableColor) Length() uint32 {
	return a.proxy.Length()
}

func (a ArrayOfImmutableColor) GetColor(index uint32) wasmtypes.ScImmutableColor {
	return wasmtypes.NewScImmutableColor(a.proxy.Index(index))
}

type MapColorToImmutableToken struct {
	proxy wasmtypes.Proxy
}

func (m MapColorToImmutableToken) GetToken(key wasmtypes.ScColor) ImmutableToken {
	return ImmutableToken{proxy: m.proxy.Key(wasmtypes.BytesFromColor(key))}
}

type ImmutableTokenRegistryState struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableTokenRegistryState) ColorList() ArrayOfImmutableColor {
	return ArrayOfImmutableColor{proxy: s.proxy.Root(StateColorList)}
}

func (s ImmutableTokenRegistryState) Registry() MapColorToImmutableToken {
	return MapColorToImmutableToken{proxy: s.proxy.Root(StateRegistry)}
}

type ArrayOfMutableColor struct {
	proxy wasmtypes.Proxy
}

func (a ArrayOfMutableColor) AppendColor() wasmtypes.ScMutableColor {
	return wasmtypes.NewScMutableColor(a.proxy.Append())
}

func (a ArrayOfMutableColor) Clear() {
	a.proxy.ClearArray()
}

func (a ArrayOfMutableColor) Length() uint32 {
	return a.proxy.Length()
}

func (a ArrayOfMutableColor) GetColor(index uint32) wasmtypes.ScMutableColor {
	return wasmtypes.NewScMutableColor(a.proxy.Index(index))
}

type MapColorToMutableToken struct {
	proxy wasmtypes.Proxy
}

func (m MapColorToMutableToken) Clear() {
	m.proxy.ClearMap()
}

func (m MapColorToMutableToken) GetToken(key wasmtypes.ScColor) MutableToken {
	return MutableToken{proxy: m.proxy.Key(wasmtypes.BytesFromColor(key))}
}

type MutableTokenRegistryState struct {
	proxy wasmtypes.Proxy
}

func (s MutableTokenRegistryState) AsImmutable() ImmutableTokenRegistryState {
	return ImmutableTokenRegistryState(s)
}

func (s MutableTokenRegistryState) ColorList() ArrayOfMutableColor {
	return ArrayOfMutableColor{proxy: s.proxy.Root(StateColorList)}
}

func (s MutableTokenRegistryState) Registry() MapColorToMutableToken {
	return MapColorToMutableToken{proxy: s.proxy.Root(StateRegistry)}
}
