package internal

import (
	"fmt"
	"strings"
)

//todo check/rework help.go file, since it was created by AI -.-'

// ------------------------------------------------------------
// --------------------------- COMMAND DOCS ---------------------------
// ------------------------------------------------------------

type CommandDoc struct {
	Name        string
	Description string
	Detail      string
	Usage       string
	Example     string
}

func helpCommand() {
	headerMsg := GetGameTextHelp("header")
	usageMsg := GetGameTextHelp("usage")
	exampleMsg := GetGameTextHelp("example")
	noteMsg := GetGameTextHelp("note")
	footerMsg := GetGameTextHelp("footer")
	availablecommandsMsg := GetGameTextGameMessage("availablecommands")

	commandDocs := GetAllGameTextCommands()

	maxNameLen := 0 // Get longest command name for padding
	for _, cmd := range commandDocs {
		if len(cmd.Name) > maxNameLen {
			maxNameLen = len(cmd.Name)
		}
	}

	fmt.Print("\n" + headerMsg + "\n\n")
	fmt.Println(availablecommandsMsg + "\n")

	for _, cmd := range commandDocs {
		fmt.Printf("%-*s", maxNameLen+2, cmd.Name)
		fmt.Printf("  %s\n", cmd.Description)
		fmt.Printf("%-*s  %s: %s\n", maxNameLen+2, "", usageMsg, cmd.Usage)
		fmt.Printf("%-*s  %s: %s\n\n", maxNameLen+2, "", exampleMsg, cmd.Example)
	}

	fmt.Print(noteMsg)
	fmt.Print("\n\n" + footerMsg + "\n")
}

// checkHelpCommand replaces the previous helpSpecificCommand
// It provides context-sensitive help based on the argument type
func checkHelpCommand(argument string) {
	// Convert to lowercase for case-insensitive comparisons
	argument = strings.ToLower(argument)

	switch {
	case isCommand(argument):
		// Help for a specific command
		helpForCommand(argument)
	case isSkillEffect(argument):
		// Help for a specific skill effect
		helpForSkillEffect(argument)
	case argument == "effects" || argument == "skilleffects":
		// Help for all skill effects categories
		helpForAllEffects()
	case argument == "talentpoints" || argument == "costs":
		// Help for talent point costs
		helpForTalentPointCosts()
	case argument == "stats":
		// Help for stats and their costs
		helpForStats()
	case argument == "difficulty":
		// Help for difficulty levels
		helpForDifficulty()
	case argument == "gameloop" || argument == "loop":
		// Help for game loop and states
		helpForGameLoop()
	default:
		// No matching help topic found
		invalidHelpMsg := GetGameTextError("invalidhelp")
		fmt.Printf("%s %s\n", invalidHelpMsg, argument)
	}
}

// Check if the argument is a valid command
func isCommand(argument string) bool {
	commandDocs := GetAllGameTextCommands()
	_, exists := commandDocs[argument]

	// Also check for command aliases
	aliases := map[string]string{
		"new":    "newplayer",
		"use":    "useskill",
		"player": "newplayer",
		"skill":  "newskill",
	}

	if alias, hasAlias := aliases[argument]; hasAlias {
		_, exists = commandDocs[alias]
	}

	return exists
}

// Check if the argument is a valid skill effect
func isSkillEffect(argument string) bool {
	effect := GetGameTextEffect(argument)
	return effect.Name != ""
}

// Display help for a specific command
func helpForCommand(command string) {
	// Handle command aliases
	aliases := map[string]string{
		"new":    "newplayer",
		"use":    "useskill",
		"player": "newplayer",
		"skill":  "newskill",
	}

	if alias, hasAlias := aliases[command]; hasAlias {
		command = alias
	}

	cmd := GetGameTextCommand(command)

	headerMsg := GetGameTextHelp("header")
	commandMsg := GetGameTextHelp("command")
	descriptionMsg := GetGameTextHelp("description")
	detailMsg := GetGameTextHelp("detail")
	usageMsg := GetGameTextHelp("usage")
	exampleMsg := GetGameTextHelp("example")
	footerMsg := GetGameTextHelp("footer")

	fmt.Print("\n" + headerMsg + "\n\n")
	fmt.Printf("%s: %s\n", commandMsg, cmd.Name)
	fmt.Printf("%s: %s\n", descriptionMsg, cmd.Description)
	fmt.Printf("%s: %s\n", detailMsg, cmd.Detail)
	fmt.Printf("%s: %s\n", usageMsg, cmd.Usage)
	fmt.Printf("%s: %s\n", exampleMsg, cmd.Example)
	fmt.Print("\n" + footerMsg + "\n")
}

// Display help for a specific skill effect
func helpForSkillEffect(effectName string) {
	effect := GetGameTextEffect(effectName)

	headerMsg := GetGameTextHelp("header")
	footerMsg := GetGameTextHelp("footer")

	effectMsg := GetGameTextGameMessage("effect")
	descriptionMsg := GetGameTextHelp("description")
	talentpointscostsMsg := GetGameTextGameMessage("talentpointscosts")

	fmt.Print("\n" + headerMsg + "\n\n")
	fmt.Printf("%s: %s\n", effectMsg, effect.Name)
	fmt.Printf("%s: %s\n", descriptionMsg, effect.Description)
	fmt.Printf("%s: %d\n", talentpointscostsMsg, effect.Cost)

	// Show which skill types can use this effect
	var validSkillTypes []string

	// Logic to determine valid skill types for this effect
	// This is a simplified approach; ideally, you would get this from your effect definitions
	switch effectName {
	case "directdamage", "pierce", "finisher", "buffturnbonusdamage", "debuffturnbonusdamage",
		"directheal", "lifeleech", "cleanse", "dispel":
		validSkillTypes = []string{"immediate"}
	case "extendbuffs", "extenddebuffs", "reducedebuffs", "reducebuffs":
		validSkillTypes = []string{"immediate", "passive"}
	case "dot", "stun", "blockbuffs", "grievouswounds":
		validSkillTypes = []string{"duration"}
	case "blockdebuffs", "healovertime", "incpower", "shield", "reflectdamage", "evasion", "critrate",
		"damagereduction":
		validSkillTypes = []string{"duration", "passive"}
	default:
		validSkillTypes = []string{"unknown"}
	}

	typeMsg := GetGameTextGameMessage("type")
	fmt.Printf("Valid %s: %s\n", typeMsg, strings.Join(validSkillTypes, ", "))

	fmt.Print("\n" + footerMsg + "\n")
}

// Display help for all skill effect categories
func helpForAllEffects() {
	headerMsg := GetGameTextHelp("header")
	footerMsg := GetGameTextHelp("footer")
	effectsMsg := GetGameTextGameMessage("effects")

	// Define effect categories
	categories := map[string][]string{
		"Direct Damage": {
			"directdamage", "pierce", "finisher", "buffturnbonusdamage", "debuffturnbonusdamage",
		},
		"Support": {
			"directheal", "lifeleech", "cleanse", "dispel", "extendbuffs", "extenddebuffs", "reducedebuffs", "reducebuffs",
		},
		"Buffs": {
			"blockdebuffs", "healovertime", "incpower", "shield", "reflectdamage", "evasion", "critrate",
		},
		"Debuffs": {
			"dot", "stun", "damagereduction", "blockbuffs", "grievouswounds",
		},
	}

	fmt.Print("\n" + headerMsg + "\n\n")
	fmt.Printf("%s Categories:\n\n", effectsMsg)

	for category, effects := range categories {
		fmt.Printf("--- %s ---\n", category)
		for _, effectName := range effects {
			effect := GetGameTextEffect(effectName)
			fmt.Printf("  %-20s - %s\n", effect.Name, effect.Description)
		}
		fmt.Println()
	}

	noteMsg := GetGameTextHelp("note")
	fmt.Printf("%s: Type 'help <effectname>' for detailed information about a specific effect\n", noteMsg)
	fmt.Print("\n" + footerMsg + "\n")
}

// Display help for talent point costs
func helpForTalentPointCosts() {
	headerMsg := GetGameTextHelp("header")
	footerMsg := GetGameTextHelp("footer")

	fmt.Print("\n" + headerMsg + "\n\n")
	fmt.Println("Talent Point Costs Overview:\n")

	// Skills costs
	fmt.Println("--- Skill Costs ---")
	fmt.Println("  Damage Multiplier: 1 talent point per 0.1 increase")
	fmt.Println("  Duration: 5 talent points per turn")
	fmt.Println("  Passive Multiplier: Effect cost x10 for passive skills\n")

	// Stats costs
	fmt.Println("--- Stats Costs ---")
	fmt.Println("  Health: 1 talent point per HP")
	fmt.Println("  Power: 1 talent point per Power")
	fmt.Println("  Speed: 2 talent points per Speed\n")

	// Effects costs - basic overview
	fmt.Println("--- Effects Costs ---")
	fmt.Println("  Basic Effect: 5 talent points")
	fmt.Println("  Duration Effect: Basic cost + (5 x duration)")
	fmt.Println("  Passive Effect: Basic cost x10\n")

	fmt.Println("Use 'help effects' to see all available effects and their descriptions.")
	fmt.Print("\n" + footerMsg + "\n")
}

// Display help for stats
func helpForStats() {
	headerMsg := GetGameTextHelp("header")
	footerMsg := GetGameTextHelp("footer")

	fmt.Print("\n" + headerMsg + "\n\n")
	fmt.Println("Character Stats:\n")

	// Stats explanation
	fmt.Println("--- Primary Stats ---")
	fmt.Printf("  %-15s - Maximum health points. Costs 1 talent point per HP.\n", "Health")
	fmt.Printf("  %-15s - Base damage for skills. Costs 1 talent point per Power.\n", "Power")
	fmt.Printf("  %-15s - Determines turn order. Costs 2 talent points per Speed.\n\n", "Speed")

	// Derived stats
	fmt.Println("--- Derived Stats ---")
	fmt.Printf("  %-15s - Your actual damage = Power x Skill Damage Multiplier.\n", "Total Power")
	fmt.Printf("  %-15s - Remaining talent points for enhancing stats or creating skills.\n\n", "Talent Points")

	// Combat mechanics
	fmt.Println("--- Combat Mechanics ---")
	fmt.Printf("  %-15s - Higher speed acts more frequently (see 'help gameloop').\n", "Turn Order")
	fmt.Printf("  %-15s - Effects apply at the start, during, or end of turns based on their type.\n", "Effect Timing")

	fmt.Print("\n" + footerMsg + "\n")
}

// Display help for difficulty levels
func helpForDifficulty() {
	headerMsg := GetGameTextHelp("header")
	footerMsg := GetGameTextHelp("footer")

	difficulties := GetGameTextDifficulty()

	fmt.Print("\n" + headerMsg + "\n\n")
	fmt.Println("Difficulty Levels and Talent Points:\n")

	// Compile difficulty information
	diffData := []struct {
		name         string
		message      string
		talentpoints int
		multiplier   string
	}{
		{difficulties.Normal.Name, difficulties.Normal.Message, getTalentpoints(normal), "x1.0"},
		{difficulties.Hard.Name, difficulties.Hard.Message, getTalentpoints(hard), "x1.3/1.2/1.1"},
		{difficulties.Expert.Name, difficulties.Expert.Message, getTalentpoints(expert), "x1.6/1.4/1.2"},
		{difficulties.Master.Name, difficulties.Master.Message, getTalentpoints(master), "x2.0/1.7/1.4"},
		{difficulties.Torment.Name, difficulties.Torment.Message, getTalentpoints(torment), "x2.5/2.0/1.6"},
	}

	for _, diff := range diffData {
		fmt.Printf("--- %s ---\n", diff.name)
		fmt.Printf("  Starting Talent Points: %d\n", diff.talentpoints)
		fmt.Printf("  Boss Multiplier (HP/Power/Speed): %s\n", diff.multiplier)
		fmt.Printf("  \"%s\"\n\n", diff.message)
	}

	fmt.Println("Each difficulty increases boss stats by an additional percentage per boss level.")
	fmt.Println("Higher difficulties provide fewer talent points but a greater challenge.")

	fmt.Print("\n" + footerMsg + "\n")
}

// Display help for game loop and states
func helpForGameLoop() {
	headerMsg := GetGameTextHelp("header")
	footerMsg := GetGameTextHelp("footer")

	fmt.Print("\n" + headerMsg + "\n\n")
	fmt.Println("Game Loop and States:\n")

	// Game states
	fmt.Println("--- Game States ---")
	fmt.Println("  Lobby - Initial state or when character is dead")
	fmt.Println("  Resting - When a character is selected and not in battle")
	fmt.Println("  Battle - During combat with a boss\n")

	// Game loop
	fmt.Println("--- Game Loop ---")
	fmt.Println("  1. Start in Lobby: Create a character or select existing one")
	fmt.Println("  2. Enter Rest state: Upgrade stats, create skills")
	fmt.Println("  3. Enter Battle state: Fight against the next boss")
	fmt.Println("  4. If victory: Gain talent points, return to Rest state")
	fmt.Println("  5. If defeat: Character dies, return to Lobby state\n")

	// Turn order explanation
	fmt.Println("--- Turn Order System ---")
	fmt.Println("  - Higher speed means more frequent turns")
	fmt.Println("  - Turn order is calculated based on player and boss speed")
	fmt.Println("  - The system creates a sequence that repeats after a cycle")
	fmt.Println("  - Example: Player(Speed 7) vs Boss(Speed 15):")
	fmt.Println("    Turn order: Boss, Player, Boss, Player, Boss, ...\n")

	fmt.Println("Your goal is to defeat all 9 bosses in sequence without dying.")

	fmt.Print("\n" + footerMsg + "\n")
}

// Legacy function kept for compatibility
func helpSpecificCommand(command string) {
	checkHelpCommand(command)
}
