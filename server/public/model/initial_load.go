package model

type InitialLoad struct {
	User        *User             `json:"user"`
	TeamMembers []*TeamMember     `json:"team_members"`
	Teams       []*Team           `json:"teams"`
	Preferences Preferences       `json:"preferences"`
	ClientCfg   map[string]string `json:"client_cfg"`
	LicenseCfg  map[string]string `json:"license_cfg"`
	NoAccounts  bool              `json:"no_accounts"`
}
