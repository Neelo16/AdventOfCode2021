package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Command struct {
	Operation string
	Step      int
}

func ReadInput() []Command {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	commands := make([]Command, 0)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		command := fields[0]
		step, _ := strconv.Atoi(fields[1])
		commands = append(commands, Command{command, step})
	}
	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	return commands
}

type Depth struct {
	X   int
	Y   int
	Aim int
}

type Operation func(depth *Depth, step int)

func main() {
	ops := map[string]Operation{
		"forward": func(depth *Depth, step int) {
			depth.X += step
		},
		"up": func(depth *Depth, step int) {
			depth.Y -= step
		},
		"down": func(depth *Depth, step int) {
			depth.Y += step
		},
	}
	commands := ReadInput()
	depth := Depth{}
	for _, command := range commands {
		ops[command.Operation](&depth, command.Step)
	}
	log.Printf("First star: %v\n", depth.X*depth.Y)

	depth = Depth{}
	ops = map[string]Operation{
		"forward": func(depth *Depth, step int) {
			depth.X += step
			depth.Y += depth.Aim * step
		},
		"up": func(depth *Depth, step int) {
			depth.Aim -= step
		},
		"down": func(depth *Depth, step int) {
			depth.Aim += step
		},
	}
	for _, command := range commands {
		ops[command.Operation](&depth, command.Step)
	}
	log.Printf("Second star: %v\n", depth.X*depth.Y)
}
