package main

import "fmt"

func main() {
	cmd := parse_cmd()
	if cmd.versionFlag {
		fmt.Println("Boiler JVM, Version 0.0.1")
	} else if cmd.helpFlag || cmd.class == ""{
		print_usage()
	} else {
		start_jvm(cmd)
	}
}

func start_jvm(cmd *Cmd) {
	fmt.Println("JVM Lanuched:")
	fmt.Printf("Classpath:%s, class:%s args:%v\n",
		cmd.cpOption, cmd.class, cmd.args)
}