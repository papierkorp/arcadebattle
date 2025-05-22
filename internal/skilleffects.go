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
	baseValue        float32 // calculate: baseValue per powerRatio, e.g. 1 damage per 3 powerRatio
	powerRatio       float32 // calculate: baseValue per powerRatio e.g. 0.1 multi per 5 powerRatio
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
)

type effectType int

const (
	etyBuff effectType = iota
	etyDebuff
)

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
		baseValue      float32
		powerRatio     float32
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
			baseValue:      2,
			powerRatio:     3,
		},
		"increasepower1": {
			internalName:   "increasepower1",
			probability:    1.0,
			effectType:     etyBuff,
			category:       ecaIncreasePower,
			execute:        effectExecuteIncreasePower1,
			checkCondition: effectCheckConditionIncreasePower1,
			usageTiming:    etiOnSkillCalculation,
			baseValue:      0.1,
			powerRatio:     5,
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
		baseValue:        effectConfig.baseValue,
		powerRatio:       effectConfig.powerRatio,
	}

	return skillEffect, nil
}

func (se SkillEffect) CalculateEffect(power int) (float32, error) {
	switch se.category {
	case ecaHeal, ecaTakeDamage, ecaRemoveEffect, ecaChangeEffectTurn:
		// e.g. 2 healing per 3 powerRatio
		// e.g. 1 effect removal per 10 powerRatio
		return float32(power) / se.powerRatio * se.baseValue, nil
	case ecaIncreasePower, ecaIncreaseOutgoingDamage, ecaDecreaseIncomingDamage,
		ecaDecreasePower, ecaDecreaseOutgoingDamage, ecaIncreaseIncomingDamage,
		ecaDecreaseHealing, ecaIncreaseHealing:
		// e.g. 0.1 increase per 5 powerRatio
		return (float32(power)/se.powerRatio)*se.baseValue + 1, nil
	default:
		internalErrorMsg := GetGameTextError("internal")
		nosuchCategoryMsg := GetGameTextError("nosuchcategory")
		return 0, fmt.Errorf("(CalculateEffect) %s - %s", internalErrorMsg, nosuchCategoryMsg)
	}
}

func effectExecuteHeal1(ae ActiveEffect) {
	ae.source.ApplyHealing(ae.rawSkillPower / 2)
}

func effectCheckConditionHeal1() bool {
	return true
}

func effectExecuteIncreasePower1(ae ActiveEffect) {
}

func effectCheckConditionIncreasePower1() bool {
	return true
}
