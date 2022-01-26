// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package testcore

import "github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib/wasmtypes"

const (
	ScName        = "testcore"
	ScDescription = "Core test for ISCP wasmlib Rust/Wasm library"
	HScName       = wasmtypes.ScHname(0x370d33ad)
)

const (
	ParamAddress         = "address"
	ParamAgentID         = "agentID"
	ParamCaller          = "caller"
	ParamChainID         = "chainID"
	ParamChainOwnerID    = "chainOwnerID"
	ParamContractCreator = "contractCreator"
	ParamContractID      = "contractID"
	ParamCounter         = "counter"
	ParamFail            = "initFailParam"
	ParamHash            = "Hash"
	ParamHname           = "Hname"
	ParamHnameContract   = "hnameContract"
	ParamHnameEP         = "hnameEP"
	ParamHnameZero       = "Hname-0"
	ParamInt64           = "int64"
	ParamInt64Zero       = "int64-0"
	ParamIntValue        = "intParamValue"
	ParamName            = "intParamName"
	ParamProgHash        = "progHash"
	ParamString          = "string"
	ParamStringZero      = "string-0"
	ParamVarName         = "varName"
)

const (
	ResultChainOwnerID = "chainOwnerID"
	ResultCounter      = "counter"
	ResultIntValue     = "intParamValue"
	ResultMintedColor  = "mintedColor"
	ResultMintedSupply = "mintedSupply"
	ResultSandboxCall  = "sandboxCall"
	ResultValues       = "this"
	ResultVars         = "this"
)

const (
	StateCounter      = "counter"
	StateHnameEP      = "hnameEP"
	StateInts         = "ints"
	StateMintedColor  = "mintedColor"
	StateMintedSupply = "mintedSupply"
)

const (
	FuncCallOnChain                 = "callOnChain"
	FuncCheckContextFromFullEP      = "checkContextFromFullEP"
	FuncDoNothing                   = "doNothing"
	FuncGetMintedSupply             = "getMintedSupply"
	FuncIncCounter                  = "incCounter"
	FuncInit                        = "init"
	FuncPassTypesFull               = "passTypesFull"
	FuncRunRecursion                = "runRecursion"
	FuncSendToAddress               = "sendToAddress"
	FuncSetInt                      = "setInt"
	FuncSpawn                       = "spawn"
	FuncTestBlockContext1           = "testBlockContext1"
	FuncTestBlockContext2           = "testBlockContext2"
	FuncTestCallPanicFullEP         = "testCallPanicFullEP"
	FuncTestCallPanicViewEPFromFull = "testCallPanicViewEPFromFull"
	FuncTestChainOwnerIDFull        = "testChainOwnerIDFull"
	FuncTestEventLogDeploy          = "testEventLogDeploy"
	FuncTestEventLogEventData       = "testEventLogEventData"
	FuncTestEventLogGenericData     = "testEventLogGenericData"
	FuncTestPanicFullEP             = "testPanicFullEP"
	FuncWithdrawToChain             = "withdrawToChain"
	ViewCheckContextFromViewEP      = "checkContextFromViewEP"
	ViewFibonacci                   = "fibonacci"
	ViewGetCounter                  = "getCounter"
	ViewGetInt                      = "getInt"
	ViewGetStringValue              = "getStringValue"
	ViewJustView                    = "justView"
	ViewPassTypesView               = "passTypesView"
	ViewTestCallPanicViewEPFromView = "testCallPanicViewEPFromView"
	ViewTestChainOwnerIDView        = "testChainOwnerIDView"
	ViewTestPanicViewEP             = "testPanicViewEP"
	ViewTestSandboxCall             = "testSandboxCall"
)

const (
	HFuncCallOnChain                 = wasmtypes.ScHname(0x95a3d123)
	HFuncCheckContextFromFullEP      = wasmtypes.ScHname(0xa56c24ba)
	HFuncDoNothing                   = wasmtypes.ScHname(0xdda4a6de)
	HFuncGetMintedSupply             = wasmtypes.ScHname(0x0c2d113c)
	HFuncIncCounter                  = wasmtypes.ScHname(0x7b287419)
	HFuncInit                        = wasmtypes.ScHname(0x1f44d644)
	HFuncPassTypesFull               = wasmtypes.ScHname(0x733ea0ea)
	HFuncRunRecursion                = wasmtypes.ScHname(0x833425fd)
	HFuncSendToAddress               = wasmtypes.ScHname(0x63ce4634)
	HFuncSetInt                      = wasmtypes.ScHname(0x62056f74)
	HFuncSpawn                       = wasmtypes.ScHname(0xec929d12)
	HFuncTestBlockContext1           = wasmtypes.ScHname(0x796d4136)
	HFuncTestBlockContext2           = wasmtypes.ScHname(0x758b0452)
	HFuncTestCallPanicFullEP         = wasmtypes.ScHname(0x4c878834)
	HFuncTestCallPanicViewEPFromFull = wasmtypes.ScHname(0xfd7e8c1d)
	HFuncTestChainOwnerIDFull        = wasmtypes.ScHname(0x2aff1167)
	HFuncTestEventLogDeploy          = wasmtypes.ScHname(0x96ff760a)
	HFuncTestEventLogEventData       = wasmtypes.ScHname(0x0efcf939)
	HFuncTestEventLogGenericData     = wasmtypes.ScHname(0x6a16629d)
	HFuncTestPanicFullEP             = wasmtypes.ScHname(0x24fdef07)
	HFuncWithdrawToChain             = wasmtypes.ScHname(0x437bc026)
	HViewCheckContextFromViewEP      = wasmtypes.ScHname(0x88ff0167)
	HViewFibonacci                   = wasmtypes.ScHname(0x7940873c)
	HViewGetCounter                  = wasmtypes.ScHname(0xb423e607)
	HViewGetInt                      = wasmtypes.ScHname(0x1887e5ef)
	HViewGetStringValue              = wasmtypes.ScHname(0xcf0a4d32)
	HViewJustView                    = wasmtypes.ScHname(0x33b8972e)
	HViewPassTypesView               = wasmtypes.ScHname(0x1a5b87ea)
	HViewTestCallPanicViewEPFromView = wasmtypes.ScHname(0x91b10c99)
	HViewTestChainOwnerIDView        = wasmtypes.ScHname(0x26586c33)
	HViewTestPanicViewEP             = wasmtypes.ScHname(0x22bc4d72)
	HViewTestSandboxCall             = wasmtypes.ScHname(0x42d72b63)
)
