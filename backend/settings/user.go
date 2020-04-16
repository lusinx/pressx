package settings

// SettingName no
type SettingName struct {
	baseSetting
	state *string
}

// Change it
func (set *SettingName) Change(newState string) {
	*set.state = newState
}
