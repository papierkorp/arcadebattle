package internal

import "fmt"

func statusPlayer() {
	alive := true
	if current_player.stats.health <= 0 {
		alive = false
	}

	status := GetGameTextStatus()
	playerStatus := status.Player

	if current_player.name == "" {
		noPlayerMsg := GetGameTextError("noplayer")
		fmt.Println("\n" + noPlayerMsg)
		fmt.Printf("%s: %s\n", playerStatus.State, current_player.state)
	} else {
		fmt.Printf("\n%s: %s\n", playerStatus.Name, current_player.name)
		fmt.Printf("%s: %d\n", playerStatus.Health, current_player.stats.health)
		fmt.Printf("%s: %d\n", playerStatus.Power, current_player.stats.power)
		fmt.Printf("%s: %d\n", playerStatus.Speed, current_player.stats.speed)

		fmt.Printf("%s:\n", playerStatus.Skilllist)
		if len(current_player.skilllist) == 0 {
			errMsg := GetGameTextError("noskill")
			fmt.Printf("  %s\n", errMsg)
		} else {
			for _, skill := range current_player.skilllist {
				fmt.Printf("  %s\n", skill)
			}
		}

		fmt.Printf("%s: %d\n", playerStatus.Talentpoints_Total, current_player.talentpoints_total)
		fmt.Printf("%s: %d\n", playerStatus.Talentpoints_Remaining, current_player.talentpoints_remaining)
		fmt.Printf("%s: %s\n", playerStatus.Difficulty, current_player.difficulty.String())
		fmt.Printf("%s: %d\n", playerStatus.Bosses, current_player.bosses)
		fmt.Printf("%s: %v\n", playerStatus.Alive, alive)
		fmt.Printf("%s: %v\n", playerStatus.State, current_player.state)
	}

}

func statusBoss() {
	status := GetGameTextStatus()
	bossStatus := status.Boss

	fmt.Printf("\n%s: %s\n", bossStatus.Name, current_boss.name)
	fmt.Printf("%s: %d\n", bossStatus.Health, current_boss.stats.health)
	fmt.Printf("%s: %d\n", bossStatus.Power, current_boss.stats.power)
	fmt.Printf("%s: %d\n", bossStatus.Speed, current_boss.stats.speed)

	fmt.Printf("%s:\n", bossStatus.Skilllist)
	if len(current_boss.skilllist) == 0 {
		errMsg := GetGameTextError("noskill")
		fmt.Printf("  %s\n", errMsg)
	} else {
		for _, skill := range current_boss.skilllist {
			fmt.Printf("  %s\n", skill)
		}
	}
}

func statusBattle() {

}

func statusDead() {

}
