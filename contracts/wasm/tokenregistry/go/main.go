// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

// +build wasm

package main

import "github.com/iotaledger/wasp/packages/wasmvm/wasmvmhost"

import "github.com/iotaledger/wasp/contracts/wasm/tokenregistry/go/tokenregistry"

func main() {
}

//export on_load
func onLoad() {
	h := &wasmvmhost.WasmVMHost{}
	h.ConnectWasmHost()
	tokenregistry.OnLoad()
}
