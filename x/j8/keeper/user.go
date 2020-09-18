package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/j/j8/x/j8/types"
)

func (k Keeper) CreateUser(ctx sdk.Context, user types.MsgUser) {
	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshalBinaryBare(&user)
	store = prefix.NewStore(store, types.KeyPrefix(types.UserKey))
	store.Set(types.KeyPrefix(types.UserKey), b)
}

func (k Keeper) GetAllUser(ctx sdk.Context) (msgs []types.MsgUser) {
	store := ctx.KVStore(k.storeKey)
	store = prefix.NewStore(store, types.KeyPrefix(types.UserKey))
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.UserKey))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var msg types.MsgUser
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &msg)
		msgs = append(msgs, msg)
	}

	return
}
