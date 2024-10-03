package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgResaleTicket{}

func NewMsgResaleTicket(creator string, amount string, eventId uint64, purchaser string, lender string, ticketId uint64) *MsgResaleTicket {
	return &MsgResaleTicket{
		Creator:   creator,
		Amount:    amount,
		EventId:   eventId,
		Purchaser: purchaser,
		Lender:    lender,
		TicketId:  ticketId,
	}
}

func (msg *MsgResaleTicket) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
