// Package internal comment
package internal

import (
	"fmt"
	"math/rand"
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

func createEffectList(args []string) ([]SkillEffect, error) {
	effectList := []SkillEffect{}

	// 5 = id where the effectName arg starts
	for i := 5; i < len(args); i++ {
		effectName := strings.ToLower(args[i])

		effect, err := newSkillEffect(effectName)
		if err != nil {
			return nil, err
		}

		effectList = append(effectList, effect)
	}

	return effectList, nil
}

// String comment
func (s Skill) String() string {
	damagemultiplierMsg := GetGameTextGameMessage("damagemultiplier")
	effectsMsg := GetGameTextGameMessage("effects")
	talentpointscostsMsg := GetGameTextGameMessage("talentpointscosts")
	separatorMsg := GetGameTextGameMessage("separator")
	effectsList := make([]string, 0, len(s.effectList))
	for _, effect := range s.effectList {
		effect := fmt.Sprintf("%s", effect.displayName)
		effectsList = append(effectsList, effect)
	}
	effectsStr := strings.Join(effectsList, ", ")

	//1: lifeleech (Damage Multiplier: 0.7x | Effects: Lifeleech | Talentpointcosts: 0)
	return fmt.Sprintf("%d: %s (%s: %.1fx %s %s: %s %s %s: %d)",
		s.id,
		s.name,
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
	// Minimum args: new0 skill1 <name2> <dmgmulti3> <duration4> <[effect effect effect5...]>
	if currentPlayer.name == "" {
		noPlayerMsg := GetGameTextError("noplayer")
		return nil, fmt.Errorf("%s", noPlayerMsg)
	}

	if len(args) < 6 {
		invalidArgsMsg := GetGameTextError("invalidargs")
		newSkillMsg := GetGameTextCommand("newskill")
		return nil, fmt.Errorf(invalidArgsMsg+" %s", newSkillMsg.Usage)
	}

	skillName := args[2]
	if skillName == "" {
		emptySkillNameMsg := GetGameTextError("emtpyskillname")
		return nil, fmt.Errorf("%s", emptySkillNameMsg)
	}

	dmgMulti, err := strconv.ParseFloat(args[3], 32)
	if err != nil {
		invalidDmgMultiMsg := GetGameTextError("invaliddmgmulti")
		return nil, fmt.Errorf("%s: %v", invalidDmgMultiMsg, err)
	}
	dmgMultiFloat32 := float32(dmgMulti)

	duration, err := strconv.Atoi(args[4])
	if err != nil {
		invalidDurationMsg := GetGameTextError("invalidduration")
		return nil, fmt.Errorf(invalidDurationMsg)
	}
	if duration <= 0 {
		durationPositiveMsg := GetGameTextError("durationpositive")
		return nil, fmt.Errorf("%s", durationPositiveMsg)
	}
	effectList, err := createEffectList(args)
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
	currentPlayer.skilllist = append(currentPlayer.skilllist, *skill)
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

// is executed with the use skill command
func (s *Skill) Use(skillSource string) error {
	var source Entity
	var target Entity

	switch skillSource {
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

	//-------------- Check if Effect is blocked --------------
	//todo: loop if an effect is blocked / target is changed

	effectIsBlocked := false

	for _, effect := range s.effectList {
		if effect.usageTiming == etiOnSkillStart {
			if effect.category == ecaBlockBuffs {
				// & s.effect == buff
				effectIsBlocked = true
			}
		}
	}

	if effectIsBlocked {
		//todo: return blocked message
		fmt.Println("effect blocked")
		return nil
	}

	//-------------- Calculate Power --------------

	fullSkillPower, err := calculateSkillPower(s, source, target)

	if err != nil {
		internalErrorMsg := GetGameTextError("internal")
		invalidskillpowercalculationMsg := GetGameTextError("invalidskillpowercalculation")
		return fmt.Errorf("%s - %s (%s)", internalErrorMsg, invalidskillpowercalculationMsg, err)
	}

	//-------------- Apply Effect --------------

	for _, effect := range s.effectList {
		// var effectTarget Entity
		// if effect.effectType == etyBuff {
		// 	effectTarget = source
		// } else {
		// 	// if effecttype == debuff
		// 	effectTarget = target
		// }

		isBlocked := false
		// effectBlockedMsg := GetGameTextBattle("effectblocked")

		if !isBlocked {
			newEffect := ActiveEffect{
				skillEffect: effect,
				totalPower:  fullSkillPower,
				turnsLeft:   s.duration,
				source:      source,
				target:      target,
			}

			fmt.Println("effect.effectType: ", effect.effectType)

			if effect.effectType == etyBuff {
				source.SetFullSkillPower(fullSkillPower)
				source.AddActiveEffect(newEffect)
			} else if effect.effectType == etyDebuff {
				source.SetFullSkillPower(fullSkillPower)
				target.AddActiveEffect(newEffect)
			}

		}

		//-------------- Handle Damage --------------

		target.ApplyDamage(fullSkillPower)

	}

	return nil
}

func calculateSkillPower(s *Skill, source Entity, target Entity) (int, error) {
	var basicSkillPower float32 = 0.0
	calculatedFullSkillPower := basicSkillPower

	for _, effect := range s.effectList {
		if effect.usageTiming == etiOnSkillCalculation {
			//todo: calculate power increase/reduction
			newPower := float32(source.GetStats().power)/5*effect.multi + 1
			basicSkillPower = newPower * s.dmgmulti

			if rand.Float32() <= effect.probability {
				calculatedFullSkillPower = calculatedFullSkillPower * effect.multi
			}
		}
	}

	for _, effect := range s.effectList {
		if effect.usageTiming == etiOnSkillCalculation {
			//todo: calculate power increase/reduction
			basicSkillPower = float32(source.GetStats().power) * s.dmgmulti
			if rand.Float32() <= effect.probability {
				calculatedFullSkillPower = calculatedFullSkillPower * effect.multi
			}
		}
	}

	fullSkillPower := int(calculatedFullSkillPower)
	return fullSkillPower, nil
}
