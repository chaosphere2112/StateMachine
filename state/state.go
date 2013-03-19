package state

const (
	FinishedState = -1
)

type StateMachine interface {
	TakeData(data byte, state int) (string, int)
	Done()
}

type StateEngine struct {
	Machine      StateMachine
	Output       []string
	CurrentState int
}

func (engine *StateEngine) ConsumeByte(data byte) {

	var result string

	result, engine.CurrentState = engine.Machine.TakeData(data, engine.CurrentState)

	if result != "" {
		engine.Output = append(engine.Output, result)
	}

	if engine.CurrentState == FinishedState {
		engine.Machine.Done()
		engine.CurrentState = 0
	}

}

func (engine *StateEngine) ConsumeBytes(data []byte) {
	for _, data_byte := range data {
		engine.ConsumeByte(data_byte)
	}
}
