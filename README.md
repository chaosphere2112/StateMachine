StateMachine
============

Golang implementation of a simple state machine
    
    package main

    import (
        "fmt"
        "bytes"
        "github.com/chaosphere2112/StateMachine"
    )

    func MyFunc(data byte) (string, int) {
        var byteslice bytes.Buffer
        byteslice.WriteByte(data)
        return byteslice.String(), 0
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