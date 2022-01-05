//nolint:dupl
package sbtests

import (
	"testing"

	"github.com/iotaledger/wasp/packages/solo"
	"github.com/iotaledger/wasp/packages/vm/core/accounts"
	"github.com/iotaledger/wasp/packages/vm/core/governance"
	"github.com/iotaledger/wasp/packages/vm/core/testcore/sbtests/sbtestsc"
	"github.com/stretchr/testify/require"
)

func TestDoNothing(t *testing.T) { run2(t, testDoNothing) }
func testDoNothing(t *testing.T, w bool) {
	env, chain := setupChain(t, nil)
	cAID, extraToken := setupTestSandboxSC(t, chain, nil, w)
	req := solo.NewCallParams(ScName, sbtestsc.FuncDoNothing.Name)
	_, err := chain.PostRequestSync(req.AddIotas(42), nil)
	require.NoError(t, err)

	t.Logf("dump accounts:\n%s", chain.DumpAccounts())
	chain.AssertL2AccountIotas(chain.OriginatorAgentID, 0)
	chain.AssertL2AccountIotas(cAID, 43)
	chain.AssertCommonAccountIotas(2 + extraToken)
	env.AssertAddressIotas(chain.OriginatorAddress, solo.Saldo-solo.ChainDustThreshold-2-1-42-extraToken)
}

func TestDoNothingUser(t *testing.T) { run2(t, testDoNothingUser) }
func testDoNothingUser(t *testing.T, w bool) {
	env, chain := setupChain(t, nil)
	cAID, extraToken := setupTestSandboxSC(t, chain, nil, w)
	user, userAddr, userAgentID := setupDeployer(t, chain)
	req := solo.NewCallParams(ScName, sbtestsc.FuncDoNothing.Name)
	_, err := chain.PostRequestSync(req.AddIotas(42), user)
	require.NoError(t, err)

	t.Logf("dump accounts:\n%s", chain.DumpAccounts())
	chain.AssertL2AccountIotas(chain.OriginatorAgentID, 0)
	chain.AssertL2AccountIotas(userAgentID, 0)
	chain.AssertL2AccountIotas(cAID, 43)
	env.AssertAddressIotas(userAddr, solo.Saldo-42)
	env.AssertAddressIotas(chain.OriginatorAddress, solo.Saldo-solo.ChainDustThreshold-4-extraToken)
	chain.AssertCommonAccountIotas(3 + extraToken)
}

func TestWithdrawToAddress(t *testing.T) { run2(t, testWithdrawToAddress) }
func testWithdrawToAddress(t *testing.T, w bool) {
	env, chain := setupChain(t, nil)
	cAID, extraToken := setupTestSandboxSC(t, chain, nil, w)
	user, userAddress, userAgentID := setupDeployer(t, chain)
	t.Logf("contract agentID: %s", cAID)

	req := solo.NewCallParams(ScName, sbtestsc.FuncDoNothing.Name)
	_, err := chain.PostRequestSync(req.AddIotas(42), user)
	require.NoError(t, err)

	t.Logf("dump accounts 1:\n%s", chain.DumpAccounts())
	chain.AssertL2AccountIotas(chain.OriginatorAgentID, 0)
	chain.AssertL2AccountIotas(userAgentID, 0)
	chain.AssertL2AccountIotas(cAID, 43)
	chain.AssertCommonAccountIotas(3 + extraToken)
	chain.AssertL2TotalIotas(46 + extraToken)
	env.AssertAddressIotas(chain.OriginatorAddress, solo.Saldo-solo.ChainDustThreshold-4-extraToken)
	env.AssertAddressIotas(userAddress, solo.Saldo-42)

	t.Logf("-------- send to address %s", userAddress.Base58())
	req = solo.NewCallParams(ScName, sbtestsc.FuncSendToAddress.Name,
		sbtestsc.ParamAddress, userAddress,
	)
	_, err = chain.PostRequestSync(req.AddIotas(1), nil)
	require.NoError(t, err)

	t.Logf("dump accounts 2:\n%s", chain.DumpAccounts())
	chain.AssertL2AccountIotas(chain.OriginatorAgentID, 0)
	chain.AssertL2AccountIotas(userAgentID, 0)
	chain.AssertL2AccountIotas(cAID, 0)
	chain.AssertCommonAccountIotas(3 + extraToken)
	chain.AssertL2TotalIotas(3 + extraToken)
	env.AssertAddressIotas(chain.OriginatorAddress, solo.Saldo-solo.ChainDustThreshold-5-extraToken)
	env.AssertAddressIotas(userAddress, solo.Saldo+2)
}

func TestDoPanicUser(t *testing.T) { run2(t, testDoPanicUser) }
func testDoPanicUser(t *testing.T, w bool) {
	env, chain := setupChain(t, nil)
	cAID, extraToken := setupTestSandboxSC(t, chain, nil, w)
	user, userAddress, userAgentID := setupDeployer(t, chain)

	t.Logf("dump accounts 1:\n%s", chain.DumpAccounts())
	chain.AssertL2AccountIotas(chain.OriginatorAgentID, 0)
	chain.AssertL2AccountIotas(userAgentID, 0)
	chain.AssertL2AccountIotas(cAID, 1)
	chain.AssertCommonAccountIotas(3 + extraToken)
	chain.AssertL2TotalIotas(4 + extraToken)
	env.AssertAddressIotas(chain.OriginatorAddress, solo.Saldo-solo.ChainDustThreshold-4-extraToken)
	env.AssertAddressIotas(userAddress, solo.Saldo)

	req := solo.NewCallParams(ScName, sbtestsc.FuncPanicFullEP.Name)
	_, err := chain.PostRequestSync(req.AddIotas(42), user)
	require.Error(t, err)

	t.Logf("dump accounts 2:\n%s", chain.DumpAccounts())
	chain.AssertL2AccountIotas(chain.OriginatorAgentID, 0)
	chain.AssertL2AccountIotas(userAgentID, 0)
	chain.AssertL2AccountIotas(cAID, 1)
	chain.AssertCommonAccountIotas(3 + extraToken)
	chain.AssertL2TotalIotas(4 + extraToken)
	env.AssertAddressIotas(chain.OriginatorAddress, solo.Saldo-solo.ChainDustThreshold-4-extraToken)
	env.AssertAddressIotas(userAddress, solo.Saldo)
}

func TestDoPanicUserFeeless(t *testing.T) { run2(t, testDoPanicUserFeeless) }
func testDoPanicUserFeeless(t *testing.T, w bool) {
	env, chain := setupChain(t, nil)
	cAID, extraToken := setupTestSandboxSC(t, chain, nil, w)
	user, userAddress, userAgentID := setupDeployer(t, chain)

	t.Logf("dump accounts 1:\n%s", chain.DumpAccounts())
	chain.AssertL2AccountIotas(chain.OriginatorAgentID, 0)
	chain.AssertL2AccountIotas(userAgentID, 0)
	chain.AssertL2AccountIotas(cAID, 1)
	chain.AssertCommonAccountIotas(3 + extraToken)
	chain.AssertL2TotalIotas(4 + extraToken)
	env.AssertAddressIotas(chain.OriginatorAddress, solo.Saldo-solo.ChainDustThreshold-4-extraToken)
	env.AssertAddressIotas(userAddress, solo.Saldo)

	req := solo.NewCallParams(ScName, sbtestsc.FuncPanicFullEP.Name)
	_, err := chain.PostRequestSync(req.AddIotas(42), user)
	require.Error(t, err)

	t.Logf("dump accounts 2:\n%s", chain.DumpAccounts())
	chain.AssertL2AccountIotas(chain.OriginatorAgentID, 0)
	chain.AssertL2AccountIotas(userAgentID, 0)
	chain.AssertL2AccountIotas(cAID, 1)
	chain.AssertCommonAccountIotas(3 + extraToken)
	chain.AssertL2TotalIotas(4 + extraToken)
	env.AssertAddressIotas(chain.OriginatorAddress, solo.Saldo-solo.ChainDustThreshold-4-extraToken)
	env.AssertAddressIotas(userAddress, solo.Saldo)

	req = solo.NewCallParams(accounts.Contract.Name, accounts.FuncWithdraw.Name)
	_, err = chain.PostRequestSync(req.AddIotas(1), user)
	require.NoError(t, err)

	chain.AssertL2AccountIotas(chain.OriginatorAgentID, 0)
	chain.AssertL2AccountIotas(userAgentID, 0)
	chain.AssertL2AccountIotas(cAID, 1)
	chain.AssertCommonAccountIotas(4 + extraToken)
	chain.AssertL2TotalIotas(5 + extraToken)
	env.AssertAddressIotas(chain.OriginatorAddress, solo.Saldo-solo.ChainDustThreshold-4-extraToken)
	env.AssertAddressIotas(userAddress, solo.Saldo-1)
}

func TestDoPanicUserFee(t *testing.T) { run2(t, testDoPanicUserFee) }
func testDoPanicUserFee(t *testing.T, w bool) {
	env, chain := setupChain(t, nil)
	cAID, extraToken := setupTestSandboxSC(t, chain, nil, w)
	user, userAddress, userAgentID := setupDeployer(t, chain)

	t.Logf("dump accounts 1:\n%s", chain.DumpAccounts())
	chain.AssertL2AccountIotas(chain.OriginatorAgentID, 0)
	chain.AssertL2AccountIotas(userAgentID, 0)
	chain.AssertL2AccountIotas(cAID, 1)
	chain.AssertCommonAccountIotas(3 + extraToken)
	chain.AssertL2TotalIotas(4 + extraToken)
	env.AssertAddressIotas(chain.OriginatorAddress, solo.Saldo-solo.ChainDustThreshold-4-extraToken)
	env.AssertAddressIotas(userAddress, solo.Saldo)

	req := solo.NewCallParams(governance.Contract.Name, governance.FuncSetContractFee.Name,
		governance.ParamHname, cAID.Hname(),
		governance.ParamOwnerFee, 10)
	_, err := chain.PostRequestSync(req.AddIotas(1), nil)
	require.NoError(t, err)

	chain.AssertL2AccountIotas(chain.OriginatorAgentID, 0)
	chain.AssertL2AccountIotas(userAgentID, 0)
	chain.AssertL2AccountIotas(cAID, 1)
	chain.AssertCommonAccountIotas(4 + extraToken)
	chain.AssertL2TotalIotas(5 + extraToken)
	env.AssertAddressIotas(chain.OriginatorAddress, solo.Saldo-solo.ChainDustThreshold-4-1-extraToken)
	env.AssertAddressIotas(userAddress, solo.Saldo)

	req = solo.NewCallParams(ScName, sbtestsc.FuncPanicFullEP.Name)
	_, err = chain.PostRequestSync(req.AddIotas(42), user)
	require.Error(t, err)

	t.Logf("dump accounts 2:\n%s", chain.DumpAccounts())
	chain.AssertL2AccountIotas(chain.OriginatorAgentID, 0)
	chain.AssertL2AccountIotas(userAgentID, 0)
	chain.AssertL2AccountIotas(cAID, 1)
	chain.AssertCommonAccountIotas(14 + extraToken)
	chain.AssertL2TotalIotas(15 + extraToken)
	env.AssertAddressIotas(chain.OriginatorAddress, solo.Saldo-solo.ChainDustThreshold-4-1-extraToken)
	env.AssertAddressIotas(userAddress, solo.Saldo-10)
}

func TestRequestToView(t *testing.T) { run2(t, testRequestToView) }
func testRequestToView(t *testing.T, w bool) {
	env, chain := setupChain(t, nil)
	cAID, extraToken := setupTestSandboxSC(t, chain, nil, w)
	user, userAddress, userAgentID := setupDeployer(t, chain)

	t.Logf("dump accounts 1:\n%s", chain.DumpAccounts())
	chain.AssertL2AccountIotas(chain.OriginatorAgentID, 0)
	chain.AssertL2AccountIotas(userAgentID, 0)
	chain.AssertL2AccountIotas(cAID, 1)
	chain.AssertCommonAccountIotas(3 + extraToken)
	chain.AssertL2TotalIotas(4 + extraToken)
	env.AssertAddressIotas(chain.OriginatorAddress, solo.Saldo-solo.ChainDustThreshold-4-extraToken)
	env.AssertAddressIotas(userAddress, solo.Saldo)

	// sending request to the view entry point should return an error and invoke fallback for tokens
	req := solo.NewCallParams(ScName, sbtestsc.FuncJustView.Name)
	_, err := chain.PostRequestSync(req.AddIotas(42), user)
	require.Error(t, err)

	t.Logf("dump accounts 2:\n%s", chain.DumpAccounts())
	chain.AssertL2AccountIotas(chain.OriginatorAgentID, 0)
	chain.AssertL2AccountIotas(userAgentID, 0)
	chain.AssertL2AccountIotas(cAID, 1)
	chain.AssertCommonAccountIotas(3 + extraToken)
	chain.AssertL2TotalIotas(4 + extraToken)
	env.AssertAddressIotas(chain.OriginatorAddress, solo.Saldo-solo.ChainDustThreshold-4-extraToken)
	env.AssertAddressIotas(userAddress, solo.Saldo)
}
