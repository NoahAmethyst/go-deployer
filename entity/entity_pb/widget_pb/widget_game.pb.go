// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.17.3
// source: widget_game.proto

package widget_pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GameBattleRoyalCallbackType int32

const (
	GameBattleRoyalCallbackType_GAME_BATTLE_ROYAL_CALLBACK_TYPE_NONE       GameBattleRoyalCallbackType = 0
	GameBattleRoyalCallbackType_GAME_BATTLE_ROYAL_CALLBACK_TYPE_INIT_GAME  GameBattleRoyalCallbackType = 1001
	GameBattleRoyalCallbackType_GAME_BATTLE_ROYAL_CALLBACK_TYPE_START_GAME GameBattleRoyalCallbackType = 1002
	GameBattleRoyalCallbackType_GAME_BATTLE_ROYAL_CALLBACK_TYPE_SIGN_UP    GameBattleRoyalCallbackType = 1003
	GameBattleRoyalCallbackType_GAME_BATTLE_ROYAL_CALLBACK_TYPE_BUY_BUFF   GameBattleRoyalCallbackType = 1004
	GameBattleRoyalCallbackType_GAME_BATTLE_ROYAL_CALLBACK_TYPE_END_GAME   GameBattleRoyalCallbackType = 1005
)

// Enum value maps for GameBattleRoyalCallbackType.
var (
	GameBattleRoyalCallbackType_name = map[int32]string{
		0:    "GAME_BATTLE_ROYAL_CALLBACK_TYPE_NONE",
		1001: "GAME_BATTLE_ROYAL_CALLBACK_TYPE_INIT_GAME",
		1002: "GAME_BATTLE_ROYAL_CALLBACK_TYPE_START_GAME",
		1003: "GAME_BATTLE_ROYAL_CALLBACK_TYPE_SIGN_UP",
		1004: "GAME_BATTLE_ROYAL_CALLBACK_TYPE_BUY_BUFF",
		1005: "GAME_BATTLE_ROYAL_CALLBACK_TYPE_END_GAME",
	}
	GameBattleRoyalCallbackType_value = map[string]int32{
		"GAME_BATTLE_ROYAL_CALLBACK_TYPE_NONE":       0,
		"GAME_BATTLE_ROYAL_CALLBACK_TYPE_INIT_GAME":  1001,
		"GAME_BATTLE_ROYAL_CALLBACK_TYPE_START_GAME": 1002,
		"GAME_BATTLE_ROYAL_CALLBACK_TYPE_SIGN_UP":    1003,
		"GAME_BATTLE_ROYAL_CALLBACK_TYPE_BUY_BUFF":   1004,
		"GAME_BATTLE_ROYAL_CALLBACK_TYPE_END_GAME":   1005,
	}
)

func (x GameBattleRoyalCallbackType) Enum() *GameBattleRoyalCallbackType {
	p := new(GameBattleRoyalCallbackType)
	*p = x
	return p
}

func (x GameBattleRoyalCallbackType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GameBattleRoyalCallbackType) Descriptor() protoreflect.EnumDescriptor {
	return file_widget_game_proto_enumTypes[0].Descriptor()
}

func (GameBattleRoyalCallbackType) Type() protoreflect.EnumType {
	return &file_widget_game_proto_enumTypes[0]
}

func (x GameBattleRoyalCallbackType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GameBattleRoyalCallbackType.Descriptor instead.
func (GameBattleRoyalCallbackType) EnumDescriptor() ([]byte, []int) {
	return file_widget_game_proto_rawDescGZIP(), []int{0}
}

type GameBattleRoyalCallbackReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OpenId       int64                       `protobuf:"varint,1,opt,name=open_id,json=openId,proto3" json:"open_id,omitempty"`
	GroupId      string                      `protobuf:"bytes,4,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	AppId        string                      `protobuf:"bytes,5,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	BotType      int32                       `protobuf:"varint,6,opt,name=bot_type,json=botType,proto3" json:"bot_type,omitempty"`
	CustomId     string                      `protobuf:"bytes,8,opt,name=custom_id,json=customId,proto3" json:"custom_id,omitempty"`
	Data         string                      `protobuf:"bytes,9,opt,name=data,proto3" json:"data,omitempty"`
	CallbackType GameBattleRoyalCallbackType `protobuf:"varint,7,opt,name=callback_type,json=callbackType,proto3,enum=proto.GameBattleRoyalCallbackType" json:"callback_type,omitempty"`
	EstimateFee  bool                        `protobuf:"varint,10,opt,name=estimate_fee,json=estimateFee,proto3" json:"estimate_fee,omitempty"`
}

func (x *GameBattleRoyalCallbackReq) Reset() {
	*x = GameBattleRoyalCallbackReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_widget_game_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameBattleRoyalCallbackReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameBattleRoyalCallbackReq) ProtoMessage() {}

func (x *GameBattleRoyalCallbackReq) ProtoReflect() protoreflect.Message {
	mi := &file_widget_game_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameBattleRoyalCallbackReq.ProtoReflect.Descriptor instead.
func (*GameBattleRoyalCallbackReq) Descriptor() ([]byte, []int) {
	return file_widget_game_proto_rawDescGZIP(), []int{0}
}

func (x *GameBattleRoyalCallbackReq) GetOpenId() int64 {
	if x != nil {
		return x.OpenId
	}
	return 0
}

func (x *GameBattleRoyalCallbackReq) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *GameBattleRoyalCallbackReq) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

func (x *GameBattleRoyalCallbackReq) GetBotType() int32 {
	if x != nil {
		return x.BotType
	}
	return 0
}

func (x *GameBattleRoyalCallbackReq) GetCustomId() string {
	if x != nil {
		return x.CustomId
	}
	return ""
}

func (x *GameBattleRoyalCallbackReq) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *GameBattleRoyalCallbackReq) GetCallbackType() GameBattleRoyalCallbackType {
	if x != nil {
		return x.CallbackType
	}
	return GameBattleRoyalCallbackType_GAME_BATTLE_ROYAL_CALLBACK_TYPE_NONE
}

func (x *GameBattleRoyalCallbackReq) GetEstimateFee() bool {
	if x != nil {
		return x.EstimateFee
	}
	return false
}

type GameBattleRoyaleCheckFeeData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GasLimit          uint64 `protobuf:"varint,1,opt,name=gas_limit,json=gasLimit,proto3" json:"gas_limit,omitempty"`
	GasPrice          uint64 `protobuf:"varint,2,opt,name=gas_price,json=gasPrice,proto3" json:"gas_price,omitempty"`
	GasDecimalStr     string `protobuf:"bytes,3,opt,name=gas_decimal_str,json=gasDecimalStr,proto3" json:"gas_decimal_str,omitempty"`
	GasSymbol         string `protobuf:"bytes,4,opt,name=gas_symbol,json=gasSymbol,proto3" json:"gas_symbol,omitempty"`
	DecimalBalance    string `protobuf:"bytes,8,opt,name=decimal_balance,json=decimalBalance,proto3" json:"decimal_balance,omitempty"`
	FeeStr            string `protobuf:"bytes,5,opt,name=fee_str,json=feeStr,proto3" json:"fee_str,omitempty"`
	FeeSymbol         string `protobuf:"bytes,6,opt,name=fee_symbol,json=feeSymbol,proto3" json:"fee_symbol,omitempty"`
	FeeDecimalStr     string `protobuf:"bytes,7,opt,name=fee_decimal_str,json=feeDecimalStr,proto3" json:"fee_decimal_str,omitempty"`
	FeeDecimalBalance string `protobuf:"bytes,9,opt,name=fee_decimal_balance,json=feeDecimalBalance,proto3" json:"fee_decimal_balance,omitempty"`
}

func (x *GameBattleRoyaleCheckFeeData) Reset() {
	*x = GameBattleRoyaleCheckFeeData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_widget_game_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameBattleRoyaleCheckFeeData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameBattleRoyaleCheckFeeData) ProtoMessage() {}

func (x *GameBattleRoyaleCheckFeeData) ProtoReflect() protoreflect.Message {
	mi := &file_widget_game_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameBattleRoyaleCheckFeeData.ProtoReflect.Descriptor instead.
func (*GameBattleRoyaleCheckFeeData) Descriptor() ([]byte, []int) {
	return file_widget_game_proto_rawDescGZIP(), []int{1}
}

func (x *GameBattleRoyaleCheckFeeData) GetGasLimit() uint64 {
	if x != nil {
		return x.GasLimit
	}
	return 0
}

func (x *GameBattleRoyaleCheckFeeData) GetGasPrice() uint64 {
	if x != nil {
		return x.GasPrice
	}
	return 0
}

func (x *GameBattleRoyaleCheckFeeData) GetGasDecimalStr() string {
	if x != nil {
		return x.GasDecimalStr
	}
	return ""
}

func (x *GameBattleRoyaleCheckFeeData) GetGasSymbol() string {
	if x != nil {
		return x.GasSymbol
	}
	return ""
}

func (x *GameBattleRoyaleCheckFeeData) GetDecimalBalance() string {
	if x != nil {
		return x.DecimalBalance
	}
	return ""
}

func (x *GameBattleRoyaleCheckFeeData) GetFeeStr() string {
	if x != nil {
		return x.FeeStr
	}
	return ""
}

func (x *GameBattleRoyaleCheckFeeData) GetFeeSymbol() string {
	if x != nil {
		return x.FeeSymbol
	}
	return ""
}

func (x *GameBattleRoyaleCheckFeeData) GetFeeDecimalStr() string {
	if x != nil {
		return x.FeeDecimalStr
	}
	return ""
}

func (x *GameBattleRoyaleCheckFeeData) GetFeeDecimalBalance() string {
	if x != nil {
		return x.FeeDecimalBalance
	}
	return ""
}

type BattleRoyaleCallbackData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TxHash            string                        `protobuf:"bytes,1,opt,name=tx_hash,json=txHash,proto3" json:"tx_hash,omitempty"`
	ChainId           uint64                        `protobuf:"varint,2,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	Fee               *GameBattleRoyaleCheckFeeData `protobuf:"bytes,3,opt,name=fee,proto3" json:"fee,omitempty"`
	FeeText           string                        `protobuf:"bytes,4,opt,name=fee_text,json=feeText,proto3" json:"fee_text,omitempty"`
	CreatorProportion uint64                        `protobuf:"varint,5,opt,name=creator_proportion,json=creatorProportion,proto3" json:"creator_proportion,omitempty"`
}

func (x *BattleRoyaleCallbackData) Reset() {
	*x = BattleRoyaleCallbackData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_widget_game_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BattleRoyaleCallbackData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BattleRoyaleCallbackData) ProtoMessage() {}

func (x *BattleRoyaleCallbackData) ProtoReflect() protoreflect.Message {
	mi := &file_widget_game_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BattleRoyaleCallbackData.ProtoReflect.Descriptor instead.
func (*BattleRoyaleCallbackData) Descriptor() ([]byte, []int) {
	return file_widget_game_proto_rawDescGZIP(), []int{2}
}

func (x *BattleRoyaleCallbackData) GetTxHash() string {
	if x != nil {
		return x.TxHash
	}
	return ""
}

func (x *BattleRoyaleCallbackData) GetChainId() uint64 {
	if x != nil {
		return x.ChainId
	}
	return 0
}

func (x *BattleRoyaleCallbackData) GetFee() *GameBattleRoyaleCheckFeeData {
	if x != nil {
		return x.Fee
	}
	return nil
}

func (x *BattleRoyaleCallbackData) GetFeeText() string {
	if x != nil {
		return x.FeeText
	}
	return ""
}

func (x *BattleRoyaleCallbackData) GetCreatorProportion() uint64 {
	if x != nil {
		return x.CreatorProportion
	}
	return 0
}

type BattleRoyaleCallbackResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error *WidgetErrorMessage       `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	Data  *BattleRoyaleCallbackData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *BattleRoyaleCallbackResp) Reset() {
	*x = BattleRoyaleCallbackResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_widget_game_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BattleRoyaleCallbackResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BattleRoyaleCallbackResp) ProtoMessage() {}

func (x *BattleRoyaleCallbackResp) ProtoReflect() protoreflect.Message {
	mi := &file_widget_game_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BattleRoyaleCallbackResp.ProtoReflect.Descriptor instead.
func (*BattleRoyaleCallbackResp) Descriptor() ([]byte, []int) {
	return file_widget_game_proto_rawDescGZIP(), []int{3}
}

func (x *BattleRoyaleCallbackResp) GetError() *WidgetErrorMessage {
	if x != nil {
		return x.Error
	}
	return nil
}

func (x *BattleRoyaleCallbackResp) GetData() *BattleRoyaleCallbackData {
	if x != nil {
		return x.Data
	}
	return nil
}

type BattleRoyale struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppId            string `protobuf:"bytes,1,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
	GameAddr         string `protobuf:"bytes,2,opt,name=game_addr,json=gameAddr,proto3" json:"game_addr,omitempty"`
	GameDataAddr     string `protobuf:"bytes,3,opt,name=game_data_addr,json=gameDataAddr,proto3" json:"game_data_addr,omitempty"`
	GameReaderAddr   string `protobuf:"bytes,4,opt,name=game_reader_addr,json=gameReaderAddr,proto3" json:"game_reader_addr,omitempty"`
	AdminAddr        string `protobuf:"bytes,5,opt,name=admin_addr,json=adminAddr,proto3" json:"admin_addr,omitempty"`
	PinCode          string `protobuf:"bytes,6,opt,name=pin_code,json=pinCode,proto3" json:"pin_code,omitempty"`
	PrepaidEndGasFee string `protobuf:"bytes,7,opt,name=prepaid_end_gas_fee,json=prepaidEndGasFee,proto3" json:"prepaid_end_gas_fee,omitempty"`
	ChainId          uint64 `protobuf:"varint,8,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
}

func (x *BattleRoyale) Reset() {
	*x = BattleRoyale{}
	if protoimpl.UnsafeEnabled {
		mi := &file_widget_game_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BattleRoyale) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BattleRoyale) ProtoMessage() {}

func (x *BattleRoyale) ProtoReflect() protoreflect.Message {
	mi := &file_widget_game_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BattleRoyale.ProtoReflect.Descriptor instead.
func (*BattleRoyale) Descriptor() ([]byte, []int) {
	return file_widget_game_proto_rawDescGZIP(), []int{4}
}

func (x *BattleRoyale) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

func (x *BattleRoyale) GetGameAddr() string {
	if x != nil {
		return x.GameAddr
	}
	return ""
}

func (x *BattleRoyale) GetGameDataAddr() string {
	if x != nil {
		return x.GameDataAddr
	}
	return ""
}

func (x *BattleRoyale) GetGameReaderAddr() string {
	if x != nil {
		return x.GameReaderAddr
	}
	return ""
}

func (x *BattleRoyale) GetAdminAddr() string {
	if x != nil {
		return x.AdminAddr
	}
	return ""
}

func (x *BattleRoyale) GetPinCode() string {
	if x != nil {
		return x.PinCode
	}
	return ""
}

func (x *BattleRoyale) GetPrepaidEndGasFee() string {
	if x != nil {
		return x.PrepaidEndGasFee
	}
	return ""
}

func (x *BattleRoyale) GetChainId() uint64 {
	if x != nil {
		return x.ChainId
	}
	return 0
}

var File_widget_game_proto protoreflect.FileDescriptor

var file_widget_game_proto_rawDesc = []byte{
	0x0a, 0x11, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x5f, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x77, 0x69, 0x64, 0x67,
	0x65, 0x74, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9f,
	0x02, 0x0a, 0x1a, 0x47, 0x61, 0x6d, 0x65, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x52, 0x6f, 0x79,
	0x61, 0x6c, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a,
	0x07, 0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x6f, 0x70, 0x65, 0x6e, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49,
	0x64, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x6f, 0x74, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x62, 0x6f, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x69, 0x64,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x47, 0x0a, 0x0d, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x52, 0x6f,
	0x79, 0x61, 0x6c, 0x43, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x0c, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a,
	0x0c, 0x65, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x5f, 0x66, 0x65, 0x65, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0b, 0x65, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x46, 0x65, 0x65,
	0x22, 0xd8, 0x02, 0x0a, 0x1c, 0x47, 0x61, 0x6d, 0x65, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x52,
	0x6f, 0x79, 0x61, 0x6c, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x46, 0x65, 0x65, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x61, 0x73, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x67, 0x61, 0x73, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x67, 0x61, 0x73, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x08, 0x67, 0x61, 0x73, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x26, 0x0a, 0x0f, 0x67,
	0x61, 0x73, 0x5f, 0x64, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x5f, 0x73, 0x74, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x67, 0x61, 0x73, 0x44, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c,
	0x53, 0x74, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x67, 0x61, 0x73, 0x5f, 0x73, 0x79, 0x6d, 0x62, 0x6f,
	0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x61, 0x73, 0x53, 0x79, 0x6d, 0x62,
	0x6f, 0x6c, 0x12, 0x27, 0x0a, 0x0f, 0x64, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x5f, 0x62, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x64, 0x65, 0x63,
	0x69, 0x6d, 0x61, 0x6c, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x66,
	0x65, 0x65, 0x5f, 0x73, 0x74, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x65,
	0x65, 0x53, 0x74, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x65, 0x65, 0x5f, 0x73, 0x79, 0x6d, 0x62,
	0x6f, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x65, 0x65, 0x53, 0x79, 0x6d,
	0x62, 0x6f, 0x6c, 0x12, 0x26, 0x0a, 0x0f, 0x66, 0x65, 0x65, 0x5f, 0x64, 0x65, 0x63, 0x69, 0x6d,
	0x61, 0x6c, 0x5f, 0x73, 0x74, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x66, 0x65,
	0x65, 0x44, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x53, 0x74, 0x72, 0x12, 0x2e, 0x0a, 0x13, 0x66,
	0x65, 0x65, 0x5f, 0x64, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x5f, 0x62, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x66, 0x65, 0x65, 0x44, 0x65, 0x63,
	0x69, 0x6d, 0x61, 0x6c, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x22, 0xcf, 0x01, 0x0a, 0x18,
	0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x52, 0x6f, 0x79, 0x61, 0x6c, 0x65, 0x43, 0x61, 0x6c, 0x6c,
	0x62, 0x61, 0x63, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x78, 0x5f, 0x68,
	0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x78, 0x48, 0x61, 0x73,
	0x68, 0x12, 0x19, 0x0a, 0x08, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x07, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x35, 0x0a, 0x03,
	0x66, 0x65, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x52, 0x6f, 0x79, 0x61,
	0x6c, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x46, 0x65, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x03,
	0x66, 0x65, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x66, 0x65, 0x65, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66, 0x65, 0x65, 0x54, 0x65, 0x78, 0x74, 0x12, 0x2d,
	0x0a, 0x12, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x72,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x11, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x80, 0x01,
	0x0a, 0x18, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x52, 0x6f, 0x79, 0x61, 0x6c, 0x65, 0x43, 0x61,
	0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2f, 0x0a, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x57, 0x69, 0x64, 0x67, 0x65, 0x74, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x33, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x52, 0x6f, 0x79, 0x61, 0x6c, 0x65, 0x43, 0x61,
	0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x22, 0x96, 0x02, 0x0a, 0x0c, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x52, 0x6f, 0x79, 0x61, 0x6c,
	0x65, 0x12, 0x15, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x61, 0x6d, 0x65,
	0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x67, 0x61, 0x6d,
	0x65, 0x41, 0x64, 0x64, 0x72, 0x12, 0x24, 0x0a, 0x0e, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x64, 0x61,
	0x74, 0x61, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x67,
	0x61, 0x6d, 0x65, 0x44, 0x61, 0x74, 0x61, 0x41, 0x64, 0x64, 0x72, 0x12, 0x28, 0x0a, 0x10, 0x67,
	0x61, 0x6d, 0x65, 0x5f, 0x72, 0x65, 0x61, 0x64, 0x65, 0x72, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x67, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x41, 0x64, 0x64, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x61,
	0x64, 0x64, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x41, 0x64, 0x64, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x69, 0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x69, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x2d, 0x0a, 0x13, 0x70, 0x72, 0x65, 0x70, 0x61, 0x69, 0x64, 0x5f, 0x65, 0x6e, 0x64, 0x5f, 0x67,
	0x61, 0x73, 0x5f, 0x66, 0x65, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x70, 0x72,
	0x65, 0x70, 0x61, 0x69, 0x64, 0x45, 0x6e, 0x64, 0x47, 0x61, 0x73, 0x46, 0x65, 0x65, 0x12, 0x19,
	0x0a, 0x08, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x07, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x64, 0x2a, 0xb4, 0x02, 0x0a, 0x1b, 0x47, 0x61,
	0x6d, 0x65, 0x42, 0x61, 0x74, 0x74, 0x6c, 0x65, 0x52, 0x6f, 0x79, 0x61, 0x6c, 0x43, 0x61, 0x6c,
	0x6c, 0x62, 0x61, 0x63, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x28, 0x0a, 0x24, 0x47, 0x41, 0x4d,
	0x45, 0x5f, 0x42, 0x41, 0x54, 0x54, 0x4c, 0x45, 0x5f, 0x52, 0x4f, 0x59, 0x41, 0x4c, 0x5f, 0x43,
	0x41, 0x4c, 0x4c, 0x42, 0x41, 0x43, 0x4b, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x4f, 0x4e,
	0x45, 0x10, 0x00, 0x12, 0x2e, 0x0a, 0x29, 0x47, 0x41, 0x4d, 0x45, 0x5f, 0x42, 0x41, 0x54, 0x54,
	0x4c, 0x45, 0x5f, 0x52, 0x4f, 0x59, 0x41, 0x4c, 0x5f, 0x43, 0x41, 0x4c, 0x4c, 0x42, 0x41, 0x43,
	0x4b, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4e, 0x49, 0x54, 0x5f, 0x47, 0x41, 0x4d, 0x45,
	0x10, 0xe9, 0x07, 0x12, 0x2f, 0x0a, 0x2a, 0x47, 0x41, 0x4d, 0x45, 0x5f, 0x42, 0x41, 0x54, 0x54,
	0x4c, 0x45, 0x5f, 0x52, 0x4f, 0x59, 0x41, 0x4c, 0x5f, 0x43, 0x41, 0x4c, 0x4c, 0x42, 0x41, 0x43,
	0x4b, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x5f, 0x47, 0x41, 0x4d,
	0x45, 0x10, 0xea, 0x07, 0x12, 0x2c, 0x0a, 0x27, 0x47, 0x41, 0x4d, 0x45, 0x5f, 0x42, 0x41, 0x54,
	0x54, 0x4c, 0x45, 0x5f, 0x52, 0x4f, 0x59, 0x41, 0x4c, 0x5f, 0x43, 0x41, 0x4c, 0x4c, 0x42, 0x41,
	0x43, 0x4b, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x49, 0x47, 0x4e, 0x5f, 0x55, 0x50, 0x10,
	0xeb, 0x07, 0x12, 0x2d, 0x0a, 0x28, 0x47, 0x41, 0x4d, 0x45, 0x5f, 0x42, 0x41, 0x54, 0x54, 0x4c,
	0x45, 0x5f, 0x52, 0x4f, 0x59, 0x41, 0x4c, 0x5f, 0x43, 0x41, 0x4c, 0x4c, 0x42, 0x41, 0x43, 0x4b,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x42, 0x55, 0x59, 0x5f, 0x42, 0x55, 0x46, 0x46, 0x10, 0xec,
	0x07, 0x12, 0x2d, 0x0a, 0x28, 0x47, 0x41, 0x4d, 0x45, 0x5f, 0x42, 0x41, 0x54, 0x54, 0x4c, 0x45,
	0x5f, 0x52, 0x4f, 0x59, 0x41, 0x4c, 0x5f, 0x43, 0x41, 0x4c, 0x4c, 0x42, 0x41, 0x43, 0x4b, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x45, 0x4e, 0x44, 0x5f, 0x47, 0x41, 0x4d, 0x45, 0x10, 0xed, 0x07,
	0x42, 0x16, 0x5a, 0x14, 0x77, 0x69, 0x64, 0x67, 0x65, 0x74, 0x5f, 0x70, 0x62, 0x2f, 0x3b, 0x77,
	0x69, 0x64, 0x67, 0x65, 0x74, 0x5f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_widget_game_proto_rawDescOnce sync.Once
	file_widget_game_proto_rawDescData = file_widget_game_proto_rawDesc
)

func file_widget_game_proto_rawDescGZIP() []byte {
	file_widget_game_proto_rawDescOnce.Do(func() {
		file_widget_game_proto_rawDescData = protoimpl.X.CompressGZIP(file_widget_game_proto_rawDescData)
	})
	return file_widget_game_proto_rawDescData
}

var file_widget_game_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_widget_game_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_widget_game_proto_goTypes = []interface{}{
	(GameBattleRoyalCallbackType)(0),     // 0: proto.GameBattleRoyalCallbackType
	(*GameBattleRoyalCallbackReq)(nil),   // 1: proto.GameBattleRoyalCallbackReq
	(*GameBattleRoyaleCheckFeeData)(nil), // 2: proto.GameBattleRoyaleCheckFeeData
	(*BattleRoyaleCallbackData)(nil),     // 3: proto.BattleRoyaleCallbackData
	(*BattleRoyaleCallbackResp)(nil),     // 4: proto.BattleRoyaleCallbackResp
	(*BattleRoyale)(nil),                 // 5: proto.BattleRoyale
	(*WidgetErrorMessage)(nil),           // 6: proto.WidgetErrorMessage
}
var file_widget_game_proto_depIdxs = []int32{
	0, // 0: proto.GameBattleRoyalCallbackReq.callback_type:type_name -> proto.GameBattleRoyalCallbackType
	2, // 1: proto.BattleRoyaleCallbackData.fee:type_name -> proto.GameBattleRoyaleCheckFeeData
	6, // 2: proto.BattleRoyaleCallbackResp.error:type_name -> proto.WidgetErrorMessage
	3, // 3: proto.BattleRoyaleCallbackResp.data:type_name -> proto.BattleRoyaleCallbackData
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_widget_game_proto_init() }
func file_widget_game_proto_init() {
	if File_widget_game_proto != nil {
		return
	}
	file_widget_error_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_widget_game_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameBattleRoyalCallbackReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_widget_game_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameBattleRoyaleCheckFeeData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_widget_game_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BattleRoyaleCallbackData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_widget_game_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BattleRoyaleCallbackResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_widget_game_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BattleRoyale); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_widget_game_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_widget_game_proto_goTypes,
		DependencyIndexes: file_widget_game_proto_depIdxs,
		EnumInfos:         file_widget_game_proto_enumTypes,
		MessageInfos:      file_widget_game_proto_msgTypes,
	}.Build()
	File_widget_game_proto = out.File
	file_widget_game_proto_rawDesc = nil
	file_widget_game_proto_goTypes = nil
	file_widget_game_proto_depIdxs = nil
}
