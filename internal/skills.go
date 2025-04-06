// Package internal comment
package internal

import (
	"fmt"
	"strings"
)

// -------------------------------------------------------------------------
// -------------------------------define skill------------------------------
// -------------------------------------------------------------------------

// Skill comment
type Skill interface {
	GetID() int
	GetName() string
	GetEffectList() []SkillEffect
	GetTalentPointCostsTotal() int
	GetSkillType() string
	GetDamageMultiplier() float32
	Use(source string) error
	String() string
}

func createEffectList(args []string, skilltype string, startIndex int) ([]SkillEffect, error) {
	effectList := []SkillEffect{}

	if len(args) <= startIndex {
		// maybe add a message that no effects were selected? idk
		return effectList, nil
	}

	for i := startIndex; i < len(args); i++ {
		effectName := strings.ToLower(args[i])

		effect, err := newSkillEffect(skilltype, effectName)
		if err != nil {
			return nil, err
		}

		effectList = append(effectList, effect)
	}

	return effectList, nil
}

func useCommand(commandArgs []string) bool {
	if len(commandArgs) < 3 || commandArgs[1] != "skill" {
		invalidargsMsg := GetGameTextError("invalidargs")
		useSkillMsg := GetGameTextCommand("useskill")
		fmt.Println(invalidargsMsg)
		fmt.Println(useSkillMsg.Usage)
		return false
	}

	skillName := strings.ToLower(commandArgs[2])
	var foundSkill Skill
	skillFound := false

	for _, skill := range current_player.skilllist {
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

	if current_boss.stats.health <= 0 {
		handleBossDefeat()
	}

	return true
}
