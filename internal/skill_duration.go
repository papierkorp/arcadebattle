// Package internal comment
package internal

import (
	"fmt"
	"strconv"
	"strings"
)

// DurationSkill comment
type DurationSkill struct {
	id                    int
	name                  string
	dmgmulti              float32
	duration              int
	effectList            []SkillEffect
	talentpointcoststotal int
	upgrade               func()
}

// String comment
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

// CreateNewDurationSkill comment
func CreateNewDurationSkill(args []string) (*DurationSkill, error) {
	// Minimum args: new0 skill1 <skilltype2> <name3> <dmgmulti4> <duration5> [effect effect effect6...]
	if current_player.name == "" {
		noPlayerMsg := GetGameTextError("noplayer")
		return nil, fmt.Errorf("%s", noPlayerMsg)
	}
	if len(args) < 7 {
		invalidArgsMsg := GetGameTextError("invalidargs")
		newdurationskillMsg := GetGameTextCommand("newdurationskill")
		return nil, fmt.Errorf(invalidArgsMsg+" %s", newdurationskillMsg.Usage)
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
	duration, err := strconv.Atoi(args[5])
	if err != nil {
		invalidDurationMsg := GetGameTextError("invalidduration")
		return nil, fmt.Errorf(invalidDurationMsg)
	}
	if duration <= 0 {
		durationPositiveMsg := GetGameTextError("durationpositive")
		return nil, fmt.Errorf("%s", durationPositiveMsg)
	}
	effectList, err := createEffectList(args, "duration", 6)
	if err != nil {
		return nil, err
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
func (ds *DurationSkill) GetID() int {
	return ds.id
}

// GetName comment
func (ds *DurationSkill) GetName() string {
	return ds.name
}

// GetEffectList comment
func (ds *DurationSkill) GetEffectList() []SkillEffect {
	return ds.effectList
}

// GetTalentPointCostsTotal comment
func (ds *DurationSkill) GetTalentPointCostsTotal() int {
	return ds.talentpointcoststotal
}

// GetSkillType comment
func (ds *DurationSkill) GetSkillType() string {
	return "duration"
}

// GetDamageMultiplier comment
func (ds *DurationSkill) GetDamageMultiplier() float32 {
	return ds.dmgmulti
}

// Use comment
func (ds *DurationSkill) Use(s string) error {
	// todo rework
	var source Entity
	var target Entity

	switch s {
	case "player":
		source = &current_player
		target = &current_boss
	case "boss":
		source = &current_boss
		target = &current_player
	default:
		invalidEntityMsg := GetGameTextError("invalidentity")
		internalErrorMsg := GetGameTextError("internal")
		return fmt.Errorf("%s - %s", internalErrorMsg, invalidEntityMsg)
	}

	//-------------- Handle Damage --------------

	fmt.Printf("baseDamage: power * skillmutli: %d * %.1fx", source.GetStats().power, ds.dmgmulti)
	baseDamageSource := int(float32(source.GetStats().power) * ds.dmgmulti)
	target.ApplyDamage(baseDamageSource)

	//-------------- Apply Effect --------------

	for _, effect := range ds.effectList {
		var effectTarget Entity
		if effect.selfTarget {
			effectTarget = source
		} else {
			effectTarget = target
		}

		isBlocked := false
		for _, activeEffect := range effectTarget.GetBattleState().activeEffectsList {
			for _, blockedBy := range activeEffect.skillEffect.isBlockedBy {
				if effect.name == blockedBy.name {
					isBlocked = true

					effectBlockedMsg := GetGameTextBattle("effectblocked")
					fmt.Printf("%s (%s, %s)\n", effectBlockedMsg,
						effect.name, blockedBy.name)
					break
				}
			}
			if isBlocked {
				break
			}
		}

		if !isBlocked {
			// apply effect
			newEffect := ActiveEffect{
				skillEffect: effect,
				totalPower:  baseDamageSource,
				turnsLeft:   ds.duration,
			}

			fmt.Println(newEffect)
		}

	}

	return nil
}
