package keeper

import (
	"testing"
	"fmt"

	"github.com/classic-terra/core/x/treasury/types"
	"github.com/stretchr/testify/require"
)

func TestBurnCoinsFromBurnAccount(t *testing.T) {
	input := CreateTestInput(t)

	burnAddress := input.AccountKeeper.GetModuleAddress(types.BurnModuleName)
	coins := input.BankKeeper.GetAllBalances(input.Ctx, burnAddress)
	require.Equal(t, InitCoins, coins)

	input.TreasuryKeeper.BurnCoinsFromBurnAccount(input.Ctx)
	coins = input.BankKeeper.GetAllBalances(input.Ctx, burnAddress)
	require.True(t, coins.IsZero())
}

func TestBurnCoinsFromBurnNoRemintAccount(t *testing.T) {
	input := CreateTestInput(t)

	burnAddress := input.AccountKeeper.GetModuleAddress(types.BurnNoRemintModuleName)
	fmt.Println(burnAddress)
	coins := input.BankKeeper.GetAllBalances(input.Ctx, burnAddress)
	require.Equal(t, InitCoins, coins)

	input.TreasuryKeeper.BurnCoinsFromBurnNoRemintAccount(input.Ctx)
	coins = input.BankKeeper.GetAllBalances(input.Ctx, burnAddress)
	require.True(t, coins.IsZero())
}
