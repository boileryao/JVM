package main

import (
	"JVM/rtdz"
	"fmt"
	"JVM/instructions/base"
	"JVM/instructions"
	"JVM/rtdz/heap"
)

func interpret(method *heap.Method) {
	thread := rtdz.NewThread()
	frame := thread.NewFrame(method)
	thread.PushFrame(frame)

	defer catch(frame)
	loop(thread, method.Code())
}

func catch(frame *rtdz.Frame) {
	if r := recover(); r != nil {
		fmt.Printf("Local Vars: %v\n", frame.LocalVars())
		fmt.Printf("Operand Stack: %v\n", frame.OperandStack())
		fmt.Println("The main method may returned")
		//panic(r)
	}
}

func loop(thread *rtdz.Thread, bytecode []byte) {
	frame := thread.PopFrame()
	reader := &base.BytecodeReader{}


	for {
		pc := frame.NextPC()
		thread.SetPC(pc)

		//decode
		reader.Reset(bytecode, pc)
		opcode := reader.ReadUint8()
		inst := instructions.NewInstruction(opcode)
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.PC())

		//exec
		fmt.Printf("PC: 0x%02x, Inst: %T %v\n", pc, inst, inst)
		inst.Execute(frame)
	}
}
