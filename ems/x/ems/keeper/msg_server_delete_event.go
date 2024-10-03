package keeper

import (
	"context"
	"fmt"

	"ems/x/ems/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) DeleteEvent(goCtx context.Context, msg *types.MsgDeleteEvent) (*types.MsgDeleteEventResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetEventInfo(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}
	k.RemoveEvent(ctx, msg.Id)

	return &types.MsgDeleteEventResponse{}, nil
}
