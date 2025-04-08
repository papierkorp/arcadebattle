// Package internal comment
package internal

// Entity comment
type Entity interface {
	GetStats() Stats
	SetHealth(hp int)
	SetSpeed(speed int)
	SetPower(power int)
	GetBattleState() *BattleState
	GetName() string
	CheckDefeat() bool
	HandleDefeat()
}
