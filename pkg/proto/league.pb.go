// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: league.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type CreateGameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VoiceChannelMembers []*VoiceChannelMember `protobuf:"bytes,1,rep,name=voice_channel_members,json=voiceChannelMembers,proto3" json:"voice_channel_members,omitempty"`
	VoiceChannelId      string                `protobuf:"bytes,2,opt,name=voice_channel_id,json=voiceChannelId,proto3" json:"voice_channel_id,omitempty"`
}

func (x *CreateGameRequest) Reset() {
	*x = CreateGameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_league_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGameRequest) ProtoMessage() {}

func (x *CreateGameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_league_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGameRequest.ProtoReflect.Descriptor instead.
func (*CreateGameRequest) Descriptor() ([]byte, []int) {
	return file_league_proto_rawDescGZIP(), []int{0}
}

func (x *CreateGameRequest) GetVoiceChannelMembers() []*VoiceChannelMember {
	if x != nil {
		return x.VoiceChannelMembers
	}
	return nil
}

func (x *CreateGameRequest) GetVoiceChannelId() string {
	if x != nil {
		return x.VoiceChannelId
	}
	return ""
}

type CreateGameResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok                   bool      `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	Error                string    `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Attackers            []*Player `protobuf:"bytes,3,rep,name=attackers,proto3" json:"attackers,omitempty"`
	Defenders            []*Player `protobuf:"bytes,4,rep,name=defenders,proto3" json:"defenders,omitempty"`
	Map                  string    `protobuf:"bytes,5,opt,name=map,proto3" json:"map,omitempty"`
	LobbyLeader          *Player   `protobuf:"bytes,6,opt,name=lobby_leader,json=lobbyLeader,proto3" json:"lobby_leader,omitempty"`
	GameId               string    `protobuf:"bytes,7,opt,name=game_id,json=gameId,proto3" json:"game_id,omitempty"`
	ExcessPlayersRemoved uint32    `protobuf:"varint,8,opt,name=excess_players_removed,json=excessPlayersRemoved,proto3" json:"excess_players_removed,omitempty"`
	RemovedPlayers       []*Player `protobuf:"bytes,9,rep,name=removed_players,json=removedPlayers,proto3" json:"removed_players,omitempty"`
}

func (x *CreateGameResponse) Reset() {
	*x = CreateGameResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_league_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGameResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGameResponse) ProtoMessage() {}

func (x *CreateGameResponse) ProtoReflect() protoreflect.Message {
	mi := &file_league_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGameResponse.ProtoReflect.Descriptor instead.
func (*CreateGameResponse) Descriptor() ([]byte, []int) {
	return file_league_proto_rawDescGZIP(), []int{1}
}

func (x *CreateGameResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

func (x *CreateGameResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *CreateGameResponse) GetAttackers() []*Player {
	if x != nil {
		return x.Attackers
	}
	return nil
}

func (x *CreateGameResponse) GetDefenders() []*Player {
	if x != nil {
		return x.Defenders
	}
	return nil
}

func (x *CreateGameResponse) GetMap() string {
	if x != nil {
		return x.Map
	}
	return ""
}

func (x *CreateGameResponse) GetLobbyLeader() *Player {
	if x != nil {
		return x.LobbyLeader
	}
	return nil
}

func (x *CreateGameResponse) GetGameId() string {
	if x != nil {
		return x.GameId
	}
	return ""
}

func (x *CreateGameResponse) GetExcessPlayersRemoved() uint32 {
	if x != nil {
		return x.ExcessPlayersRemoved
	}
	return 0
}

func (x *CreateGameResponse) GetRemovedPlayers() []*Player {
	if x != nil {
		return x.RemovedPlayers
	}
	return nil
}

type CreatePlayerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Player *Player `protobuf:"bytes,1,opt,name=player,proto3" json:"player,omitempty"`
}

func (x *CreatePlayerRequest) Reset() {
	*x = CreatePlayerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_league_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePlayerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePlayerRequest) ProtoMessage() {}

func (x *CreatePlayerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_league_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePlayerRequest.ProtoReflect.Descriptor instead.
func (*CreatePlayerRequest) Descriptor() ([]byte, []int) {
	return file_league_proto_rawDescGZIP(), []int{2}
}

func (x *CreatePlayerRequest) GetPlayer() *Player {
	if x != nil {
		return x.Player
	}
	return nil
}

type CreatePlayerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *CreatePlayerResponse) Reset() {
	*x = CreatePlayerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_league_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePlayerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePlayerResponse) ProtoMessage() {}

func (x *CreatePlayerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_league_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePlayerResponse.ProtoReflect.Descriptor instead.
func (*CreatePlayerResponse) Descriptor() ([]byte, []int) {
	return file_league_proto_rawDescGZIP(), []int{3}
}

func (x *CreatePlayerResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type VoiceChannelMember struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Username    string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Roles       []string `protobuf:"bytes,3,rep,name=roles,proto3" json:"roles,omitempty"`
	DisplayName string   `protobuf:"bytes,4,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
}

func (x *VoiceChannelMember) Reset() {
	*x = VoiceChannelMember{}
	if protoimpl.UnsafeEnabled {
		mi := &file_league_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VoiceChannelMember) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VoiceChannelMember) ProtoMessage() {}

func (x *VoiceChannelMember) ProtoReflect() protoreflect.Message {
	mi := &file_league_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VoiceChannelMember.ProtoReflect.Descriptor instead.
func (*VoiceChannelMember) Descriptor() ([]byte, []int) {
	return file_league_proto_rawDescGZIP(), []int{4}
}

func (x *VoiceChannelMember) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *VoiceChannelMember) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *VoiceChannelMember) GetRoles() []string {
	if x != nil {
		return x.Roles
	}
	return nil
}

func (x *VoiceChannelMember) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

type Player struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DiscordId   string   `protobuf:"bytes,1,opt,name=discord_id,json=discordId,proto3" json:"discord_id,omitempty"`
	DisplayName string   `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	UserId      string   `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	RiotId      string   `protobuf:"bytes,4,opt,name=riot_id,json=riotId,proto3" json:"riot_id,omitempty"`
	RiotTag     string   `protobuf:"bytes,5,opt,name=riot_tag,json=riotTag,proto3" json:"riot_tag,omitempty"`
	Roles       []string `protobuf:"bytes,6,rep,name=roles,proto3" json:"roles,omitempty"`
}

func (x *Player) Reset() {
	*x = Player{}
	if protoimpl.UnsafeEnabled {
		mi := &file_league_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Player) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Player) ProtoMessage() {}

func (x *Player) ProtoReflect() protoreflect.Message {
	mi := &file_league_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Player.ProtoReflect.Descriptor instead.
func (*Player) Descriptor() ([]byte, []int) {
	return file_league_proto_rawDescGZIP(), []int{5}
}

func (x *Player) GetDiscordId() string {
	if x != nil {
		return x.DiscordId
	}
	return ""
}

func (x *Player) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *Player) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Player) GetRiotId() string {
	if x != nil {
		return x.RiotId
	}
	return ""
}

func (x *Player) GetRiotTag() string {
	if x != nil {
		return x.RiotTag
	}
	return ""
}

func (x *Player) GetRoles() []string {
	if x != nil {
		return x.Roles
	}
	return nil
}

var File_league_proto protoreflect.FileDescriptor

var file_league_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6c, 0x65, 0x61, 0x67, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x86,
	0x01, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x47, 0x0a, 0x15, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x56, 0x6f, 0x69, 0x63, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x13, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x43,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x12, 0x28, 0x0a,
	0x10, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x5f, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x43, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x22, 0xc7, 0x02, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x12, 0x25, 0x0a, 0x09, 0x61, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x65, 0x72,
	0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x52, 0x09, 0x61, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x73, 0x12, 0x25, 0x0a, 0x09, 0x64,
	0x65, 0x66, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x07,
	0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x09, 0x64, 0x65, 0x66, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x61, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6d, 0x61, 0x70, 0x12, 0x2a, 0x0a, 0x0c, 0x6c, 0x6f, 0x62, 0x62, 0x79, 0x5f, 0x6c, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x52, 0x0b, 0x6c, 0x6f, 0x62, 0x62, 0x79, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x12, 0x17, 0x0a, 0x07, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x67, 0x61, 0x6d, 0x65, 0x49, 0x64, 0x12, 0x34, 0x0a, 0x16, 0x65, 0x78, 0x63,
	0x65, 0x73, 0x73, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x5f, 0x72, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x14, 0x65, 0x78, 0x63, 0x65, 0x73,
	0x73, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x12,
	0x30, 0x0a, 0x0f, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x52, 0x0e, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x64, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x73, 0x22, 0x36, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x06, 0x70, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x52, 0x06, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x22, 0x2f, 0x0a, 0x14, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x79, 0x0a, 0x12, 0x56, 0x6f,
	0x69, 0x63, 0x65, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x72, 0x6f, 0x6c,
	0x65, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xad, 0x01, 0x0a, 0x06, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x72, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x72, 0x64, 0x49, 0x64, 0x12,
	0x21, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x72,
	0x69, 0x6f, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x69,
	0x6f, 0x74, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x69, 0x6f, 0x74, 0x5f, 0x74, 0x61, 0x67,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x69, 0x6f, 0x74, 0x54, 0x61, 0x67, 0x12,
	0x14, 0x0a, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05,
	0x72, 0x6f, 0x6c, 0x65, 0x73, 0x32, 0x87, 0x01, 0x0a, 0x0d, 0x4c, 0x65, 0x61, 0x67, 0x75, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x37, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x47, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x61,
	0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x3d, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x12, 0x14, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x0a, 0x5a, 0x08, 0x2e, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_league_proto_rawDescOnce sync.Once
	file_league_proto_rawDescData = file_league_proto_rawDesc
)

func file_league_proto_rawDescGZIP() []byte {
	file_league_proto_rawDescOnce.Do(func() {
		file_league_proto_rawDescData = protoimpl.X.CompressGZIP(file_league_proto_rawDescData)
	})
	return file_league_proto_rawDescData
}

var file_league_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_league_proto_goTypes = []interface{}{
	(*CreateGameRequest)(nil),    // 0: CreateGameRequest
	(*CreateGameResponse)(nil),   // 1: CreateGameResponse
	(*CreatePlayerRequest)(nil),  // 2: CreatePlayerRequest
	(*CreatePlayerResponse)(nil), // 3: CreatePlayerResponse
	(*VoiceChannelMember)(nil),   // 4: VoiceChannelMember
	(*Player)(nil),               // 5: Player
}
var file_league_proto_depIdxs = []int32{
	4, // 0: CreateGameRequest.voice_channel_members:type_name -> VoiceChannelMember
	5, // 1: CreateGameResponse.attackers:type_name -> Player
	5, // 2: CreateGameResponse.defenders:type_name -> Player
	5, // 3: CreateGameResponse.lobby_leader:type_name -> Player
	5, // 4: CreateGameResponse.removed_players:type_name -> Player
	5, // 5: CreatePlayerRequest.player:type_name -> Player
	0, // 6: LeagueService.CreateGame:input_type -> CreateGameRequest
	2, // 7: LeagueService.CreatePlayer:input_type -> CreatePlayerRequest
	1, // 8: LeagueService.CreateGame:output_type -> CreateGameResponse
	3, // 9: LeagueService.CreatePlayer:output_type -> CreatePlayerResponse
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_league_proto_init() }
func file_league_proto_init() {
	if File_league_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_league_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGameRequest); i {
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
		file_league_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGameResponse); i {
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
		file_league_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePlayerRequest); i {
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
		file_league_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePlayerResponse); i {
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
		file_league_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VoiceChannelMember); i {
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
		file_league_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Player); i {
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
			RawDescriptor: file_league_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_league_proto_goTypes,
		DependencyIndexes: file_league_proto_depIdxs,
		MessageInfos:      file_league_proto_msgTypes,
	}.Build()
	File_league_proto = out.File
	file_league_proto_rawDesc = nil
	file_league_proto_goTypes = nil
	file_league_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// LeagueServiceClient is the client API for LeagueService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LeagueServiceClient interface {
	CreateGame(ctx context.Context, in *CreateGameRequest, opts ...grpc.CallOption) (*CreateGameResponse, error)
	CreatePlayer(ctx context.Context, in *CreatePlayerRequest, opts ...grpc.CallOption) (*CreatePlayerResponse, error)
}

type leagueServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLeagueServiceClient(cc grpc.ClientConnInterface) LeagueServiceClient {
	return &leagueServiceClient{cc}
}

func (c *leagueServiceClient) CreateGame(ctx context.Context, in *CreateGameRequest, opts ...grpc.CallOption) (*CreateGameResponse, error) {
	out := new(CreateGameResponse)
	err := c.cc.Invoke(ctx, "/LeagueService/CreateGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *leagueServiceClient) CreatePlayer(ctx context.Context, in *CreatePlayerRequest, opts ...grpc.CallOption) (*CreatePlayerResponse, error) {
	out := new(CreatePlayerResponse)
	err := c.cc.Invoke(ctx, "/LeagueService/CreatePlayer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LeagueServiceServer is the server API for LeagueService service.
type LeagueServiceServer interface {
	CreateGame(context.Context, *CreateGameRequest) (*CreateGameResponse, error)
	CreatePlayer(context.Context, *CreatePlayerRequest) (*CreatePlayerResponse, error)
}

// UnimplementedLeagueServiceServer can be embedded to have forward compatible implementations.
type UnimplementedLeagueServiceServer struct {
}

func (*UnimplementedLeagueServiceServer) CreateGame(context.Context, *CreateGameRequest) (*CreateGameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGame not implemented")
}
func (*UnimplementedLeagueServiceServer) CreatePlayer(context.Context, *CreatePlayerRequest) (*CreatePlayerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePlayer not implemented")
}

func RegisterLeagueServiceServer(s *grpc.Server, srv LeagueServiceServer) {
	s.RegisterService(&_LeagueService_serviceDesc, srv)
}

func _LeagueService_CreateGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LeagueServiceServer).CreateGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/LeagueService/CreateGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LeagueServiceServer).CreateGame(ctx, req.(*CreateGameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LeagueService_CreatePlayer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePlayerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LeagueServiceServer).CreatePlayer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/LeagueService/CreatePlayer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LeagueServiceServer).CreatePlayer(ctx, req.(*CreatePlayerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LeagueService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "LeagueService",
	HandlerType: (*LeagueServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateGame",
			Handler:    _LeagueService_CreateGame_Handler,
		},
		{
			MethodName: "CreatePlayer",
			Handler:    _LeagueService_CreatePlayer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "league.proto",
}
