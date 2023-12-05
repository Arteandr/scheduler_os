package command

type Status uint8

const (
	ErrorCmd Status = iota
	SuccessCmd
)

type Command interface {
	Execute(args []string)
}

type Args struct {
	Alias string
	Args  []string
}

func NewCommandArgs(cmd string, args []string) *Args {
	return &Args{
		Alias: cmd,
		Args:  args,
	}
}
