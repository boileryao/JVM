package main

import "fmt"
import "strings"
import "JVM/classpath"

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
	classData, _, err := cp.ReadClass(className)

	if err != nil {
		fmt.Println("Load main class failed!")
		return
	}

	fmt.Printf("Class Content:\n%v", classData)
}
