package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgNewEscrow = "new_escrow"

var _ sdk.Msg = &MsgNewEscrow{}

func NewMsgNewEscrow(creator string, strategy string, instigator string, instigatorWager string, rider string, riderWager string) *MsgNewEscrow {
	return &MsgNewEscrow{
		Creator:         creator,
		Strategy:        strategy,
		Instigator:      instigator,
		InstigatorWager: instigatorWager,
		Rider:           rider,
		RiderWager:      riderWager,
	}
}

func (msg *MsgNewEscrow) Route() string {
	return RouterKey
}

func (msg *MsgNewEscrow) Type() string {
	return TypeMsgNewEscrow
}

func (msg *MsgNewEscrow) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgNewEscrow) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgNewEscrow) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
