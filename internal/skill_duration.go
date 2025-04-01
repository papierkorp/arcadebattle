package internal

import (
	"fmt"
	"strconv"
	"strings"
)

type DurationSkill struct {
	id                    int
	name                  string
	dmgmulti              float32
	duration              int
	effectList            []SkillEffect
	talentpointcoststotal int
	upgrade               func()
}

func (ds DurationSkill) String() string {
	effectsWithDuration := make([]string, 0, len(ds.effectList))
	for _, effect := range ds.effectList {
		effectWithDuration := fmt.Sprintf("%s (%d)", effect.name, ds.duration)
		effectsWithDuration = append(effectsWithDuration, effectWithDuration)
	}
	effectsStr := strings.Join(effectsWithDuration, ", ")

	typeMsg := GetGameTextGameMessage("type")
	damagemultiplierMsg := GetGameTextGameMessage("damagemultiplier")
	effectsMsg := GetGameTextGameMessage("effects")
	talentpointscostsMsg := GetGameTextGameMessage("talentpointscosts")
	durationskillMsg := GetGameTextGameMessage("durationskill")
	separatorMsg := GetGameTextGameMessage("separator")

	//1: fireball (Type: Duration, DMG: 1.5x, : Damage Over Time (3), Stun (5), Talentpointcosts: 64)
	return fmt.Sprintf("%d: %s (%s: %s %s %s: %.1fx %s %s: %s %s %s: %d)",
		ds.id,
		ds.name,
		typeMsg,
		durationskillMsg,
		separatorMsg,
		damagemultiplierMsg,
		ds.dmgmulti,
		separatorMsg,
		effectsMsg,
		effectsStr,
		separatorMsg,
		talentpointscostsMsg,
		ds.talentpointcoststotal)
}

func CreateNewDurationSkill(args []string) (error, *DurationSkill) {
	// Minimum args: new0 skill1 <skilltype2> <name3> <dmgmulti4> <duration5> [effect effect effect6...]

	if current_player.name == "" {
		noPlayerMsg := GetGameTextError("noplayer")
		return fmt.Errorf("%s", noPlayerMsg), nil
	}

	if len(args) < 7 {
		invalidArgsMsg := GetGameTextError("invalidargs")
		newdurationskillMsg := GetGameTextCommand("newdurationskill")
		return fmt.Errorf(invalidArgsMsg+" %s", newdurationskillMsg.Usage), nil
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

	duration, err := strconv.Atoi(args[5])

	if err != nil {
		invalidDurationMsg := GetGameTextError("invalidduration")
		return fmt.Errorf(invalidDurationMsg), nil
	}

	if duration <= 0 {
		durationPositiveMsg := GetGameTextError("durationpositive")
		return fmt.Errorf("%s", durationPositiveMsg), nil
	}

	err, effectList := createEffectList(args, 6)
	if err != nil {
		return err, nil
	}

	skill := &DurationSkill{
		id:                    tempSkillCounter,
		name:                  skillName,
		duration:              duration,
		dmgmulti:              dmgMultiFloat32,
		effectList:            effectList,
		talentpointcoststotal: 0,
	}

	usedTalentpoints := calculateDurationSkillCost(skill)
	err2 := changeTalentpoints(usedTalentpoints)
	if err2 != nil {
		return err2, nil
	}

	skill.talentpointcoststotal = usedTalentpoints

	current_player.skilllist = append(current_player.skilllist, skill)

	tempSkillCounter++
	return nil, skill
}

func (ds *DurationSkill) GetID() int {
	return ds.id
}

func (ds *DurationSkill) GetName() string {
	return ds.name
}

func (ds *DurationSkill) GetEffectList() []SkillEffect {
	return ds.effectList
}

func (ds *DurationSkill) GetTalentPointCostsTotal() int {
	return ds.talentpointcoststotal
}

func (ds *DurationSkill) GetSkillType() string {
	return "duration"
}

func (ds *DurationSkill) GetDamageMultiplier() float32 {
	return ds.dmgmulti
}

func (ds *DurationSkill) Use(user *Player, target *Boss) error {
	// todo rework
	// ---------------------------------------------------------------------------------
	// --------------------- EXAMPLE IMPLEMENTATION OF AI - REWORK ---------------------
	// ---------------------------------------------------------------------------------

	// Calculate base damage using power and damage multiplier
	baseDamage := int(float32(user.stats.power) * ds.dmgmulti)

	// Apply damage to target
	target.stats.health -= baseDamage

	// Log damage dealt
	damagedealtMsg := GetGameTextBattle("damagedealt")
	damageMsg := GetGameTextBattle("damage")
	fmt.Printf("%s %d %s\n", damagedealtMsg, baseDamage, damageMsg)

	// Apply effects (simplified)
	for _, effect := range ds.effectList {
		// Create an active effect in target's battlestate
		activeEffect := ActiveEffect{
			skill_effect: effect,
			total_power:  float32(user.stats.power) * ds.dmgmulti,
			turns_left:   ds.duration,
		}

		// Add to target's active effects
		target.battlestate.active_effects_list = append(target.battlestate.active_effects_list, activeEffect)
	}

	return nil
}
