package base

import "JVM/rtdz"

func Branch(frame *rtdz.Frame, offset int) {
	pc := frame.Thread().PC()
	nextPC := pc + offset
	frame.SetNextPC(nextPC)
}
