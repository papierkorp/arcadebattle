// Package internal comment
package internal

import (
	"fmt"
)

// BattleState comment
type BattleState struct {
	currentHealth         int
	currentPower          int
	totalBuffTurnsCount   int
	totalBuffCount        int
	totalDebuffTurnCount  int
	totalDebuffCount      int
	activeEffectsList     []ActiveEffect
	currentBattlePhase    BattlePhase
	lastRawSkillPowerUsed int
	lastActualDamageTaken int
	lastIncomingDamage    int
	lastOutgoingDamage    int
	lastSkillUsed         *Skill
	currentSkillUsed      *Skill
}

// BattlePhase comment
type BattlePhase int

const (
	turnStart BattlePhase = iota
	turnAction
	turnEnd
)

// -------------------------------------------------------------------------
// -------------------------------battle loop------------------------------
// -------------------------------------------------------------------------

func startBattle() {
	currentPlayer.ResetBattleState()
	currentBoss.ResetBattleState()
	currentPlayer.state = battle
	turnOrder = initBattle()

	getAllTurnsBeforeLoop(currentPlayer.stats.speed, currentBoss.stats.speed)

	battleStartedMsg := GetGameTextBattle("start")
	currentEnemyMsg := GetGameTextBattle("currentenemy")
	fmt.Println(battleStartedMsg)
	fmt.Println(currentEnemyMsg)

	statusBoss()
	printTurnOrderSequence()
	checkCurrentTurn()
}

func checkCurrentTurn() {
	turnIndex := (turnOrder.currentLoopTurn - 1) % turnOrder.turnsBeforeLoop

	if turnIndex < len(turnOrder.turnSequence) {
		turnType := turnOrder.turnSequence[turnIndex]

		turnOrder.currentLoopTurn++

		if turnOrder.currentLoopTurn > turnOrder.turnsBeforeLoop {
			turnOrder.currentLoopTurn = 1
		}

		if currentPlayer.state != battle {
			return
		}

		if currentPlayer.CheckDefeat() {
			currentPlayer.HandleDefeat()
			return
		}

		if currentBoss.CheckDefeat() {
			currentBoss.HandleDefeat()
			return
		}

		if turnType == "player" {
			playerTurn()
		} else {
			bossTurn()
		}
	} else {
		internal := GetGameTextError("internal")
		internalturnoutofbounds := GetGameTextError("internalturnoutofbounds")
		fmt.Printf("%s: %s", internal, internalturnoutofbounds)
	}
}

// -------------------------------------------------------------------------
// -------------------------------Battle State------------------------------
// -------------------------------------------------------------------------

// TurnOrder comment
type TurnOrder struct {
	currentLoopTurn int
	turnsBeforeLoop int
	turnSequence    []string
	currentTurn     int
}

func (to TurnOrder) String() string {
	return fmt.Sprintf("TurnOrder {current_turn: %d, current_loop_turn: %d, turns_before_loop: %d,turn_sequence: %s}",
		to.currentTurn, to.currentLoopTurn, to.turnsBeforeLoop, to.turnSequence)
}

func initBattle() TurnOrder {
	to := TurnOrder{
		currentLoopTurn: 1,
		turnsBeforeLoop: 0,
		turnSequence:    make([]string, 0),
		currentTurn:     1,
	}
	return to
}

func updateTurnOrderCurrentTurn() {
	turnOrder.currentTurn++
}

// -------------------------------------------------------------------------
// -------------------------------Battle Order Calculation------------------------------
// -------------------------------------------------------------------------

func getAllTurnsBeforeLoop(playerSpeed, bossSpeed int) {
	turnOrder.currentLoopTurn = 1
	lcm := LCM(playerSpeed, bossSpeed)
	turnOrder.turnsBeforeLoop = (lcm / playerSpeed) + (lcm / bossSpeed)
	turnOrder.turnSequence = make([]string, 0, turnOrder.turnsBeforeLoop)

	playerTurnCounter := bossSpeed
	bossTurnCounter := playerSpeed

	for i := 0; i < turnOrder.turnsBeforeLoop; i++ {
		if playerSpeed == bossSpeed { // if equal player starts
			if i%2 == 0 {
				turnOrder.turnSequence = append(turnOrder.turnSequence, "player")
			} else {
				turnOrder.turnSequence = append(turnOrder.turnSequence, "boss")
			}
		} else {
			if bossTurnCounter < playerTurnCounter {
				turnOrder.turnSequence = append(turnOrder.turnSequence, "boss")
				bossTurnCounter += playerSpeed
			} else if playerTurnCounter < bossTurnCounter {
				turnOrder.turnSequence = append(turnOrder.turnSequence, "player")
				playerTurnCounter += bossSpeed
			} else { // if equal
				if bossSpeed > playerSpeed {
					turnOrder.turnSequence = append(turnOrder.turnSequence, "boss")
					bossTurnCounter = playerSpeed
					turnOrder.turnSequence = append(turnOrder.turnSequence, "player")
					playerTurnCounter = bossSpeed
					i++ // 2 turns
				} else {
					turnOrder.turnSequence = append(turnOrder.turnSequence, "player")
					playerTurnCounter = bossSpeed
					turnOrder.turnSequence = append(turnOrder.turnSequence, "boss")
					bossTurnCounter = playerSpeed
					i++ // 2 turns
				}
			}
		}
	}
}

func printTurnOrderSequence() {
	turnMsg := GetGameTextBattle("turn")
	turnOrdnerMsg := GetGameTextBattle("turnorder")
	separator2Msg := GetGameTextGameMessage("separator2")
	loopMsg := GetGameTextBattle("loop")

	fmt.Print("\n" + separator2Msg + "\n")
	fmt.Printf("\n%s: \n", turnOrdnerMsg)

	for i, turn := range turnOrder.turnSequence {
		fmt.Printf("  %s %d: %s\n", turnMsg, i+1, turn)
	}

	fmt.Println("  --- " + loopMsg + " ---\n\n")

	// fmt.Print("\n" + separator2Msg + "\n\n")
}

// GCD greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// LCM find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

/*
playerTurn    7 14 21 28 35 42 49 56 63 70 77 84 91 98 105 112 119 126 133
bossTurn         15     30    45    60    75    90     105      120     135
lcm                                                    105

playerSpeed = 7
bossSpeed = 15
lcm = 105
turns = lcm / playerSpeed + lcm / bossSpeed = 22

Turn  1: bossTurn --- playerTurn = 15 bossTurn = 7
Turn  2: bossTurn --- playerTurn = 15 bossTurn = 14
Turn  3: playerTurn - playerTurn = 15 bossTurn = 21
Turn  4: bossTurn --- playerTurn = 30 bossTurn = 21
Turn  5: bossTurn --- playerTurn = 30 bossTurn = 28
Turn  6: playerTurn - playerTurn = 30 bossTurn = 35
Turn  7: bossTurn --- playerTurn = 45 bossTurn = 35
Turn  8: bossTurn --- playerTurn = 45 bossTurn = 42
Turn  9: playerTurn - playerTurn = 45 bossTurn = 49
Turn 10: bossTurn --- playerTurn = 60 bossTurn = 49
Turn 11: bossTurn --- playerTurn = 60 bossTurn = 56
Turn 12: playerTurn - playerTurn = 60 bossTurn = 63
Turn 13: bossTurn --- playerTurn = 75 bossTurn = 63
Turn 14: bossTurn --- playerTurn = 75 bossTurn = 70
Turn 15: playerTurn - playerTurn = 75 bossTurn = 77
Turn 16: bossTurn --- playerTurn = 90 bossTurn = 77
Turn 17: bossTurn --- playerTurn = 90 bossTurn = 84
Turn 18: playerTurn - playerTurn = 90 bossTurn = 91
Turn 19: bossTurn --- playerTurn = 105 bossTurn = 91
Turn 20: bossTurn --- playerTurn = 105 bossTurn = 98
Turn 21: bossTurn --- playerTurn = 105 bossTurn = 105
Turn 22: playerTurn - playerTurn =  bossTurn = 7

*/
