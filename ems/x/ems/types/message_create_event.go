package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateEvent{}

func NewMsgCreateEvent(creator string, eventName string, eventDescription string, startDate string, endDate string, location string, numFtTickets uint64, organizer string, ticketPrice uint64, ticketsLeft uint64) *MsgCreateEvent {
	return &MsgCreateEvent{
		Creator:          creator,
		EventName:        eventName,
		EventDescription: eventDescription,
		StartDate:        startDate,
		EndDate:          endDate,
		Location:         location,
		NumFtTickets:     numFtTickets,
		Organizer:        organizer,
		TicketPrice:      ticketPrice,
		TicketsLeft:      ticketsLeft,
	}
}

func (msg *MsgCreateEvent) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
