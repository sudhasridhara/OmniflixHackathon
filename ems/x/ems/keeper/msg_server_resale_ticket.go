package keeper

import (
	"context"

	"ems/x/ems/types"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) ResaleTicket(goCtx context.Context, msg *types.MsgResaleTicket) (*types.MsgResaleTicketResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	ticket, _ := k.GetTicketInfo(ctx, msg.TicketId)

	//event, _ := k.GetEventInfo(ctx, msg.EventId)

	// TODO: Handling the message
	purchaser, err := sdk.AccAddressFromBech32(msg.Purchaser)
	if err != nil {
		return nil, errors.Wrap(err, "invalid lender address")
	}

	lender, err := sdk.AccAddressFromBech32(msg.Lender)
	if err != nil {
		return nil, errors.Wrap(err, "invalid borrower address")
	}

	amount, err := sdk.ParseCoinsNormalized(msg.Amount)
	if err != nil {
		return nil, errors.Wrap(types.ErrCoinParse, "cannot parse coins in loan amount")
	}

	// Check lender's balance
	lenderBalance := k.bankKeeper.SpendableCoins(ctx, lender)
	if !lenderBalance.IsAllGTE(amount) {
		return nil, errors.Wrap(types.ErrInsufficientFunds, "lender does not have enough funds")
	}

	// Send coins from lender to borrower
	err = k.bankKeeper.SendCoins(ctx, purchaser, lender, amount)
	if err != nil {
		return nil, errors.Wrap(err, "failed to send coins")
	}

	ticket.Creator = msg.Purchaser
	k.UpdateTicketInfo(ctx, ticket)

	return &types.MsgResaleTicketResponse{}, nil
}
