package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/alice/checkers/x/checkers/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type (
	Keeper struct {
		bank     types.BankKeeper
		cdc      codec.BinaryCodec
		storeKey sdk.StoreKey
		memKey   sdk.StoreKey
	}
)

func NewKeeper(
	bank types.BankKeeper,
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,

) *Keeper {
	return &Keeper{
		bank:     bank,
		cdc:      cdc,
		storeKey: storeKey,
		memKey:   memKey,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
