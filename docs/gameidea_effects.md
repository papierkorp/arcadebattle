# Effects new grouping

**all**

- PrimaryFunction
  - increase
  - decrease
  - block
  - remove
- Target
  - self
  - enemy
- EffectTiming
  - OnTurnStart
  - OnSkillUse
  - OnTurnEnd
  - OnDurationEnd
- probability
- effectMultiplier
  - totalBuffTurnCount 
  - totalBuffCount      
  - totalDebuffTurnCount
  - totalDebuffCount   
- Cost (does it cost something to activate, can be null, dynamically available for everything, can be used to reduce talentpointcosts)
  - nothing (default, just cast the effect)
  - remove a buff
  - decrease some health
- CostValue (integer, default 0, can only be used if Cost != nothing)
- Category
  - buff
  - debuff
- talentpointCosts

**Increase/Decrease**

- PrimaryFunction
  - increase
  - decrease
- Target
  - self
  - enemy
- affectedStat
    - calculatedDamage
    - currentHealth
    - currentPower
    - activeEffect
    - randomActiveEffect
    - allActiveEffects
- EffectTiming
  - OnTurnStart
  - OnSkillUse
  - OnTurnEnd
  - OnDurationEnd
- effectMultiplier
  - totalBuffTurnCount 
  - totalBuffCount      
  - totalDebuffTurnCount
  - totalDebuffCount   
- Cost (does it cost something to activate, can be null, dynamically available for everything, can be used to reduce talentpointcosts)
  - nothing (default, just cast the effect)
  - remove a buff
  - decrease some health
- CostValue (integer, default 0, can only be used if Cost != nothing)
- Category
  - buff
  - debuff
- talentpointCosts



**Block**

- blockType
  - Increase
  - Decrease
  - Remove
- affectedStat
    - calculatedDamage
    - currentHealth
    - currentPower
    - activeEffect
    - randomActiveEffect
    - newActiveEffect
    - allActiveEffects
    - useSkill
    - (randomSkill)
- EffectTiming
  - OnSkillUse
- Category
  - buff
  - debuff


**Remove**

- affectedStat
    - currentHealth
    - activeEffect
    - randomActiveEffect
    - allActiveEffects
- EffectTiming
  - OnSkillUse
- probability
- effectMultiplier
- Cost
- CostValue
- Category
  - buff
  - debuff
- talentpointCosts



|           stat          | increase | decrease | block | remove |
|-------------------------|----------|----------|-------|--------|
| currentHealth           | +        | +        | +     | +      |
| currentPower            | +        | +        | +     | +      |
| activeEffectTurns       | +        | +        | +     | +      |
| randomActiveEffectTurns | +        | +        | +     | +      |
| newActiveEffect         | -        | -        | +     | -      |
| allActiveEffectsTurns   | +        | +        | +     | +      |
| randomSkill             | -        | -        | +     | -      |


\pagebreak




# Effects new grouping - table

<br />

**template**

| name | PrimaryFunction | affectedStat | Target | Category | EffectTiming | Trigger | effectMultiplier | Cost | CostValue | blockType | talentpointCosts | fehlt |
|------|-----------------|--------------|--------|----------|--------------|---------|------------------|------|-----------|-----------|------------------|-------|
|      |                 |              |        |          |              |         |                  |      |           |           |                  |       |

---
---
---

|          name         | PrimaryFunction |   affectedStat   | Target | Category |   EffectTiming  |        Trigger         |   effectMultiplier   | Cost | CostValue | blockType |     fehlt      |
|-----------------------|-----------------|------------------|--------|----------|-----------------|------------------------|----------------------|------|-----------|-----------|----------------|
| Execution             | remove          | currentHealth    | enemy  | debuff   | OnOtherSkillUse | enemy below 10% health | -                    | -    | -         | -         | -              |
| Finisher              | increase        | calculatedDamage | self   | buff     | OnOtherSkillUse | enemy below 10% health | 100%                 | -    | -         | -         | -              |
| BuffTurnBonusDamage   | increase        | calculatedDamage | self   | buff     | OnOtherSkillUse | -                      | totalBuffTurnCount   | -    | -         | -         | -              |
| DebuffTurnBonusDamage | increase        | calculatedDamage | self   | buff     | OnOtherSkillUse | -                      | totalDebuffTurnCount | -    | -         | -         | -              |
| Heal                  | increase        | currentHealth    | self   | buff     | OnThisSkillUse  | -                      | -                    | -    | -         | -         | wieviel health |
| Heal Over Time        | increase        | currentHealth    | self   | buff     | OnTurnEnd       | -                      | -                    | -    | -         | -         | wieviel health |
| Lifeleech             | increase        | currentHealth    |        |          |                 |                        |                      |      |           |           |                |
| BuffHeal              |                 |                  |        |          |                 |                        |                      |      |           |           |                |
| HealthCleanse         |                 |                  |        |          |                 |                        |                      |      |           |           |                |
| Cleanse               |                 |                  |        |          |                 |                        |                      |      |           |           |                |
| Dispel                |                 |                  |        |          |                 |                        |                      |      |           |           |                |
| ExtendBuffs           |                 |                  |        |          |                 |                        |                      |      |           |           |                |
| ExtendDebuffs         |                 |                  |        |          |                 |                        |                      |      |           |           |                |
| ReduceBuffs           |                 |                  |        |          |                 |                        |                      |      |           |           |                |
| ReduceDebuffs         |                 |                  |        |          |                 |                        |                      |      |           |           |                |
| Block Debuffs         |                 |                  |        |          |                 |                        |                      |      |           |           |                |
| Evasion               |                 |                  |        |          |                 |                        |                      |      |           |           |                |
| Shield                |                 |                  |        |          |                 |                        |                      |      |           |           |                |
| Resistance            |                 |                  |        |          |                 |                        |                      |      |           |           |                |
| Immunity              |                 |                  |        |          |                 |                        |                      |      |           |           |                |
| Adaptation            |                 |                  |        |          |                 |                        |                      |      |           |           |                |
| Critical Strike       |                 |                  |        |          |                 |                        |                      |      |           |           |                |
| IncPower              |                 |                  |        |          |                 |                        |                      |      |           |           |                |
| Fury                  |                 |                  |        |          |                 |                        |                      |      |           |           |                |
| Reflect Damage        |                 |                  |        |          |                 |                        |                      |      |           |           |                |
| Bleeding              |                 |                  |        |          |                 |                        |                      |      |           |           |                |
| Weaken                |                 |                  |        |          |                 |                        |                      |      |           |           |                |
| BlockBuffs            |                 |                  |        |          |                 |                        |                      |      |           |           |                |
| ReduceHealing         |                 |                  |        |          |                 |                        |                      |      |           |           |                |
| Vulnerability         |                 |                  |        |          |                 |                        |                      |      |           |           |                |
| Confusion             |                 |                  |        |          |                 |                        |                      |      |           |           |                |
| Distraction           |                 |                  |        |          |                 |                        |                      |      |           |           |                |


\pagebreak

# Effects

https://github.com/papierkorp/arcadebattle/blob/1c62a25f4e3e0624f7643347efaf1ef25b9b21d0/docs/gameidea.md

- Direct Damage
    - Execution: `Immediately kill the enemy while below 10% health`
    - Finisher: `adds bonus damage if enemy is low`
    - BuffTurnBonusDamage: `Bonus Damage to amount of all remaining Buff Turns`
    - DebuffTurnBonusDamage: `Bonus Damage to Amount of all remaining Debuff Turns of Enemy`
    - Pierce: `ingore shield effects`
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
    - Shield: `Create a shield that absorbs damage equal to 25% of max health`
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


