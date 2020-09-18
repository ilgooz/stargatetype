package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
    cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
    // this line is used by starport scaffolding
sdk "github.com/cosmos/cosmos-sdk/types"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
    // this line is used by starport scaffolding # 2
cdc.RegisterConcrete(MsgUser{}, "j8/CreateUser", nil)
} 

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
    // this line is used by starport scaffolding # 3
registry.RegisterImplementations((*sdk.Msg)(nil),
	&MsgUser{},
)
}

var (
	amino = codec.New()
	ModuleCdc = codec.NewAminoCodec(amino)
)
