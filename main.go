package main

import "fmt"
import "strings"
import (
	"JVM/classpath"
	"JVM/classfile"
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
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	fmt.Printf("Classpath:%s, class:%s args:%v\n",
		cmd.cpOption, cmd.class, cmd.args)

	className := strings.Replace(cmd.class, ".", "/", -1)
	cf := loadClass(className, cp)
	fmt.Println(cmd.class)
	printKlsInfo(cf)
}

func loadClass(klsName string, cp *classpath.Classpath) *classfile.ClassFile {
	classData, _, err := cp.ReadClass(klsName)
	if err != nil {
		panic(err)
	}

	cf, err := classfile.Parse(classData)
	if err != nil {
		panic(err)
	}
	return cf // abstraction of class file
}

func printKlsInfo(cf *classfile.ClassFile) {
	fmt.Printf("Version\t\t: %v.%v\n", cf.MajorVersion(), cf.MinorVersion())
	fmt.Printf("AccessFlag\t: 0x%x\n", cf.AccessFlags())
	fmt.Printf("This Class\t: %v\n", cf.ClassName())
	fmt.Printf("Super Class\t: %v\n", cf.SuperClassName())
	fmt.Printf("Interfaces\t: %v\n", cf.InterfaceNames())
	fmt.Printf("Constant Count\t: %v\n", len(cf.ConstantPool()))
	fmt.Printf("Fields Count\t:%v\n", len(cf.Fields()))
	for _, f := range cf.Fields() {
		fmt.Printf("\t%s - %s\n", f.Name(), f.Descriptor())
	}

	fmt.Printf("Method Count\t:%v\n", len(cf.Methods()))
	for _, m := range cf.Methods() {
		fmt.Printf("\t%s - %s\n", m.Name(), m.Descriptor())
	}
}
