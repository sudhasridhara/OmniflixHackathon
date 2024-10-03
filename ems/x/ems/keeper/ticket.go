package keeper

import (
	"encoding/binary"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"ems/x/ems/types"
)

func (k Keeper) AppendTicket(ctx sdk.Context, ticket types.Ticket) uint64 {
	count := k.GetTicketCount(ctx)
	ticket.Id = count
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.TicketStoreKey))
	appendedValue := k.cdc.MustMarshal(&ticket)
	store.Set(GetTicketIDBytes(ticket.Id), appendedValue)
	k.SetTicketCount(ctx, count+1)
	return count
}

func (k Keeper) GetTicketCount(ctx sdk.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.TicketCountKey)
	bz := store.Get(byteKey)
	if bz == nil {
		return 0
	}
	return binary.BigEndian.Uint64(bz)
}

func GetTicketIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

func (k Keeper) SetTicketCount(ctx sdk.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.TicketCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

func (k Keeper) GetTicketInfo(ctx sdk.Context, id uint64) (val types.Ticket, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.TicketStoreKey))
	b := store.Get(GetTicketIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) UpdateTicketInfo(ctx sdk.Context, ticket types.Ticket) {

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.TicketStoreKey))

	b := k.cdc.MustMarshal(&ticket)
	store.Set(GetEventIDBytes(ticket.Id), b)
}

func (k Keeper) RemoveTicket(ctx sdk.Context, id uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.TicketStoreKey))
	store.Delete(GetTicketIDBytes(id))
}
