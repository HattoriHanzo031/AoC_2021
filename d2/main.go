package main

import (
	"bufio"
	"fmt"
	"os"
)

type Command struct {
	name  string
	value int
}

type Possition struct {
	x     int
	depth int
}

type Submarine struct {
	pos Possition
	aim int
}

func (sub *Submarine) MoveP1(command Command) {
	switch command.name {
	case "forward":
		sub.pos.x += command.value
	case "down":
		sub.pos.depth += command.value
	case "up":
		sub.pos.depth -= command.value
	default:
		panic("Unknown command")
	}
}

func (sub *Submarine) MoveP2(command Command) {
	switch command.name {
	case "forward":
		sub.pos.depth += sub.aim * command.value
		sub.pos.x += command.value
	case "down":
		sub.aim += command.value
	case "up":
		sub.aim -= command.value
	default:
		panic("Unknown command")
	}
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scaner := bufio.NewScanner(f)
	var commands []Command
	for scaner.Scan() {
		var command Command
		n, err := fmt.Sscan(scaner.Text(), &command.name, &command.value)
		if err != nil || n != 2 {
			panic(err)
		}
		commands = append(commands, command)
	}
	sub := Submarine{}

	for _, com := range commands {
		sub.MoveP2(com)
	}

	fmt.Println(sub, sub.pos.x*sub.pos.depth)
}
