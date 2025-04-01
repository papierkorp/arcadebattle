package internal

import (
	"fmt"
	"strings"
)

type BattleState struct {
	current_health      int
	total_turns_buffs   int
	total_turns_debuff  int
	is_stunned          bool
	active_effects_list []ActiveEffect
}

type ActiveEffect struct {
	skill_effect SkillEffect
	total_power  float32 //power*dmgmultiplier
	turns_left   int
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

func battleLoopPlayerTurn() {
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

type TurnOrder struct {
	current_loop_turn int
	turns_before_loop int
	turn_sequence     []string
	current_turn      int
}

func (bs TurnOrder) String() string {
	return fmt.Sprintf("TurnOrder {current_turn: %d, current_loop_turn: %d, turns_before_loop: %d,turn_sequence: %s}",
		bs.current_turn, bs.current_loop_turn, bs.turns_before_loop, bs.turn_sequence)
}

func initBattle() TurnOrder {
	bs := TurnOrder{
		current_loop_turn: 1,
		turns_before_loop: 0,
		turn_sequence:     make([]string, 0),
		current_turn:      1,
	}
	return bs
}

func updateTurnOrderCurrentTurn() {
	turn_order.current_turn += 1
}

// -------------------------------------------------------------------------
// -------------------------------Battle Order Calculation------------------------------
// -------------------------------------------------------------------------

func getAllTurnsBeforeLoop(playerSpeed, bossSpeed int) {
	turn_order.current_loop_turn = 1
	lcm := LCM(playerSpeed, bossSpeed)
	turn_order.turns_before_loop = (lcm / playerSpeed) + (lcm / bossSpeed)
	turn_order.turn_sequence = make([]string, 0, turn_order.turns_before_loop)

	playerTurnCounter := bossSpeed
	bossTurnCounter := playerSpeed

	fmt.Println("Player Speed:", playerSpeed)
	fmt.Println("Boss Speed:", bossSpeed)

	for i := 0; i < turn_order.turns_before_loop; i++ {
		if playerSpeed == bossSpeed { // if equal player starts
			if i%2 == 0 {
				turn_order.turn_sequence = append(turn_order.turn_sequence, "player")
			} else {
				turn_order.turn_sequence = append(turn_order.turn_sequence, "boss")
			}
		} else {
			if bossTurnCounter < playerTurnCounter {
				turn_order.turn_sequence = append(turn_order.turn_sequence, "boss")
				bossTurnCounter += playerSpeed
			} else if playerTurnCounter < bossTurnCounter {
				turn_order.turn_sequence = append(turn_order.turn_sequence, "player")
				playerTurnCounter += bossSpeed
			} else { // if equal
				if bossSpeed > playerSpeed {
					turn_order.turn_sequence = append(turn_order.turn_sequence, "boss")
					bossTurnCounter = playerSpeed
					turn_order.turn_sequence = append(turn_order.turn_sequence, "player")
					playerTurnCounter = bossSpeed
					i++ // 2 turns
				} else {
					turn_order.turn_sequence = append(turn_order.turn_sequence, "player")
					playerTurnCounter = bossSpeed
					turn_order.turn_sequence = append(turn_order.turn_sequence, "boss")
					bossTurnCounter = playerSpeed
					i++ // 2 turns
				}
			}
		}
	}

	fmt.Printf("Initial Battle State: %s\n", turn_order)
}

func checkCurrentTurn() {
	turnIndex := (turn_order.current_loop_turn - 1) % turn_order.turns_before_loop

	if turnIndex < len(turn_order.turn_sequence) {
		turnType := turn_order.turn_sequence[turnIndex]

		turn_order.current_loop_turn++

		if turn_order.current_loop_turn > turn_order.turns_before_loop {
			turn_order.current_loop_turn = 1
		}

		if turnType == "player" {
			battleLoopPlayerTurn() // => playerAction
		} else {
			bossAction()
		}
	} else {
		internal := GetGameTextError("internal")
		internalturnoutofbounds := GetGameTextError("internalturnoutofbounds")
		fmt.Errorf("%s: %s", internal, internalturnoutofbounds)
	}
}

func printTurnOrderSequence() {
	turnMsg := GetGameTextBattle("turn")
	turnOrdnerMsg := GetGameTextBattle("turnorder")
	separator2Msg := GetGameTextGameMessage("separator2")
	loopMsg := GetGameTextBattle("loop")

	fmt.Print("\n" + separator2Msg + "\n")
	fmt.Printf("\n%s: \n", turnOrdnerMsg)

	for i, turn := range turn_order.turn_sequence {
		fmt.Printf("  %s %d: %s\n", turnMsg, i+1, turn)
	}

	fmt.Println("  --- " + loopMsg + " ---")

	fmt.Print("\n" + separator2Msg + "\n\n")
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
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
