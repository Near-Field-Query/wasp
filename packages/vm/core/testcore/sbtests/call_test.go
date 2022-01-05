package sbtests

import (
	"testing"

	"github.com/iotaledger/wasp/packages/kv/codec"
	"github.com/iotaledger/wasp/packages/solo"
	"github.com/iotaledger/wasp/packages/vm/core/testcore/sbtests/sbtestsc"
	"github.com/stretchr/testify/require"
)

func TestGetSet(t *testing.T) { run2(t, testGetSet) }
func testGetSet(t *testing.T, w bool) {
	_, chain := setupChain(t, nil)
	setupTestSandboxSC(t, chain, nil, w)

	req := solo.NewCallParams(ScName, sbtestsc.FuncSetInt.Name,
		sbtestsc.ParamIntParamName, "ppp",
		sbtestsc.ParamIntParamValue, 314)
	_, err := chain.PostRequestSync(req.AddIotas(1), nil)
	require.NoError(t, err)

	ret, err := chain.CallView(ScName, sbtestsc.FuncGetInt.Name,
		sbtestsc.ParamIntParamName, "ppp")
	require.NoError(t, err)

	retInt, err := codec.DecodeInt64(ret.MustGet("ppp"))
	require.NoError(t, err)
	require.EqualValues(t, 314, retInt)
}

func TestCallRecursive(t *testing.T) { run2(t, testCallRecursive) }
func testCallRecursive(t *testing.T, w bool) {
	_, chain := setupChain(t, nil)
	setupTestSandboxSC(t, chain, nil, w)

	req := solo.NewCallParams(ScName, sbtestsc.FuncCallOnChain.Name,
		sbtestsc.ParamIntParamValue, 31,
		sbtestsc.ParamHnameContract, HScName,
		sbtestsc.ParamHnameEP, sbtestsc.FuncRunRecursion.Hname())
	_, err := chain.PostRequestSync(req.AddIotas(1), nil)
	require.NoError(t, err)

	ret, err := chain.CallView(ScName, sbtestsc.FuncGetCounter.Name)
	require.NoError(t, err)

	r, err := codec.DecodeInt64(ret.MustGet(sbtestsc.VarCounter))
	require.NoError(t, err)
	require.EqualValues(t, 32, r)
}

const n = 10

func fibo(n int64) int64 {
	if n == 0 || n == 1 {
		return n
	}
	return fibo(n-1) + fibo(n-2)
}

func TestCallFibonacci(t *testing.T) { run2(t, testCallFibonacci) }
func testCallFibonacci(t *testing.T, w bool) {
	_, chain := setupChain(t, nil)
	setupTestSandboxSC(t, chain, nil, w)

	ret, err := chain.CallView(ScName, sbtestsc.FuncGetFibonacci.Name,
		sbtestsc.ParamIntParamValue, n,
	)
	require.NoError(t, err)
	val, err := codec.DecodeInt64(ret.MustGet(sbtestsc.ParamIntParamValue))
	require.NoError(t, err)
	require.EqualValues(t, fibo(n), val)
}

func TestCallFibonacciIndirect(t *testing.T) { run2(t, testCallFibonacciIndirect) }
func testCallFibonacciIndirect(t *testing.T, w bool) {
	_, chain := setupChain(t, nil)
	setupTestSandboxSC(t, chain, nil, w)

	req := solo.NewCallParams(ScName, sbtestsc.FuncCallOnChain.Name,
		sbtestsc.ParamIntParamValue, n,
		sbtestsc.ParamHnameContract, HScName,
		sbtestsc.ParamHnameEP, sbtestsc.FuncGetFibonacci.Hname())
	ret, err := chain.PostRequestSync(req.AddIotas(1), nil)
	require.NoError(t, err)
	r, err := codec.DecodeInt64(ret.MustGet(sbtestsc.ParamIntParamValue))
	require.NoError(t, err)
	require.EqualValues(t, fibo(n), r)

	ret, err = chain.CallView(ScName, sbtestsc.FuncGetCounter.Name)
	require.NoError(t, err)

	r, err = codec.DecodeInt64(ret.MustGet(sbtestsc.VarCounter))
	require.NoError(t, err)
	require.EqualValues(t, 1, r)
}
