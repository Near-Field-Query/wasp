// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package timestamp

import "github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib/wasmtypes"

const (
	ScName        = "timestamp"
	ScDescription = "Extremely simple timestamp server"
	HScName       = wasmtypes.ScHname(0x3988002e)
)

const (
	ResultTimestamp = "timestamp"
)

const (
	StateTimestamp = "timestamp"
)

const (
	FuncNow          = "now"
	ViewGetTimestamp = "getTimestamp"
)

const (
	HFuncNow          = wasmtypes.ScHname(0xd73b7fc9)
	HViewGetTimestamp = wasmtypes.ScHname(0x40c6376a)
)
