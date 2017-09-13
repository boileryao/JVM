package main

import (
	"JVM/rtdz"
	"fmt"
	"JVM/instructions/base"
	"JVM/instructions"
	"JVM/rtdz/heap"
)

func interpret(method *heap.Method, logInst bool, args []string) {
	thread := rtdz.NewThread()
	frame := thread.NewFrame(method)
	thread.PushFrame(frame)

	jArgs := createArgsArray(method.Class().Loader(), args)
	frame.LocalVars().SetRef(0, jArgs)

	defer catch(thread)
	loop(thread, logInst)
}
func createArgsArray(loader *heap.ClassLoader, args []string) *heap.Object {
	stringKls := loader.LoadClass("java/lang/String")
	argsArr := stringKls.ArrayClass().NewArray(uint(len(args)))
	jArgs := argsArr.Refs()
	for i, arg := range args {
		jArgs[i] = heap.JString(loader, arg)
	}
	return argsArr
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
