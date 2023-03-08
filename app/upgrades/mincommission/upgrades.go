package mincommission

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
)

func CreateV2UpgradeHandler(stakingKeeper *stakingkeeper.Keeper) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, _ upgradetypes.Plan, fromVM module.VersionMap) (module.VersionMap, error) {
		allValidators := stakingKeeper.GetAllValidators(ctx)
		for _, validator := range allValidators {
			// increase commission rate
			if validator.Commission.CommissionRates.Rate.LT(sdk.NewDecWithPrec(5,2)) {
				commission, err := stakingKeeper.UpdateValidatorCommission(ctx, validator, sdk.NewDecWithPrec(5, 2))
				if err != nil {
					return nil, err
				}

				// call the before-modification hook since we're about to update the commission
				stakingKeeper.BeforeValidatorModified(ctx, validator.GetOperator())

				validator.Commission = commission
			}

			// increase max commission rate
			if validator.Commission.CommissionRates.MaxRate.LT(sdk.NewDecWithPrec(5, 2)) {
				validator.Commission.MaxRate = sdk.NewDecWithPrec(5, 2)
			}

			stakingKeeper.SetValidator(ctx, validator)
		}

		return fromVM, nil
	}
}
