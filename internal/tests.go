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

	"skill":   testSkillImmediate,
	"skill2":  testSkillImmediateMulti,
	"skill3":  testSkillImmediateEmpty,
	"skill4":  testSkillImmediateNoArg,
	"skill5":  testSkillImmediateInvalidEffect,
	"skill6":  testSkillImmediateNegMulti,
	"skill7":  testSkillImmediateHighMulti,
	"skill8":  testSkillImmediateEffectMismatch,
	"skill9":  testSkillImmediateNoPlayer,
	"skill10": testSkillImmediateSpecialChar,

	"skill20": testSkillDuration,
	"skill21": testSkillDurationMulti,
	"skill22": testSkillDurationNegDur,
	"skill23": testSkillDurationZeroDur,
	"skill24": testSkillDurationHighDur,
	"skill25": testSkillDurationEffectMismatch,

	"skill40": testSkillPassive,
	"skill41": testSkillPassiveMulti,
	"skill42": testSkillPassiveEffectMismatch,
	"skill43": testSkillPassiveEffectDMGMulti,
	"skill44": testSkillPassiveEffectDMGMultiDuration,

	"battle1": testBattleStart,
	"battle2": testBattleNoSkills,
	"battle3": testBattleUseSkill,

	"seq":  testValidSkillTypesSequence,
	"seq2": testInvalidSkillTypesSequence,
	"seq3": testFullGameSequence,
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

func testSkillImmediate() {
	testInput := "new skill immediate Fireball 1.5 directdamage"
	ExecuteTest(testInput)
}

func testSkillImmediateMulti() {
	testInput := "new skill immediate MegaBlast 2.0 directdamage pierce finisher"
	ExecuteTest(testInput)
}

func testSkillImmediateEmpty() {
	testInput := "new skill immediate \"\" 1.0 directdamage"
	ExecuteTest(testInput)
}

func testSkillImmediateNoArg() {
	testInput := "new skill immediate"
	ExecuteTest(testInput)
}

func testSkillImmediateInvalidEffect() {
	testInput := "new skill immediate Sparkle 1.0 madeupeffect"
	ExecuteTest(testInput)
}

func testSkillImmediateNegMulti() {
	testInput := "new skill immediate WeakAttack -0.5 directdamage"
	ExecuteTest(testInput)
}

func testSkillImmediateHighMulti() {
	testInput := "new skill immediate SuperNova 100.0 directdamage"
	ExecuteTest(testInput)
}

func testSkillImmediateEffectMismatch() {
	testInput := "new skill immediate WrongType 1.0 stun" // stun is for duration skills
	ExecuteTest(testInput)
}

func testSkillImmediateNoPlayer() {
	// Reset current player to test no player state
	oldPlayer := current_player
	current_player = Player{}

	testInput := "new skill immediate NoPlayerSkill 1.0 directdamage"
	ExecuteTest(testInput)

	// Restore player after test
	current_player = oldPlayer
}

func testSkillImmediateSpecialChar() {
	testInput := "new skill immediate Skill!@#$ 1.0 directdamage"
	ExecuteTest(testInput)
}

// ----------------------------------------------------------------
// Duration skill tests
// ----------------------------------------------------------------

func testSkillDuration() {
	testInput := "new skill duration Poison 0.7 3 dot"
	ExecuteTest(testInput)
}

func testSkillDurationMulti() {
	testInput := "new skill duration Nightmare 1.2 4 dot stun"
	ExecuteTest(testInput)
}

func testSkillDurationNegDur() {
	testInput := "new skill duration NegativeDur 1.0 -3 dot"
	ExecuteTest(testInput)
}

func testSkillDurationZeroDur() {
	testInput := "new skill duration ZeroDur 1.0 0 dot"
	ExecuteTest(testInput)
}

func testSkillDurationHighDur() {
	testInput := "new skill duration LongEffect 0.5 100 dot"
	ExecuteTest(testInput)
}

func testSkillDurationEffectMismatch() {
	testInput := "new skill duration WrongType 1.0 3 pierce" // pierce is for immediate skills
	ExecuteTest(testInput)
}

// ----------------------------------------------------------------
// Passive skill tests
// ----------------------------------------------------------------

func testSkillPassive() {
	testInput := "new skill passive Protection shield"
	ExecuteTest(testInput)
}

func testSkillPassiveMulti() {
	testInput := "new skill passive MegaDefense shield blockdebuffs damagereduction"
	ExecuteTest(testInput)
}

func testSkillPassiveEffectMismatch() {
	testInput := "new skill passive BadPassive dot" // dot is for duration skills
	ExecuteTest(testInput)
}

func testSkillPassiveEffectDMGMulti() {
	testInput := "new skill passive PassiveDuration 1 shield"
	ExecuteTest(testInput)
}

func testSkillPassiveEffectDMGMultiDuration() {
	testInput := "new skill passive PassiveDuration 1 1 shield"
	ExecuteTest(testInput)
}

// ----------------------------------------------------------------
// Battle tests
// ----------------------------------------------------------------

func testBattleStart() {
	// First create a player if none exists
	if current_player.name == "" {
		ExecuteTest("new player TestWarrior normal 100 50 20")
	}

	// Then start battle
	ExecuteTest("battle")
}

func testBattleNoSkills() {
	// Create new player with no skills
	ExecuteTest("new player NoSkillsGuy normal 100 50 20")

	// Start battle without creating any skills
	ExecuteTest("battle")
}

func testBattleUseSkill() {
	// Create player
	ExecuteTest("new player SkillUser normal 100 50 20")

	// Create a skill
	ExecuteTest("new skill immediate QuickAttack 1.5 directdamage")

	// Start battle
	ExecuteTest("battle")

	// Use the skill
	ExecuteTest("use skill QuickAttack")
}

// ----------------------------------------------------------------
// Sequence tests
// ----------------------------------------------------------------

func testFullGameSequence() {
	// Create player
	ExecuteTest("new player Adventurer normal 100 50 30")

	// Check status
	ExecuteTest("status")

	// Create skills
	ExecuteTest("new skill immediate Slash 1.2 directdamage")
	ExecuteTest("new skill duration Bleed 0.8 3 dot")
	ExecuteTest("new skill passive Toughness blockdebuffs")

	// Check status again
	ExecuteTest("status")

	// Start battle
	ExecuteTest("battle")

	// Use skills in battle
	ExecuteTest("use skill Slash")
}

func testValidSkillTypesSequence() {
	// Create player
	testPlayer()

	// Create one valid skill of each type
	testSkillImmediateMulti()
	testSkillDurationMulti()
	testSkillPassiveMulti()

	// Check status with created skills
	statusCommand([]string{"status"})
}

func testInvalidSkillTypesSequence() {
	// Create player
	testPlayer()

	// Create two invalid immediate skills
	testSkillImmediateEmpty()
	testSkillImmediateEffectMismatch()

	// Create two invalid duration skills
	testSkillDurationZeroDur()
	testSkillDurationEffectMismatch()

	// Create two invalid passive skills
	testSkillPassiveEffectMismatch()
	testSkillPassiveEffectDMGMulti()
	testSkillPassiveEffectDMGMultiDuration()

}
