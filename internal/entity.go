// Package internal comment
package internal

// Entity comment
type Entity interface {
	// Stats
	GetStats() Stats
	SetHealth(hp int)
	SetPower(power int)
	SetSpeed(speed int)

	// BattleState
	GetBattleState() *BattleState
	SetCurrentHealth(health int)
	SetTotalTurnsBuffs(turns int)
	SetTotalTurnsDebuff(turns int)
	AddActiveEffect(effect ActiveEffect)
	RemoveActiveEffect(effect SkillEffect)
	ClearActiveEffects()
	SetBattlePhase(phase BattlePhase)

	// Status
	GetName() string
	CheckDefeat() bool
	HandleDefeat()

	// Additonal Utility
	ResetBattleState()
	ApplyDamage(amount int)
	ApplyHealing(amount int)
	HasActiveEffect(effectType string) bool
}
