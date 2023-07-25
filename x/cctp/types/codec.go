package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func init() {
	RegisterLegacyAminoCodec(amino)
	amino.Seal()
}

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgAddTokenMessenger{}, "cctp/AddTokenMessenger", nil)
	cdc.RegisterConcrete(&MsgDepositForBurnWithCaller{}, "cctp/DepositForBurnWithCaller", nil)
	cdc.RegisterConcrete(&MsgDepositForBurn{}, "cctp/DepositForBurn", nil)
	cdc.RegisterConcrete(&MsgDisableAttester{}, "cctp/DisableAttester", nil)
	cdc.RegisterConcrete(&MsgEnableAttester{}, "cctp/EnableAttester", nil)
	cdc.RegisterConcrete(&MsgLinkTokenPair{}, "cctp/LinkTokenPair", nil)
	cdc.RegisterConcrete(&MsgPauseBurningAndMinting{}, "cctp/PauseBurningAndMinting", nil)
	cdc.RegisterConcrete(&MsgPauseSendingAndReceivingMessages{}, "cctp/PauseSendingAndReceivingMessages", nil)
	cdc.RegisterConcrete(&MsgReceiveMessage{}, "cctp/ReceiveMessage", nil)
	cdc.RegisterConcrete(&MsgRemoveTokenMessenger{}, "cctp/RemoveTokenMessenger", nil)
	cdc.RegisterConcrete(&MsgReplaceDepositForBurn{}, "cctp/ReplaceDepositForBurn", nil)
	cdc.RegisterConcrete(&MsgReplaceMessage{}, "cctp/ReplaceMessage", nil)
	cdc.RegisterConcrete(&MsgSendMessageWithCaller{}, "cctp/SendMessageWithCaller", nil)
	cdc.RegisterConcrete(&MsgSendMessage{}, "cctp/SendMessage", nil)
	cdc.RegisterConcrete(&MsgUnlinkTokenPair{}, "cctp/UnlinkTokenPair", nil)
	cdc.RegisterConcrete(&MsgUnpauseBurningAndMinting{}, "cctp/UnpauseBurningAndMinting", nil)
	cdc.RegisterConcrete(&MsgUnpauseSendingAndReceivingMessages{}, "cctp/UnpauseSendingAndReceivingMessages", nil)
	cdc.RegisterConcrete(&MsgUpdateAuthority{}, "cctp/UpdateAuthority", nil)
	cdc.RegisterConcrete(&MsgUpdateMaxMessageBodySize{}, "cctp/UpdateMaxMessageBodySize", nil)
	cdc.RegisterConcrete(&MsgUpdatePerMessageBurnLimit{}, "cctp/UpdatePerMessageBurnLimit", nil)
	cdc.RegisterConcrete(&MsgUpdateSignatureThreshold{}, "cctp/UpdateSignatureThreshold", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgAddTokenMessenger{},
		&MsgDepositForBurnWithCaller{},
		&MsgDepositForBurn{},
		&MsgDisableAttester{},
		&MsgEnableAttester{},
		&MsgLinkTokenPair{},
		&MsgPauseBurningAndMinting{},
		&MsgPauseSendingAndReceivingMessages{},
		&MsgReceiveMessage{},
		&MsgRemoveTokenMessenger{},
		&MsgReplaceDepositForBurn{},
		&MsgReplaceMessage{},
		&MsgSendMessageWithCaller{},
		&MsgSendMessage{},
		&MsgUnlinkTokenPair{},
		&MsgUnpauseBurningAndMinting{},
		&MsgUnpauseSendingAndReceivingMessages{},
		&MsgUpdateAuthority{},
		&MsgUpdateMaxMessageBodySize{},
		&MsgUpdatePerMessageBurnLimit{},
		&MsgUpdateSignatureThreshold{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)