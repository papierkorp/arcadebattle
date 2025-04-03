package internal

import (
	"fmt"
	"slices"
	"strings"
)

type SkillEffect struct {
	name            string
	description     string
	cost            int
	usage           func()
	usageTiming     EffectTiming
	removalTiming   EffectTiming
	validSkillTypes []string
}

type EffectTiming int

const (
	OnTurnStart EffectTiming = iota
	OnSkillUse
	OnTurnEnd
	OnDurationEnd
)

type effectFunctions map[string]func()

func newSkillEffect(skillType string, effectName string) (error, SkillEffect) {
	//todo use this function to create the docs while also passing the skilltype

	type partSkillEffect struct {
		usage           func()
		usageTiming     EffectTiming
		removalTiming   EffectTiming
		validSkillTypes []string
	}

	partSkillEffectMap := map[string]partSkillEffect{
		// Damage Effects
		"directdamage": {
			usage:           effectUseDirectDamage,
			usageTiming:     OnSkillUse,
			removalTiming:   OnSkillUse,
			validSkillTypes: []string{"immediate", "duration"},
		},
		"pierce": {
			usage:           effectUsePierce,
			usageTiming:     OnSkillUse,
			removalTiming:   OnSkillUse,
			validSkillTypes: []string{"immediate"},
		},
		"finisher": {
			usage:           effectUseFinisher,
			usageTiming:     OnSkillUse,
			removalTiming:   OnSkillUse,
			validSkillTypes: []string{"immediate"},
		},
		"buffturnbonusdamage": {
			usage:           effectUseBuffTurnBonusDamage,
			usageTiming:     OnSkillUse,
			removalTiming:   OnSkillUse,
			validSkillTypes: []string{"immediate"},
		},
		"debuffturnbonusdamage": {
			usage:           effectUseDebuffTurnBonusDamage,
			usageTiming:     OnSkillUse,
			removalTiming:   OnSkillUse,
			validSkillTypes: []string{"immediate"},
		},

		// Support Effects
		"directheal": {
			usage:           effectUseDirectHeal,
			usageTiming:     OnSkillUse,
			removalTiming:   OnSkillUse,
			validSkillTypes: []string{"immediate"},
		},
		"lifeleech": {
			usage:           effectUseLifeleech,
			usageTiming:     OnSkillUse,
			removalTiming:   OnSkillUse,
			validSkillTypes: []string{"immediate"},
		},
		"cleanse": {
			usage:           effectUseCleanse,
			usageTiming:     OnSkillUse,
			removalTiming:   OnSkillUse,
			validSkillTypes: []string{"immediate"},
		},
		"dispel": {
			usage:           effectUseDispel,
			usageTiming:     OnSkillUse,
			removalTiming:   OnSkillUse,
			validSkillTypes: []string{"immediate"},
		},
		"extendbuffs": {
			usage:           effectUseExtendBuffs,
			usageTiming:     OnSkillUse,
			removalTiming:   OnSkillUse,
			validSkillTypes: []string{"duration", "passive"},
		},
		"extenddebuffs": {
			usage:           effectUseExtendDebuffs,
			usageTiming:     OnSkillUse,
			removalTiming:   OnSkillUse,
			validSkillTypes: []string{"duration", "passive"},
		},
		"reducedebuffs": {
			usage:           effectUseReduceDebuffs,
			usageTiming:     OnSkillUse,
			removalTiming:   OnSkillUse,
			validSkillTypes: []string{"duration", "passive"},
		},
		"reducebuffs": {
			usage:           effectUseReduceBuffs,
			usageTiming:     OnSkillUse,
			removalTiming:   OnSkillUse,
			validSkillTypes: []string{"duration", "passive"},
		},

		// Over Time - Buff Effects
		"blockdebuffs": {
			usage:           effectUseBlockDebuffs,
			usageTiming:     OnTurnStart,
			removalTiming:   OnDurationEnd,
			validSkillTypes: []string{"duration", "passive"},
		},
		"healovertime": {
			usage:           effectUseHealOverTime,
			usageTiming:     OnTurnStart,
			removalTiming:   OnDurationEnd,
			validSkillTypes: []string{"duration"},
		},
		"incpower": {
			usage:           effectUseIncPower,
			usageTiming:     OnTurnStart,
			removalTiming:   OnDurationEnd,
			validSkillTypes: []string{"duration", "passive"},
		},
		"shield": {
			usage:           effectUseShield,
			usageTiming:     OnTurnStart,
			removalTiming:   OnDurationEnd,
			validSkillTypes: []string{"duration", "passive"},
		},
		"reflectdamage": {
			usage:           effectUseReflectDamage,
			usageTiming:     OnTurnStart,
			removalTiming:   OnDurationEnd,
			validSkillTypes: []string{"duration", "passive"},
		},
		"evasion": {
			usage:           effectUseEvasion,
			usageTiming:     OnTurnStart,
			removalTiming:   OnDurationEnd,
			validSkillTypes: []string{"duration", "passive"},
		},
		"criticalstrike": {
			usage:           effectUseCriticalRate,
			usageTiming:     OnTurnStart,
			removalTiming:   OnDurationEnd,
			validSkillTypes: []string{"duration", "passive"},
		},

		// Over Time - Debuff Effects
		"dot": {
			usage:           effectUseDOT,
			usageTiming:     OnTurnStart,
			removalTiming:   OnDurationEnd,
			validSkillTypes: []string{"duration"},
		},
		"stun": {
			usage:           effectUseStun,
			usageTiming:     OnTurnStart,
			removalTiming:   OnDurationEnd,
			validSkillTypes: []string{"duration"},
		},
		"damagereduction": {
			usage:           effectUseDamageReduction,
			usageTiming:     OnTurnStart,
			removalTiming:   OnDurationEnd,
			validSkillTypes: []string{"duration", "passive"},
		},
		"blockbuffs": {
			usage:           effectUseBlockBuffs,
			usageTiming:     OnTurnStart,
			removalTiming:   OnDurationEnd,
			validSkillTypes: []string{"duration"},
		},
		"grievouswounds": {
			usage:           effectUseGrievousWounds,
			usageTiming:     OnTurnStart,
			removalTiming:   OnDurationEnd,
			validSkillTypes: []string{"duration"},
		},
	}

	skillType = strings.ToLower(skillType)
	effectName = strings.ToLower(effectName)

	effectConfig, exists := partSkillEffectMap[effectName]
	if !exists {
		invalidEffectMsg := GetGameTextError("invalideffect")
		return fmt.Errorf("%s: %s", invalidEffectMsg, effectName), SkillEffect{}
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
		return fmt.Errorf("%s: %s (%s)", invalidskilltypeeffect, skillType, effectName), SkillEffect{}
	}

	err, effectCost := getEffectCost(effectName)
	if err != nil {
		return err, SkillEffect{}
	}

	gameTextEffectMsg := GetGameTextEffect(effectName)

	skillEffect := SkillEffect{
		name:            gameTextEffectMsg.Name,
		description:     gameTextEffectMsg.Description,
		cost:            effectCost,
		usage:           effectConfig.usage,
		usageTiming:     effectConfig.usageTiming,
		removalTiming:   effectConfig.removalTiming,
		validSkillTypes: effectConfig.validSkillTypes,
	}

	return nil, skillEffect
}

func upgradeSkillEffect(effectName SkillEffect) (error, SkillEffect) {
	//todo implement
	err, effectCost := getEffectCost(effectName.name)
	if err != nil {
		return err, SkillEffect{}
	}

	updatedEffect := SkillEffect{
		name:        effectName.name,
		description: effectName.description,
		cost:        effectCost,
	}

	return nil, updatedEffect
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
