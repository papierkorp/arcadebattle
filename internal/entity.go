// Package internal comment
package internal

// Entity comment
type Entity interface {
	GetStats() Stats
	SetStats(stats Stats)
	GetBattleState() *BattleState
	GetName() string
}
