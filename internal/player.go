package internal

import (
	"fmt"
	"strconv"
)

// todo talentpoints_total, talentpoints_remaining
type Player struct {
	name                   string
	stats                  stats
	talentpoints_total     int
	talentpoints_remaining int
	state                  state
	difficulty             difficulty
	bosses                 int
	skilllist              []Skill
	battlestate            BattleState
}

type stats struct {
	health int
	power  int
	speed  int
}

func createPlayer(args []string) (*Player, error) {

	newplayerMsg := GetGameTextCommand("newplayer")
	if len(args) != 7 {
		invalidArgsMsg := GetGameTextError("invalidargs")
		return nil, fmt.Errorf(invalidArgsMsg, newplayerMsg.Usage)
	}

	diff, err := ParseDifficulty(args[3])
	if err != nil {
		return nil, err
	}

	initialTalentpoints := getTalentpoints(diff)

	health, err := strconv.Atoi(args[4])
	if err != nil {
		invalidHealthMsg := GetGameTextError("invalidhealth")
		return nil, fmt.Errorf(invalidHealthMsg, err)
	}

	power, err := strconv.Atoi(args[5])
	if err != nil {
		invalidPowerMsg := GetGameTextError("invalidpower")
		return nil, fmt.Errorf(invalidPowerMsg, err)
	}

	speed, err := strconv.Atoi(args[6])
	if err != nil {
		invalidSpeedMsg := GetGameTextError("invalidspeed")
		return nil, fmt.Errorf(invalidSpeedMsg, err)
	}

	stats := stats{
		health: health,
		power:  power,
		speed:  speed,
	}

	usedTalentpoints := calculateStatsCost(stats)

	if usedTalentpoints > initialTalentpoints {
		exceededMsg := GetGameTextError("exceededtalentpoints")
		separatorMsg := GetGameTextGameMessage("separator")
		usedMsg := GetGameTextGameMessage("used")
		availableMsg := GetGameTextGameMessage("available")
		return nil, fmt.Errorf("%s, %s: %d %s %s: %d",
			exceededMsg,         // Used too much Talentpoints
			availableMsg,        // Available:
			initialTalentpoints, // available value
			separatorMsg,        //  |
			usedMsg,             // Used:
			usedTalentpoints)    // used value
	} else {
		beforeMsg := GetGameTextGameMessage("talentpointsbefore")
		afterMsg := GetGameTextGameMessage("talentpointsafter")
		usedMsg := GetGameTextGameMessage("talentpointsused")
		remainingTalentpoints := initialTalentpoints - usedTalentpoints

		fmt.Printf("\n%s: %d\n%s: %d\n%s: %d\n\n",
			beforeMsg, initialTalentpoints,
			usedMsg, usedTalentpoints,
			afterMsg, remainingTalentpoints)
	}

	player := &Player{
		name:                   args[2],
		stats:                  stats,
		talentpoints_total:     initialTalentpoints,
		talentpoints_remaining: initialTalentpoints - usedTalentpoints,
		state:                  idle,
		difficulty:             diff,
		bosses:                 0,
	}

	return player, nil
}
