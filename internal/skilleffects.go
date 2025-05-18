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
	multi            float32
}

type effectCategory int

const (
	ecaHeal effectCategory = iota
	ecaDoDamage
	ecaIncreasePower
	ecaIncreaseDamageDone
	ecaDecreaseDamageTaken
	ecaIncreaseHealing
	ecaaBlockDebuffs
	ecaBlockDamage
	ecaStopSkill
	ecaChangeTarget
	ecaTakeDamage
	ecaDecreasePower
	ecaDecreaseDamageDone
	ecaIncreaseDamageTaken
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
	skillEffect SkillEffect
	totalPower  int
	turnsLeft   int
	source      Entity
	target      Entity
}

func (ae ActiveEffect) String() string {
	return fmt.Sprintf("Effectname: %s | Power: %d | Duration: %d",
		ae.skillEffect.displayName, ae.totalPower, ae.turnsLeft)
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
		multi          float32
	}

	partSkillEffectMap := map[string]partSkillEffect{
		"heal1": {
			internalName:   "heal1",
			probability:    1.0,
			effectType:     etyBuff,
			category:       ecaHeal,
			execute:        effectExecuteHeal1,
			checkCondition: effectCheckConditionHeal1,
			usageTiming:    etiOnTurnEnd,
			multi:          1,
		},
	}

	effectName = strings.ToLower(effectName)
	effectConfig, exists := partSkillEffectMap[effectName]

	if !exists {
		invalidEffectMsg := GetGameTextError("invalideffect")
		return SkillEffect{}, fmt.Errorf("%s: %s", invalidEffectMsg, effectName)
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
		multi:            effectConfig.multi,
	}

	return skillEffect, nil
}

func effectExecuteHeal1(ae ActiveEffect) {
	ae.source.ApplyHealing(ae.totalPower / 2)
}

func effectCheckConditionHeal1() bool {
	return true
}
