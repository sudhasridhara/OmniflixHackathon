package keeper

import (
	"context"

	"ems/x/ems/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateEvent(goCtx context.Context, msg *types.MsgCreateEvent) (*types.MsgCreateEventResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var event = types.Event{
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
	id := k.AppendEvent(
		ctx,
		event,
	)
	// TODO: Handling the message
	_ = ctx

	return &types.MsgCreateEventResponse{Id: id}, nil
}
