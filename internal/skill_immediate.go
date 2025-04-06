// Package internal comment
package internal

import (
	"fmt"
	"strconv"
	"strings"
)

// ImmediateSkill comment
type ImmediateSkill struct {
	id                    int
	name                  string
	dmgmulti              float32
	effectList            []SkillEffect
	talentpointcoststotal int
}

// CreateNewImmediateSkill comment
func CreateNewImmediateSkill(args []string) (*ImmediateSkill, error) {
	// Minimum args: new0 skill1 immediate2 <name3> <dmgmulti4> [effect effect effect5...]
	if current_player.name == "" {
		noPlayerMsg := GetGameTextError("noplayer")
		return nil, fmt.Errorf("%s", noPlayerMsg)
	}
	if len(args) < 5 {
		invalidArgsMsg := GetGameTextError("invalidargs")
		newimmediateskillMsg := GetGameTextCommand("newimmediateskill")
		return nil, fmt.Errorf(invalidArgsMsg+" %s", newimmediateskillMsg.Usage)
	}
	skillName := args[3]
	if skillName == "" {
		emptySkillNameMsg := GetGameTextError("emtpyskillname")
		return nil, fmt.Errorf("%s", emptySkillNameMsg)
	}
	dmgMulti, err := strconv.ParseFloat(args[4], 32)
	if err != nil {
		invalidDmgMultiMsg := GetGameTextError("invaliddmgmulti")
		return nil, fmt.Errorf("%s: %v", invalidDmgMultiMsg, err)
	}
	dmgMultiFloat32 := float32(dmgMulti)
	effectList, err := createEffectList(args, "immediate", 5)
	if err != nil {
		return nil, err
	}
	skill := &ImmediateSkill{
		id:                    tempSkillCounter,
		name:                  skillName,
		dmgmulti:              dmgMultiFloat32,
		effectList:            effectList,
		talentpointcoststotal: 0,
	}
	usedTalentpoints := calculateImmediateSkillCost(skill)
	err = changeTalentpoints(usedTalentpoints)
	if err != nil {
		return nil, err
	}
	skill.talentpointcoststotal = usedTalentpoints
	current_player.skilllist = append(current_player.skilllist, skill)
	tempSkillCounter++
	return skill, nil
}

// GetID comment
func (is *ImmediateSkill) GetID() int {
	return is.id
}

// GetName comment
func (is *ImmediateSkill) GetName() string {
	return is.name
}

// GetEffectList comment
func (is *ImmediateSkill) GetEffectList() []SkillEffect {
	return is.effectList
}

// GetTalentPointCostsTotal comment
func (is *ImmediateSkill) GetTalentPointCostsTotal() int {
	return is.talentpointcoststotal
}

// String comment
func (is ImmediateSkill) String() string {
	effectsList := make([]string, 0, len(is.effectList))
	for _, effect := range is.effectList {
		effectsList = append(effectsList, effect.name)
	}
	effectsStr := strings.Join(effectsList, ", ")

	typeMsg := GetGameTextGameMessage("type")
	damagemultiplierMsg := GetGameTextGameMessage("damagemultiplier")
	effectsMsg := GetGameTextGameMessage("effects")
	talentpointscostsMsg := GetGameTextGameMessage("talentpointscosts")
	immediateskillMsg := GetGameTextGameMessage("immediateskill")
	separatorMsg := GetGameTextGameMessage("separator")

	// 1: Fireball (Type: Immediate, DMG: 1.5x, Effects: Direct Damage, Pierce, Talentpointcosts: 45)
	return fmt.Sprintf("%d: %s (%s: %s %s %s: %.1fx %s %s: %s %s %s: %d)",
		is.id,
		is.name,
		typeMsg,
		immediateskillMsg,
		separatorMsg,
		damagemultiplierMsg,
		is.dmgmulti,
		separatorMsg,
		effectsMsg,
		effectsStr,
		separatorMsg,
		talentpointscostsMsg,
		is.talentpointcoststotal)
}

// GetSkillType comment
func (is *ImmediateSkill) GetSkillType() string {
	return "immediate"
}

// GetDamageMultiplier comment
func (is *ImmediateSkill) GetDamageMultiplier() float32 {
	return is.dmgmulti
}

// Use comment
func (is *ImmediateSkill) Use(source string) error {
	// todo rework
	// ---------------------------------------------------------------------------------
	// --------------------- EXAMPLE IMPLEMENTATION OF AI - REWORK ---------------------
	// ---------------------------------------------------------------------------------

	user := current_player
	target := current_boss

	// Calculate damage
	baseDamage := int(float32(user.stats.power) * is.dmgmulti)

	// Apply damage to target
	target.stats.health -= baseDamage

	// Log damage dealt
	damagedealtMsg := GetGameTextBattle("damagedealt")
	damageMsg := GetGameTextBattle("damage")
	fmt.Printf("%s %d %s\n", damagedealtMsg, baseDamage, damageMsg)

	// Apply immediate effects (simplified)
	for _, effect := range is.effectList {
		// Just call the effect's usage function
		effect.usage()
	}

	return nil
}
