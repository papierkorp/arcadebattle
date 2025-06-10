// Package internal comment
package internal

import "fmt"

// Entity comment
type Entity interface {
	// Stats
	GetStats() Stats
	SetHealth(hp int)
	SetStrength(strength int)
	SetSpeed(speed int)

	// BattleState
	GetBattleState() *BattleState
	Turn()
	SetCurrentHealth(health int)
	SetTotalTurnsBuffs(turns int)
	SetTotalTurnsDebuff(turns int)
	AddActiveEffect(effect ActiveEffect)
	RemoveActiveEffect(effect SkillEffect)
	ClearActiveEffects()
	SetBattlePhase(phase BattlePhase)
	SetLastRawSkillPowerUsed(power int)
	SetLastActualDamageTaken(damage int)
	SetLastIncomingDamage(damage int)
	SetLastOutgoingDamage(damage int)
	SetLastSkillUsed(skill *Skill)
	SetCurrentSkillUsed(skill *Skill)

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

// EntityType comment
type EntityType string

const (
	PlayerEntity EntityType = "player"
	BossEntity   EntityType = "boss"
)

// GetEntity get current Entity based on EntityType
func GetEntity(entityType EntityType) Entity {
	switch entityType {
	case PlayerEntity:
		return &currentPlayer
	case BossEntity:
		return &currentBoss
	default:
		// This should hopefully never happen
		internalerrorMsg := GetGameTextError("internal")
		panic(fmt.Sprintf("%s - Invalid entity type: %s", internalerrorMsg, entityType))
	}
}
