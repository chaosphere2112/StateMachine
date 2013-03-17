package state

type StateTarget interface {
}

type StateMachine struct {
	Target       StateTarget
	States       []func(*StateTarget, byte) (string, int)
	Output       []string
	CurrentState int
}

func (machine *StateMachine) AddState(state func(*StateTarget, byte) (string, int)) {

	machine.States = append(machine.States, state)

}

func (machine *StateMachine) ConsumeByte(data byte) {

	var result string

	result, machine.CurrentState = machine.States[machine.CurrentState](machine.target, data)

	if result != "" {
		machine.Output = append(machine.Output, result)
	}

}

func (machine *StateMachine) ConsumeBytes(data []byte) {
	for _, data_byte := range data {
		machine.ConsumeByte(data_byte)
	}
}
