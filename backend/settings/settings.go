package settings

type Setting interface {
	Change(new_state string)
}

func Exec(s Setting, new_state string) {
	s.Change(new_state)
}

type Settings = []Setting

type baseSetting struct {
	name string
}
