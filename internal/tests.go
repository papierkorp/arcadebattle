package internal

import (
	"fmt"
	"strings"
)

/*
Example playthrough:

1. Start game
2. Create player:
   > test player
   Creates: John (normal) with HP:10, ATK:5, SPD:5

3. Check status:
   > status
   Shows player stats and available talent points

4. Create skill:
   > test skill
   Creates: Fireball (DMG:1.5x, Duration:3, Effect:Direct Damage)

5. Create more skills:
   > new skill Snowstorm 0.7 2 stun
   > new skill Shield 1.0 2 shield
   > new skill PowerUp 1.0 3 buff

6. Start battle:
   > battle
   Shows boss stats and begins combat

7. Use skills in battle:
   > use fireball
   > use shield
   > use snowstorm

8. Win battle, gain talent points, repeat from step 4
*/

var TestCases = map[string]func(){
	"player":     testValidPlayer,
	"player2":    testValidPlayer2,
	"player3":    testValidPlayer3,
	"playerf":    testInvalidPlayer,
	"skill":      testValidSkill,
	"skill2":     testValidSkill2,
	"skill3":     testValidSkill3,
	"skill4":     testValidSkill4,
	"skill5":     testValidSkill5,
	"skillf":     testInvalidSkill,
	"playerfail": testInvalidPlayerStats,
	"skillfail":  testInvalidSkillDuration,
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

func testValidPlayer() {
	testInput := "new player John normal 250 150 80"
	ExecuteTest(testInput)
}

func testValidPlayer2() {
	testInput := "new player Sammy hard 10 5 7"
	ExecuteTest(testInput)
}

func testValidPlayer3() {
	testInput := "new player Luna torment 1 1 1"
	ExecuteTest(testInput)
}

func testValidSkill() {
	//new0 skill1 <skilltype2> <name3> <dmgmulti4> <duration5> [effect effect effect6...]
	testInput := "new skill duration Fire 1.5 6 dot stun"
	ExecuteTest(testInput)
}

func testValidSkill2() {
	testInput := "new skill duration poison 5 2 blockdebuffs healovertime"
	ExecuteTest(testInput)
}

func testValidSkill3() {
	testInput := "new skill passive hardening reducebuffs dispel"
	ExecuteTest(testInput)
}

func testValidSkill4() {
	testInput := "new skill immediate 3 blockdebuffs"
	ExecuteTest(testInput)
}

func testValidSkill5() {
	testInput := "new skill empty immediate"
	ExecuteTest(testInput)
}

func testInvalidPlayer() {
	testInput := "new player"
	ExecuteTest(testInput)
}

func testInvalidSkill() {
	testInput := "new skill Fireball -1.5 dot 3"
	ExecuteTest(testInput)
}

func testInvalidPlayerStats() {
	testInput := "new player John normal 1000 1000 1000"
	ExecuteTest(testInput)
}

func testInvalidSkillDuration() {
	testInput := "new skill Fireball 1.5 dot 0"
	ExecuteTest(testInput)
}

/* Additional test cases that could be implemented:

1. Player Creation Edge Cases:
   - Empty name
   - Invalid difficulty levels
   - Negative stat values
   - Zero stat values
   - Special characters in name

2. Skill Creation Edge Cases:
   - Empty skill name
   - Invalid effect types
   - Extremely high multipliers
   - Very long durations
   - Special characters in skill name

3. Battle System Tests:
   - Starting battle without skills
   - Starting battle with low health
   - Using non-existent skills
   - Using skills on cooldown

4. State Management Tests:
   - Transitioning between states
   - Death state handling
   - Victory state handling
   - State persistence

5. Resource Management Tests:
   - Talent point calculations
   - Skill cost validations
   - Health/damage calculations
*/
