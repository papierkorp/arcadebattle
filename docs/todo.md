# ToDo

## battle / skill usage

**how battle works**

in checkCurrentTurn sets either `battleLoopPlayerTurn()` or `bossAction()`
in PlayerTurn i have to call player.battleState.currentBattlePhase = turn.start


**turn.start**
- on player
    - check all effects of player/boss (battleState)
- on skill
- on effect
    - use effect if necessary

**turn.mid**
- on player
    - add/delete effect
- on skill
    - use skill
    - apply dmgmulti
- on effect
    - check if effect can be applied (isBlockedBy)

**turn.end**
- on player
    - check duration of effects
    - reduce duration of effects
    - delete effect from player/boss
- on skill
- on effect

# ToDo AI

analyze the current available code and the docs and act like a scrum master to create a sprint to finish this projekt.
break the tasks left into a few Epics but dont go too deep since i will know what to do with the epics

Based on the current code and the read me file give me a guideline on how i proceed with the next steps in form of a short to do list to progress in the game development:

I've analyzed the codebase and documentation for your Arcadebattle text-based RPG. Here's a sprint plan that organizes the remaining work into several focused epics to help you complete the project efficiently.

-----

If I were to identify the most critical tasks to focus on first, they would be:

1. **Complete the skill effect implementations** - This is crucial for battle mechanics
- Complete the skill cost calculation functions in `balancing_calculation.go`
- Implement the effect functions in `skilleffects.go` (currently they all just print "asdf")
- Finish the skill upgrade functionality
- Test skill effects in battle scenarios
2. **Implement boss AI and battle logic** - Without this, the core gameplay loop doesn't function
- Complete the boss AI in `bosses.go` (currently does nothing substantive)
- Implement proper skill usage and effect application during battles
- Add status effect handling for both player and bosses
- Create the battle UI to clearly show what's happening each turn
- Implement proper battle conclusion logic (victory/defeat)
3. **Balance the talent point system** - Important for game progression and difficulty
- Finalize all skill effect costs in `balancing_calculation.go`
- Balance boss stat progressions by difficulty level
- Tune talent point awards after boss victories
- Create appropriate starting talent points for each difficulty

# General Plan AI

- Project/Game Structure
  - basic game loop with exit handling
  - command parser with error handling
- State System
  - state machine
  - 3 states
  - handle transitions
- Player Management
  - create a character system
  - stat management (health, power, speed)
  - talent point system
  - character persistence (saving/loading)
  - character selection
- Skill Framework
  - basic skill struct
  - create a skill system (interface?)
    - skill validation
    - cost calculations
    - cooldown system
  - implement effect system
  - skill upgrading system
- Battle State
  - implement turn based mechanics
  - skill usage implementation
  - status effect handling
  - victory/defeat conditions
  - basic boss ai behavior
  - effect duration tracking
  - effect stacking rules
  - effect cancellation/cleansing
- Boss Progression
  - create 9 boss templates
  - difficulty scaling
  - reward system
- Statistics & Progression Tracking
  - persistent time tracking
    - overall
    - per boss
  - persistent input tracking
    - commands overall
    - characters typed overall
  - persistent game tracking
    - amount of skills used
    - amount of skills created
    - talent point usage
- QOL Features
  - help system with command documentation
  - game state saving/loading
  - character management (multiple character)
  - Tutorial
  - Statistics Export
- Game Balance/Configuration
  - json templates
  - balance talent point costs
  - tune damage multipliers
  - balance boss difficulty


# Future Features

- Elementpoints => 8 total, 1 at the beginning, 1 after each defeated boss

- Types
  - 1 Element (5 level)
    - fire
    - earth
    - air
    - water
  - 2 elements (3 level)
    - fire + earth
    - fire + air
    - fire + water
    - earth + air
    - earth + water
    - air + water
  - 3/4 elements (2 level)
    - fire + earth + air
    - fire + earth + water
    - fire + air + water
    - earth + air + water
    - fire + earth + air + water

```json
{
  "elements": {
    "1": {
      "base_multiplier": 0.0,
      "elements": ["fire", "earth", "air", "water"],
      "allowed_level_upgrades": ["level1", "level2", "level3", "level4", "level5"]
    },
    "2": {
      "base_multiplier": 0.25,
      "combinations": [
        ["fire", "earth"],
        ["fire", "air"],
        ["fire", "water"],
        ["earth", "air"],
        ["earth", "water"],
        ["air", "water"]
      ],
      "allowed_level_upgrades": ["level1", "level2", "level3"]
    },
    "3": {
      "base_multiplier": 0.5,
      "combinations": [
        ["fire", "earth", "air"],
        ["fire", "earth", "water"],
        ["fire", "air", "water"],
        ["earth", "air", "water"]
      ],
      "allowed_level_upgrades": ["level1", "level2"]
    },
    "4": {
      "base_multiplier": 1.0,
      "combinations": [["fire", "earth", "air", "water"]],
      "allowed_level_upgrades": ["level1", "level2"]
    }
  },

  "levelupgrades": {
    "level1": {
      "cost": 1,
      "value": 0.5
    },
    "level2": {
      "cost": 2,
      "value": 1.0
    },
    "level3": {
      "cost": 3,
      "value": 2.0
    },
    "level4": {
      "cost": 5,
      "value": 4.0
    },
    "level5": {
      "cost": 10,
      "value": 8.0
    }
  }
}
```
