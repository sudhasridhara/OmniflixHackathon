package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/ems module sentinel errors
var (
	ErrInvalidSigner     = sdkerrors.Register(ModuleName, 1100, "expected gov account as only signer for proposal message")
	ErrSample            = sdkerrors.Register(ModuleName, 1101, "sample error")
	ErrCoinParse         = sdkerrors.Register(ModuleName, 1102, "Cannot parse coins")
	ErrInsufficientFunds = sdkerrors.Register(ModuleName, 1103, "Insufficient Funds")
)
