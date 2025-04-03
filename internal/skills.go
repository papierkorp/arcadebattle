package internal

import (
	"fmt"
	"strings"
)

// -------------------------------------------------------------------------
// -------------------------------define skill------------------------------
// -------------------------------------------------------------------------

type Skill interface {
	GetID() int
	GetName() string
	GetEffectList() []SkillEffect
	GetTalentPointCostsTotal() int
	GetSkillType() string
	GetDamageMultiplier() float32
	Use() error
	String() string
}

func createEffectList(args []string, skilltype string, startIndex int) (error, []SkillEffect) {
	effectList := []SkillEffect{}

	if len(args) <= startIndex {
		// maybe add a message that no effects were selected? idk
		return nil, effectList
	}

	for i := startIndex; i < len(args); i++ {
		effectName := strings.ToLower(args[i])

		err, effect := newSkillEffect(skilltype, effectName)
		if err != nil {
			return err, nil
		}

		effectList = append(effectList, effect)
	}

	return nil, effectList
}

func useCommand(commandArgs []string) bool {
	if commandArgs[1] == "skill" && len(commandArgs) > 1 {
		skillName := commandArgs[2]

		useSkill(skillName)

	} else {
		invalidargsMsg := GetGameTextError("invalidargs")
		useSkillCommand := GetGameTextCommand("useskill")

		fmt.Println(invalidargsMsg)
		fmt.Println(useSkillCommand)
		return false
	}

	return true
}

func useSkill(skillName string) error {
	//todo implement skill usage
	searchName := strings.ToLower(skillName)
	skillNameValid := false

	for _, skill := range current_player.skilllist {
		if strings.ToLower(skill.GetName()) == searchName {
			skillNameValid = true
			// Here you would call the skill's usage function
			// This will depend on your implementation of skill effects
		}
	}

	if !skillNameValid {
		invalidskillnameMsg := GetGameTextError("invalidskillname")
		return fmt.Errorf(invalidskillnameMsg)
	}

	useskillMsg := GetGameTextBattle("useskill")
	fmt.Printf("%s: %s\n", useskillMsg, skillName)

	if current_boss.stats.health <= 0 {
		handleBossDefeat()
	}

	return nil
}
