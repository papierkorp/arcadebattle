package internal

import "github.com/chzyer/readline"

var rl *readline.Instance
var current_player Player
var current_boss Boss
var tempSkillCounter int = 1
var gameText GameText
var turn_order TurnOrder
