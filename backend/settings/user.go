package settings

// SettingName structure
type SettingName struct {
	baseSetting
	state *string
}

// Change methods
func (set *SettingName) Change(newState string) {
	*set.state = newState
}
