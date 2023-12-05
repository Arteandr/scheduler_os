package command

import "kurs_scheduler/internal/scheduler"

type Executor struct {
	Scheduler *scheduler.Scheduler
	CMDs      map[string]Command
}

func NewCommandExecutor(scheduler *scheduler.Scheduler) *Executor {
	cmdExecutor := &Executor{
		CMDs:      make(map[string]Command),
		Scheduler: scheduler,
	}

	cmdExecutor.CMDs["run"] = NewRunCommand()
	cmdExecutor.CMDs["clear"] = NewClearCommand()
	cmdExecutor.CMDs["quantum"] = NewQuantumCommand(scheduler)

	return cmdExecutor
}

func (e *Executor) ExecuteCMD(input *Args) Status {
	cmd, ok := e.CMDs[input.Alias]
	if !ok {
		return ErrorCmd
	}

	cmd.Execute(input.Args)

	return SuccessCmd
}
