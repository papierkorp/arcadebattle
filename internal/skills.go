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

	rawSkillPower, err := calculateRawSkillPower(s, source, target)

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
				skillEffect:   effect,
				rawSkillPower: rawSkillPower,
				turnsLeft:     s.duration,
				source:        source,
				target:        target,
			}

			fmt.Println("effect.effectType: ", effect.effectType)

			if effect.effectType == etyBuff {
				source.SetLastRawSkillPowerUsed(rawSkillPower)
				source.AddActiveEffect(newEffect)
			} else if effect.effectType == etyDebuff {
				source.SetLastRawSkillPowerUsed(rawSkillPower)
				target.AddActiveEffect(newEffect)
			}
		}

		//-------------- Handle Damage --------------

		calculatedDamage, err := calculateCalculatedDamage(s, rawSkillPower)

		if err != nil {
			internalErrorMsg := GetGameTextError("internal")
			invalidskillpowercalculationMsg := GetGameTextError("invalidskillpowercalculation")
			return fmt.Errorf("%s - %s (%s)", internalErrorMsg, invalidskillpowercalculationMsg, err)
		}

		outgoingDamage, err := calculateOutgoingDamage(s, source, calculatedDamage)

		if err != nil {
			internalErrorMsg := GetGameTextError("internal")
			invalidskillpowercalculationMsg := GetGameTextError("invalidskillpowercalculation")
			return fmt.Errorf("%s - %s (%s)", internalErrorMsg, invalidskillpowercalculationMsg, err)
		}

		incomingDamage := outgoingDamage
		actualDamageTaken, err := calculateActualDamageTaken(s, target, incomingDamage)

		if err != nil {
			internalErrorMsg := GetGameTextError("internal")
			invalidskillpowercalculationMsg := GetGameTextError("invalidskillpowercalculation")
			return fmt.Errorf("%s - %s (%s)", internalErrorMsg, invalidskillpowercalculationMsg, err)
		}

		source.SetLastOutgoingDamage(outgoingDamage)
		target.SetLastIncomingDamage(incomingDamage)
		target.SetLastActualDamageTaken(actualDamageTaken)
		target.ApplyDamage(rawSkillPower)

	}

	return nil
}

func calculateRawSkillPower(s *Skill, source Entity, target Entity) (int, error) {
	//todo: return error
	var basePower float32 = float32(source.GetBattleState().currentPower)
	modifiedPower := basePower

	for _, effect := range s.effectList {
		if effect.usageTiming == etiOnSkillCalculation {
			if effect.category == ecaIncreasePower || effect.category == ecaDecreasePower {
				if rand.Float32() <= effect.probability {
					modifiedPower = float32(basePower/effect.powerRatio*effect.baseValue + 1)
				}
			}
		}
	}

	for _, effect := range source.GetBattleState().activeEffectsList {
		if effect.skillEffect.usageTiming == etiOnSkillCalculation {
			if effect.skillEffect.category == ecaIncreasePower || effect.skillEffect.category == ecaDecreasePower {
				if rand.Float32() <= effect.skillEffect.probability {
					modifiedPower = float32(basePower/effect.skillEffect.powerRatio*effect.skillEffect.baseValue + 1)
				}
			}
		}
	}

	rawSkillPower := int(modifiedPower * s.dmgmulti)

	return rawSkillPower, nil
}

func calculateCalculatedDamage(s *Skill, rawSkillpower int) (int, error) {
	//todo: return error
	calculatedDamage := int(float32(rawSkillpower) * s.dmgmulti)
	return calculatedDamage, nil
}

func calculateOutgoingDamage(s *Skill, source Entity, calculatedDamage int) (int, error) {
	var outgoingDamage int = calculatedDamage

	for _, effect := range source.GetBattleState().activeEffectsList {
		if effect.skillEffect.usageTiming == etiOnSkillCalculation {
			if effect.skillEffect.category == ecaIncreasePower || effect.skillEffect.category == ecaDecreasePower {
				if rand.Float32() <= effect.skillEffect.probability {
					outgoingDamage = float32(basePower/5*effect.skillEffect.baseValue + 1)
				}
			}
		}
	}
	return outgoingDamage, nil
}

func calculateActualDamageTaken(s *Skill, target Entity, incomingDamage int) (int, error) {
	var actualDamageTaken int = 0

	for _, effect := range target.GetBattleState().activeEffectsList {
		if effect.skillEffect.usageTiming == etiOnSkillCalculation {
			if effect.skillEffect.category == ecaIncreasePower || effect.skillEffect.category == ecaDecreasePower {
				if rand.Float32() <= effect.skillEffect.probability {
					modifiedPower = float32(basePower/5*effect.skillEffect.baseValue + 1)
				}
			}
		}
	}
	return actualDamageTaken, nil
}
