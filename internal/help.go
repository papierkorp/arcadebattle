package internal

import "fmt"

// ------------------------------------------------------------
// --------------------------- COMMAND DOCS ---------------------------
// ------------------------------------------------------------

type CommandDoc struct {
	Name        string
	Description string
	Detail      string
	Usage       string
	Example     string
}

func helpCommand() {
	headerMsg := GetGameTextHelp("header")
	usageMsg := GetGameTextHelp("usage")
	exampleMsg := GetGameTextHelp("example")
	noteMsg := GetGameTextHelp("note")
	footerMsg := GetGameTextHelp("footer")
	availablecommandsMsg := GetGameTextGameMessage("availablecommands")

	commandDocs := GetAllGameTextCommands()

	maxNameLen := 0 // Get longest command name for padding
	for _, cmd := range commandDocs {
		if len(cmd.Name) > maxNameLen {
			maxNameLen = len(cmd.Name)
		}
	}

	fmt.Print("\n" + headerMsg + "\n\n")
	fmt.Println(availablecommandsMsg + "\n")

	for _, cmd := range commandDocs {
		fmt.Printf("%-*s", maxNameLen+2, cmd.Name)
		fmt.Printf("  %s\n", cmd.Description)
		fmt.Printf("%-*s  %s: %s\n", maxNameLen+2, "", usageMsg, cmd.Usage)
		fmt.Printf("%-*s  %s: %s\n\n", maxNameLen+2, "", exampleMsg, cmd.Example)
	}

	fmt.Print(noteMsg)
	fmt.Print("\n\n" + footerMsg + "\n")
}

func helpSpecificCommand(command string) {
	cmd := GetGameTextCommand(command)

	headerMsg := GetGameTextHelp("header")
	commandMsg := GetGameTextHelp("command")
	descriptionMsg := GetGameTextHelp("description")
	detailMsg := GetGameTextHelp("detail")
	usageMsg := GetGameTextHelp("usage")
	exampleMsg := GetGameTextHelp("example")
	footerMsg := GetGameTextHelp("footer")

	fmt.Print("\n" + headerMsg + "\n\n")
	fmt.Printf("%s: %s\n", commandMsg, cmd.Name)
	fmt.Printf("%s: %s\n", descriptionMsg, cmd.Description)
	fmt.Printf("%s: %s\n", detailMsg, cmd.Detail)
	fmt.Printf("%s: %s\n", usageMsg, cmd.Usage)
	fmt.Printf("%s: %s\n", exampleMsg, cmd.Example)
	fmt.Print("\n" + footerMsg + "\n")
}
