package keeper

import (
	"context"

	"ems/x/ems/types"

	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) PurchaseTicket(goCtx context.Context, msg *types.MsgPurchaseTicket) (*types.MsgPurchaseTicketResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	event, _ := k.GetEventInfo(ctx, msg.EventId)

	ticket := types.Ticket{
		Eventid: msg.EventId,
		Creator: msg.Purchaser,
		Owner:   event.EventName,
	}
	// TODO: Handling the message
	lender, err := sdk.AccAddressFromBech32(msg.Purchaser)
	if err != nil {
		return nil, errors.Wrap(err, "invalid lender address")
	}

	borrower, err := sdk.AccAddressFromBech32(event.Creator)
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
	err = k.bankKeeper.SendCoins(ctx, lender, borrower, amount)
	if err != nil {
		return nil, errors.Wrap(err, "failed to send coins")
	}

	k.AppendTicket(ctx, ticket)

	event.TicketsLeft = event.TicketsLeft - 1
	k.UpdateEventInfo(ctx, event)

	return &types.MsgPurchaseTicketResponse{}, nil
}
