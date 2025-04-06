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
	usage           func()
	usageTiming     EffectTiming
	isBlockedby     []SkillEffect
	selfTarget      bool //if false target=enemy
	validSkillTypes []string
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
		usage           func()
		usageTiming     EffectTiming
		validSkillTypes []string
	}

	partSkillEffectMap := map[string]partSkillEffect{
		// Damage Effects
		"directdamage": {
			usage:           effectUseDirectDamage,
			usageTiming:     OnSkillUse,
			validSkillTypes: []string{"immediate", "duration"},
		},
		"pierce": {
			usage:           effectUsePierce,
			usageTiming:     OnSkillUse,
			validSkillTypes: []string{"immediate"},
		},
		"finisher": {
			usage:           effectUseFinisher,
			usageTiming:     OnSkillUse,
			validSkillTypes: []string{"immediate"},
		},
		"buffturnbonusdamage": {
			usage:           effectUseBuffTurnBonusDamage,
			usageTiming:     OnSkillUse,
			validSkillTypes: []string{"immediate"},
		},
		"debuffturnbonusdamage": {
			usage:           effectUseDebuffTurnBonusDamage,
			usageTiming:     OnSkillUse,
			validSkillTypes: []string{"immediate"},
		},

		// Support Effects
		"directheal": {
			usage:           effectUseDirectHeal,
			usageTiming:     OnSkillUse,
			validSkillTypes: []string{"immediate"},
		},
		"lifeleech": {
			usage:           effectUseLifeleech,
			usageTiming:     OnSkillUse,
			validSkillTypes: []string{"immediate"},
		},
		"cleanse": {
			usage:           effectUseCleanse,
			usageTiming:     OnSkillUse,
			validSkillTypes: []string{"immediate"},
		},
		"dispel": {
			usage:           effectUseDispel,
			usageTiming:     OnSkillUse,
			validSkillTypes: []string{"immediate"},
		},
		"extendbuffs": {
			usage:           effectUseExtendBuffs,
			usageTiming:     OnSkillUse,
			validSkillTypes: []string{"duration", "passive"},
		},
		"extenddebuffs": {
			usage:           effectUseExtendDebuffs,
			validSkillTypes: []string{"duration", "passive"},
		},
		"reducedebuffs": {
			usage:           effectUseReduceDebuffs,
			usageTiming:     OnSkillUse,
			validSkillTypes: []string{"duration", "passive"},
		},
		"reducebuffs": {
			usage:           effectUseReduceBuffs,
			usageTiming:     OnSkillUse,
			validSkillTypes: []string{"duration", "passive"},
		},

		// Over Time - Buff Effects
		"blockdebuffs": {
			usage:           effectUseBlockDebuffs,
			usageTiming:     OnTurnStart,
			validSkillTypes: []string{"duration", "passive"},
		},
		"healovertime": {
			usage:           effectUseHealOverTime,
			usageTiming:     OnTurnStart,
			validSkillTypes: []string{"duration"},
		},
		"incpower": {
			usage:           effectUseIncPower,
			usageTiming:     OnTurnStart,
			validSkillTypes: []string{"duration", "passive"},
		},
		"shield": {
			usage:           effectUseShield,
			usageTiming:     OnTurnStart,
			validSkillTypes: []string{"duration", "passive"},
		},
		"reflectdamage": {
			usage:           effectUseReflectDamage,
			usageTiming:     OnTurnStart,
			validSkillTypes: []string{"duration", "passive"},
		},
		"evasion": {
			usage:           effectUseEvasion,
			usageTiming:     OnTurnStart,
			validSkillTypes: []string{"duration", "passive"},
		},
		"criticalstrike": {
			usage:           effectUseCriticalRate,
			usageTiming:     OnTurnStart,
			validSkillTypes: []string{"duration", "passive"},
		},

		// Over Time - Debuff Effects
		"dot": {
			usage:           effectUseDOT,
			usageTiming:     OnTurnStart,
			validSkillTypes: []string{"duration"},
		},
		"stun": {
			usage:           effectUseStun,
			usageTiming:     OnTurnStart,
			validSkillTypes: []string{"duration"},
		},
		"damagereduction": {
			usage:           effectUseDamageReduction,
			usageTiming:     OnTurnStart,
			validSkillTypes: []string{"duration", "passive"},
		},
		"blockbuffs": {
			usage:           effectUseBlockBuffs,
			usageTiming:     OnTurnStart,
			validSkillTypes: []string{"duration"},
		},
		"grievouswounds": {
			usage:           effectUseGrievousWounds,
			usageTiming:     OnTurnStart,
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

func effectUseDirectDamage() {
	fmt.Println("asdf")
}

func effectUsePierce() {
	fmt.Println("asdf")
}

func effectUseFinisher() {
	fmt.Println("asdf")
}

func effectUseBuffTurnBonusDamage() {
	fmt.Println("asdf")
}

func effectUseDebuffTurnBonusDamage() {
	fmt.Println("asdf")
}

// -------------------------------------------------------------------------
// ------------------------------direct support-----------------------------
// -------------------------------------------------------------------------

func effectUseDirectHeal() {
	fmt.Println("asdf")
}

func effectUseLifeleech() {
	fmt.Println("asdf")
}

func effectUseCleanse() {
	fmt.Println("asdf")
}

func effectUseDispel() {
	fmt.Println("asdf")
}

func effectUseExtendBuffs() {
	fmt.Println("asdf")
}

func effectUseExtendDebuffs() {
	fmt.Println("asdf")
}

func effectUseReduceDebuffs() {
	fmt.Println("asdf")
}

func effectUseReduceBuffs() {
	fmt.Println("asdf")
}

// -------------------------------------------------------------------------
// ------------------------------over time buffs-----------------------------
// -------------------------------------------------------------------------

func effectUseBlockDebuffs() {
	fmt.Println("asdf")
}

func effectUseHealOverTime() {
	fmt.Println("asdf")
}

func effectUseIncPower() {
	fmt.Println("asdf")
}

func effectUseShield() {
	fmt.Println("asdf")
}

func effectUseReflectDamage() {
	fmt.Println("asdf")
}

func effectUseEvasion() {
	fmt.Println("asdf")
}

func effectUseCriticalRate() {
	fmt.Println("asdf")
}

// -------------------------------------------------------------------------
// ------------------------------overtime debuffs-----------------------------
// -------------------------------------------------------------------------

func effectUseDOT() {
	fmt.Println("asdf")
}

func effectUseStun() {
	fmt.Println("asdf")
}

func effectUseDamageReduction() {
	fmt.Println("asdf")
}

func effectUseBlockBuffs() {
	fmt.Println("asdf")
}

func effectUseGrievousWounds() {
	fmt.Println("asdf")
}
