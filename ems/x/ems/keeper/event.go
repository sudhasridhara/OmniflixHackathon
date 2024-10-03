package keeper

import (
	"encoding/binary"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"ems/x/ems/types"
)

func (k Keeper) AppendEvent(ctx sdk.Context, event types.Event) uint64 {
	count := k.GetEventCount(ctx)
	event.Id = count
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.MemStoreKey))
	appendedValue := k.cdc.MustMarshal(&event)
	store.Set(GetEventIDBytes(event.Id), appendedValue)
	k.SetEventCount(ctx, count+1)
	return count
}

func (k Keeper) GetEventCount(ctx sdk.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.MemCountKey)
	bz := store.Get(byteKey)
	if bz == nil {
		return 0
	}
	return binary.BigEndian.Uint64(bz)
}

func GetEventIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

func (k Keeper) SetEventCount(ctx sdk.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.MemCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

func (k Keeper) GetEventInfo(ctx sdk.Context, id uint64) (val types.Event, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.MemStoreKey))
	b := store.Get(GetEventIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) UpdateEventInfo(ctx sdk.Context, event types.Event) {
	logger := ctx.Logger()
	logger.Info("Attempting to update event", "event_id", event.Id)
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.MemStoreKey))
	val, found := k.GetEventInfo(ctx, event.Id)
	if !found {
		logger.Error("Event not found for update", "event_id", val.Id)
		//return fmt.Errorf("event with ID %d does not exist", event.Id)
	}
	b := k.cdc.MustMarshal(&event)
	store.Set(GetEventIDBytes(event.Id), b)
	logger.Info("Event updated successfully", "event_id", event.Id)
}

func (k Keeper) RemoveEvent(ctx sdk.Context, id uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.MemStoreKey))
	store.Delete(GetEventIDBytes(id))
}
