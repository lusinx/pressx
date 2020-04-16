package settings

// Setting could be anything
type Setting interface {
	Change(newState string)
}

// Exec the change?
func Exec(s Setting, newState string) {
	s.Change(newState)
}

// Settings are settings
type Settings = []Setting

type baseSetting struct {
	name string
}
