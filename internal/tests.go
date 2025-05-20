package internal

import (
	"fmt"
	"strings"
)

var TestCases = map[string]func(){
	"player":   testPlayer,
	"player2":  testPlayer2,
	"player3":  testPlayer3,
	"player4":  testPlayerEmpty,
	"player5":  testPlayerNoName,
	"player6":  testPlayerInvalidDiff,
	"player7":  testPlayerNegativeHP,
	"player8":  testPlayerNegativePWR,
	"player9":  testPlayerNegativeSPD,
	"player10": testPlayerZeroStats,
	"player11": testPlayerTooManyTP,
	"player12": testPlayerSpecialChar,

	"skill":  testSkill,
	"skill2": testSkill2,
	"skill3": testSkill3,
	"skill4": testSkill4,

	"seq": testSeq,
}

func RunTest(testCase string) {
	if testFunc, exists := TestCases[testCase]; exists {
		testFunc()
	} else {
		fmt.Printf("Unknown test case: %s\n", testCase)
	}
}

func ExecuteTest(testInput string) {
	fmt.Printf("Testinput: `%s`\n\n", testInput)
	newCommand(strings.Fields(strings.ToLower(testInput)))
}

// ----------------------------------------------------------------
// Player creation tests
// ----------------------------------------------------------------

func testPlayer() {
	testInput := "new player John normal 25 15 8"
	ExecuteTest(testInput)
}

func testPlayer2() {
	testInput := "new player Sammy hard 10 5 7"
	ExecuteTest(testInput)
}

func testPlayer3() {
	testInput := "new player Luna torment 1 1 1"
	ExecuteTest(testInput)
}

func testPlayerEmpty() {
	testInput := "new player"
	ExecuteTest(testInput)
}

func testPlayerNoName() {
	testInput := "new player \"\" normal 10 10 10"
	ExecuteTest(testInput)
}

func testPlayerInvalidDiff() {
	testInput := "new player Bob impossible 10 10 10"
	ExecuteTest(testInput)
}

func testPlayerNegativeHP() {
	testInput := "new player Evil normal -50 10 10"
	ExecuteTest(testInput)
}

func testPlayerNegativePWR() {
	testInput := "new player Weak normal 50 -10 10"
	ExecuteTest(testInput)
}

func testPlayerNegativeSPD() {
	testInput := "new player Slow normal 50 10 -5"
	ExecuteTest(testInput)
}

func testPlayerZeroStats() {
	testInput := "new player Zero normal 0 0 0"
	ExecuteTest(testInput)
}

func testPlayerTooManyTP() {
	testInput := "new player Greedy normal 1000 1000 1000"
	ExecuteTest(testInput)
}

func testPlayerSpecialChar() {
	testInput := "new player Test!@#$% normal 10 10 10"
	ExecuteTest(testInput)
}

// ----------------------------------------------------------------
// Immediate skill tests
// ----------------------------------------------------------------

func testSkill() {
	testInput := "new skill Lifeleech 0.7 3 heal1"
	ExecuteTest(testInput)
}

func testSkill2() {
	testInput := "new skill Lifeleech2 0.7 3 Lifeleech"
	ExecuteTest(testInput)
}

func testSkill3() {
	testInput := "new skill PowerUp 0.7 3 increasepower1"
	ExecuteTest(testInput)
}

func testSkill4() {
	testInput := "new skill Poison 0.7 3 dot"
	ExecuteTest(testInput)
}

// ----------------------------------------------------------------
// Sequence tests
// ----------------------------------------------------------------

func testSeq() {
	ExecuteTest("new player John normal 25 15 8")
	ExecuteTest("new skill Lifeleech 0.7 9 heal1")
	// ExecuteTest("status")
	// ExecuteTest("battle")
	// ExecuteTest("use skill lifeleech")
}
