// Package internal comment
package internal

import (
	"fmt"
	"slices"
	"strings"
)

// SkillEffect comment
type SkillEffect struct {
	name            string
	description     string
	cost            int
	usageTiming     EffectTiming
	isBlockedBy     []SkillEffect
	target          EffectTarget
	validSkillTypes []string
	modifier        int
	probability     float32
	affectedStat    EffectModifiableStats
	costType        EffectCostType
	costValue       int
	primaryFunction EffectFunction
	category        EffectCategory
}

// EffectTarget comment
type EffectTarget int

const (
	Self EffectTarget = iota
	Enemy
)

// EffectFunction comment
type EffectFunction int

const (
	Increase EffectFunction = iota
	Decrease
	Block
)

// EffectCategory which stats can be modified with effects
type EffectCategory int

const (
	Buff EffectCategory = iota
	Debuff
	Damage
)

// EffectModifiableStats which stats can be modified with effects
type EffectModifiableStats int

const (
	Health EffectModifiableStats = iota
	Power
	BuffTurn
	DebuffTurn
)

// EffectCostType comment
type EffectCostType int

const (
	Health EffectCostType = iota
	ReamainingTurnCount
	TotalEffectsCount
)

// ActiveEffect comment
type ActiveEffect struct {
	skillEffect SkillEffect
	totalPower  int
	turnsLeft   int
	source      Entity
	target      Entity
}

func (ae ActiveEffect) String() string {

	return fmt.Sprintf("Effectname: %s | Power: %d | Duration: %d",
		ae.skillEffect.name, ae.totalPower, ae.turnsLeft)
}

// EffectTiming comment
type EffectTiming int

const (
	OnTurnStart EffectTiming = iota
	OnSkillUse
	OnTurnEnd
	OnDurationEnd
)

type effectFunctions map[string]func()

func newSkillEffect(skillType string, effectName string) (SkillEffect, error) {
	//todo use this function to create the docs while also passing the skilltype

	type partSkillEffect struct {
		usage           func(ae ActiveEffect)
		usageTiming     EffectTiming
		isBlockedBy     []SkillEffect
		selfTarget      bool //if false target=enemy
		validSkillTypes []string
	}

	partSkillEffectMap := map[string]partSkillEffect{
		// Damage Effects
		"finisher": {
			usage:           effectUseFinisher,
			usageTiming:     OnSkillUse,
			isBlockedBy:     []SkillEffect{},
			selfTarget:      false,
			validSkillTypes: []string{"immediate"},
		},
		"buffturnbonusdamage": {
			usage:           effectUseBuffTurnBonusDamage,
			usageTiming:     OnSkillUse,
			isBlockedBy:     []SkillEffect{},
			selfTarget:      false,
			validSkillTypes: []string{"immediate"},
		},
		"debuffturnbonusdamage": {
			usage:           effectUseDebuffTurnBonusDamage,
			usageTiming:     OnSkillUse,
			isBlockedBy:     []SkillEffect{},
			selfTarget:      false,
			validSkillTypes: []string{"immediate"},
		},
		// Support Effects
		"directheal": {
			usage:           effectUseDirectHeal,
			usageTiming:     OnSkillUse,
			isBlockedBy:     []SkillEffect{},
			selfTarget:      true,
			validSkillTypes: []string{"immediate"},
		},
		"lifeleech": {
			usage:           effectUseLifeleech,
			usageTiming:     OnSkillUse,
			isBlockedBy:     []SkillEffect{},
			selfTarget:      true,
			validSkillTypes: []string{"duration", "immediate", "passive"},
		},
		"cleanse": {
			usage:           effectUseCleanse,
			usageTiming:     OnSkillUse,
			isBlockedBy:     []SkillEffect{},
			selfTarget:      true,
			validSkillTypes: []string{"immediate"},
		},
		"dispel": {
			usage:           effectUseDispel,
			usageTiming:     OnSkillUse,
			isBlockedBy:     []SkillEffect{},
			selfTarget:      false,
			validSkillTypes: []string{"immediate"},
		},
		"extendbuffs": {
			usage:           effectUseExtendBuffs,
			usageTiming:     OnSkillUse,
			isBlockedBy:     []SkillEffect{},
			selfTarget:      true,
			validSkillTypes: []string{"immediate"},
		},
		"extenddebuffs": {
			usage:           effectUseExtendDebuffs,
			usageTiming:     OnSkillUse,
			isBlockedBy:     []SkillEffect{},
			selfTarget:      false,
			validSkillTypes: []string{"immediate"},
		},
		"reducedebuffs": {
			usage:           effectUseReduceDebuffs,
			usageTiming:     OnSkillUse,
			isBlockedBy:     []SkillEffect{},
			selfTarget:      true,
			validSkillTypes: []string{"immediate"},
		},
		"reducebuffs": {
			usage:           effectUseReduceBuffs,
			usageTiming:     OnSkillUse,
			isBlockedBy:     []SkillEffect{},
			selfTarget:      false,
			validSkillTypes: []string{"immediate"},
		},
		// Over Time - Buff Effects
		"blockdebuffs": {
			usage:           effectUseBlockDebuffs,
			usageTiming:     OnTurnStart,
			isBlockedBy:     []SkillEffect{},
			selfTarget:      true,
			validSkillTypes: []string{"duration"},
		},
		"healovertime": {
			usage:           effectUseHealOverTime,
			usageTiming:     OnTurnStart,
			isBlockedBy:     []SkillEffect{},
			selfTarget:      true,
			validSkillTypes: []string{"duration"},
		},
		"incpower": {
			usage:           effectUseIncPower,
			usageTiming:     OnTurnStart,
			isBlockedBy:     []SkillEffect{},
			selfTarget:      true,
			validSkillTypes: []string{"duration", "passive"},
		},
		"reflectdamage": {
			usage:           effectUseReflectDamage,
			usageTiming:     OnTurnStart,
			isBlockedBy:     []SkillEffect{},
			selfTarget:      true,
			validSkillTypes: []string{"duration", "passive"},
		},
		"evasion": {
			usage:           effectUseEvasion,
			usageTiming:     OnTurnStart,
			isBlockedBy:     []SkillEffect{},
			selfTarget:      true,
			validSkillTypes: []string{"duration", "passive"},
		},
		"criticalstrike": {
			usage:           effectUseCriticalRate,
			usageTiming:     OnTurnStart,
			isBlockedBy:     []SkillEffect{},
			selfTarget:      true,
			validSkillTypes: []string{"duration", "passive"},
		},
		// Over Time - Debuff Effects
		"dot": {
			usage:           effectUseDOT,
			usageTiming:     OnTurnStart,
			isBlockedBy:     []SkillEffect{},
			selfTarget:      false,
			validSkillTypes: []string{"duration"},
		},
		"damagereduction": {
			usage:           effectUseDamageReduction,
			usageTiming:     OnTurnStart,
			isBlockedBy:     []SkillEffect{},
			selfTarget:      false,
			validSkillTypes: []string{"duration", "passive"},
		},
		"blockbuffs": {
			usage:           effectUseBlockBuffs,
			usageTiming:     OnTurnStart,
			isBlockedBy:     []SkillEffect{},
			selfTarget:      false,
			validSkillTypes: []string{"duration"},
		},
		"reducehealing": {
			usage:           effectUseReduceHealing,
			usageTiming:     OnTurnStart,
			isBlockedBy:     []SkillEffect{},
			selfTarget:      false,
			validSkillTypes: []string{"duration"},
		},
	}

	skillType = strings.ToLower(skillType)
	effectName = strings.ToLower(effectName)

	effectConfig, exists := partSkillEffectMap[effectName]
	if !exists {
		invalidEffectMsg := GetGameTextError("invalideffect")
		return SkillEffect{}, fmt.Errorf("%s: %s", invalidEffectMsg, effectName)
	}

	isValidType := slices.Contains(effectConfig.validSkillTypes, skillType)

	if !isValidType {
		invalidskilltypeeffect := GetGameTextError("invalidskilltypeeffect")
		return SkillEffect{}, fmt.Errorf("%s: %s (%s)", invalidskilltypeeffect, skillType, effectName)
	}

	effectCost, err := getEffectCost(effectName)
	if err != nil {
		return SkillEffect{}, err
	}

	gameTextEffectMsg := GetGameTextEffect(effectName)

	skillEffect := SkillEffect{
		name:            gameTextEffectMsg.Name,
		description:     gameTextEffectMsg.Description,
		cost:            effectCost,
		usage:           effectConfig.usage,
		usageTiming:     effectConfig.usageTiming,
		validSkillTypes: effectConfig.validSkillTypes,
	}

	return skillEffect, nil
}

func upgradeSkillEffect(effectName SkillEffect) (SkillEffect, error) {
	//todo implement
	effectCost, err := getEffectCost(effectName.name)
	if err != nil {
		return SkillEffect{}, err
	}

	updatedEffect := SkillEffect{
		name:        effectName.name,
		description: effectName.description,
		cost:        effectCost,
	}

	return updatedEffect, nil
}

func handleEffect(ae ActiveEffect) {
	primaryFunction := ae.skillEffect.primaryFunction
	costType := ae.skillEffect.costType

	switch primaryFunction {
	case Increase:
	case Decrease:
	case Block:
	default:
		internalErrMsg := GetGameTextError("internal")
		fmt.Printf("%s - %s doest not exist", internalErrMsg, primaryFunction)
	}

	switch costType {
	case Health:
	case ReamainingTurnCount:
	case TotalEffectsCount:
	default:
		internalErrMsg := GetGameTextError("internal")
		fmt.Printf("%s - %s doest not exist", internalErrMsg, costType)
	}
}
