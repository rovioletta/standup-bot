package domain

type Team struct {
	TeamMembers      []string
	TeamName         string
	ManagerID        string
	ChannelID        string
	ID               uint64
	NotificationHour uint16
}
