package state

type StateMachine struct {
	States       []func(byte) (string, int)
	Output       []string
	CurrentState int
}

func (machine *StateMachine) AddState(state func(byte) (string, int)) {
    
    machine.States = append(machine.States, state)

}

func (machine *StateMachine) Consume(data byte) {

	var result string
	result, machine.CurrentState = machine.States[machine.CurrentState](data)
	machine.Output = append(machine.Output, result)

}
