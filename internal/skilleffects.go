package internal

import (
	"fmt"
	"strings"
)

type SkillEffect struct {
	name        string
	description string
	cost        int
	usage       func()
}

type effectFunctions map[string]func()

func newSkillEffect(effectName string) (error, SkillEffect) {
	var usageEffectFunctions = effectFunctions{
		// Damage Effects
		"directdamage":          effectUseDirectDamage,
		"pierce":                effectUsePierce,
		"finisher":              effectUseFinisher,
		"buffturnbonusdamage":   effectUseBuffTurnBonusDamage,
		"debuffturnbonusdamage": effectUseDebuffTurnBonusDamage,

		// Support Effects
		"directheal":    effectUseDirectHeal,
		"lifeleech":     effectUseLifeleech,
		"cleanse":       effectUseCleanse,
		"dispel":        effectUseDispel,
		"extendbuffs":   effectUseExtendBuffs,
		"extenddebuffs": effectUseExtendDebuffs,
		"reducedebuffs": effectUseReduceDebuffs,
		"reducebuffs":   effectUseReduceBuffs,

		// Over Time - Buff Effects
		"blockdebuffs":   effectUseBlockDebuffs,
		"healovertime":   effectUseHealOverTime,
		"incpower":       effectUseIncPower,
		"shield":         effectUseShield,
		"reflectdamage":  effectUseReflectDamage,
		"evasion":        effectUseEvasion,
		"criticalstrike": effectUseCriticalStrike,

		// Over Time - Debuff Effects
		"dot":             effectUseDOT,
		"stun":            effectUseStun,
		"damagereduction": effectUseDamageReduction,
		"blockbuffs":      effectUseBlockBuffs,
		"grievouswounds":  effectUseGrievousWounds,
	}

	gameTextEffectMsg := GetGameTextEffect(effectName)
	err, effectCost := getEffectCost(effectName)
	if err != nil {
		return err, SkillEffect{}
	}

	functionName := strings.ToLower(effectName)
	newUsageFunc, exists := usageEffectFunctions[functionName]
	invalideffectMsg := GetGameTextError("invalideffect")
	if !exists {
		return fmt.Errorf(invalideffectMsg, ": ", effectName), SkillEffect{}
	}

	newName := gameTextEffectMsg.Name
	newCost := effectCost
	newDescription := gameTextEffectMsg.Description

	skillEffect := SkillEffect{
		name:        newName,
		cost:        newCost,
		description: newDescription,
		usage:       newUsageFunc,
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

func effectUseCriticalStrike() {
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
