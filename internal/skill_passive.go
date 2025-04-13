// Package internal comment
package internal

import (
	"fmt"
	"strings"
)

// PassiveSkill comment
type PassiveSkill struct {
	id                    int
	name                  string
	effectList            []SkillEffect
	talentpointcoststotal int
}

// CreateNewPassiveSkill comment
func CreateNewPassiveSkill(args []string) (*PassiveSkill, error) {
	// Minimum args: new0 skill1 passive2 <name3> [effect effect effect4...]
	if currentPlayer.name == "" {
		noPlayerMsg := GetGameTextError("noplayer")
		return nil, fmt.Errorf("%s", noPlayerMsg)
	}
	if len(args) < 4 {
		invalidArgsMsg := GetGameTextError("invalidargs")
		newpassiveskillMsg := GetGameTextCommand("newpassiveskill")
		return nil, fmt.Errorf(invalidArgsMsg+" %s", newpassiveskillMsg.Usage)
	}
	skillName := args[3]
	if skillName == "" {
		emptySkillNameMsg := GetGameTextError("emtpyskillname")
		return nil, fmt.Errorf("%s", emptySkillNameMsg)
	}
	effectList, err := createEffectList(args, "passive", 4)
	if err != nil {
		return nil, err
	}
	skill := &PassiveSkill{
		id:                    tempSkillCounter,
		name:                  skillName,
		effectList:            effectList,
		talentpointcoststotal: 0,
	}
	usedTalentpoints := calculatePassiveSkillCost(skill)
	err = changeTalentpoints(usedTalentpoints)
	if err != nil {
		return nil, err
	}
	skill.talentpointcoststotal = usedTalentpoints
	currentPlayer.skilllist = append(currentPlayer.skilllist, skill)
	tempSkillCounter++
	return skill, nil
}

// GetID comment
func (ps *PassiveSkill) GetID() int {
	return ps.id
}

// GetName comment
func (ps *PassiveSkill) GetName() string {
	return ps.name
}

// GetEffectList comment
func (ps *PassiveSkill) GetEffectList() []SkillEffect {
	return ps.effectList
}

// GetTalentPointCostsTotal comment
func (ps *PassiveSkill) GetTalentPointCostsTotal() int {
	return ps.talentpointcoststotal
}

// String comment
func (ps PassiveSkill) String() string {
	effectsList := make([]string, 0, len(ps.effectList))
	for _, effect := range ps.effectList {
		effectsList = append(effectsList, effect.name)
	}
	effectsStr := strings.Join(effectsList, ", ")

	typeMsg := GetGameTextGameMessage("type")
	effectsMsg := GetGameTextGameMessage("effects")
	talentpointscostsMsg := GetGameTextGameMessage("talentpointscosts")
	passiveskillMsg := GetGameTextGameMessage("passiveskill")
	separatorMsg := GetGameTextGameMessage("separator")

	// 1: Resilience (Type: Passive, Effects: Block Debuffs, Damage Reduction, Talentpointcosts: 70)
	return fmt.Sprintf("%d: %s (%s: %s %s %s: %s %s %s: %d)",
		ps.id,
		ps.name,
		typeMsg,
		passiveskillMsg,
		separatorMsg,
		effectsMsg,
		effectsStr,
		separatorMsg,
		talentpointscostsMsg,
		ps.talentpointcoststotal)
}

// GetSkillType comment
func (ps *PassiveSkill) GetSkillType() string {
	return "passive"
}

// GetDamageMultiplier comment
func (ps *PassiveSkill) GetDamageMultiplier() float32 {
	return 0
}

// Use comment
func (ps *PassiveSkill) Use(s string) error {
	// todo rework
	// ---------------------------------------------------------------------------------
	// --------------------- EXAMPLE IMPLEMENTATION OF AI - REWORK ---------------------
	// ---------------------------------------------------------------------------------

	// Passive skills don't deal direct damage or have typical "use" behavior
	// Instead, they apply their effects to the player's stats/state

	// For simplicity in this implementation, we'll just log that the passive is active
	fmt.Printf("Passive skill '%s' is active with effects: ", ps.name)

	for i, effect := range ps.effectList {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(effect.name)
	}
	fmt.Println()

	// In a real implementation, you might modify player stats or set flags
	// This is typically done when the skill is equipped, not when "used"

	return nil
}
