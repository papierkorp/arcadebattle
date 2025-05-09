// Package internal comment
package internal

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

// todo talentpoints_total, talentpoints_remaining

// -------------------------------------------------------------------------
// -------------------------------define player------------------------------
// -------------------------------------------------------------------------

// Player comment
type Player struct {
	name                  string
	stats                 Stats
	talentpointsTotal     int
	talentpointsRemaining int
	state                 state
	difficulty            difficulty
	bosses                int
	skilllist             []Skill
	battlestate           BattleState
}

// Stats comment
type Stats struct {
	health int
	power  int
	speed  int
}

func createPlayer(args []string) (*Player, error) {

	newplayerMsg := GetGameTextCommand("newplayer")
	if len(args) != 7 {
		invalidArgsMsg := GetGameTextError("invalidargs")
		return nil, fmt.Errorf(invalidArgsMsg, newplayerMsg.Usage)
	}

	diff, err := ParseDifficulty(args[3])
	if err != nil {
		return nil, err
	}

	initialTalentpoints := getTalentpoints(diff)

	health, err := strconv.Atoi(args[4])
	if err != nil {
		invalidHealthMsg := GetGameTextError("invalidhealth")
		return nil, fmt.Errorf(invalidHealthMsg, err)
	}

	power, err := strconv.Atoi(args[5])
	if err != nil {
		invalidPowerMsg := GetGameTextError("invalidpower")
		return nil, fmt.Errorf(invalidPowerMsg, err)
	}

	speed, err := strconv.Atoi(args[6])
	if err != nil {
		invalidSpeedMsg := GetGameTextError("invalidspeed")
		return nil, fmt.Errorf(invalidSpeedMsg, err)
	}

	stats := Stats{
		health: health,
		power:  power,
		speed:  speed,
	}

	usedTalentpoints := calculateStatsCost(stats)

	if usedTalentpoints > initialTalentpoints {
		exceededMsg := GetGameTextError("exceededtalentpoints")
		separatorMsg := GetGameTextGameMessage("separator")
		usedMsg := GetGameTextGameMessage("used")
		availableMsg := GetGameTextGameMessage("available")
		return nil, fmt.Errorf("%s, %s: %d %s %s: %d",
			exceededMsg,         // Used too much Talentpoints
			availableMsg,        // Available:
			initialTalentpoints, // available value
			separatorMsg,        //  |
			usedMsg,             // Used:
			usedTalentpoints)    // used value
	}

	beforeMsg := GetGameTextGameMessage("talentpointsbefore")
	afterMsg := GetGameTextGameMessage("talentpointsafter")
	usedMsg := GetGameTextGameMessage("talentpointsused")
	remainingTalentpoints := initialTalentpoints - usedTalentpoints

	fmt.Printf("\n%s: %d\n%s: %d\n%s: %d\n\n",
		beforeMsg, initialTalentpoints,
		usedMsg, usedTalentpoints,
		afterMsg, remainingTalentpoints)

	player := &Player{
		name:                  args[2],
		stats:                 stats,
		talentpointsTotal:     initialTalentpoints,
		talentpointsRemaining: initialTalentpoints - usedTalentpoints,
		state:                 idle,
		difficulty:            diff,
		bosses:                0,
	}

	player.ResetBattleState()

	return player, nil
}

// GetStats comment
func (p *Player) GetStats() Stats {
	return p.stats
}

// SetHealth set the health stat for the player
func (p *Player) SetHealth(hp int) {
	p.stats.health = hp
}

// SetPower sets the power stat for the player
func (p *Player) SetPower(power int) {
	p.stats.power = power
}

// SetSpeed sets the speed stat for the player
func (p *Player) SetSpeed(speed int) {
	p.stats.speed = speed
}

// GetBattleState returns the player's battle state
func (p *Player) GetBattleState() *BattleState {
	return &p.battlestate
}

// SetCurrentHealth sets the current health in battle state
func (p *Player) SetCurrentHealth(health int) {
	p.battlestate.currentHealth = health
}

// SetTotalTurnsBuffs sets the total turns for buffs in battle state
func (p *Player) SetTotalTurnsBuffs(turns int) {
	p.battlestate.totalTurnsBuffs = turns
}

// SetTotalTurnsDebuff sets the total turns for debuffs in battle state
func (p *Player) SetTotalTurnsDebuff(turns int) {
	p.battlestate.totalTurnsDebuff = turns
}

// AddActiveEffect adds an effect to the active effects list
func (p *Player) AddActiveEffect(effect ActiveEffect) {
	p.battlestate.activeEffectsList = append(p.battlestate.activeEffectsList, effect)
}

// RemoveActiveEffect removes an effect from the active effects list
func (p *Player) RemoveActiveEffect(effect SkillEffect) {
	p.battlestate.activeEffectsList = slices.DeleteFunc(p.battlestate.activeEffectsList, func(ae ActiveEffect) bool {
		return ae.skillEffect.name == effect.name
	})
}

// ClearActiveEffects removes all active effects
func (p *Player) ClearActiveEffects() {
	p.battlestate.activeEffectsList = []ActiveEffect{}
}

// SetBattlePhase sets the current battle phase
func (p *Player) SetBattlePhase(phase BattlePhase) {
	p.battlestate.currentBattlePhase = phase
}

// GetName comment
func (p *Player) GetName() string {
	return p.name
}

// CheckDefeat checks if the player is defeated
func (p *Player) CheckDefeat() bool {
	if p.battlestate.currentHealth <= 0 {
		p.state = dead
		return true
	}
	return false
}

// HandleDefeat comment
func (p *Player) HandleDefeat() {
	// todo: add statistics, e.g. defeated bosses..
	// todo: show statistic summary

	currentPlayer.state = dead

	StateDead()

	currentPlayer.state = idle

	StateIdle()
}

// ResetBattleState resets the battle state to initial values
func (p *Player) ResetBattleState() {
	p.battlestate = BattleState{
		currentHealth:      p.stats.health,
		totalTurnsBuffs:    0,
		totalTurnsDebuff:   0,
		activeEffectsList:  []ActiveEffect{},
		currentBattlePhase: turnStart,
	}
}

// ApplyDamage applies damage to the player
func (p *Player) ApplyDamage(amount int) {
	// todo: implement damage application logic with shields and effects

	currentHealth := p.battlestate.currentHealth
	newHealth := max(currentHealth-amount, 0)

	damagedealtMsg := GetGameTextBattle("damagedealt")
	damageMsg := GetGameTextBattle("damage")
	fmt.Printf("%s %d %s\n", damagedealtMsg, amount, damageMsg)
	p.battlestate.currentHealth = newHealth
}

// ApplyHealing applies healing to the player
func (p *Player) ApplyHealing(amount int) {
	// todo: implement healing logic with effects
	currentHealth := p.battlestate.currentHealth
	maxHealth := p.stats.health
	newHealth := min(maxHealth, currentHealth+amount)
	p.battlestate.currentHealth = newHealth
}

// HasActiveEffect checks if the player has a specific active effect
func (p *Player) HasActiveEffect(effectType string) bool {
	// todo: implement effect checking logic
	return false
}

// -------------------------------------------------------------------------
// -------------------------------handle state------------------------------
// -------------------------------------------------------------------------

// -------------------------------------------------------------------------
// -------------------------------battle------------------------------
// -------------------------------------------------------------------------

func playerTurn() {
	fmt.Print("player turn\n")

	// -----------------------turnStart-----------------------
	currentPlayer.SetBattlePhase(turnStart)
	var remainingEffects []ActiveEffect

	for _, activeEffect := range currentPlayer.GetBattleState().activeEffectsList {
		activeEffect.skillEffect.usage(activeEffect)

		activeEffect.turnsLeft--

		if activeEffect.turnsLeft > 0 {
			remainingEffects = append(remainingEffects, activeEffect)
		}
	}

	currentPlayer.GetBattleState().activeEffectsList = remainingEffects

	// -----------------------turnAction-----------------------
	currentPlayer.SetBattlePhase(turnAction)

	// is handled in start.go in the useCommand

	battlepromptMsg := GetGameTextBattle("battleprompt")

	for {
		rl.SetPrompt(battlepromptMsg)

		input, err := rl.Readline()
		if err != nil {
			break
		}

		commandArgs := strings.Fields(strings.ToLower(input))

		if len(commandArgs) == 0 {
			continue
		}

		command := commandArgs[0]
		validCommand := true

		switch command {
		case "help", "?":
			if len(commandArgs) > 1 {
				helpSpecificCommand(commandArgs[1])
			} else {
				helpCommand()
			}

			validCommand = false
		case "status":
			statusCommand(commandArgs)

			validCommand = false
		case "quit", "exit", "run":
			leaveBattle()
			return
		case "use":
			validCommand = useCommand(commandArgs)

		default:
			invalidCommandMsg := GetGameTextError("invalidcommand")
			gamestarthelpMsg := GetGameTextGameMessage("gamestarthelp")
			fmt.Println(invalidCommandMsg)
			fmt.Println(gamestarthelpMsg)
			validCommand = false
		}

		if validCommand {
			break
		}
	}

	promptMsg := GetGameTextGameMessage("prompt")
	rl.SetPrompt(promptMsg)

	// -----------------------turnEnd-----------------------

	currentPlayer.SetBattlePhase(turnEnd)
}
