package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/briandowns/spinner"
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
		getArgCommand(arg[1], "kw")
	}
}
func getArgCommand(command string, user string) {
	help := `
Valid commands
-------

hello			responds with hello back
somethingelse		responds with Not sure what to say
help			prints this message
query			Prompts for a question, has limited answers in kw_user mode
elevate			Elevates kw_user to root access, opening permissions to the entire system
init			Initializes the kwctl CLI
exit			Exits the CLI, escapes root mode if elevated`

	switch command {
	case "query":
		fmt.Print("Please specify your query: ")
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		text = strings.ToLower(text)
		queried := text
		answerQuery(queried)
	case "init":
		initializeKwill()
	case "hello":
		fmt.Println("Hello Back")
	case "somethingelse":
		fmt.Println("Not sure what to say")
	case "help":
		fmt.Println(help)
	case "elevate":
		fmt.Print("Enter encrypted passcode to elevate privileges: ")
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		text = strings.ToLower(text)
		user := "su"
		openPrompt(user)
	case "exit":
		if user == "su" {
			openPrompt("kw")
		} else {
			os.Exit(0)
		}
	default:
		fmt.Println("Please provide a valid argument. Use `kwctl help` for a list of valid arguments/commands.")
	}
	openPrompt(user)
}
func initializeKwill() {
	s := spinner.New(spinner.CharSets[25], 100*time.Millisecond) // Build our new spinner
	s.Color("red")                                               // Set the spinner color to red
	s.Start()                                                    // Start the spinner
	time.Sleep(3 * time.Second)                                  // Run for some time to simulate work
	fmt.Println("Booting Custom Knowledge Worker Subroutines...")
	time.Sleep(2 * time.Second)
	fmt.Println("Loading Zion Lattice Interface...")
	time.Sleep(2 * time.Second)
	fmt.Println("Boot sequence Complete...")
	s.Stop()
	fmt.Println("Welcome Kwi-II.  Commands are now accessible via kwctl interface.")
	user := "kw"
	openPrompt(user)
}
func openPrompt(user string) {
	const colorReset = "\033[0m"

	const colorRed = "\033[31m"
	const colorGreen = "\033[32m"
	const colorYellow = "\033[33m"
	const colorBlue = "\033[34m"
	const colorPurple = "\033[35m"
	const colorCyan = "\033[36m"
	const colorWhite = "\033[37m"
	if user == "kw" {
		fmt.Print("$ kw_user/~ > ")
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		text = strings.ToLower(text)
		command := text
		if command == "init" {
			fmt.Println(string(colorRed), "Knowledge Worker control interface already active.", string(colorReset))
			openPrompt(user)
		} else {
			getArgCommand(command, user)
		}
	} else if user == "su" {
		fmt.Print(string(colorRed), "# root/~ > ", string(colorReset))
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		text = strings.ToLower(text)
		command := text
		if command == "init" {
			fmt.Println(string(colorRed), "Knowledge Worker control interface already active.", string(colorReset))
			openPrompt(user)
		} else {
			getArgCommand(command, user)
		}
	}
}
func answerQuery(queried string) {
	const colorReset = "\033[0m"

	const colorRed = "\033[31m"
	if strings.Contains(queried, "zion") {
		fmt.Println("The Zion Lattice Interface is Blocked for now")
	} else if strings.Contains(queried, "garble") {
		fmt.Println("BEGINNING OF ENTRY\n----------")
		fmt.Println(string(colorRed), "Garble", string(colorReset), " is a goblin of unusual small stature, clad in not but a few scraps of leather and cloth with black beady eyes and yellowed broken teeth.\nWorking under the employ of Thali Bumwhistle,", string(colorRed), " Garble", string(colorReset), " uses his innate talents of engineering\nand perhaps a bit of magic to craft and repair the machinery used \nby the Faldihrai in their scientific endeavors.", string(colorRed), "Garble", string(colorReset), " is a warm and welcoming creature,\nknowing no hatred or judgement towards anyone. \nHis simple speech is a disguise for his true depth of wisdom and ability. \nHe owes no allegiance to any particular master, but seems to defer to whom he currently works with, \ntaking a back seat to support those around him however he can.")
		fmt.Println("----------\nEND OF ENTRY")
	} else {
		fmt.Println("Error accessing database.")
	}
}
