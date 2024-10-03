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

func (k Keeper) ListTicket(goCtx context.Context, req *types.QueryListTicketRequest) (*types.QueryListTicketResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.TicketStoreKey))

	var tickets []types.Ticket
	pageRes, err := query.Paginate(store, req.Pagination, func(key []byte, value []byte) error {
		var ticket types.Ticket
		if err := k.cdc.Unmarshal(value, &ticket); err != nil {
			return err
		}

		tickets = append(tickets, ticket)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &types.QueryListTicketResponse{Ticket: tickets, Pagination: pageRes}, nil
}
