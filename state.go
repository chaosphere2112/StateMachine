package state

type StateMachine struct {
	States       []func(byte) (string, StateChange)
	Output       []string
	CurrentState int
}

func (machine *StateMachine) AddState(state func(byte) (string, int)) {

	machine.States = append(machine.States, state)

}

func (machine *StateMachine) ConsumeByte(data byte) {

	var result string

	result, machine.CurrentState = machine.States[machine.CurrentState](data)

	if result != "" {
		machine.Output = append(machine.Output, result)
	}

}

func (machine *StateMachine) ConsumeBytes(data []byte) {
	for index, data_byte := range data {
		machine.ConsumeByte(data_byte)
	}
}
