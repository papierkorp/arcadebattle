# Arcade Battle

This game should be a Textbased Game which completly works in the Console/Terminal with Commands written in pure GOlang.

At the beginning the player has to create his own Character. Depending on the difficulty he gets a different amount of Talentpoints.
The Talentpoints are used to: upgrade Player stats (e.g. health, strength, speed), create skills (each skill has stats and effects) or buy talismans (one time use items). The Player can create as many skills as he wants/has Talentpoints.
The game is based on the effects which should enable strategic depth and a paper, rock, scissors system.
After winning against a boss he gets new Talentpoints.
The Game has 9 bosses, after beating the last boss you won the game. If you loose one fight your Character is dead and you have to create new one.


# Gameloop

- lobby state => start the game
  - create new character => rest state
  - select existing character => rest state
  - exit => close game
- rest state
  - create skill
  - buy talisman
  - update skill
  - update stats
  - start battle => battle state
  - lobby => lobby state
- battle state
  - `use skill1` ...
  - `use talisman` ...
  - win against boss: get new talentpoints => rest state
  - loose against boss => lobby state

# States + Commands

- lobby // when the game starts, when no character is selected, when character is dead, when `lobby` command
- resting // when a character is selected
- battle

**command list**

- `status`
  - Current Stage:
  - Current Selected Character:
  - ... different output depending on current stage (see each stage)
- `lobby`
- `exit`
- `create`
- `use`
- `list`
- `update`
- `battle`

**lobby**

- `status`
  - List of Characters
- `create character` => [Create a Character](#creation)
- `list character`
- `use characterx/idx`

**resting**

- `status`
  - Current stats
  - Current available talentpoints
  - List of Skills
  - Current stage (boss)
- `update stat x`
- `create skill` => [Create a Skill](#creation)
- `battle`

**battle**

- `status`
  - Current stats
  - List of Skills
  - Current enemy stats
  - Enemy skill list
- `use skill`
- `use talisman`

# Creation

**character_creation**

step by step:

- Name: string
- Difficulty: string
- for each stat

**skill_creation**

step by step:

- Name: string
- DamageMultiplier: float
- Effect: string (see list Effects)
- Duration: int (turns for each effect)

# Stats

each fixed stat should cost a certain amount of talentpoints

- Name
- Health
- strength
- Speed
- Talentpoints


# Skills

eacht stat should cost a certain amout of talentpoints
each effect should cost a certain amount of talentpoints

## Effects

see [gameidea_effects](#gameide_effects)

# Statistics

i want to track:

- overall time
- time for each boss (entered fight until defeat), time until each boss is beaten from starting the game
- amount of commands overall
- amount of characters written
- amount of skills created
- amount of talentpoints used
