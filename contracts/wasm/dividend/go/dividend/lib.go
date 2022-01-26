// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package dividend

import "github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib"

func OnLoad() {
	exports := wasmlib.NewScExports()
	exports.AddFunc(FuncDivide, funcDivideThunk)
	exports.AddFunc(FuncInit, funcInitThunk)
	exports.AddFunc(FuncMember, funcMemberThunk)
	exports.AddFunc(FuncSetOwner, funcSetOwnerThunk)
	exports.AddView(ViewGetFactor, viewGetFactorThunk)
	exports.AddView(ViewGetOwner, viewGetOwnerThunk)
}

type DivideContext struct {
	State MutableDividendState
}

func funcDivideThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("dividend.funcDivide")
	f := &DivideContext{
		State: MutableDividendState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	funcDivide(ctx, f)
	ctx.Log("dividend.funcDivide ok")
}

type InitContext struct {
	Params ImmutableInitParams
	State  MutableDividendState
}

func funcInitThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("dividend.funcInit")
	f := &InitContext{
		Params: ImmutableInitParams{
			proxy: wasmlib.NewParamsProxy(),
		},
		State: MutableDividendState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	funcInit(ctx, f)
	ctx.Log("dividend.funcInit ok")
}

type MemberContext struct {
	Params ImmutableMemberParams
	State  MutableDividendState
}

func funcMemberThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("dividend.funcMember")
	f := &MemberContext{
		Params: ImmutableMemberParams{
			proxy: wasmlib.NewParamsProxy(),
		},
		State: MutableDividendState{
			proxy: wasmlib.NewStateProxy(),
		},
	}

	// only defined owner of contract can add members
	access := f.State.Owner()
	ctx.Require(access.Exists(), "access not set: owner")
	ctx.Require(ctx.Caller() == access.Value(), "no permission")

	ctx.Require(f.Params.Address().Exists(), "missing mandatory address")
	ctx.Require(f.Params.Factor().Exists(), "missing mandatory factor")
	funcMember(ctx, f)
	ctx.Log("dividend.funcMember ok")
}

type SetOwnerContext struct {
	Params ImmutableSetOwnerParams
	State  MutableDividendState
}

func funcSetOwnerThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("dividend.funcSetOwner")
	f := &SetOwnerContext{
		Params: ImmutableSetOwnerParams{
			proxy: wasmlib.NewParamsProxy(),
		},
		State: MutableDividendState{
			proxy: wasmlib.NewStateProxy(),
		},
	}

	// only defined owner of contract can change owner
	access := f.State.Owner()
	ctx.Require(access.Exists(), "access not set: owner")
	ctx.Require(ctx.Caller() == access.Value(), "no permission")

	ctx.Require(f.Params.Owner().Exists(), "missing mandatory owner")
	funcSetOwner(ctx, f)
	ctx.Log("dividend.funcSetOwner ok")
}

type GetFactorContext struct {
	Params  ImmutableGetFactorParams
	Results MutableGetFactorResults
	State   ImmutableDividendState
}

func viewGetFactorThunk(ctx wasmlib.ScViewContext) {
	ctx.Log("dividend.viewGetFactor")
	results := wasmlib.NewScDict()
	f := &GetFactorContext{
		Params: ImmutableGetFactorParams{
			proxy: wasmlib.NewParamsProxy(),
		},
		Results: MutableGetFactorResults{
			proxy: results.AsProxy(),
		},
		State: ImmutableDividendState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	ctx.Require(f.Params.Address().Exists(), "missing mandatory address")
	viewGetFactor(ctx, f)
	ctx.Log("dividend.viewGetFactor ok")
	ctx.Results(results)
}

type GetOwnerContext struct {
	Results MutableGetOwnerResults
	State   ImmutableDividendState
}

func viewGetOwnerThunk(ctx wasmlib.ScViewContext) {
	ctx.Log("dividend.viewGetOwner")
	results := wasmlib.NewScDict()
	f := &GetOwnerContext{
		Results: MutableGetOwnerResults{
			proxy: results.AsProxy(),
		},
		State: ImmutableDividendState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	viewGetOwner(ctx, f)
	ctx.Log("dividend.viewGetOwner ok")
	ctx.Results(results)
}
