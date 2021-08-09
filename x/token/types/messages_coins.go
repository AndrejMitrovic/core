package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateCoins{}

func NewMsgCreateCoins(
	creator string,
	user string,
	amount string,

) *MsgCreateCoins {
	return &MsgCreateCoins{
		Creator: creator,
		User:    user,
		Amount:  amount,
	}
}

func (msg *MsgCreateCoins) Route() string {
	return RouterKey
}

func (msg *MsgCreateCoins) Type() string {
	return "CreateCoins"
}

func (msg *MsgCreateCoins) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateCoins) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateCoins) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateCoins{}

func NewMsgUpdateCoins(
	creator string,
	user string,
	amount string,

) *MsgUpdateCoins {
	return &MsgUpdateCoins{
		Creator: creator,
		User:    user,
		Amount:  amount,
	}
}

func (msg *MsgUpdateCoins) Route() string {
	return RouterKey
}

func (msg *MsgUpdateCoins) Type() string {
	return "UpdateCoins"
}

func (msg *MsgUpdateCoins) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateCoins) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateCoins) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteCoins{}

func NewMsgDeleteCoins(
	creator string,
	user string,

) *MsgDeleteCoins {
	return &MsgDeleteCoins{
		Creator: creator,
		User:    user,
	}
}
func (msg *MsgDeleteCoins) Route() string {
	return RouterKey
}

func (msg *MsgDeleteCoins) Type() string {
	return "DeleteCoins"
}

func (msg *MsgDeleteCoins) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteCoins) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteCoins) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
