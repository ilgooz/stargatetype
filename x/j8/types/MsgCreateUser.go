package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgUser{}

func NewMsgUser(creator sdk.AccAddress, name string, email string) *MsgUser {
  return &MsgUser{
    Id: uuid.New().String(),
		Creator: creator,
    Name: name,
    Email: email,
	}
}

func (msg *MsgUser) Route() string {
  return RouterKey
}

func (msg *MsgUser) Type() string {
  return "CreateUser"
}

func (msg *MsgUser) GetSigners() []sdk.AccAddress {
  return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg *MsgUser) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgUser) ValidateBasic() error {
  if msg.Creator.Empty() {
    return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
  }
  return nil
}
