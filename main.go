package main

import "fmt"
import (
	"JVM/rtdz"
)

func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Println("Boiler JVM, Version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJvm(cmd)
	}
}

func startJvm(cmd *Cmd) {
	fmt.Println("JVM Lanuched:")
	frame := rtdz.NewFrame(100, 100)
	testLocalVars(frame.LocalVars())
	testOperandStack(frame.OperandStack())
}

func testLocalVars(vars rtdz.LocalVars) {
	vars.SetInt(0, 100)
	vars.SetInt(1, -100)
	vars.SetLong(2, 172331623333)
	vars.SetLong(4, -2997924580)
	vars.SetFloat(6, 3.14)
	vars.SetDouble(7, 2.7113124141)
	vars.SetRef(9, nil)
	fmt.Println("Local Vars POP: ", vars.GetInt(0))
	fmt.Println("Local Vars POP: ", vars.GetInt(1))
	fmt.Println("Local Vars POP: ", vars.GetLong(2))
	fmt.Println("Local Vars POP: ", vars.GetLong(4))
	fmt.Println("Local Vars POP: ", vars.GetFloat(6))
	fmt.Println("Local Vars POP: ", vars.GetDouble(7))
	fmt.Println("Local Vars POP: ", vars.GetRef(9))
}

func testOperandStack(stack *rtdz.OperandStack) {
	stack.PushInt(1)
	stack.PushLong(23232235)
	stack.PushDouble(2.71828182845)
	stack.PushFloat(1.2323)
	stack.PushRef(nil)
	fmt.Println("Operand Stack POP: ", stack.PopRef())
	fmt.Println("Operand Stack POP: ", stack.PopFloat())
	fmt.Println("Operand Stack POP: ", stack.PopDouble())
	fmt.Println("Operand Stack POP: ", stack.PopLong())
	fmt.Println("Operand Stack POP: ", stack.PopInt())
}
