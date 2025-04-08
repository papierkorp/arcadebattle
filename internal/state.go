// Package internal comment
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

// CheckCurrentState comment
func CheckCurrentState(s state) {
	switch s {
	case idle:
		StateIdle()
	case battle:
		StateBattle()
	case dead:
		StateDead()
	default:
		StateIdle()
	}

	err := checkCurrentBoss()
	if err != nil {
		fmt.Errorf("%s", err)
	}
}

// StateIdle comment
func StateIdle() {
	msg := GetGameTextState()
	fmt.Println("\n" + msg.Idle.Message + "\n")

}

// StateBattle comment
func StateBattle() {
	msg := GetGameTextState()
	fmt.Println("\n" + msg.Battle.Message + "\n")
}

// StateDead comment
func StateDead() {
	msg := GetGameTextState()
	fmt.Println("\n" + msg.Dead.Message + "\n")
}
