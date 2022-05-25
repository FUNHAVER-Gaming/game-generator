package service

type ValRole string

const (
	Controller ValRole = "978812366454136832"
	Flex       ValRole = "978812462730207242"
	Duelist    ValRole = "978812420040573039"
	Sentinel   ValRole = "978812511451250718"

	OptimalFlex       = 4
	OptimalDuelist    = 2
	OptimalController = 2
	OptimalSentinel   = 2
)

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
	return ""
}
