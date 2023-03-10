package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgConfirm = "confirm"

var _ sdk.Msg = &MsgConfirm{}

func NewMsgConfirm(creator string, escrowIndex string, winner string) *MsgConfirm {
	return &MsgConfirm{
		Creator:     creator,
		EscrowIndex: escrowIndex,
		Winner:      winner,
	}
}

func (msg *MsgConfirm) Route() string {
	return RouterKey
}

func (msg *MsgConfirm) Type() string {
	return TypeMsgConfirm
}

func (msg *MsgConfirm) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgConfirm) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgConfirm) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
