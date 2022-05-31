package service

import "github.com/FUNHAVER-Gaming/game-generator/pkg/consts"

type ValRole int

const (
	Controller  ValRole = 1
	Initiator   ValRole = 2
	Duelist     ValRole = 3
	Sentinel    ValRole = 4
	InvalidRole ValRole = -1

	OptimalInitiator  = 4
	OptimalDuelist    = 2
	OptimalController = 2
	OptimalSentinel   = 2
)

func (vr ValRole) getRoleId() string {
	switch vr {
	case Initiator:
		return consts.InitiatorRoleId
	case Controller:
		return consts.ControllerRoleId
	case Sentinel:
		return consts.SentinelRoleId
	case Duelist:
		return consts.DuelistRoleId
	}
	return ""
}

func (vr ValRole) getName() string {
	switch vr {
	case Initiator:
		return "Initiator"
	case Controller:
		return "Controller"
	case Sentinel:
		return "Sentinel"
	case Duelist:
		return "Duelist"
	}
	return ""
}

func getValRoleFromRoleID(roleId string) ValRole {
	switch roleId {
	case consts.ControllerRoleId:
		return Controller
	case consts.InitiatorRoleId:
		return Initiator
	case consts.DuelistRoleId:
		return Duelist
	case consts.SentinelRoleId:
		return Sentinel
	}
	return InvalidRole
}
