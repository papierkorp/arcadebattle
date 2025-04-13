package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

// Usage examples:
/*
   gameMsg := GetGameTextGameMessage("welcome")
   helpMsg := GetGameTextHelp("header")
   errorMsg := GetGameTextError("invalidcommand")
   cmdMsg := GetGameTextCommand("help")
   effectMsg := GetGameTextEffect("stun")
   battleMsg := GetGameTextBattle("start")
   invalid := GetGameTextEffect("nonexistent")
*/

// -------------------------- Init --------------------------

type GameTextEffect struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Cost        int    `json:"cost"`
}

type GameTextCommand struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Detail      string `json:"detail"`
	Usage       string `json:"usage"`
	Example     string `json:"example"`
}

type GameTextStatus struct {
	Player GameTextStatusPlayer `json:"player"`
	Boss   GameTextStatusBoss   `json:"boss"`
}

type GameTextStatusPlayer struct {
	Name                   string `json:"name"`
	Health                 string `json:"health"`
	Power                  string `json:"power"`
	Speed                  string `json:"speed"`
	Skilllist              string `json:"skilllist"`
	Talentpoints_Total     string `json:"talentpointstotal"`
	Talentpoints_Remaining string `json:"talentpointsremaining"`
	Difficulty             string `json:"difficulty"`
	Bosses                 string `json:"bosses"`
	Alive                  string `json:"alive"`
	State                  string `json:"state"`
	ActiveEffectList       string `json:"activeeffectlist"`
}

type GameTextStatusBoss struct {
	Name             string `json:"name"`
	Health           string `json:"health"`
	Power            string `json:"power"`
	Speed            string `json:"speed"`
	Skilllist        string `json:"skilllist"`
	ActiveEffectList string `json:"activeeffectlist"`
}

type GameTextNamedMessage struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

type GameTextState struct {
	Idle   GameTextNamedMessage `json:"idle"`
	Battle GameTextNamedMessage `json:"battle"`
	Dead   GameTextNamedMessage `json:"dead"`
}

type GameTextDifficulty struct {
	Normal  GameTextNamedMessage `json:"normal"`
	Hard    GameTextNamedMessage `json:"hard"`
	Expert  GameTextNamedMessage `json:"expert"`
	Master  GameTextNamedMessage `json:"master"`
	Torment GameTextNamedMessage `json:"torment"`
}

type GameText struct {
	Error        map[string]string          `json:"error"`
	Help         map[string]string          `json:"help"`
	GameMessage  map[string]string          `json:"gamemessage"`
	Status       GameTextStatus             `json:"status"`
	SkillEffects map[string]GameTextEffect  `json:"skilleffects"`
	CommandDocs  map[string]GameTextCommand `json:"commanddocs"`
	State        GameTextState              `json:"state"`
	Difficulty   GameTextDifficulty         `json:"difficulty"`
	Battle       map[string]string          `json:"battle"`
}

func initGametext() error {
	jsonFile, err := os.ReadFile("../internal/gametext_en.json")
	if err != nil {
		return fmt.Errorf("unable to open gametext_en.json file: %s", err)
	}

	if len(jsonFile) == 0 {
		return fmt.Errorf("gametext_en.json file is empty")
	}

	if !json.Valid(jsonFile) {
		return fmt.Errorf("gametext_en.json contains invalid JSON")
	}

	decoder := json.NewDecoder(bytes.NewBuffer(jsonFile))
	if err := decoder.Decode(&gameText); err != nil {
		return fmt.Errorf("failed to decode gametext_en.json: %s", err)
	}

	if gameText.Error == nil {
		return fmt.Errorf("error messages are missing")
	}
	if gameText.Help == nil {
		return fmt.Errorf("help messages are missing")
	}
	if gameText.GameMessage == nil {
		return fmt.Errorf("game messages are missing")
	}
	if gameText.Battle == nil {
		return fmt.Errorf("battle messages are missing")
	}
	if gameText.SkillEffects == nil {
		return fmt.Errorf("skill effects are missing")
	}
	if gameText.CommandDocs == nil {
		return fmt.Errorf("command documentation is missing")
	}

	return nil
}

// -------------------------- Get All --------------------------

func GetGameTextDifficulty() GameTextDifficulty {
	return gameText.Difficulty
}

func GetGameTextState() GameTextState {
	return gameText.State
}

func GetGameTextStatus() GameTextStatus {
	return gameText.Status
}

func GetGameTextStatusPlayer() GameTextStatusPlayer {
	return gameText.Status.Player
}

func GetGameTextStatusBoss() GameTextStatusBoss {
	return gameText.Status.Boss
}

func GetAllGameTextCommands() map[string]GameTextCommand {
	return gameText.CommandDocs
}

// -------------------------- Get key --------------------------

func GetGameTextGameMessage(key string) string {
	return gameText.GameMessage[key]
}

func GetGameTextHelp(key string) string {
	return gameText.Help[key]
}

func GetGameTextError(key string) string {
	return gameText.Error[key]
}

func GetGameTextBattle(key string) string {
	return gameText.Battle[key]
}

func GetGameTextCommand(key string) *GameTextCommand {
	cmd := gameText.CommandDocs[key]

	return &GameTextCommand{
		Name:        cmd.Name,
		Description: cmd.Description,
		Detail:      cmd.Detail,
		Usage:       cmd.Usage,
		Example:     cmd.Example,
	}
}

func GetGameTextEffect(key string) *GameTextEffect {
	effect := gameText.SkillEffects[key]

	return &GameTextEffect{
		Name:        effect.Name,
		Description: effect.Description,
		Cost:        effect.Cost,
	}
}

// todo commanddocs newskills => text update
