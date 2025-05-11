// Package internal comment
package internal

import (
	"fmt"
	"strings"
)

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
		currentPlayer = *newPlayer
		playercreatedMsg := GetGameTextGameMessage("playercreated")
		fmt.Println(playercreatedMsg)
	case "skill":
		invalidskillcreationMsg := GetGameTextError("invalidskillcreation")

		skill, err := CreateNewSkill(commandArgs)

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
		switch currentPlayer.state.String() {
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

func useCommand(commandArgs []string) bool {
	if len(commandArgs) < 3 || commandArgs[1] != "skill" {
		invalidargsMsg := GetGameTextError("invalidargs")
		useSkillMsg := GetGameTextCommand("useskill")
		fmt.Println(invalidargsMsg)
		fmt.Println(useSkillMsg.Usage)
		return false
	}

	if currentPlayer.state != battle {
		noskillusageMsg := GetGameTextError("noskillusage")
		fmt.Println(noskillusageMsg)
	}

	skillName := strings.ToLower(commandArgs[2])
	var foundSkill Skill
	skillFound := false

	for _, skill := range currentPlayer.skilllist {
		if strings.ToLower(skill.GetName()) == skillName {
			skillFound = true
			foundSkill = skill
			break
		}
	}

	if !skillFound {
		invalidSkillNameMsg := GetGameTextError("invalidskillname")
		fmt.Println(invalidSkillNameMsg)
		return false
	}

	useskillMsg := GetGameTextBattle("useskill")
	fmt.Printf("%s: %s\n", useskillMsg, foundSkill.GetName())

	err := foundSkill.Use("player")
	if err != nil {
		fmt.Println(err)
		return false
	}

	currentPlayer.SetBattlePhase(turnEnd)

	updateTurnOrderCurrentTurn()
	checkCurrentTurn()

	return true
}
