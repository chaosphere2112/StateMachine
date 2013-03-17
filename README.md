StateMachine
============

Golang implementation of a simple state machine
    
    import (
        "fmt"
        "github.com/chaosphere2112/StateMachine"
    )

    func MyFunc(data byte) (string, int) {
        return append("",byte), 0     
    }
    
    func main() {

        stringbuilder := new(state.StateMachine)
        
        stringbuilder.AddState(MyFunc)
        
        data := "hello, world"

        for ind := 0; ind < len(data); ind++ {
            stringbuilder.Consume(data[ind])
        }
        
        fmt.Println(stringbuilder.Output)

    }