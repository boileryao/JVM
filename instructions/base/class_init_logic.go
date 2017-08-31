package base

import (
	"JVM/rtdz"
	"JVM/rtdz/heap"
)

// jvm spec 5.5
func InitClass(thread *rtdz.Thread, class *heap.Class) {
	class.SetInited()
	scheduleClinit(thread, class)
	initSuperClass(thread, class)
}

func scheduleClinit(thread *rtdz.Thread, class *heap.Class) {
	clinit := class.GetClinitMethod()
	if clinit != nil {
		// exec <clinit>
		newFrame := thread.NewFrame(clinit)
		thread.PushFrame(newFrame)
	}
}

func initSuperClass(thread *rtdz.Thread, class *heap.Class) {
	if !class.IsInterface() {
		superClass := class.SuperClass()
		if superClass != nil && !superClass.Inited() {
			InitClass(thread, superClass)
		}
	}
}