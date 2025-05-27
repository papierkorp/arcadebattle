// Package internal comment
package internal

import "fmt"

func statusPlayer() {
	// todo: different status for differen states

	alive := true
	if currentPlayer.battlestate.currentHealth <= 0 {
		alive = false
	}

	status := GetGameTextStatus()
	playerStatus := status.Player

	if currentPlayer.name == "" {
		noPlayerMsg := GetGameTextError("noplayer")
		fmt.Println("\n" + noPlayerMsg)
		fmt.Printf("%s: %s\n", playerStatus.State, currentPlayer.state)
	} else {
		fmt.Printf("\n%s: %s\n", playerStatus.Name, currentPlayer.name)
		fmt.Printf("%s: %d\n", playerStatus.Health, currentPlayer.battlestate.currentHealth)
		fmt.Printf("%s: %d\n", playerStatus.Strength, currentPlayer.stats.strength)
		fmt.Printf("%s: %d\n", playerStatus.Speed, currentPlayer.stats.speed)

		fmt.Printf("%s:\n", playerStatus.Skilllist)
		if len(currentPlayer.skilllist) == 0 {
			errMsg := GetGameTextError("noskill")
			fmt.Printf("  %s\n", errMsg)
		} else {
			for _, skill := range currentPlayer.skilllist {
				fmt.Printf("  %s\n", skill)
			}
		}

		fmt.Printf("%s: %d\n", playerStatus.Talentpoints_Total, currentPlayer.talentpointsTotal)
		fmt.Printf("%s: %d\n", playerStatus.Talentpoints_Remaining, currentPlayer.talentpointsRemaining)
		fmt.Printf("%s: %s\n", playerStatus.Difficulty, currentPlayer.difficulty.String())
		fmt.Printf("%s: %d\n", playerStatus.Bosses, currentPlayer.bosses)
		fmt.Printf("%s: %v\n", playerStatus.Alive, alive)
		fmt.Printf("%s: %v\n", playerStatus.State, currentPlayer.state)

		playerBattlestate := currentPlayer.battlestate
		fmt.Printf("%s:\n", playerStatus.ActiveEffectList)
		if len(playerBattlestate.activeEffectsList) == 0 {
			errMsg := GetGameTextError("noactiveeffect")
			fmt.Printf("  %s\n", errMsg)
		} else {
			for _, effect := range playerBattlestate.activeEffectsList {
				fmt.Printf("  %s\n", effect)
			}
		}
	}

}

func statusBoss() {
	status := GetGameTextStatus()
	bossStatus := status.Boss
	bossBattlestate := currentBoss.battlestate

	fmt.Printf("\n%s: %s\n", bossStatus.Name, currentBoss.name)
	fmt.Printf("%s: %d\n", bossStatus.Health, currentBoss.battlestate.currentHealth)
	fmt.Printf("%s: %d\n", bossStatus.Strength, currentBoss.stats.strength)
	fmt.Printf("%s: %d\n", bossStatus.Speed, currentBoss.stats.speed)

	fmt.Printf("%s:\n", bossStatus.Skilllist)
	if len(currentBoss.skilllist) == 0 {
		errMsg := GetGameTextError("noskill")
		fmt.Printf("  %s\n", errMsg)
	} else {
		for _, skill := range currentBoss.skilllist {
			fmt.Printf("  %s\n", skill)
		}
	}

	fmt.Printf("%s:\n", bossStatus.ActiveEffectList)
	if len(bossBattlestate.activeEffectsList) == 0 {
		errMsg := GetGameTextError("noactiveeffect")
		fmt.Printf("  %s\n", errMsg)
	} else {
		for _, effect := range bossBattlestate.activeEffectsList {
			fmt.Printf("  %s\n", effect)
		}
	}
}

func statusBattle() {

}

func statusDead() {

}
