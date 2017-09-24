package main 

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"smp/mlib"
	"smp/mp"
)

var lib *library.MusicManager
var id int = 1
var ctrl, signal chan int

func handleLibCommands(tokens []string) {
	switch tokens[1] {
		case "list": {
			for i := 0; i < lib.Len(); i++ {
				e, _ := lib.Get(i)
				fmt.Println(i + 1, ":", e.Name, e.Artist, e.Source, e.Type)
			}
		}
		case "add": {
			if len(tokens) == 6 {
				id++
				lib.Add(&mlib.MusicEntry{strconv.Itoa(id),
					tokens[2], tokens[3], tokens[4], tokens[5]})
			} else {
				fmt.Println("USAGE: lib add <name><artist><source><type>")
			}
		}
		case "remove": {
			if len(tokens) == 3 {
				lib.RemoveByName(tokens[2])
			} else {
				fmt.Println("USAGE: lib remove<name>")
			}
		}
		default: {
			fmt.Println("no command", tokens[1])
		}
	}
}

func handlePlayCommand(tokens []string) {
	if len(tokens) != 2 {
		fmt.Println("USAGE: play <name>")
		return
	}
	e := lib.Find(tokens[1])
	if e == nil {
		fmt.Println("The misic", tokens[1], "dos not exist")
		return
	}

	mp.Play(e.Source, e.Type, ctrl, signal)
}

func main() {
	fmt.Println(
		
	)
}