package internal

import (
	"fmt"
	"strings"
)

type PassiveSkill struct {
	id                    int
	name                  string
	effectList            []SkillEffect
	talentpointcoststotal int
}

func CreateNewPassiveSkill(args []string) (error, *PassiveSkill) {
	// Minimum args: new0 skill1 passive2 <name3> [effect effect effect4...]

	if current_player.name == "" {
		noPlayerMsg := GetGameTextError("noplayer")
		return fmt.Errorf("%s", noPlayerMsg), nil
	}

	if len(args) < 4 {
		invalidArgsMsg := GetGameTextError("invalidargs")
		newpassiveskillMsg := GetGameTextCommand("newpassiveskill")
		return fmt.Errorf(invalidArgsMsg+" %s", newpassiveskillMsg.Usage), nil
	}

	skillName := args[3]
	if skillName == "" {
		emptySkillNameMsg := GetGameTextError("emtpyskillname")
		return fmt.Errorf("%s", emptySkillNameMsg), nil
	}

	err, effectList := createEffectList(args, 4)
	if err != nil {
		return err, nil
	}

	skill := &PassiveSkill{
		id:                    tempSkillCounter,
		name:                  skillName,
		effectList:            effectList,
		talentpointcoststotal: 0,
	}

	usedTalentpoints := calculatePassiveSkillCost(skill)
	err2 := changeTalentpoints(usedTalentpoints)
	if err2 != nil {
		return err2, nil
	}

	skill.talentpointcoststotal = usedTalentpoints

	current_player.skilllist = append(current_player.skilllist, skill)

	tempSkillCounter++
	return nil, skill
}

func (ps *PassiveSkill) GetID() int {
	return ps.id
}

func (ps *PassiveSkill) GetName() string {
	return ps.name
}

func (ps *PassiveSkill) GetEffectList() []SkillEffect {
	return ps.effectList
}

func (ps *PassiveSkill) GetTalentPointCostsTotal() int {
	return ps.talentpointcoststotal
}

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

func (ps *PassiveSkill) GetSkillType() string {
	return "passive"
}

func (ps *PassiveSkill) GetDamageMultiplier() float32 {
	return 0
}

func (ps *PassiveSkill) Use(user *Player, target *Boss) error {
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
