package main

import (
	"JVM/rtdz"
	"fmt"
	"JVM/instructions/base"
	"JVM/instructions"
	"JVM/rtdz/heap"
)

func interpret(method *heap.Method, logInst bool) {
	thread := rtdz.NewThread()
	frame := thread.NewFrame(method)
	thread.PushFrame(frame)

	defer catch(thread)
	loop(thread, logInst)
}

func catch(thread *rtdz.Thread) {
	if r := recover(); r != nil {
		logFrames(thread)
		panic(r)
	}
}
func loop(thread *rtdz.Thread, logInst bool) {
	reader := &base.BytecodeReader{}

	for {
		frame := thread.CurrentFrame()
		pc := frame.NextPC()
		thread.SetPC(pc)

		//decode
		reader.Reset(frame.Method().Code(), pc)
		opcode := reader.ReadUint8()
		inst := instructions.NewInstruction(opcode)
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.PC())

		if logInst {
			logInstruction(frame, inst)
		}

		//exec
		inst.Execute(frame)
		if thread.IsStackEmpty() {
			break
		}
	}
}

func logInstruction(frame *rtdz.Frame, inst base.Instruction) {
	method := frame.Method()
	className := method.Class().Name()
	methodName := method.Name()
	pc := frame.Thread().PC()
	fmt.Printf("%v.%v() #%2d %T %v\n", className, methodName, pc, inst, inst)
}

func logFrames(thread *rtdz.Thread) {
	for !thread.IsStackEmpty() {
		frame := thread.PopFrame()
		method := frame.Method()
		className := method.Class().Name()
		fmt.Printf(">> pc:%4d %v.%v%v \n",
			frame.NextPC(), className, method.Name(), method.Descriptor())
	}
}
