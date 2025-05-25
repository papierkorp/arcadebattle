// Package internal comment
package internal

import (
	"fmt"
	"strings"
)

// SkillEffect to be used in the Skillcreation
type SkillEffect struct {
	internalName     string
	displayName      string
	description      string
	talentpointCosts int
	probability      float32
	effectType       effectType
	category         effectCategory
	execute          func(ae ActiveEffect)
	checkCondition   func() bool
	usageTiming      effectTiming
	outputValue      func() effectOutputValue
}

type effectCategory int

const (
	ecaHeal effectCategory = iota
	ecaDoDamage
	ecaIncreasePower
	ecaIncreaseOutgoingDamage
	ecaDecreaseIncomingDamage
	ecaIncreaseHealing
	ecaaBlockDebuffs
	ecaBlockDamage
	ecaStopSkill
	ecaChangeTarget
	ecaTakeDamage
	ecaDecreasePower
	ecaDecreaseOutgoingDamage
	ecaIncreaseIncomingDamage
	ecaDecreaseHealing
	ecaBlockBuffs
	ecaBlockHealing
	ecaRemoveEffect
	ecaChangeEffectTurn
)

type effectTiming int

const (
	etiOnTurnStart effectTiming = iota
	etiOnSkillStart
	etiOnSkillCalculation
	etiOnTurnEnd
	etiOnEffectRemoval
	etiOnSkillEnd
)

type effectType int

const (
	etyBuff effectType = iota
	etyDebuff
)

type effectOutputValue struct {
	multiplier float32
	output     int
}

// ActiveEffect is on the Entity in the BattleState
type ActiveEffect struct {
	skillEffect   SkillEffect
	rawSkillPower int
	turnsLeft     int
	source        Entity
	target        Entity
}

func (ae ActiveEffect) String() string {
	return fmt.Sprintf("Effectname: %s | Power: %d | Duration: %d",
		ae.skillEffect.displayName, ae.rawSkillPower, ae.turnsLeft)
}

func newSkillEffect(effectName string) (SkillEffect, error) {
	//todo use this function to create the docs while also passing the skilltype
	//todo convert displayname to internalname so skill can be created with either

	type partSkillEffect struct {
		internalName   string
		probability    float32
		effectType     effectType
		category       effectCategory
		execute        func(ae ActiveEffect)
		checkCondition func() bool
		usageTiming    effectTiming
		damageOutput   func() effectOutputValue
	}

	partSkillEffectMap := map[string]partSkillEffect{
		"heal1": {
			internalName:   "heal1",
			probability:    1.0,
			effectType:     etyDebuff,
			category:       ecaHeal,
			execute:        effectExecuteHeal1,
			checkCondition: effectCheckConditionHeal1,
			usageTiming:    etiOnTurnEnd,
			damageOutput:   effectOutputValueHeal1,
		},
		"increasepower1": {
			internalName:   "increasepower1",
			probability:    1.0,
			effectType:     etyBuff,
			category:       ecaIncreasePower,
			execute:        effectExecuteIncreasePower1,
			checkCondition: effectCheckConditionIncreasePower1,
			usageTiming:    etiOnSkillCalculation,
			damageOutput:   effectOutputValueIncreasePower1,
		},
	}

	effectNameOld := effectName
	effectName = strings.ToLower(effectName)
	effectConfig, exists := partSkillEffectMap[effectName]

	if !exists {
		allEffects := GetGameTextSkilleffects()

		for internalName, effect := range allEffects {

			if effectName == strings.ToLower(effect.Name) {
				effectName = strings.ToLower(internalName)
				effectConfig, exists = partSkillEffectMap[effectName]
			}
		}
	}

	if !exists {
		invalidEffectMsg := GetGameTextError("invalideffect")
		return SkillEffect{}, fmt.Errorf("%s: %s", invalidEffectMsg, effectNameOld)
	}

	effectCost, err := getEffectCost(effectName)
	if err != nil {
		return SkillEffect{}, err
	}

	gameTextEffectMsg := GetGameTextEffect(effectName)

	skillEffect := SkillEffect{
		internalName:     effectConfig.internalName,
		displayName:      gameTextEffectMsg.Name,
		description:      gameTextEffectMsg.Description,
		talentpointCosts: effectCost,
		probability:      effectConfig.probability,
		effectType:       effectConfig.effectType,
		category:         effectConfig.category,
		execute:          effectConfig.execute,
		checkCondition:   effectConfig.checkCondition,
		usageTiming:      effectConfig.usageTiming,
	}

	return skillEffect, nil
}

func effectExecuteHeal1(ae ActiveEffect) {
	ae.source.ApplyHealing(ae.rawSkillPower / 2)
}

func effectCheckConditionHeal1() bool {
	return true
}

func effectOutputValueHeal1() effectOutputValue {
	return effectOutputValue{
		multiplier: 1.0,
		baseValue:  2.0,
		powerRatio: 3.0,
		damage:     0,
	}
}

func effectExecuteIncreasePower1(ae ActiveEffect) {
}

func effectCheckConditionIncreasePower1() bool {
	return true
}

func effectOutputValueIncreasePower1() effectOutputValue {
	return effectOutputValue{
		multiplier: 1.0,
		baseValue:  0.1,
		powerRatio: 5.0,
		damage:     0,
	}
}
