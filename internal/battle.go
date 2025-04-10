// Package internal comment
package internal

import (
	"fmt"
	"strings"
)

// BattleState comment
type BattleState struct {
	currentHealth      int
	totalTurnsBuffs    int
	totalTurnsDebuff   int
	activeEffectsList  []ActiveEffect
	currentBattlePhase BattlePhase
}

// BattlePhase comment
type BattlePhase int

const (
	turnStart BattlePhase = iota
	turnAction
	turnEnd
)

// ActiveEffect comment
type ActiveEffect struct {
	skillEffect SkillEffect
	totalPower  int //power*dmgmultiplier
	turnsLeft   int
}

// -------------------------------------------------------------------------
// -------------------------------battle loop------------------------------
// -------------------------------------------------------------------------

func startBattle() {
	current_player.state = battle
	turn_order = initBattle()

	getAllTurnsBeforeLoop(current_player.stats.speed, current_boss.stats.speed)

	battleStartedMsg := GetGameTextBattle("start")
	currentEnemyMsg := GetGameTextBattle("currentenemy")
	fmt.Println(battleStartedMsg)
	fmt.Println(currentEnemyMsg)

	statusBoss()
	printTurnOrderSequence()
	checkCurrentTurn()
}

func playerTurn() {
	invalidCommandMsg := GetGameTextError("invalidcommand")
	gamestarthelpMsg := GetGameTextGameMessage("gamestarthelp")
	battlepromptMsg := GetGameTextBattle("battleprompt")

	fmt.Print("player turn - ")
	fmt.Printf("%s\n", turn_order)

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
		case "quit", "exit":
			leaveBattle()
			return
		case "use":
			validCommand = useCommand(commandArgs)

		default:
			fmt.Println(invalidCommandMsg)
			fmt.Println(gamestarthelpMsg)
			validCommand = false
		}

		if validCommand {
			break
		}
	}

	updateTurnOrderCurrentTurn()
	promptMsg := GetGameTextGameMessage("prompt")
	rl.SetPrompt(promptMsg)
	checkCurrentTurn()
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
	turn_order.currentTurn++
}

// -------------------------------------------------------------------------
// -------------------------------Battle Order Calculation------------------------------
// -------------------------------------------------------------------------

func getAllTurnsBeforeLoop(playerSpeed, bossSpeed int) {
	turn_order.currentLoopTurn = 1
	lcm := LCM(playerSpeed, bossSpeed)
	turn_order.turnsBeforeLoop = (lcm / playerSpeed) + (lcm / bossSpeed)
	turn_order.turnSequence = make([]string, 0, turn_order.turnsBeforeLoop)

	playerTurnCounter := bossSpeed
	bossTurnCounter := playerSpeed

	fmt.Println("Player Speed:", playerSpeed)
	fmt.Println("Boss Speed:", bossSpeed)

	for i := 0; i < turn_order.turnsBeforeLoop; i++ {
		if playerSpeed == bossSpeed { // if equal player starts
			if i%2 == 0 {
				turn_order.turnSequence = append(turn_order.turnSequence, "player")
			} else {
				turn_order.turnSequence = append(turn_order.turnSequence, "boss")
			}
		} else {
			if bossTurnCounter < playerTurnCounter {
				turn_order.turnSequence = append(turn_order.turnSequence, "boss")
				bossTurnCounter += playerSpeed
			} else if playerTurnCounter < bossTurnCounter {
				turn_order.turnSequence = append(turn_order.turnSequence, "player")
				playerTurnCounter += bossSpeed
			} else { // if equal
				if bossSpeed > playerSpeed {
					turn_order.turnSequence = append(turn_order.turnSequence, "boss")
					bossTurnCounter = playerSpeed
					turn_order.turnSequence = append(turn_order.turnSequence, "player")
					playerTurnCounter = bossSpeed
					i++ // 2 turns
				} else {
					turn_order.turnSequence = append(turn_order.turnSequence, "player")
					playerTurnCounter = bossSpeed
					turn_order.turnSequence = append(turn_order.turnSequence, "boss")
					bossTurnCounter = playerSpeed
					i++ // 2 turns
				}
			}
		}
	}

	fmt.Printf("Initial Battle State: %s\n", turn_order)
}

func checkCurrentTurn() {
	turnIndex := (turn_order.currentLoopTurn - 1) % turn_order.turnsBeforeLoop

	if turnIndex < len(turn_order.turnSequence) {
		turnType := turn_order.turnSequence[turnIndex]

		turn_order.currentLoopTurn++

		if turn_order.currentLoopTurn > turn_order.turnsBeforeLoop {
			turn_order.currentLoopTurn = 1
		}

		if current_player.state != battle {
			return
		}

		if current_player.state == dead {
			current_player.HandleDefeat()
			return
		}

		if current_boss.CheckDefeat() {
			current_boss.HandleDefeat()
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

func printTurnOrderSequence() {
	turnMsg := GetGameTextBattle("turn")
	turnOrdnerMsg := GetGameTextBattle("turnorder")
	separator2Msg := GetGameTextGameMessage("separator2")
	loopMsg := GetGameTextBattle("loop")

	fmt.Print("\n" + separator2Msg + "\n")
	fmt.Printf("\n%s: \n", turnOrdnerMsg)

	for i, turn := range turn_order.turnSequence {
		fmt.Printf("  %s %d: %s\n", turnMsg, i+1, turn)
	}

	fmt.Println("  --- " + loopMsg + " ---")

	fmt.Print("\n" + separator2Msg + "\n\n")
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
