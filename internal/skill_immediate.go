package internal

import (
	"fmt"
	"strconv"
	"strings"
)

type ImmediateSkill struct {
	id                    int
	name                  string
	dmgmulti              float32
	effectList            []SkillEffect
	talentpointcoststotal int
}

func CreateNewImmediateSkill(args []string) (error, *ImmediateSkill) {
	// Minimum args: new0 skill1 immediate2 <name3> <dmgmulti4> [effect effect effect5...]

	if current_player.name == "" {
		noPlayerMsg := GetGameTextError("noplayer")
		return fmt.Errorf("%s", noPlayerMsg), nil
	}

	if len(args) < 5 {
		invalidArgsMsg := GetGameTextError("invalidargs")
		newimmediateskillMsg := GetGameTextCommand("newimmediateskill")
		return fmt.Errorf(invalidArgsMsg+" %s", newimmediateskillMsg.Usage), nil
	}

	skillName := args[3]
	if skillName == "" {
		emptySkillNameMsg := GetGameTextError("emtpyskillname")
		return fmt.Errorf("%s", emptySkillNameMsg), nil
	}

	dmgMulti, err := strconv.ParseFloat(args[4], 32)
	if err != nil {
		invalidDmgMultiMsg := GetGameTextError("invaliddmgmulti")
		return fmt.Errorf("%s: %v", invalidDmgMultiMsg, err), nil
	}
	dmgMultiFloat32 := float32(dmgMulti)

	err, effectList := createEffectList(args, 5)
	if err != nil {
		return err, nil
	}

	skill := &ImmediateSkill{
		id:                    tempSkillCounter,
		name:                  skillName,
		dmgmulti:              dmgMultiFloat32,
		effectList:            effectList,
		talentpointcoststotal: 0,
	}

	usedTalentpoints := calculateImmediateSkillCost(skill)
	err2 := changeTalentpoints(usedTalentpoints)
	if err2 != nil {
		return err2, nil
	}

	skill.talentpointcoststotal = usedTalentpoints

	current_player.skilllist = append(current_player.skilllist, skill)

	tempSkillCounter++
	return nil, skill
}

func (is *ImmediateSkill) GetID() int {
	return is.id
}

func (is *ImmediateSkill) GetName() string {
	return is.name
}

func (is *ImmediateSkill) GetEffectList() []SkillEffect {
	return is.effectList
}

func (is *ImmediateSkill) GetTalentPointCostsTotal() int {
	return is.talentpointcoststotal
}

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

func (is *ImmediateSkill) GetSkillType() string {
	return "immediate"
}

func (is *ImmediateSkill) GetDamageMultiplier() float32 {
	return is.dmgmulti
}

func (is *ImmediateSkill) Use(user *Player, target *Boss) error {
	// todo rework
	// ---------------------------------------------------------------------------------
	// --------------------- EXAMPLE IMPLEMENTATION OF AI - REWORK ---------------------
	// ---------------------------------------------------------------------------------

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
