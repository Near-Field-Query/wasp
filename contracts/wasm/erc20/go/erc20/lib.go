// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package erc20

import "github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib"

func OnLoad() {
	exports := wasmlib.NewScExports()
	exports.AddFunc(FuncApprove, funcApproveThunk)
	exports.AddFunc(FuncInit, funcInitThunk)
	exports.AddFunc(FuncTransfer, funcTransferThunk)
	exports.AddFunc(FuncTransferFrom, funcTransferFromThunk)
	exports.AddView(ViewAllowance, viewAllowanceThunk)
	exports.AddView(ViewBalanceOf, viewBalanceOfThunk)
	exports.AddView(ViewTotalSupply, viewTotalSupplyThunk)
}

type ApproveContext struct {
	Events Erc20Events
	Params ImmutableApproveParams
	State  MutableErc20State
}

func funcApproveThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("erc20.funcApprove")
	f := &ApproveContext{
		Params: ImmutableApproveParams{
			proxy: wasmlib.NewParamsProxy(),
		},
		State: MutableErc20State{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	ctx.Require(f.Params.Amount().Exists(), "missing mandatory amount")
	ctx.Require(f.Params.Delegation().Exists(), "missing mandatory delegation")
	funcApprove(ctx, f)
	ctx.Log("erc20.funcApprove ok")
}

type InitContext struct {
	Events Erc20Events
	Params ImmutableInitParams
	State  MutableErc20State
}

func funcInitThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("erc20.funcInit")
	f := &InitContext{
		Params: ImmutableInitParams{
			proxy: wasmlib.NewParamsProxy(),
		},
		State: MutableErc20State{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	ctx.Require(f.Params.Creator().Exists(), "missing mandatory creator")
	ctx.Require(f.Params.Supply().Exists(), "missing mandatory supply")
	funcInit(ctx, f)
	ctx.Log("erc20.funcInit ok")
}

type TransferContext struct {
	Events Erc20Events
	Params ImmutableTransferParams
	State  MutableErc20State
}

func funcTransferThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("erc20.funcTransfer")
	f := &TransferContext{
		Params: ImmutableTransferParams{
			proxy: wasmlib.NewParamsProxy(),
		},
		State: MutableErc20State{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	ctx.Require(f.Params.Account().Exists(), "missing mandatory account")
	ctx.Require(f.Params.Amount().Exists(), "missing mandatory amount")
	funcTransfer(ctx, f)
	ctx.Log("erc20.funcTransfer ok")
}

type TransferFromContext struct {
	Events Erc20Events
	Params ImmutableTransferFromParams
	State  MutableErc20State
}

func funcTransferFromThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("erc20.funcTransferFrom")
	f := &TransferFromContext{
		Params: ImmutableTransferFromParams{
			proxy: wasmlib.NewParamsProxy(),
		},
		State: MutableErc20State{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	ctx.Require(f.Params.Account().Exists(), "missing mandatory account")
	ctx.Require(f.Params.Amount().Exists(), "missing mandatory amount")
	ctx.Require(f.Params.Recipient().Exists(), "missing mandatory recipient")
	funcTransferFrom(ctx, f)
	ctx.Log("erc20.funcTransferFrom ok")
}

type AllowanceContext struct {
	Params  ImmutableAllowanceParams
	Results MutableAllowanceResults
	State   ImmutableErc20State
}

func viewAllowanceThunk(ctx wasmlib.ScViewContext) {
	ctx.Log("erc20.viewAllowance")
	results := wasmlib.NewScDict()
	f := &AllowanceContext{
		Params: ImmutableAllowanceParams{
			proxy: wasmlib.NewParamsProxy(),
		},
		Results: MutableAllowanceResults{
			proxy: results.AsProxy(),
		},
		State: ImmutableErc20State{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	ctx.Require(f.Params.Account().Exists(), "missing mandatory account")
	ctx.Require(f.Params.Delegation().Exists(), "missing mandatory delegation")
	viewAllowance(ctx, f)
	ctx.Log("erc20.viewAllowance ok")
	ctx.Results(results)
}

type BalanceOfContext struct {
	Params  ImmutableBalanceOfParams
	Results MutableBalanceOfResults
	State   ImmutableErc20State
}

func viewBalanceOfThunk(ctx wasmlib.ScViewContext) {
	ctx.Log("erc20.viewBalanceOf")
	results := wasmlib.NewScDict()
	f := &BalanceOfContext{
		Params: ImmutableBalanceOfParams{
			proxy: wasmlib.NewParamsProxy(),
		},
		Results: MutableBalanceOfResults{
			proxy: results.AsProxy(),
		},
		State: ImmutableErc20State{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	ctx.Require(f.Params.Account().Exists(), "missing mandatory account")
	viewBalanceOf(ctx, f)
	ctx.Log("erc20.viewBalanceOf ok")
	ctx.Results(results)
}

type TotalSupplyContext struct {
	Results MutableTotalSupplyResults
	State   ImmutableErc20State
}

func viewTotalSupplyThunk(ctx wasmlib.ScViewContext) {
	ctx.Log("erc20.viewTotalSupply")
	results := wasmlib.NewScDict()
	f := &TotalSupplyContext{
		Results: MutableTotalSupplyResults{
			proxy: results.AsProxy(),
		},
		State: ImmutableErc20State{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	viewTotalSupply(ctx, f)
	ctx.Log("erc20.viewTotalSupply ok")
	ctx.Results(results)
}
