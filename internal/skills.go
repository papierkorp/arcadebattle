// Package internal comment
package internal

import (
	"fmt"
	"strconv"
	"strings"
)

// -------------------------------------------------------------------------
// -------------------------------define skill------------------------------
// -------------------------------------------------------------------------

// Skill comment
type Skill struct {
	id                    int
	name                  string
	dmgmulti              float32
	duration              int
	effectList            []SkillEffect
	talentpointcoststotal int
	upgrade               func()
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

// String comment
func (s Skill) String() string {
	effectsWithDuration := make([]string, 0, len(s.effectList))
	for _, effect := range s.effectList {
		effectWithDuration := fmt.Sprintf("%s (%d)", effect.affectedStat, s.duration)
		effectsWithDuration = append(effectsWithDuration, effectWithDuration)
	}
	effectsStr := strings.Join(effectsWithDuration, ", ")

	damagemultiplierMsg := GetGameTextGameMessage("damagemultiplier")
	effectsMsg := GetGameTextGameMessage("effects")
	talentpointscostsMsg := GetGameTextGameMessage("talentpointscosts")
	separatorMsg := GetGameTextGameMessage("separator")

	//1: fireball (Type: Duration, DMG: 1.5x, : Damage Over Time (3), Stun (5), Talentpointcosts: 64)
	return fmt.Sprintf("%d: %s (%s %s: %.1fx %s %s: %s %s %s: %d)",
		s.id,
		s.name,
		separatorMsg,
		damagemultiplierMsg,
		s.dmgmulti,
		separatorMsg,
		effectsMsg,
		effectsStr,
		separatorMsg,
		talentpointscostsMsg,
		s.talentpointcoststotal)
}

// CreateNewSkill comment
func CreateNewSkill(args []string) (*Skill, error) {
	// Minimum args: new0 skill1 <skilltype2> <name3> <dmgmulti4> <duration5> [effect effect effect6...]
	if currentPlayer.name == "" {
		noPlayerMsg := GetGameTextError("noplayer")
		return nil, fmt.Errorf("%s", noPlayerMsg)
	}
	if len(args) < 7 {
		invalidArgsMsg := GetGameTextError("invalidargs")
		newSkillMsg := GetGameTextCommand("newSkill")
		return nil, fmt.Errorf(invalidArgsMsg+" %s", newSkillMsg.Usage)
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
	skill := &Skill{
		id:                    tempSkillCounter,
		name:                  skillName,
		duration:              duration,
		dmgmulti:              dmgMultiFloat32,
		effectList:            effectList,
		talentpointcoststotal: 0,
	}
	usedTalentpoints := calculateSkillCost(skill)
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
func (ds *Skill) GetID() int {
	return ds.id
}

// GetName comment
func (ds *Skill) GetName() string {
	return ds.name
}

// GetEffectList comment
func (ds *Skill) GetEffectList() []SkillEffect {
	return ds.effectList
}

// GetTalentPointCostsTotal comment
func (ds *Skill) GetTalentPointCostsTotal() int {
	return ds.talentpointcoststotal
}

// GetDamageMultiplier comment
func (ds *Skill) GetDamageMultiplier() float32 {
	return ds.dmgmulti
}

// Use comment
func (ds *Skill) Use(s string) error {
	var source Entity
	var target Entity

	switch s {
	case "player":
		source = GetEntity(PlayerEntity)
		target = GetEntity(BossEntity)
	case "boss":
		source = GetEntity(BossEntity)
		target = GetEntity(PlayerEntity)
	default:
		invalidEntityMsg := GetGameTextError("invalidentity")
		internalErrorMsg := GetGameTextError("internal")
		return fmt.Errorf("%s - %s", internalErrorMsg, invalidEntityMsg)
	}

	//-------------- Handle Damage --------------

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
			newEffect := ActiveEffect{
				skillEffect: effect,
				totalPower:  baseDamageSource,
				turnsLeft:   ds.duration,
				source:      source,
				target:      target,
			}

			target.AddActiveEffect(newEffect)
		}

	}

	return nil
}
