// Package internal comment
package internal

import (
	"strings"
)

// -------------------------------------------------------------------------
// -------------------------------define skill------------------------------
// -------------------------------------------------------------------------

// Skill comment
type Skill interface {
	GetID() int
	GetName() string
	GetEffectList() []SkillEffect
	GetTalentPointCostsTotal() int
	GetSkillType() string
	GetDamageMultiplier() float32
	Use(source string) error
	String() string
}

func createEffectList(args []string, skilltype string, startIndex int) ([]SkillEffect, error) {
	effectList := []SkillEffect{}

	if len(args) <= startIndex {
		// maybe add a message that no effects were selected? idk
		return effectList, nil
	}

	for i := startIndex; i < len(args); i++ {
		effectName := strings.ToLower(args[i])

		effect, err := newSkillEffect(skilltype, effectName)
		if err != nil {
			return nil, err
		}

		effectList = append(effectList, effect)
	}

	return effectList, nil
}
