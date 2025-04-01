package internal

import "fmt"

type state int

const (
	idle state = iota + 1
	battle
	dead
)

func (s state) String() string {
	state := GetGameTextState()
	switch s {
	case idle:
		return state.Idle.Name
	case battle:
		return state.Battle.Name
	case dead:
		return state.Dead.Name
	default:
		return state.Idle.Name
	}
}

func (s state) EnumStateIndex() int {
	return int(s)
}

func check_current_state(s state) {
	switch s {
	case idle:
		state_idle()
	case battle:
		state_battle()
	case dead:
		state_dead()
	default:
		state_idle()
	}

	err := checkCurrentBoss()
	if err != nil {
		fmt.Errorf("%s", err)
	}
}

func state_idle() {
	msg := GetGameTextState()
	fmt.Println("\n" + msg.Idle.Message + "\n")

}

func state_battle() {
	msg := GetGameTextState()
	fmt.Println("\n" + msg.Battle.Message + "\n")
}

func state_dead() {
	msg := GetGameTextState()
	fmt.Println("\n" + msg.Dead.Message + "\n")
}
