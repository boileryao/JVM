package main

import "flag"
import "fmt"
import "os"

type Cmd struct {
	helpFlag bool
	versionFlag bool
	cpOption string
	class string
	args []string
}

func parse_cmd() *Cmd{
	cmd := &Cmd{}

	flag.Usage = print_usage
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help messages")
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help messages")	
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version messages")
	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath is xx.")
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath is xx.")
	flag.Parse()

	args := flag.Args()
	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}

	return cmd
}

func print_usage() {
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
}