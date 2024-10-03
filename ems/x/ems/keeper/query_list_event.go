package keeper

import (
	"context"

	"ems/x/ems/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ListEvent(goCtx context.Context, req *types.QueryListEventRequest) (*types.QueryListEventResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.MemStoreKey))

	var events []types.Event
	pageRes, err := query.Paginate(store, req.Pagination, func(key []byte, value []byte) error {
		var event types.Event
		if err := k.cdc.Unmarshal(value, &event); err != nil {
			return err
		}

		events = append(events, event)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryListEventResponse{Event: events, Pagination: pageRes}, nil
}
