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

	cmdExecutor.CMDs["run"] = NewRunCommand(scheduler)
	cmdExecutor.CMDs["clear"] = NewClearCommand()
	cmdExecutor.CMDs["exit"] = NewExitCommand()
	cmdExecutor.CMDs["quantum"] = NewQuantumCommand(scheduler)
	cmdExecutor.CMDs["burst"] = NewBurstCommand(scheduler)
	cmdExecutor.CMDs["ps"] = NewProcessListCommand(scheduler)

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
