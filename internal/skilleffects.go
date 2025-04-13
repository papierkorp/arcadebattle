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
	usage           func(ae ActiveEffect)
	usageTiming     EffectTiming
	isBlockedBy     []SkillEffect
	selfTarget      bool //if false target=enemy
	validSkillTypes []string
}

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
		"directdamage": {
			usage:           effectUseDirectDamage,
			usageTiming:     OnSkillUse,
			isBlockedBy:     []SkillEffect{},
			selfTarget:      false,
			validSkillTypes: []string{"immediate", "duration"},
		},
		"pierce": {
			usage:           effectUsePierce,
			usageTiming:     OnSkillUse,
			isBlockedBy:     []SkillEffect{},
			selfTarget:      false,
			validSkillTypes: []string{"immediate"},
		},
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
			validSkillTypes: []string{"immediate"},
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
			validSkillTypes: []string{"duration", "passive"},
		},
		"extenddebuffs": {
			usage:           effectUseExtendDebuffs,
			usageTiming:     OnSkillUse,
			isBlockedBy:     []SkillEffect{},
			selfTarget:      false,
			validSkillTypes: []string{"duration", "passive"},
		},
		"reducedebuffs": {
			usage:           effectUseReduceDebuffs,
			usageTiming:     OnSkillUse,
			isBlockedBy:     []SkillEffect{},
			selfTarget:      true,
			validSkillTypes: []string{"duration", "passive"},
		},
		"reducebuffs": {
			usage:           effectUseReduceBuffs,
			usageTiming:     OnSkillUse,
			isBlockedBy:     []SkillEffect{},
			selfTarget:      false,
			validSkillTypes: []string{"duration", "passive"},
		},
		// Over Time - Buff Effects
		"blockdebuffs": {
			usage:           effectUseBlockDebuffs,
			usageTiming:     OnTurnStart,
			isBlockedBy:     []SkillEffect{},
			selfTarget:      true,
			validSkillTypes: []string{"duration", "passive"},
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
		"shield": {
			usage:           effectUseShield,
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
		"stun": {
			usage:           effectUseStun,
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
		"grievouswounds": {
			usage:           effectUseGrievousWounds,
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
	// isValidType := false
	// for _, validType := range effectConfig.validSkillTypes {
	// 	if validType == skillType {
	// 		isValidType = true
	// 		break
	// 	}
	// }

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

// -------------------------------------------------------------------------
// -------------------------------direct damage------------------------------
// -------------------------------------------------------------------------
func effectUseDirectDamage(ae ActiveEffect) {
	fmt.Println("Direct damage effect used")
}

func effectUsePierce(ae ActiveEffect) {
	fmt.Println("Pierce effect used")
}

func effectUseFinisher(ae ActiveEffect) {
	fmt.Println("Finisher effect used")
}

func effectUseBuffTurnBonusDamage(ae ActiveEffect) {
	fmt.Println("Buff turn bonus damage effect used")
}

func effectUseDebuffTurnBonusDamage(ae ActiveEffect) {
	fmt.Println("Debuff turn bonus damage effect used")
}

// -------------------------------------------------------------------------
// ------------------------------direct support-----------------------------
// -------------------------------------------------------------------------

func effectUseDirectHeal(ae ActiveEffect) {
	fmt.Println("Direct heal effect used")
}

func effectUseLifeleech(ae ActiveEffect) {
	fmt.Println("Life leech effect used")
}

func effectUseCleanse(ae ActiveEffect) {
	fmt.Println("Cleanse effect used")
}

func effectUseDispel(ae ActiveEffect) {
	fmt.Println("Dispel effect used")
}

func effectUseExtendBuffs(ae ActiveEffect) {
	fmt.Println("Extend buffs effect used")
}

func effectUseExtendDebuffs(ae ActiveEffect) {
	fmt.Println("Extend debuffs effect used")
}

func effectUseReduceDebuffs(ae ActiveEffect) {
	fmt.Println("Reduce debuffs effect used")
}

func effectUseReduceBuffs(ae ActiveEffect) {
	fmt.Println("Reduce buffs effect used")
}

// -------------------------------------------------------------------------
// ------------------------------over time buffs-----------------------------
// -------------------------------------------------------------------------

func effectUseBlockDebuffs(ae ActiveEffect) {
	fmt.Println("Block debuffs effect used")
}

func effectUseHealOverTime(ae ActiveEffect) {
	fmt.Println("Heal over time effect used")
}

func effectUseIncPower(ae ActiveEffect) {
	fmt.Println("Increase power effect used")
}

func effectUseShield(ae ActiveEffect) {
	fmt.Println("Shield effect used")
}

func effectUseReflectDamage(ae ActiveEffect) {
	fmt.Println("Reflect damage effect used")
}

func effectUseEvasion(ae ActiveEffect) {
	fmt.Println("Evasion effect used")
}

func effectUseCriticalRate(ae ActiveEffect) {
	fmt.Println("Critical rate effect used")
}

// -------------------------------------------------------------------------
// ------------------------------overtime debuffs-----------------------------
// -------------------------------------------------------------------------

func effectUseDOT(ae ActiveEffect) {
	fmt.Println("DOT effect used")
	ae.target.ApplyDamage(int(float64(ae.totalPower) * 0.25))
}

func effectUseStun(ae ActiveEffect) {
	fmt.Println("Stun effect used")
}

func effectUseDamageReduction(ae ActiveEffect) {
	fmt.Println("Damage reduction effect used")
}

func effectUseBlockBuffs(ae ActiveEffect) {
	fmt.Println("Block buffs effect used")
}

func effectUseGrievousWounds(ae ActiveEffect) {
	fmt.Println("Grievous wounds effect used")
}
