// Package internal comment
package internal

import (
	"fmt"
	"strings"

	"github.com/chzyer/readline"
)

// Start comment
func Start() {
	err := initGametext()

	if err != nil {
		fmt.Println(err)
	}

	welcomeMsg := GetGameTextGameMessage("welcome")
	gamestartMsg := GetGameTextGameMessage("gamestart")
	gamestarthelpMsg := GetGameTextGameMessage("gamestarthelp")

	promptMsg := GetGameTextGameMessage("prompt")

	var initErr error
	rl, initErr = readline.New(promptMsg)
	if initErr != nil {
		panic(initErr)
	}
	defer rl.Close()

	fmt.Println(welcomeMsg + "\n")

	fmt.Println(gamestartMsg)
	fmt.Println(gamestarthelpMsg)

	gameLoop()
}

func gameLoop() {
	invalidCommandMsg := GetGameTextError("invalidcommand")
	goodbyeMsg := GetGameTextGameMessage("goodbye")
	gamestarthelpMsg := GetGameTextGameMessage("gamestarthelp")

	for {
		input, err := rl.Readline()
		if err != nil {
			break
		}
		commandArgs := strings.Fields(strings.ToLower(input))

		if len(commandArgs) == 0 {
			continue
		}
		command := commandArgs[0]

		switch command {
		case "help", "?":
			if len(commandArgs) > 1 {
				helpSpecificCommand(commandArgs[1])
			} else {
				helpCommand()
			}
		case "new":
			newCommand(commandArgs)
		case "upgrade":
			upgradeCommand(commandArgs)
		case "status":
			statusCommand(commandArgs)
		case "quit", "exit":
			fmt.Println(goodbyeMsg)
			return
		case "battle":
			startBattle()
		case "test":
			if len(commandArgs) > 1 {
				RunTest(commandArgs[1])
			} else {
				fmt.Println("Please specify a test case")
			}
		default:
			fmt.Println(invalidCommandMsg)
			fmt.Println(gamestarthelpMsg)
		}

		CheckCurrentState(current_player.state)
	}
}
