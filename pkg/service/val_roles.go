package service

type ValRole int

const (
	Controller ValRole = 1
	Flex       ValRole = 2
	Duelist    ValRole = 3
	Sentinel   ValRole = 4

	OptimalFlex       = 4
	OptimalDuelist    = 2
	OptimalController = 2
	OptimalSentinel   = 2
)

func (vr ValRole) getRoleId() string {
	switch vr {
	case Flex:
		return "978812462730207242"
	case Controller:
		return "978812366454136832"
	case Sentinel:
		return "978812511451250718"
	case Duelist:
		return "978812420040573039"
	}
	return ""
}

func getValRoleFromRoleID(roleId string) ValRole {
	switch roleId {
	case "978812366454136832":
		return Controller
	case "978812462730207242":
		return Flex
	case "978812420040573039":
		return Duelist
	case "978812511451250718":
		return Sentinel
	}
	return -1
}
