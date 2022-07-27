package tracking

import (
	"github.com/archway-network/archway/x/tracking/keeper"
	"github.com/archway-network/archway/x/tracking/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
	"time"
)

// EndBlocker modifies tracked transactions using tracked contract operations.
func EndBlocker(ctx sdk.Context, k keeper.Keeper) []abci.ValidatorUpdate {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyEndBlocker)

	k.FinalizeBlockTxTracking(ctx)

	return []abci.ValidatorUpdate{}
}
