package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	getArgs()
}
func getArgs() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("panic occurred:", err)
		}
	}()
	arg := os.Args
	if len(arg) > 2 || len(arg) < 2 {
		fmt.Println("Knowledge Worker Control Interface Requires One Argument")
	} else {
		getArgCommand(arg[1])
	}
}
func getArgCommand(command string) {
	help := "Valid commands \n-------\nhello -\tresponds with hello back\nsomethingelse -\tresponds with Not sure what to say\nhelp -\tprints this message"

	switch command {
	case "query":
		fmt.Print("Please specify your query: ")
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		text = strings.ToLower(text)
		if strings.Contains(text, "zion") {
			fmt.Println("The Zion Lattice Interface is Blocked for now")
		} else {
			fmt.Println("Error accessing database.")
		}
	case "init":
		initializeKwill()
	case "hello":
		fmt.Println("Hello Back")
	case "somethingelse":
		fmt.Println("Not sure what to say")
	case "help":
		fmt.Println(help)
	default:
		fmt.Println("Please provide a valid argument. Use `kwctl help` for a list of valid arguments/commands.")
	}
}
func initializeKwill(){
	fmt.Println("Booting Custom Knowledge Worker Subroutines...")
	time.Sleep(2 * time.Second)
	fmt.Println("Loading Zion Lattice Interface...")
	time.Sleep(2 * time.Second)
	fmt.Println("Boot sequence Complete...")
	fmt.Println("Welcome Kwi-II.  Commands are now accessible via kwctl interface.")
}
