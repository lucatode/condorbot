package dispacher

type Dispacher interface {
	GetActionFunc(input string) (bool, func([]string) string)
}

type CommandDispacher struct {
	actions map[string]func([]string) string
}

func (cd CommandDispacher) GetActionFunc(input string) (bool, func([]string) string) {
	f, ok := cd.actions[input]
	return ok, f
}

func NewCommandDispacher(dict map[string]func([]string) string) Dispacher {
	return CommandDispacher{dict}
}
