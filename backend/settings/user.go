package settings

// User settings
type SettingName struct {
	baseSetting
	state *string
}

func (set *SettingName) Change(new_state string) {
	*set.state = new_state
}
