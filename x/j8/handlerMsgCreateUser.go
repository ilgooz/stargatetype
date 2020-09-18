package j8

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/j/j8/x/j8/types"
	"github.com/j/j8/x/j8/keeper"
)

func handleMsgCreateUser(ctx sdk.Context, k keeper.Keeper, user *types.MsgUser) (*sdk.Result, error) {
	k.CreateUser(ctx, *user)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
