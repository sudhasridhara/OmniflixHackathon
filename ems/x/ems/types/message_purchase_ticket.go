package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgPurchaseTicket{}

func NewMsgPurchaseTicket(creator string, amount string, eventId uint64, purchaser string) *MsgPurchaseTicket {
	return &MsgPurchaseTicket{
		Creator:   creator,
		Amount:    amount,
		EventId:   eventId,
		Purchaser: purchaser,
	}
}

func (msg *MsgPurchaseTicket) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
