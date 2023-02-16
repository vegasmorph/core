package keeper

import (
	"fmt"
	"github.com/classic-terra/core/x/treasury/types"

	core "github.com/classic-terra/core/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// BurnCoinsFromBurnAccount burn all coins from burn account
func (k Keeper) BurnCoinsFromBurnAccount(ctx sdk.Context) {
	burnAddress := k.accountKeeper.GetModuleAddress(types.BurnModuleName)
	fmt.Println(burnAddress.String())
	if coins := k.bankKeeper.GetAllBalances(ctx, burnAddress); !coins.IsZero() {
		fmt.Println(fmt.Sprintf("NORMAL BURN:BURNED AND RECORDED %s",coins))
		fmt.Println(fmt.Sprintf("NORMAL BURN:TOTAL SUPPLY BEFORE BURN %s",k.bankKeeper.GetSupply(ctx, core.MicroLunaDenom)))
		err := k.bankKeeper.BurnCoins(ctx, types.BurnModuleName, coins)
		fmt.Println(fmt.Sprintf("NORMAL BURN:TOTAL SUPPLY AFTER BURN %s",k.bankKeeper.GetSupply(ctx, core.MicroLunaDenom)))
		fmt.Println(fmt.Sprintf("NORMAL BURN:SEIGNIORAGE %s",k.PeekEpochSeigniorage(ctx)))
		if err != nil {
			panic(err)
		}
	}

	return
}
// BurnCoinsFromBurnNoRemintAccount burn all coins from burn account
func (k Keeper) BurnCoinsFromBurnNoRemintAccount(ctx sdk.Context) {
	burnAddress := k.accountKeeper.GetModuleAddress(types.BurnNoRemintModuleName)
	fmt.Println(burnAddress.String())

	if amount := k.bankKeeper.GetBalance(ctx, burnAddress, core.MicroLunaDenom); !amount.IsZero() {
		k.RecordEpochBNRProceeds(ctx, amount.Amount)
		fmt.Println(fmt.Sprintf("NO REMINT:BURNED AND RECORDED %s",amount.Amount))
		fmt.Println(fmt.Sprintf("NO REMINT:RECORDED THIS EPOCH %s",k.GetBNR(ctx)))
		fmt.Println(fmt.Sprintf("No REMINT:TOTAL SUPPLY BEFORE BURN %s",k.bankKeeper.GetSupply(ctx, core.MicroLunaDenom)))
	}

	if coins := k.bankKeeper.GetAllBalances(ctx, burnAddress); !coins.IsZero() {
		err := k.bankKeeper.BurnCoins(ctx, types.BurnNoRemintModuleName, coins)
		fmt.Println(fmt.Sprintf("No REMINT:TOTAL SUPPLY AFTER BURN %s",k.bankKeeper.GetSupply(ctx, core.MicroLunaDenom)))
		fmt.Println(fmt.Sprintf("NO REMINT:SEIGNIORAGE %s",k.PeekEpochSeigniorage(ctx)))
		if err != nil {
			panic(err)
		}
	}

	return
}
