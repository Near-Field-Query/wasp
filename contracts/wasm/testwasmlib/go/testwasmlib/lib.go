// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

//nolint:dupl
package testwasmlib

import "github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib"

var exportMap = wasmlib.ScExportMap{
	Names: []string{
    	FuncArrayOfArraysAddrAppend,
    	FuncArrayOfArraysAddrClear,
    	FuncArrayOfArraysAddrSet,
    	FuncArrayOfArraysAppend,
    	FuncArrayOfArraysClear,
    	FuncArrayOfArraysSet,
    	FuncArrayOfMapsClear,
    	FuncArrayOfMapsSet,
    	FuncMapOfArraysAddrAppend,
    	FuncMapOfArraysAddrClear,
    	FuncMapOfArraysAddrSet,
    	FuncMapOfArraysAppend,
    	FuncMapOfArraysClear,
    	FuncMapOfArraysSet,
    	FuncMapOfMapsAddrClear,
    	FuncMapOfMapsAddrSet,
    	FuncMapOfMapsClear,
    	FuncMapOfMapsSet,
    	FuncParamTypes,
    	FuncRandom,
    	FuncTakeAllowance,
    	FuncTakeBalance,
    	FuncTriggerEvent,
    	ViewArrayOfArraysAddrLength,
    	ViewArrayOfArraysAddrValue,
    	ViewArrayOfArraysLength,
    	ViewArrayOfArraysValue,
    	ViewArrayOfMapsValue,
    	ViewBlockRecord,
    	ViewBlockRecords,
    	ViewGetRandom,
    	ViewIotaBalance,
    	ViewMapOfArraysAddrLength,
    	ViewMapOfArraysAddrValue,
    	ViewMapOfArraysLength,
    	ViewMapOfArraysValue,
    	ViewMapOfMapsAddrValue,
    	ViewMapOfMapsValue,
	},
	Funcs: []wasmlib.ScFuncContextFunction{
    	funcArrayOfArraysAddrAppendThunk,
    	funcArrayOfArraysAddrClearThunk,
    	funcArrayOfArraysAddrSetThunk,
    	funcArrayOfArraysAppendThunk,
    	funcArrayOfArraysClearThunk,
    	funcArrayOfArraysSetThunk,
    	funcArrayOfMapsClearThunk,
    	funcArrayOfMapsSetThunk,
    	funcMapOfArraysAddrAppendThunk,
    	funcMapOfArraysAddrClearThunk,
    	funcMapOfArraysAddrSetThunk,
    	funcMapOfArraysAppendThunk,
    	funcMapOfArraysClearThunk,
    	funcMapOfArraysSetThunk,
    	funcMapOfMapsAddrClearThunk,
    	funcMapOfMapsAddrSetThunk,
    	funcMapOfMapsClearThunk,
    	funcMapOfMapsSetThunk,
    	funcParamTypesThunk,
    	funcRandomThunk,
    	funcTakeAllowanceThunk,
    	funcTakeBalanceThunk,
    	funcTriggerEventThunk,
	},
	Views: []wasmlib.ScViewContextFunction{
    	viewArrayOfArraysAddrLengthThunk,
    	viewArrayOfArraysAddrValueThunk,
    	viewArrayOfArraysLengthThunk,
    	viewArrayOfArraysValueThunk,
    	viewArrayOfMapsValueThunk,
    	viewBlockRecordThunk,
    	viewBlockRecordsThunk,
    	viewGetRandomThunk,
    	viewIotaBalanceThunk,
    	viewMapOfArraysAddrLengthThunk,
    	viewMapOfArraysAddrValueThunk,
    	viewMapOfArraysLengthThunk,
    	viewMapOfArraysValueThunk,
    	viewMapOfMapsAddrValueThunk,
    	viewMapOfMapsValueThunk,
	},
}

func OnLoad(index int32) {
	if index >= 0 {
		wasmlib.ScExportsCall(index, &exportMap)
		return
	}

	wasmlib.ScExportsExport(&exportMap)
}

type ArrayOfArraysAddrAppendContext struct {
	Events  TestWasmLibEvents
	Params  ImmutableArrayOfArraysAddrAppendParams
	State   MutableTestWasmLibState
}

func funcArrayOfArraysAddrAppendThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("testwasmlib.funcArrayOfArraysAddrAppend")
	f := &ArrayOfArraysAddrAppendContext{
		Params: ImmutableArrayOfArraysAddrAppendParams{
			proxy: wasmlib.NewParamsProxy(),
		},
		State: MutableTestWasmLibState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	ctx.Require(f.Params.Index().Exists(), "missing mandatory index")
	funcArrayOfArraysAddrAppend(ctx, f)
	ctx.Log("testwasmlib.funcArrayOfArraysAddrAppend ok")
}

type ArrayOfArraysAddrClearContext struct {
	Events  TestWasmLibEvents
	State   MutableTestWasmLibState
}

func funcArrayOfArraysAddrClearThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("testwasmlib.funcArrayOfArraysAddrClear")
	f := &ArrayOfArraysAddrClearContext{
		State: MutableTestWasmLibState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	funcArrayOfArraysAddrClear(ctx, f)
	ctx.Log("testwasmlib.funcArrayOfArraysAddrClear ok")
}

type ArrayOfArraysAddrSetContext struct {
	Events  TestWasmLibEvents
	Params  ImmutableArrayOfArraysAddrSetParams
	State   MutableTestWasmLibState
}

func funcArrayOfArraysAddrSetThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("testwasmlib.funcArrayOfArraysAddrSet")
	f := &ArrayOfArraysAddrSetContext{
		Params: ImmutableArrayOfArraysAddrSetParams{
			proxy: wasmlib.NewParamsProxy(),
		},
		State: MutableTestWasmLibState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	ctx.Require(f.Params.Index0().Exists(), "missing mandatory index0")
	ctx.Require(f.Params.Index1().Exists(), "missing mandatory index1")
	ctx.Require(f.Params.ValueAddr().Exists(), "missing mandatory valueAddr")
	funcArrayOfArraysAddrSet(ctx, f)
	ctx.Log("testwasmlib.funcArrayOfArraysAddrSet ok")
}

type ArrayOfArraysAppendContext struct {
	Events  TestWasmLibEvents
	Params  ImmutableArrayOfArraysAppendParams
	State   MutableTestWasmLibState
}

func funcArrayOfArraysAppendThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("testwasmlib.funcArrayOfArraysAppend")
	f := &ArrayOfArraysAppendContext{
		Params: ImmutableArrayOfArraysAppendParams{
			proxy: wasmlib.NewParamsProxy(),
		},
		State: MutableTestWasmLibState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	ctx.Require(f.Params.Index().Exists(), "missing mandatory index")
	funcArrayOfArraysAppend(ctx, f)
	ctx.Log("testwasmlib.funcArrayOfArraysAppend ok")
}

type ArrayOfArraysClearContext struct {
	Events  TestWasmLibEvents
	State   MutableTestWasmLibState
}

func funcArrayOfArraysClearThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("testwasmlib.funcArrayOfArraysClear")
	f := &ArrayOfArraysClearContext{
		State: MutableTestWasmLibState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	funcArrayOfArraysClear(ctx, f)
	ctx.Log("testwasmlib.funcArrayOfArraysClear ok")
}

type ArrayOfArraysSetContext struct {
	Events  TestWasmLibEvents
	Params  ImmutableArrayOfArraysSetParams
	State   MutableTestWasmLibState
}

func funcArrayOfArraysSetThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("testwasmlib.funcArrayOfArraysSet")
	f := &ArrayOfArraysSetContext{
		Params: ImmutableArrayOfArraysSetParams{
			proxy: wasmlib.NewParamsProxy(),
		},
		State: MutableTestWasmLibState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	ctx.Require(f.Params.Index0().Exists(), "missing mandatory index0")
	ctx.Require(f.Params.Index1().Exists(), "missing mandatory index1")
	ctx.Require(f.Params.Value().Exists(), "missing mandatory value")
	funcArrayOfArraysSet(ctx, f)
	ctx.Log("testwasmlib.funcArrayOfArraysSet ok")
}

type ArrayOfMapsClearContext struct {
	Events  TestWasmLibEvents
	State   MutableTestWasmLibState
}

func funcArrayOfMapsClearThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("testwasmlib.funcArrayOfMapsClear")
	f := &ArrayOfMapsClearContext{
		State: MutableTestWasmLibState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	funcArrayOfMapsClear(ctx, f)
	ctx.Log("testwasmlib.funcArrayOfMapsClear ok")
}

type ArrayOfMapsSetContext struct {
	Events  TestWasmLibEvents
	Params  ImmutableArrayOfMapsSetParams
	State   MutableTestWasmLibState
}

func funcArrayOfMapsSetThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("testwasmlib.funcArrayOfMapsSet")
	f := &ArrayOfMapsSetContext{
		Params: ImmutableArrayOfMapsSetParams{
			proxy: wasmlib.NewParamsProxy(),
		},
		State: MutableTestWasmLibState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	ctx.Require(f.Params.Index().Exists(), "missing mandatory index")
	ctx.Require(f.Params.Key().Exists(), "missing mandatory key")
	ctx.Require(f.Params.Value().Exists(), "missing mandatory value")
	funcArrayOfMapsSet(ctx, f)
	ctx.Log("testwasmlib.funcArrayOfMapsSet ok")
}

type MapOfArraysAddrAppendContext struct {
	Events  TestWasmLibEvents
	Params  ImmutableMapOfArraysAddrAppendParams
	State   MutableTestWasmLibState
}

func funcMapOfArraysAddrAppendThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("testwasmlib.funcMapOfArraysAddrAppend")
	f := &MapOfArraysAddrAppendContext{
		Params: ImmutableMapOfArraysAddrAppendParams{
			proxy: wasmlib.NewParamsProxy(),
		},
		State: MutableTestWasmLibState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	ctx.Require(f.Params.NameAddr().Exists(), "missing mandatory nameAddr")
	ctx.Require(f.Params.ValueAddr().Exists(), "missing mandatory valueAddr")
	funcMapOfArraysAddrAppend(ctx, f)
	ctx.Log("testwasmlib.funcMapOfArraysAddrAppend ok")
}

type MapOfArraysAddrClearContext struct {
	Events  TestWasmLibEvents
	Params  ImmutableMapOfArraysAddrClearParams
	State   MutableTestWasmLibState
}

func funcMapOfArraysAddrClearThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("testwasmlib.funcMapOfArraysAddrClear")
	f := &MapOfArraysAddrClearContext{
		Params: ImmutableMapOfArraysAddrClearParams{
			proxy: wasmlib.NewParamsProxy(),
		},
		State: MutableTestWasmLibState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	ctx.Require(f.Params.NameAddr().Exists(), "missing mandatory nameAddr")
	funcMapOfArraysAddrClear(ctx, f)
	ctx.Log("testwasmlib.funcMapOfArraysAddrClear ok")
}

type MapOfArraysAddrSetContext struct {
	Events  TestWasmLibEvents
	Params  ImmutableMapOfArraysAddrSetParams
	State   MutableTestWasmLibState
}

func funcMapOfArraysAddrSetThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("testwasmlib.funcMapOfArraysAddrSet")
	f := &MapOfArraysAddrSetContext{
		Params: ImmutableMapOfArraysAddrSetParams{
			proxy: wasmlib.NewParamsProxy(),
		},
		State: MutableTestWasmLibState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	ctx.Require(f.Params.Index().Exists(), "missing mandatory index")
	ctx.Require(f.Params.NameAddr().Exists(), "missing mandatory nameAddr")
	ctx.Require(f.Params.ValueAddr().Exists(), "missing mandatory valueAddr")
	funcMapOfArraysAddrSet(ctx, f)
	ctx.Log("testwasmlib.funcMapOfArraysAddrSet ok")
}

type MapOfArraysAppendContext struct {
	Events  TestWasmLibEvents
	Params  ImmutableMapOfArraysAppendParams
	State   MutableTestWasmLibState
}

func funcMapOfArraysAppendThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("testwasmlib.funcMapOfArraysAppend")
	f := &MapOfArraysAppendContext{
		Params: ImmutableMapOfArraysAppendParams{
			proxy: wasmlib.NewParamsProxy(),
		},
		State: MutableTestWasmLibState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	ctx.Require(f.Params.Name().Exists(), "missing mandatory name")
	ctx.Require(f.Params.Value().Exists(), "missing mandatory value")
	funcMapOfArraysAppend(ctx, f)
	ctx.Log("testwasmlib.funcMapOfArraysAppend ok")
}

type MapOfArraysClearContext struct {
	Events  TestWasmLibEvents
	Params  ImmutableMapOfArraysClearParams
	State   MutableTestWasmLibState
}

func funcMapOfArraysClearThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("testwasmlib.funcMapOfArraysClear")
	f := &MapOfArraysClearContext{
		Params: ImmutableMapOfArraysClearParams{
			proxy: wasmlib.NewParamsProxy(),
		},
		State: MutableTestWasmLibState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	ctx.Require(f.Params.Name().Exists(), "missing mandatory name")
	funcMapOfArraysClear(ctx, f)
	ctx.Log("testwasmlib.funcMapOfArraysClear ok")
}

type MapOfArraysSetContext struct {
	Events  TestWasmLibEvents
	Params  ImmutableMapOfArraysSetParams
	State   MutableTestWasmLibState
}

func funcMapOfArraysSetThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("testwasmlib.funcMapOfArraysSet")
	f := &MapOfArraysSetContext{
		Params: ImmutableMapOfArraysSetParams{
			proxy: wasmlib.NewParamsProxy(),
		},
		State: MutableTestWasmLibState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	ctx.Require(f.Params.Index().Exists(), "missing mandatory index")
	ctx.Require(f.Params.Name().Exists(), "missing mandatory name")
	ctx.Require(f.Params.Value().Exists(), "missing mandatory value")
	funcMapOfArraysSet(ctx, f)
	ctx.Log("testwasmlib.funcMapOfArraysSet ok")
}

type MapOfMapsAddrClearContext struct {
	Events  TestWasmLibEvents
	Params  ImmutableMapOfMapsAddrClearParams
	State   MutableTestWasmLibState
}

func funcMapOfMapsAddrClearThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("testwasmlib.funcMapOfMapsAddrClear")
	f := &MapOfMapsAddrClearContext{
		Params: ImmutableMapOfMapsAddrClearParams{
			proxy: wasmlib.NewParamsProxy(),
		},
		State: MutableTestWasmLibState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	ctx.Require(f.Params.NameAddr().Exists(), "missing mandatory nameAddr")
	funcMapOfMapsAddrClear(ctx, f)
	ctx.Log("testwasmlib.funcMapOfMapsAddrClear ok")
}

type MapOfMapsAddrSetContext struct {
	Events  TestWasmLibEvents
	Params  ImmutableMapOfMapsAddrSetParams
	State   MutableTestWasmLibState
}

func funcMapOfMapsAddrSetThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("testwasmlib.funcMapOfMapsAddrSet")
	f := &MapOfMapsAddrSetContext{
		Params: ImmutableMapOfMapsAddrSetParams{
			proxy: wasmlib.NewParamsProxy(),
		},
		State: MutableTestWasmLibState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	ctx.Require(f.Params.KeyAddr().Exists(), "missing mandatory keyAddr")
	ctx.Require(f.Params.NameAddr().Exists(), "missing mandatory nameAddr")
	ctx.Require(f.Params.ValueAddr().Exists(), "missing mandatory valueAddr")
	funcMapOfMapsAddrSet(ctx, f)
	ctx.Log("testwasmlib.funcMapOfMapsAddrSet ok")
}

type MapOfMapsClearContext struct {
	Events  TestWasmLibEvents
	Params  ImmutableMapOfMapsClearParams
	State   MutableTestWasmLibState
}

func funcMapOfMapsClearThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("testwasmlib.funcMapOfMapsClear")
	f := &MapOfMapsClearContext{
		Params: ImmutableMapOfMapsClearParams{
			proxy: wasmlib.NewParamsProxy(),
		},
		State: MutableTestWasmLibState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	ctx.Require(f.Params.Name().Exists(), "missing mandatory name")
	funcMapOfMapsClear(ctx, f)
	ctx.Log("testwasmlib.funcMapOfMapsClear ok")
}

type MapOfMapsSetContext struct {
	Events  TestWasmLibEvents
	Params  ImmutableMapOfMapsSetParams
	State   MutableTestWasmLibState
}

func funcMapOfMapsSetThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("testwasmlib.funcMapOfMapsSet")
	f := &MapOfMapsSetContext{
		Params: ImmutableMapOfMapsSetParams{
			proxy: wasmlib.NewParamsProxy(),
		},
		State: MutableTestWasmLibState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	ctx.Require(f.Params.Key().Exists(), "missing mandatory key")
	ctx.Require(f.Params.Name().Exists(), "missing mandatory name")
	ctx.Require(f.Params.Value().Exists(), "missing mandatory value")
	funcMapOfMapsSet(ctx, f)
	ctx.Log("testwasmlib.funcMapOfMapsSet ok")
}

type ParamTypesContext struct {
	Events  TestWasmLibEvents
	Params  ImmutableParamTypesParams
	State   MutableTestWasmLibState
}

func funcParamTypesThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("testwasmlib.funcParamTypes")
	f := &ParamTypesContext{
		Params: ImmutableParamTypesParams{
			proxy: wasmlib.NewParamsProxy(),
		},
		State: MutableTestWasmLibState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	funcParamTypes(ctx, f)
	ctx.Log("testwasmlib.funcParamTypes ok")
}

type RandomContext struct {
	Events  TestWasmLibEvents
	State   MutableTestWasmLibState
}

func funcRandomThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("testwasmlib.funcRandom")
	f := &RandomContext{
		State: MutableTestWasmLibState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	funcRandom(ctx, f)
	ctx.Log("testwasmlib.funcRandom ok")
}

type TakeAllowanceContext struct {
	Events  TestWasmLibEvents
	State   MutableTestWasmLibState
}

func funcTakeAllowanceThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("testwasmlib.funcTakeAllowance")
	f := &TakeAllowanceContext{
		State: MutableTestWasmLibState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	funcTakeAllowance(ctx, f)
	ctx.Log("testwasmlib.funcTakeAllowance ok")
}

type TakeBalanceContext struct {
	Events  TestWasmLibEvents
	Results MutableTakeBalanceResults
	State   MutableTestWasmLibState
}

func funcTakeBalanceThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("testwasmlib.funcTakeBalance")
	results := wasmlib.NewScDict()
	f := &TakeBalanceContext{
		Results: MutableTakeBalanceResults{
			proxy: results.AsProxy(),
		},
		State: MutableTestWasmLibState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	funcTakeBalance(ctx, f)
	ctx.Results(results)
	ctx.Log("testwasmlib.funcTakeBalance ok")
}

type TriggerEventContext struct {
	Events  TestWasmLibEvents
	Params  ImmutableTriggerEventParams
	State   MutableTestWasmLibState
}

func funcTriggerEventThunk(ctx wasmlib.ScFuncContext) {
	ctx.Log("testwasmlib.funcTriggerEvent")
	f := &TriggerEventContext{
		Params: ImmutableTriggerEventParams{
			proxy: wasmlib.NewParamsProxy(),
		},
		State: MutableTestWasmLibState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	ctx.Require(f.Params.Address().Exists(), "missing mandatory address")
	ctx.Require(f.Params.Name().Exists(), "missing mandatory name")
	funcTriggerEvent(ctx, f)
	ctx.Log("testwasmlib.funcTriggerEvent ok")
}

type ArrayOfArraysAddrLengthContext struct {
	Results MutableArrayOfArraysAddrLengthResults
	State   ImmutableTestWasmLibState
}

func viewArrayOfArraysAddrLengthThunk(ctx wasmlib.ScViewContext) {
	ctx.Log("testwasmlib.viewArrayOfArraysAddrLength")
	results := wasmlib.NewScDict()
	f := &ArrayOfArraysAddrLengthContext{
		Results: MutableArrayOfArraysAddrLengthResults{
			proxy: results.AsProxy(),
		},
		State: ImmutableTestWasmLibState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	viewArrayOfArraysAddrLength(ctx, f)
	ctx.Results(results)
	ctx.Log("testwasmlib.viewArrayOfArraysAddrLength ok")
}

type ArrayOfArraysAddrValueContext struct {
	Params  ImmutableArrayOfArraysAddrValueParams
	Results MutableArrayOfArraysAddrValueResults
	State   ImmutableTestWasmLibState
}

func viewArrayOfArraysAddrValueThunk(ctx wasmlib.ScViewContext) {
	ctx.Log("testwasmlib.viewArrayOfArraysAddrValue")
	results := wasmlib.NewScDict()
	f := &ArrayOfArraysAddrValueContext{
		Params: ImmutableArrayOfArraysAddrValueParams{
			proxy: wasmlib.NewParamsProxy(),
		},
		Results: MutableArrayOfArraysAddrValueResults{
			proxy: results.AsProxy(),
		},
		State: ImmutableTestWasmLibState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	ctx.Require(f.Params.Index0().Exists(), "missing mandatory index0")
	ctx.Require(f.Params.Index1().Exists(), "missing mandatory index1")
	viewArrayOfArraysAddrValue(ctx, f)
	ctx.Results(results)
	ctx.Log("testwasmlib.viewArrayOfArraysAddrValue ok")
}

type ArrayOfArraysLengthContext struct {
	Results MutableArrayOfArraysLengthResults
	State   ImmutableTestWasmLibState
}

func viewArrayOfArraysLengthThunk(ctx wasmlib.ScViewContext) {
	ctx.Log("testwasmlib.viewArrayOfArraysLength")
	results := wasmlib.NewScDict()
	f := &ArrayOfArraysLengthContext{
		Results: MutableArrayOfArraysLengthResults{
			proxy: results.AsProxy(),
		},
		State: ImmutableTestWasmLibState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	viewArrayOfArraysLength(ctx, f)
	ctx.Results(results)
	ctx.Log("testwasmlib.viewArrayOfArraysLength ok")
}

type ArrayOfArraysValueContext struct {
	Params  ImmutableArrayOfArraysValueParams
	Results MutableArrayOfArraysValueResults
	State   ImmutableTestWasmLibState
}

func viewArrayOfArraysValueThunk(ctx wasmlib.ScViewContext) {
	ctx.Log("testwasmlib.viewArrayOfArraysValue")
	results := wasmlib.NewScDict()
	f := &ArrayOfArraysValueContext{
		Params: ImmutableArrayOfArraysValueParams{
			proxy: wasmlib.NewParamsProxy(),
		},
		Results: MutableArrayOfArraysValueResults{
			proxy: results.AsProxy(),
		},
		State: ImmutableTestWasmLibState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	ctx.Require(f.Params.Index0().Exists(), "missing mandatory index0")
	ctx.Require(f.Params.Index1().Exists(), "missing mandatory index1")
	viewArrayOfArraysValue(ctx, f)
	ctx.Results(results)
	ctx.Log("testwasmlib.viewArrayOfArraysValue ok")
}

type ArrayOfMapsValueContext struct {
	Params  ImmutableArrayOfMapsValueParams
	Results MutableArrayOfMapsValueResults
	State   ImmutableTestWasmLibState
}

func viewArrayOfMapsValueThunk(ctx wasmlib.ScViewContext) {
	ctx.Log("testwasmlib.viewArrayOfMapsValue")
	results := wasmlib.NewScDict()
	f := &ArrayOfMapsValueContext{
		Params: ImmutableArrayOfMapsValueParams{
			proxy: wasmlib.NewParamsProxy(),
		},
		Results: MutableArrayOfMapsValueResults{
			proxy: results.AsProxy(),
		},
		State: ImmutableTestWasmLibState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	ctx.Require(f.Params.Index().Exists(), "missing mandatory index")
	ctx.Require(f.Params.Key().Exists(), "missing mandatory key")
	viewArrayOfMapsValue(ctx, f)
	ctx.Results(results)
	ctx.Log("testwasmlib.viewArrayOfMapsValue ok")
}

type BlockRecordContext struct {
	Params  ImmutableBlockRecordParams
	Results MutableBlockRecordResults
	State   ImmutableTestWasmLibState
}

func viewBlockRecordThunk(ctx wasmlib.ScViewContext) {
	ctx.Log("testwasmlib.viewBlockRecord")
	results := wasmlib.NewScDict()
	f := &BlockRecordContext{
		Params: ImmutableBlockRecordParams{
			proxy: wasmlib.NewParamsProxy(),
		},
		Results: MutableBlockRecordResults{
			proxy: results.AsProxy(),
		},
		State: ImmutableTestWasmLibState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	ctx.Require(f.Params.BlockIndex().Exists(), "missing mandatory blockIndex")
	ctx.Require(f.Params.RecordIndex().Exists(), "missing mandatory recordIndex")
	viewBlockRecord(ctx, f)
	ctx.Results(results)
	ctx.Log("testwasmlib.viewBlockRecord ok")
}

type BlockRecordsContext struct {
	Params  ImmutableBlockRecordsParams
	Results MutableBlockRecordsResults
	State   ImmutableTestWasmLibState
}

func viewBlockRecordsThunk(ctx wasmlib.ScViewContext) {
	ctx.Log("testwasmlib.viewBlockRecords")
	results := wasmlib.NewScDict()
	f := &BlockRecordsContext{
		Params: ImmutableBlockRecordsParams{
			proxy: wasmlib.NewParamsProxy(),
		},
		Results: MutableBlockRecordsResults{
			proxy: results.AsProxy(),
		},
		State: ImmutableTestWasmLibState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	ctx.Require(f.Params.BlockIndex().Exists(), "missing mandatory blockIndex")
	viewBlockRecords(ctx, f)
	ctx.Results(results)
	ctx.Log("testwasmlib.viewBlockRecords ok")
}

type GetRandomContext struct {
	Results MutableGetRandomResults
	State   ImmutableTestWasmLibState
}

func viewGetRandomThunk(ctx wasmlib.ScViewContext) {
	ctx.Log("testwasmlib.viewGetRandom")
	results := wasmlib.NewScDict()
	f := &GetRandomContext{
		Results: MutableGetRandomResults{
			proxy: results.AsProxy(),
		},
		State: ImmutableTestWasmLibState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	viewGetRandom(ctx, f)
	ctx.Results(results)
	ctx.Log("testwasmlib.viewGetRandom ok")
}

type IotaBalanceContext struct {
	Results MutableIotaBalanceResults
	State   ImmutableTestWasmLibState
}

func viewIotaBalanceThunk(ctx wasmlib.ScViewContext) {
	ctx.Log("testwasmlib.viewIotaBalance")
	results := wasmlib.NewScDict()
	f := &IotaBalanceContext{
		Results: MutableIotaBalanceResults{
			proxy: results.AsProxy(),
		},
		State: ImmutableTestWasmLibState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	viewIotaBalance(ctx, f)
	ctx.Results(results)
	ctx.Log("testwasmlib.viewIotaBalance ok")
}

type MapOfArraysAddrLengthContext struct {
	Params  ImmutableMapOfArraysAddrLengthParams
	Results MutableMapOfArraysAddrLengthResults
	State   ImmutableTestWasmLibState
}

func viewMapOfArraysAddrLengthThunk(ctx wasmlib.ScViewContext) {
	ctx.Log("testwasmlib.viewMapOfArraysAddrLength")
	results := wasmlib.NewScDict()
	f := &MapOfArraysAddrLengthContext{
		Params: ImmutableMapOfArraysAddrLengthParams{
			proxy: wasmlib.NewParamsProxy(),
		},
		Results: MutableMapOfArraysAddrLengthResults{
			proxy: results.AsProxy(),
		},
		State: ImmutableTestWasmLibState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	ctx.Require(f.Params.NameAddr().Exists(), "missing mandatory nameAddr")
	viewMapOfArraysAddrLength(ctx, f)
	ctx.Results(results)
	ctx.Log("testwasmlib.viewMapOfArraysAddrLength ok")
}

type MapOfArraysAddrValueContext struct {
	Params  ImmutableMapOfArraysAddrValueParams
	Results MutableMapOfArraysAddrValueResults
	State   ImmutableTestWasmLibState
}

func viewMapOfArraysAddrValueThunk(ctx wasmlib.ScViewContext) {
	ctx.Log("testwasmlib.viewMapOfArraysAddrValue")
	results := wasmlib.NewScDict()
	f := &MapOfArraysAddrValueContext{
		Params: ImmutableMapOfArraysAddrValueParams{
			proxy: wasmlib.NewParamsProxy(),
		},
		Results: MutableMapOfArraysAddrValueResults{
			proxy: results.AsProxy(),
		},
		State: ImmutableTestWasmLibState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	ctx.Require(f.Params.Index().Exists(), "missing mandatory index")
	ctx.Require(f.Params.NameAddr().Exists(), "missing mandatory nameAddr")
	viewMapOfArraysAddrValue(ctx, f)
	ctx.Results(results)
	ctx.Log("testwasmlib.viewMapOfArraysAddrValue ok")
}

type MapOfArraysLengthContext struct {
	Params  ImmutableMapOfArraysLengthParams
	Results MutableMapOfArraysLengthResults
	State   ImmutableTestWasmLibState
}

func viewMapOfArraysLengthThunk(ctx wasmlib.ScViewContext) {
	ctx.Log("testwasmlib.viewMapOfArraysLength")
	results := wasmlib.NewScDict()
	f := &MapOfArraysLengthContext{
		Params: ImmutableMapOfArraysLengthParams{
			proxy: wasmlib.NewParamsProxy(),
		},
		Results: MutableMapOfArraysLengthResults{
			proxy: results.AsProxy(),
		},
		State: ImmutableTestWasmLibState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	ctx.Require(f.Params.Name().Exists(), "missing mandatory name")
	viewMapOfArraysLength(ctx, f)
	ctx.Results(results)
	ctx.Log("testwasmlib.viewMapOfArraysLength ok")
}

type MapOfArraysValueContext struct {
	Params  ImmutableMapOfArraysValueParams
	Results MutableMapOfArraysValueResults
	State   ImmutableTestWasmLibState
}

func viewMapOfArraysValueThunk(ctx wasmlib.ScViewContext) {
	ctx.Log("testwasmlib.viewMapOfArraysValue")
	results := wasmlib.NewScDict()
	f := &MapOfArraysValueContext{
		Params: ImmutableMapOfArraysValueParams{
			proxy: wasmlib.NewParamsProxy(),
		},
		Results: MutableMapOfArraysValueResults{
			proxy: results.AsProxy(),
		},
		State: ImmutableTestWasmLibState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	ctx.Require(f.Params.Index().Exists(), "missing mandatory index")
	ctx.Require(f.Params.Name().Exists(), "missing mandatory name")
	viewMapOfArraysValue(ctx, f)
	ctx.Results(results)
	ctx.Log("testwasmlib.viewMapOfArraysValue ok")
}

type MapOfMapsAddrValueContext struct {
	Params  ImmutableMapOfMapsAddrValueParams
	Results MutableMapOfMapsAddrValueResults
	State   ImmutableTestWasmLibState
}

func viewMapOfMapsAddrValueThunk(ctx wasmlib.ScViewContext) {
	ctx.Log("testwasmlib.viewMapOfMapsAddrValue")
	results := wasmlib.NewScDict()
	f := &MapOfMapsAddrValueContext{
		Params: ImmutableMapOfMapsAddrValueParams{
			proxy: wasmlib.NewParamsProxy(),
		},
		Results: MutableMapOfMapsAddrValueResults{
			proxy: results.AsProxy(),
		},
		State: ImmutableTestWasmLibState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	ctx.Require(f.Params.KeyAddr().Exists(), "missing mandatory keyAddr")
	ctx.Require(f.Params.NameAddr().Exists(), "missing mandatory nameAddr")
	viewMapOfMapsAddrValue(ctx, f)
	ctx.Results(results)
	ctx.Log("testwasmlib.viewMapOfMapsAddrValue ok")
}

type MapOfMapsValueContext struct {
	Params  ImmutableMapOfMapsValueParams
	Results MutableMapOfMapsValueResults
	State   ImmutableTestWasmLibState
}

func viewMapOfMapsValueThunk(ctx wasmlib.ScViewContext) {
	ctx.Log("testwasmlib.viewMapOfMapsValue")
	results := wasmlib.NewScDict()
	f := &MapOfMapsValueContext{
		Params: ImmutableMapOfMapsValueParams{
			proxy: wasmlib.NewParamsProxy(),
		},
		Results: MutableMapOfMapsValueResults{
			proxy: results.AsProxy(),
		},
		State: ImmutableTestWasmLibState{
			proxy: wasmlib.NewStateProxy(),
		},
	}
	ctx.Require(f.Params.Key().Exists(), "missing mandatory key")
	ctx.Require(f.Params.Name().Exists(), "missing mandatory name")
	viewMapOfMapsValue(ctx, f)
	ctx.Results(results)
	ctx.Log("testwasmlib.viewMapOfMapsValue ok")
}
