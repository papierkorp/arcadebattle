// Package internal comment
package internal

import (
	"fmt"
	"time"
)

// -------------------------------------------------------------------------
// -------------------------------define boss------------------------------
// -------------------------------------------------------------------------

// Boss comment
type Boss struct {
	name        string
	stats       Stats
	skilllist   []Skill
	battlestate BattleState
}

func createBoss(oldstats Stats) (*Boss, error) {
	currentBossNumber := current_player.bosses + 1
	if currentBossNumber > 9 {
		return nil, fmt.Errorf("game completed - all bosses defeated")
	}

	difficulty := current_player.difficulty

	baseHealth := 100
	if oldstats.health > 0 {
		baseHealth = oldstats.health
	}

	var basePower = 25
	if oldstats.power > 0.0 {
		basePower = oldstats.power
	}

	baseSpeed := 15
	if oldstats.speed > 0 {
		baseSpeed = oldstats.speed
	}

	healthMultiplier := 1.0 + (float64(currentBossNumber-1) * 0.4) // +40% per boss
	powerMultiplier := 1.0 + (float64(currentBossNumber-1) * 0.35) // +35% per boss
	speedMultiplier := 1.0 + (float64(currentBossNumber-1) * 0.2)  // +20% per boss

	switch difficulty {
	case hard:
		healthMultiplier *= 1.3
		powerMultiplier *= 1.2
		speedMultiplier *= 1.1
	case expert:
		healthMultiplier *= 1.6
		powerMultiplier *= 1.4
		speedMultiplier *= 1.2
	case master:
		healthMultiplier *= 2.0
		powerMultiplier *= 1.7
		speedMultiplier *= 1.4
	case torment:
		healthMultiplier *= 2.5
		powerMultiplier *= 2.0
		speedMultiplier *= 1.6
	}

	bossHealth := int(float64(baseHealth) * healthMultiplier)
	bossPower := int(float64(basePower) * powerMultiplier)
	bossSpeed := int(float64(baseSpeed) * speedMultiplier)

	newStats := Stats{
		health: bossHealth,
		power:  bossPower,
		speed:  bossSpeed,
	}

	bossName := fmt.Sprintf("%s-boss-%d", difficulty, currentBossNumber)

	boss := &Boss{
		name:  bossName,
		stats: newStats,
	}

	return boss, nil
}

// GetStats comment
func (b *Boss) GetStats() Stats {
	return b.stats
}

// SetHealth sets the health stat for the boss
func (b *Boss) SetHealth(hp int) {
	b.stats.health = hp
}

// SetPower sets the power stat for the boss
func (b *Boss) SetPower(power int) {
	b.stats.power = power
}

// SetSpeed sets the speed stat for the boss
func (b *Boss) SetSpeed(speed int) {
	b.stats.speed = speed
}

// GetBattleState comment
func (b *Boss) GetBattleState() *BattleState {
	return &b.battlestate
}

// SetCurrentHealth sets the current health in battle state
func (b *Boss) SetCurrentHealth(health int) {
	b.battlestate.currentHealth = health
}

// SetTotalTurnsBuffs sets the total turns for buffs
func (b *Boss) SetTotalTurnsBuffs(turns int) {
	b.battlestate.totalTurnsBuffs = turns
}

// SetTotalTurnsDebuff sets the total turns for debuffs
func (b *Boss) SetTotalTurnsDebuff(turns int) {
	b.battlestate.totalTurnsDebuff = turns
}

// AddActiveEffect adds an effect to the active effects list
func (b *Boss) AddActiveEffect(effect ActiveEffect) {
	b.battlestate.activeEffectsList = append(b.battlestate.activeEffectsList, effect)
}

// RemoveActiveEffect removes an effect from the active effects list
func (b *Boss) RemoveActiveEffect(effect SkillEffect) {
	// Build a new slice excluding the effect to remove
	newEffects := []ActiveEffect{}
	for _, activeEffect := range b.battlestate.activeEffectsList {
		if activeEffect.skillEffect.name != effect.name {
			newEffects = append(newEffects, activeEffect)
		}
	}
	b.battlestate.activeEffectsList = newEffects
}

// ClearActiveEffects removes all active effects
func (b *Boss) ClearActiveEffects() {
	b.battlestate.activeEffectsList = []ActiveEffect{}
}

// SetBattlePhase sets the current battle phase
func (b *Boss) SetBattlePhase(phase BattlePhase) {
	b.battlestate.currentBattlePhase = phase
}

// GetName comment
func (b *Boss) GetName() string {
	return b.name
}

// CheckDefeat checks if the boss is defeated and returns true/false
func (b *Boss) CheckDefeat() bool {
	if b.stats.health <= 0 {
		return true
	}
	return false
}

// HandleDefeat comment
func (b *Boss) HandleDefeat() {
	bossdefeatedMsg := GetGameTextBattle("bossdefeated")
	newtalentpointsMsg := GetGameTextGameMessage("newtalentpoints")

	fmt.Printf("%s: %s!\n", bossdefeatedMsg, b.name)

	current_player.bosses++
	// todo add to balancing file
	rewardTalentPoints := 10 + (5 * current_player.bosses)
	current_player.talentpointsRemaining += rewardTalentPoints

	fmt.Printf("%s: %d\n", newtalentpointsMsg, rewardTalentPoints)

	current_player.state = idle
}

// ResetBattleState resets the battle state to initial values
func (b *Boss) ResetBattleState() {
	b.battlestate = BattleState{
		currentHealth:      b.stats.health,
		totalTurnsBuffs:    0,
		totalTurnsDebuff:   0,
		activeEffectsList:  []ActiveEffect{},
		currentBattlePhase: turnStart,
	}
}

// ApplyDamage applies damage to the boss
func (b *Boss) ApplyDamage(amount int) {
	currentHealth := b.battlestate.currentHealth
	newHealth := currentHealth - amount
	if newHealth < 0 {
		newHealth = 0
	}
	b.battlestate.currentHealth = newHealth
}

// ApplyHealing applies healing to the boss
func (b *Boss) ApplyHealing(amount int) {
	currentHealth := b.battlestate.currentHealth
	maxHealth := b.stats.health
	newHealth := currentHealth + amount
	if newHealth > maxHealth {
		newHealth = maxHealth
	}
	b.battlestate.currentHealth = newHealth
}

// HasActiveEffect checks if the boss has a specific active effect
func (b *Boss) HasActiveEffect(effectType string) bool {
	for _, effect := range b.battlestate.activeEffectsList {
		if effect.skillEffect.name == effectType {
			return true
		}
	}
	return false
}

// -------------------------------------------------------------------------
// -------------------------------handle state------------------------------
// -------------------------------------------------------------------------

func checkCurrentBoss() error {
	stats := Stats{
		health: 0,
		power:  0.0,
		speed:  0,
	}
	boss, err := createBoss(stats)
	if err != nil {
		internalErrMsg := GetGameTextError("internal")
		invalidBossMsg := GetGameTextError("invalidboss")
		return fmt.Errorf(internalErrMsg + ": " + invalidBossMsg)
	}

	current_boss = *boss
	return nil
}

func leaveBattle() {
	current_player.state = idle
	goodbyeMsg := GetGameTextGameMessage("goodbyebattle")
	promptMsg := GetGameTextGameMessage("prompt")

	fmt.Println(goodbyeMsg)
	rl.SetPrompt(promptMsg)
}

// -------------------------------------------------------------------------
// -------------------------------boss ai------------------------------
// -------------------------------------------------------------------------

func bossTurn() {
	fmt.Print("boss turn - ")
	fmt.Printf("%s\n", turn_order)

	time.Sleep(1 * time.Second)
	updateTurnOrderCurrentTurn()
	checkCurrentTurn()
}
