package keeper

import (
	"context"
	"fmt"

	"ems/x/ems/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) UpdateEvent(goCtx context.Context, msg *types.MsgUpdateEvent) (*types.MsgUpdateEventResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	event := types.Event{
		EventName:        msg.EventName,
		EventDescription: msg.EventDescription,
		StartDate:        msg.StartDate,
		EndDate:          msg.EndDate,
		Location:         msg.Location,
		NumFtTickets:     msg.NumFtTickets,
		Organizer:        msg.Organizer,
		Creator:          msg.Creator,
		TicketPrice:      msg.TicketPrice,
		TicketsLeft:      msg.NumFtTickets,
	}
	val, found := k.GetEventInfo(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}
	k.UpdateEventInfo(ctx, event)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgUpdateEventResponse{}, nil
}
