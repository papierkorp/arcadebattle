package internal

import "fmt"

func newCommand(commandArgs []string) {
	newtype := commandArgs[1]

	switch newtype {
	case "player":
		newPlayer, err := createPlayer(commandArgs)
		invalidPlayerCreationMsg := GetGameTextError("invalidplayercreation")

		if err != nil {
			fmt.Println(invalidPlayerCreationMsg+" ", err)
			return
		}
		current_player = *newPlayer
		playercreatedMsg := GetGameTextGameMessage("playercreated")
		fmt.Println(playercreatedMsg)
	case "skill":
		skilltype := commandArgs[2]
		invalidskillcreationMsg := GetGameTextError("invalidskillcreation")
		var err error
		var skill Skill

		switch skilltype {
		case "duration":
			err, skill = CreateNewDurationSkill(commandArgs)
		case "immediate":
			err, skill = CreateNewImmediateSkill(commandArgs)
		case "passive":
			err, skill = CreateNewPassiveSkill(commandArgs)
		default:
			invalidskilltypeMsg := GetGameTextError("invalidskilltype")
			fmt.Println(invalidskillcreationMsg + " - " + invalidskilltypeMsg + "\n")
			return
		}

		if err != nil {
			fmt.Println(invalidskillcreationMsg, err)
			return
		}
		skillcreatedMsg := GetGameTextGameMessage("skillcreated")
		fmt.Println(skillcreatedMsg+"\n", skill)
	}

}

func statusCommand(commandArgs []string) {
	if len(commandArgs) == 0 {
		statusPlayer()
		return
	}

	if len(commandArgs) == 1 {
		switch current_player.state.String() {
		case "Idle":
			statusPlayer()
		case "Battle":
			statusBoss()
		case "Dead":
			statusPlayer()
		default:
			statusPlayer()
		}
		return
	}

	statusType := commandArgs[1]
	switch statusType {
	case "p":
		statusPlayer()
	case "b":
		statusBoss()
	default:
		statusPlayer()
	}
}

func upgradeCommand(commandArgs []string) {
	//todo implement upgradeCommand
	newtype := commandArgs[1]

	switch newtype {
	case "player":
		// same command as new but check the current player stats if upgrade is possible
		fmt.Println("upgrade player")
	case "skill":
		// same command as new but check the skillname/ide if upgrade is possible
		fmt.Println("upgrade skill")
	}

}
