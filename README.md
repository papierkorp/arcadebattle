# Arcadebattle

A text-based RPG adventure entirely in your console, written in pure Go.

## Game Overview

Arcadebattle is a console-based RPG where you create and customize your own hero. In this dangerous world, one defeat means death - so build your character wisely! Strategically allocate talent points to strengthen your stats and craft unique skills with various effects. Battle through increasingly difficult bosses, earning new talent points with each victory.

## Features

- **Character Creation**: Build your unique hero with custom stats
- **Dynamic Difficulty**: Choose from 5 difficulty levels, each offering different starting talent points
- **Skill Crafting**: Create and customize skills with various effects:
  - Direct damage, healing, and support effects
  - Duration-based effects for damage-over-time and buffs
  - Passive skills for constant bonuses
- **Boss Progression**: Fight through 9 unique bosses with increasing difficulty
- **Permadeath**: One defeat means your character is gone forever

## How to Play

### Installation

```bash
# Clone the repository
git clone https://github.com/yourusername/arcadebattle.git
cd arcadebattle

# Run the game
go run main.go

# Or build an executable
go build -o arcadebattle.exe
```

### Game Loop

1. **Lobby State**: Create your character or select an existing one
2. **Rest State**: Upgrade stats, create skills, or prepare for battle
3. **Battle State**: Strategically use your skills to defeat bosses

### Basic Commands

- `help` - Display available commands
- `status` - Check your current stats and skills
- `new player <name> <difficulty> <health> <power> <speed>` - Create a character
- `new skill <type> <name> [parameters]` - Create a new skill
- `battle` - Start a battle with the next boss
- `use <skill>` - Use a skill during battle
- `exit` - Exit the game

## Skill System

Skills come in three types:
- **Immediate**: Direct effects that happen instantly
- **Duration**: Effects that last several turns
- **Passive**: Permanent bonuses that are always active

Each skill costs talent points based on its power and effects. Choose wisely!

## Development

```bash
# Run the game with hot reloading
go run main.go

# Build for Windows
GOOS=windows GOARCH=amd64 go build -o game.exe

# Run tests
go test ./...
```

## License

This project is completely free and open source. You are free to:
- Use, share, and copy the game for any purpose
- Modify and adapt the code as you wish
- Distribute your own versions
- Use it as a basis for your own projects

No attribution required. No restrictions whatsoever. Enjoy the game and do whatever you want with it!

---

Created by [https://github.com/papierkorp] - Happy gaming!
