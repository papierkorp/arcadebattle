# Effect Usage

**Turn Process for use skill**

1. check effects for turn start
- execution
- Bleeding
2. `use skill` command
3. loop if an effect is blocked (effect.type block)
- Block Debuffs
- BlockBuffs
4. loop for special effects (`effect.probability`)
- evasion
- Resistance
- ReduceHealing
- Confusion
- Distraction
5. determine `basicSkillPower` (`currentPower * skillMulti`)
6. loop effects of both entities to create `fullSkillPower` (`currentPower * skillMulti * effectMulti`) (can be more effect multis, e.g. 0,5 for damage reduction or 2,0 for crit) with `effect.multi`
- Finisher
- Immunity
- Adapation
- critical strike
- IncPower
- Fury
- Weaken
- Vulnerability
7. loop effects of both enitites for special effects (e.g. Lifeleech, confusion, reflect damage) and excecute the effects (`effect.execute()`)
- Lifeleech
- reflect damage
8. use fullSkillPower to do the damage
9. loop effects for turn end and reduction of turns of effects
- Heal Over Time (HOT)

**Turn Process for use talisman**

1. check effects for turn start
- execution
2. `use talisman` command
- BuffTurnBonusDamage
- DebuffTurnBonusDamage
- Heal
- BuffHeal
- HealthCleanse
- Cleanse
- Dispel
- ExtendBuffs
- ExtendDebuffs
- ReduceBuffs
- ReduceDebuffs
...

**example struct**

```go
type SkillEffect struct {
  name                 string
  description          string
  talentpointCosts     int
  probability          float32
  category             EffectCategory
  execute              func()
  checkCondition       func() bool
  usageTiming          EffectTiming
  multi                float32
}

type EffectCategory int
const (
  ecfIncrease EffectCategory = iota
  ecDecrease
  ecBlock
  ecRemove
)

type EffectTiming int
const (
  etOnTurnStart EffectTiming = iota
  etOnSkillUse
  etOnTurnEnd
  etOnDurationEnd
)

```


# Effectslist

https://github.com/papierkorp/arcadebattle/blob/1c62a25f4e3e0624f7643347efaf1ef25b9b21d0/docs/gameidea.md

- Direct Damage
    - Execution: `Immediately kill the enemy while below 10% health`
    - Finisher: `adds bonus damage if enemy is low`
    - BuffTurnBonusDamage: `Bonus Damage to amount of all remaining Buff Turns`
    - DebuffTurnBonusDamage: `Bonus Damage to Amount of all remaining Debuff Turns of Enemy`
- Recovery
    - Heal: `Immediately restores health based on power stat`
    - Heal Over Time (HOT): `Restores health at the start of each turn`
    - Lifeleech: `Gain 50% of your Damage in Health`
    - BuffHeal: `Stronger direct heal but remove one random buff`
    - HealthCleanse: `Pay 10% of your current health to remove a random Debuff`
- Status Management
    - Cleanse: `Removes all negative effects from the target`
    - Dispel: `Removes all positive effects from the target`
    - ExtendBuffs: `+1 Turn for each 10 Total Power`
    - ExtendDebuffs: `+1 Turn for each 10 Total Power`
    - ReduceBuffs: `-1 Turn for each 10 Total Power`
    - ReduceDebuffs: `-1 Turn for each 10 Total Power`
- Defensive Buffs
    - Block Debuffs: `Prevents new debuffs from being applied while active`
    - evasion: `50% Chance to dont get damage`
    - Resistance: `50% Chance to block an incoming Debuff`
    - Immunity: `Receive 50% less damage`
    - Adapation: `Receive 10% less Damage from repeated sources`
- Offensive Buffs
    - critical strike: `50% Chance to double the damage`
    - IncPower: `Increases damage output by 50%`
    - Fury: `each attack increases damage of next attack`
    - reflect damage: `Reflect 50% of the damage you Receive`
- Debuffs
    - Bleeding: `Applies a damaging effect that deals damage at the start of each turn`
    - Weaken: `Reduces target's damage output by 50%`
    - BlockBuffs: `Prevents the target from receiving buffs and healing effects`
    - ReduceHealing: `Reduces all healing received by 50%`
    - Vulnerability: `Receive 50% more damage`
    - Confusion: `50% Chance to miss the skill`
    - Distraction: `50% Chance to attack itself`


