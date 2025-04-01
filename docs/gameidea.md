# Arcade Battle

This game should be a Textbased Game which completly works in the Console with Commands written in pure GOlang.

At the beginning the player has to create his own Character. Depending on the difficulty he gets a different amount of Talentpoints.
The Talentpoints are used to: upgrade Player stats (e.g. health, power) and create skills (each skill has stats and effects). The Player can create as many skills as he wants/has Talentpoints.
After winning against a boss he gets new Talentpoints.
The Game has 9 bosses, after beating the last boss you won the game. If you loose one fight your Character is dead and you have to create new one.


# Gameloop

- lobby state => start the game
  - create new character => rest state
  - select existing character => rest state
  - exit => close game
- rest state
  - create skill
  - update skill
  - update stats
  - start battle => battle state
  - lobby => lobby state
- battle state
  - `use skill1` ...
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
- `use skillx/slotx`

# Creation

**character_creation**

step by step:

- Name: string
- for each stat

**skill_creation**

step by step:

- Name: string
- DamageMultiplier: float
- Effect: string (see list Effects)
- Duration: int (turns for each effect)
- Cooldown: int (turns)

# Stats

_Fixed_

each fixed stat should cost a certain amount of talentpoints

- Name => set by player
- Health => set by player
- Power => set by player
- Speed => set by player
- Talentpoints => based on difficulty
  - normal
  - hard
  - expert
  - master
  - torment

_Dynamic_

- currentHealth => at the start the same as Health
- Alive

# Skills

eacht stat should cost a certain amout of talentpoints
each effect should cost a certain amount of talentpoints

- Skills
  - json template
    - Name: string
    - DamageMultiplier: float
    - Effect: string (see list Effects)
    - Duration: int (turns for each effect)
    - Cooldown: int (turns)
  - Effects
    - Direct
      - Damage
        - Direct Damage: `Deals immediate damage based on power stat and multiplier`
        - Pierce: `Ignores shield effects and deals partial damage directly to health`
        - Finisher: `adds bonus damage if enemy is low`
        - BuffTurnBonusDamage: `Bonus Damage to amount of all remaining Buff Turns`
        - DebuffTurnBonusDamage: `Bonus Damage to Amount of all remaining Debuff Turns of Enemy`
      - Support
        - Direct Heal: `Immediately restores health based on power stat`
        - Lifeleech: `Gain 50% of your Damage in Health`
        - Cleanse: `Removes all negative effects from the target`
        - Dispel: `Removes all positive effects from the target`
        - ExtendBuffs: `+1 Turn for each 10 Total Power`
        - ExtendDebuffs: `+1 Turn for each 10 Total Power`
        - ReduceDebuffs: `-1 Turn for each 10 Total Power`
        - ReduceBuffs: `-1 Turn for each 10 Total Power`
    - Over Time
      - Buff
        - Block Debuffs: `Prevents new debuffs from being applied while active`
        - Heal Over Time (HOT): `Restores health at the start of each turn`
        - IncPower: `Increases damage output by 50%`
        - Shield: `Creates a defensive barrier that blocks direct damage and some effects`
        - reflect damage: `Reflect 50% of the damage you Receive`
        - evasion: `50% Chance to dont get damage`
        - critical strike: `50% Chance to doubl`
      - Debuff
        - Damage Over Time (DOT): `Applies a damaging effect that deals damage at the start of each turn`
        - Stun: `Prevents the target from taking any actions`
        - Damage Reduction: `Reduces target's damage output by 50%`
        - Block Buffs: `Prevents the target from receiving buffs and healing effects`
        - Grievous Wounds: `Reduces all healing received by 50%`
      - ...
        - taunt/provoke: `basic skill`
        - Counterattack: `basic skill`
        - silence: `basic skill`
  
Grouping:

- Damage
- Heal
- Debuff
- Buff
- Special


- Damage (DD, DOT, Pierce)
- Healing (HOT, Direct Heal, Lifeleech)
- Recovery (Cleanse, Block Debuffs)
- Protection (Shield, Counterattack)
- Enhancement (Buff, Damage Reduction)
- Control (Stun, Block Buffs)
- Disruption (Dispel, Grievous Wounds)
- Special


# Statistics

i want to track:

- overall time
- time for each boss (entered fight until defeat), time until each boss is beaten from starting the game
- amount of commands overall
- amount of characters written
- amount of skills created
- amount of talentpoints used


# Terms

Total Power = `power * dmgMultiplier * [activeSkillMultiplier] * [passiveSkillMultiplier]`
