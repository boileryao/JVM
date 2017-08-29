package main

import "fmt"
import (
	"JVM/classpath"
	"strings"
	"JVM/classfile"
	"JVM/rtdz/heap"
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
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	classLoader := heap.NewClassLoader(cp, cmd.verboseClassFlag)
	className := strings.Replace(cmd.class, ".", "/", -1)
	mainClass := classLoader.LoadClass(className)
	mainMethod := mainClass.GetMainMethod()

	if mainMethod != nil {
		interpret(mainMethod, cmd.verboseInstFlag)
	} else {
		fmt.Printf("Not Found Main Method in %s.", className)
	}
}

func getMainMethod(file *classfile.ClassFile) *classfile.MemberInfo {
	for _, m := range file.Methods() {
		if m.Name() == "main" && m.Descriptor() == "([Ljava/lang/String;)V" {
			return m
		}
	}
	return nil
}

func loadClass(name string, cp *classpath.Classpath) *classfile.ClassFile {
	classData, _, err := cp.ReadClass(name)
	if err != nil {
		panic(err)
	}

	cf, err := classfile.Parse(classData)
	if err != nil {
		panic(err)
	}

	return cf
}
