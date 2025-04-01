package internal

import (
	"fmt"
	"strings"
)

type difficulty int

const (
	normal difficulty = iota + 1
	hard
	expert
	master
	torment
)

func (d difficulty) String() string {
	difficulties := GetGameTextDifficulty()

	switch d {
	case normal:
		return difficulties.Normal.Name
	case hard:
		return difficulties.Hard.Name
	case expert:
		return difficulties.Expert.Name
	case master:
		return difficulties.Master.Name
	case torment:
		return difficulties.Torment.Name
	default:
		return difficulties.Normal.Name
	}
}

func (d difficulty) EnumDifficultyIndex() int {
	return int(d)
}

func ParseDifficulty(s string) (difficulty, error) {
	errorMsg := GetGameTextError("invaliddifficulty")
	switch strings.ToLower(s) {
	case "normal":
		return normal, nil
	case "hard":
		return hard, nil
	case "expert":
		return expert, nil
	case "master":
		return master, nil
	case "torment":
		return torment, nil
	default:
		return normal, fmt.Errorf(errorMsg+": %s", s)
	}
}

func setDifficulty(p *Player, difficulty difficulty) error {
	p.difficulty = difficulty
	return nil
}
