// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package fairauction

import "github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib"

func OnLoad() {
	exports := wasmlib.NewScExports()
	exports.AddFunc(FuncFinalizeAuction, funcFinalizeAuctionThunk)
	exports.AddFunc(FuncPlaceBid, funcPlaceBidThunk)
	exports.AddFunc(FuncSetOwnerMargin, funcSetOwnerMarginThunk)
	exports.AddFunc(FuncStartAuction, funcStartAuctionThunk)
	exports.AddView(ViewGetInfo, viewGetInfoThunk)
}

type FinalizeAuctionContext struct {
	Params ImmutableFinalizeAuctionParams
	State  MutableFairAuctionState
}

func funcFinalizeAuctionThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("fairauction.funcFinalizeAuction")
	f := &FinalizeAuctionContext{
		Params: ImmutableFinalizeAuctionParams{
			proxy: wasmlib.NewParamsProxy(),
		},
		State: MutableFairAuctionState{
			proxy: wasmlib.NewStateProxy(),
		},
	}

	// only SC itself can invoke this function
	ctx.Require(ctx.Caller() == ctx.AccountID(), "no permission")

	ctx.Require(f.Params.Color().Exists(), "missing mandatory color")
	funcFinalizeAuction(ctx, f)
	ctx.Log("fairauction.funcFinalizeAuction ok")
}

type PlaceBidContext struct {
	Params ImmutablePlaceBidParams
	State  MutableFairAuctionState
}

func funcPlaceBidThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("fairauction.funcPlaceBid")
	f := &PlaceBidContext{
		Params: ImmutablePlaceBidParams{
			proxy: wasmlib.NewParamsProxy(),
		},
		State: MutableFairAuctionState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	ctx.Require(f.Params.Color().Exists(), "missing mandatory color")
	funcPlaceBid(ctx, f)
	ctx.Log("fairauction.funcPlaceBid ok")
}

type SetOwnerMarginContext struct {
	Params ImmutableSetOwnerMarginParams
	State  MutableFairAuctionState
}

func funcSetOwnerMarginThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("fairauction.funcSetOwnerMargin")
	f := &SetOwnerMarginContext{
		Params: ImmutableSetOwnerMarginParams{
			proxy: wasmlib.NewParamsProxy(),
		},
		State: MutableFairAuctionState{
			proxy: wasmlib.NewStateProxy(),
		},
	}

	// only SC creator can set owner margin
	ctx.Require(ctx.Caller() == ctx.ContractCreator(), "no permission")

	ctx.Require(f.Params.OwnerMargin().Exists(), "missing mandatory ownerMargin")
	funcSetOwnerMargin(ctx, f)
	ctx.Log("fairauction.funcSetOwnerMargin ok")
}

type StartAuctionContext struct {
	Params ImmutableStartAuctionParams
	State  MutableFairAuctionState
}

func funcStartAuctionThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("fairauction.funcStartAuction")
	f := &StartAuctionContext{
		Params: ImmutableStartAuctionParams{
			proxy: wasmlib.NewParamsProxy(),
		},
		State: MutableFairAuctionState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	ctx.Require(f.Params.Color().Exists(), "missing mandatory color")
	ctx.Require(f.Params.MinimumBid().Exists(), "missing mandatory minimumBid")
	funcStartAuction(ctx, f)
	ctx.Log("fairauction.funcStartAuction ok")
}

type GetInfoContext struct {
	Params  ImmutableGetInfoParams
	Results MutableGetInfoResults
	State   ImmutableFairAuctionState
}

func viewGetInfoThunk(ctx wasmlib.ScViewContext) {
	ctx.Log("fairauction.viewGetInfo")
	results := wasmlib.NewScDict()
	f := &GetInfoContext{
		Params: ImmutableGetInfoParams{
			proxy: wasmlib.NewParamsProxy(),
		},
		Results: MutableGetInfoResults{
			proxy: results.AsProxy(),
		},
		State: ImmutableFairAuctionState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	ctx.Require(f.Params.Color().Exists(), "missing mandatory color")
	viewGetInfo(ctx, f)
	ctx.Log("fairauction.viewGetInfo ok")
	ctx.Results(results)
}
