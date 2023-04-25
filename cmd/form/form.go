package form



type Argument struct {
	argMap map[string][]string
}

func New() *Argument {
	return &Argument{argMap: make(map[string][]string)}
}

func NewWithParsing(args []string) *Argument {

	if (len(args) == 0) {
		return New()
	}

	var prevArg []string
	var argMap map[string][]string

	for _, arg := range args {
		if arg[0] == '-' {
			argMap[arg] = make([]string, 0)
			prevArg = argMap[arg]
		} else {
			prevArg = append(prevArg, arg)
		}
	}

	return &Argument{argMap: argMap}
}