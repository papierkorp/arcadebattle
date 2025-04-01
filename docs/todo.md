# ToDo

Based on the current code and the read me file give me a guideline on how i proceed with the next steps in form of a short to do list to progress in the game development:

LIST OF TO DOs

- [Project/Game Structure](#project-game-structure)
- [State System](#state-system)
- [Player Management](#player-management)
- [Skill Framework](#skill-framework)
- [Battle State](#battle-state)
- [Boss Progression](#boss-progression)
- [Statistics & Progression Tracking](#statistics-progression-tracking)
- [QOL Features](#qol-features)
- [Game Balance/Configuration](#game-balance-json-configuration)

## Project / Game Structure

- basic game loop with exit handling
- command parser with error handling


## State System

- state machine
- 3 states
- handle transitions

## Player Management

- create a character system
- stat management (health, power, speed)
- talent point system
- character persistence (saving/loading)
- character selection

## Skill Framework

- basic skill struct
- create a skill system (interface?)
  - skill validation
  - cost calculations
  - cooldown system
- implement effect system
- skill upgrading system

## Battle State

- implement turn based mechanics
- skill usage implementation
- status effect handling
- victory/defeat conditions
- basic boss ai behavior
- effect duration tracking
- effect stacking rules
- effect cancellation/cleansing

## Boss Progression

- create 9 boss templates
- difficulty scaling
- reward system

## Statistics & Progression Tracking

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

## QOL Features

- help system with command documentation
- game state saving/loading
- character management (multiple character)
- Tutorial
- Statistics Export

## Game Balance / Json Configuration

- json templates
- balance talent point costs
- tune damage multipliers
- balance boss difficulty

## Future Features

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
