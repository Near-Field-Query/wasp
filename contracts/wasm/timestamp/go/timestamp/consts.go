// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package timestamp

import "github.com/iotaledger/wasp/packages/vm/wasmlib/go/wasmlib"

const (
	ScName        = "timestamp"
	ScDescription = "Extremely simple timestamp server"
	HScName       = wasmlib.ScHname(0x3988002e)
)

const (
	ResultTimestamp = wasmlib.Key("timestamp")
)

const (
	StateTimestamp = wasmlib.Key("timestamp")
)

const (
	FuncNow          = "now"
	ViewGetTimestamp = "getTimestamp"
)

const (
	HFuncNow          = wasmlib.ScHname(0xd73b7fc9)
	HViewGetTimestamp = wasmlib.ScHname(0x40c6376a)
)
