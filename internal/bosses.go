// Package internal comment
package internal

import (
	"fmt"
	"slices"
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
	currentBossNumber := currentPlayer.bosses + 1
	if currentBossNumber > 9 {
		return nil, fmt.Errorf("game completed - all bosses defeated")
	}

	difficulty := currentPlayer.difficulty

	baseHealth := 100
	if oldstats.health > 0 {
		baseHealth = oldstats.health
	}

	var baseStrength = 25
	if oldstats.strength > 0.0 {
		baseStrength = oldstats.strength
	}

	baseSpeed := 15
	if oldstats.speed > 0 {
		baseSpeed = oldstats.speed
	}

	healthMultiplier := 1.0 + (float64(currentBossNumber-1) * 0.4)    // +40% per boss
	strengthMultiplier := 1.0 + (float64(currentBossNumber-1) * 0.35) // +35% per boss
	speedMultiplier := 1.0 + (float64(currentBossNumber-1) * 0.2)     // +20% per boss

	switch difficulty {
	case hard:
		healthMultiplier *= 1.3
		strengthMultiplier *= 1.2
		speedMultiplier *= 1.1
	case expert:
		healthMultiplier *= 1.6
		strengthMultiplier *= 1.4
		speedMultiplier *= 1.2
	case master:
		healthMultiplier *= 2.0
		strengthMultiplier *= 1.7
		speedMultiplier *= 1.4
	case torment:
		healthMultiplier *= 2.5
		strengthMultiplier *= 2.0
		speedMultiplier *= 1.6
	}

	bossHealth := int(float64(baseHealth) * healthMultiplier)
	bossStrength := int(float64(baseStrength) * strengthMultiplier)
	bossSpeed := int(float64(baseSpeed) * speedMultiplier)

	newStats := Stats{
		health:   bossHealth,
		strength: bossStrength,
		speed:    bossSpeed,
	}

	bossName := fmt.Sprintf("%s-boss-%d", difficulty, currentBossNumber)

	boss := &Boss{
		name:  bossName,
		stats: newStats,
	}

	boss.ResetBattleState()

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

// SetStrength sets the strength stat for the boss
func (b *Boss) SetStrength(strength int) {
	b.stats.strength = strength
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
	b.battlestate.totalBuffTurnsCount = turns
}

// SetTotalTurnsDebuff sets the total turns for debuffs
func (b *Boss) SetTotalTurnsDebuff(turns int) {
	b.battlestate.totalDebuffTurnCount = turns
}

// AddActiveEffect adds an effect to the active effects list
func (b *Boss) AddActiveEffect(effect ActiveEffect) {
	b.battlestate.activeEffectsList = append(b.battlestate.activeEffectsList, effect)
}

// RemoveActiveEffect removes an effect from the active effects list
func (b *Boss) RemoveActiveEffect(effect SkillEffect) {
	b.battlestate.activeEffectsList = slices.DeleteFunc(b.battlestate.activeEffectsList, func(ae ActiveEffect) bool {
		return ae.skillEffect.internalName == effect.internalName
	})
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
	if b.battlestate.currentHealth <= 0 {
		return true
	}
	return false
}

// HandleDefeat comment
func (b *Boss) HandleDefeat() {
	bossdefeatedMsg := GetGameTextBattle("bossdefeated")
	newtalentpointsMsg := GetGameTextGameMessage("newtalentpoints")

	fmt.Printf("%s: %s!\n", bossdefeatedMsg, b.name)

	currentPlayer.bosses++
	// todo add to balancing file
	rewardTalentPoints := 10 + (5 * currentPlayer.bosses)
	currentPlayer.talentpointsRemaining += rewardTalentPoints

	fmt.Printf("%s: %d\n", newtalentpointsMsg, rewardTalentPoints)

	currentPlayer.state = idle
}

// ResetBattleState resets the battle state to initial values
func (b *Boss) ResetBattleState() {
	b.battlestate = BattleState{
		currentHealth:        b.stats.health,
		totalBuffTurnsCount:  0,
		totalDebuffTurnCount: 0,
		activeEffectsList:    []ActiveEffect{},
		currentBattlePhase:   turnStart,
	}
}

// ApplyDamage applies damage to the boss
func (b *Boss) ApplyDamage(amount int) {
	currentHealth := b.battlestate.currentHealth
	newHealth := max(0, currentHealth-amount)

	damagedealtMsg := GetGameTextBattle("damagedealt")
	damageMsg := GetGameTextBattle("damage")
	newhealthMsg := GetGameTextBattle("newhealth")
	fmt.Printf("%s %d %s\n", damagedealtMsg, amount, damageMsg)
	fmt.Printf("%s: %d\n", newhealthMsg, newHealth)

	b.battlestate.currentHealth = newHealth
}

// ApplyHealing applies healing to the boss
func (b *Boss) ApplyHealing(amount int) {
	currentHealth := b.battlestate.currentHealth
	maxHealth := b.stats.health
	newHealth := min(maxHealth, currentHealth+amount)
	b.battlestate.currentHealth = newHealth
}

// HasActiveEffect checks if the boss has a specific active effect
func (b *Boss) HasActiveEffect(effectType string) bool {
	for _, effect := range b.battlestate.activeEffectsList {
		if effect.skillEffect.internalName == effectType {
			return true
		}
	}
	return false
}

func (b *Boss) SetLastRawSkillPowerUsed(power int) {
	b.battlestate.lastRawSkillPowerUsed = power
}

func (b *Boss) SetLastActualDamageTaken(damage int) {
	b.battlestate.lastActualDamageTaken = damage
}

func (b *Boss) SetLastOutgoingDamage(damage int) {
	b.battlestate.lastOutgoingDamage = damage
}

func (b *Boss) SetLastIncomingDamage(damage int) {
	b.battlestate.lastIncomingDamage = damage
}

func (b *Boss) SetLastSkillUsed(skill *Skill) {
	b.battlestate.lastSkillUsed = skill
}

func (b *Boss) SetCurrentSkillUsed(skill *Skill) {
	b.battlestate.currentSkillUsed = skill
}

// -------------------------------------------------------------------------
// -------------------------------handle state------------------------------
// -------------------------------------------------------------------------

func checkCurrentBoss() error {
	stats := Stats{
		health:   0,
		strength: 0.0,
		speed:    0,
	}
	boss, err := createBoss(stats)
	if err != nil {
		internalErrMsg := GetGameTextError("internal")
		invalidBossMsg := GetGameTextError("invalidboss")
		return fmt.Errorf(internalErrMsg + ": " + invalidBossMsg)
	}

	currentBoss = *boss
	return nil
}

func leaveBattle() {
	currentPlayer.state = idle
	goodbyeMsg := GetGameTextGameMessage("goodbyebattle")
	promptMsg := GetGameTextGameMessage("prompt")

	fmt.Println(goodbyeMsg)
	rl.SetPrompt(promptMsg)
}

// -------------------------------------------------------------------------
// -------------------------------boss ai------------------------------
// -------------------------------------------------------------------------

func bossTurn() {
	separator2Msg := GetGameTextGameMessage("separator2")
	fmt.Println(separator2Msg)
	fmt.Print("boss turn\n")

	// -----------------------turnStart-----------------------
	currentBoss.SetBattlePhase(turnStart)
	var remainingEffects []ActiveEffect

	for _, activeEffect := range currentBoss.GetBattleState().activeEffectsList {
		activeEffect.skillEffect.execute(activeEffect)
		activeEffect.turnsLeft--

		if activeEffect.turnsLeft > 0 {
			remainingEffects = append(remainingEffects, activeEffect)
		}
	}

	currentBoss.GetBattleState().activeEffectsList = remainingEffects

	// -----------------------turnAction-----------------------
	currentBoss.SetBattlePhase(turnAction)

	fmt.Println("Boss dealt 6 Damage")
	currentPlayer.ApplyDamage(6)

	// todo: add use skills

	// -----------------------turnEnd-----------------------
	currentBoss.SetBattlePhase(turnEnd)

	time.Sleep(1 * time.Second)
	updateTurnOrderCurrentTurn()
	checkCurrentTurn()
}
