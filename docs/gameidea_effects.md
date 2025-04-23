# Effects new grouping

- name
- description
- talentpointCosts
- isBlockedBy (can the effect be blocked by another effect)
- probability (if there is a probability involved else 100%)
- affectedStat (see in PrimaryFunction which stats can be affected)
    - calculatedDamage
    - currentHealth
    - currentPower
    - activeEffect
    - randomActiveEffect
    - newActiveEffect
    - randomSkill
- damageBase (how is the damage calculated)
  - currentPower
  - totalBuffTurnCount 
  - totalBuffCount      
  - totalDebuffTurnCount
  - totalDebuffCount    
  - calculatedDamage
- PrimaryFunctionValue (always a % value)
- PrimaryFunction
  - increase
  - decrease
  - block
  - remove
- Target (who to target)
  - self
  - enemy
- Cost (does it cost something to activate, can be null)
  - nothing (just cast the effect)
  - remove a buff
  - decrease some health
- CostValue (can be 0)
- Category (to which category does it belong)
  - buff
  - debuff
  - damage
- EffectTiming (when is the effect applied) / (immediate, duration, passive)
  - OnTurnStart
  - OnSkillUse
  - OnTurnEnd
  - OnDurationEnd

\pagebreak

# Effects new grouping - table

<br />

**currentPower**

|      Name      | TP | IsBlockedBy | Probability | affectedStat |  damageBase  | PrimaryFunction | PrimaryFunctionValue | Target | CostType | CostValue | Category | EffectTiming |
|----------------|----|-------------|-------------|--------------|--------------|-----------------|----------------------|--------|----------|-----------|----------|--------------|
| Increase Power |  5 | Dispel      | 100%        | currentPower | currentPower | increase        | 50%                  | self   | -        | -         | buff     | OnTurnStart  |

**calculatedDamage**

|           Name           | TP |      IsBlockedBy      | Probability |   affectedStat   |      damageBase      | PrimaryFunction | PrimaryFunctionValue | Target | CostType | CostValue | Category | EffectTiming |
|--------------------------|----|-----------------------|-------------|------------------|----------------------|-----------------|----------------------|--------|----------|-----------|----------|--------------|
| Shield                   |  5 | Pierce, Dispel        | 100%        | calculatedDamage | currentPower         | absorb          | 25%                  | self   | -        | -         | buff     | OnTurnStart  |
| Evasion                  |  5 | Dispel                | 50%         | calculatedDamage | damagereceived       | decrease        | 100%                 | self   | -        | -         | buff     | OnTurnStart  |
| Direct Damage            |  5 | -                     | 100%        | calculatedDamage | currentPower         | increase        | 100%                 | enemy  | -        | -         | damage   | OnSkillUse   |
| Finisher                 |  5 | -                     | 100%        | calculatedDamage | currentPower         | increase        | 50%                  | enemy  | -        | -         | damage   | OnSkillUse   |
| Buff Turn Bonus Damage   |  5 | -                     | 100%        | calculatedDamage | totalBuffTurnsCount  | increase        | 5% per turn          | enemy  | -        | -         | damage   | OnSkillUse   |
| Debuff Turn Bonus Damage |  5 | -                     | 100%        | calculatedDamage | totalDebuffTurnCount | increase        | 5% per turn          | enemy  | -        | -         | damage   | OnSkillUse   |
| Critical Rate            |  5 | Dispel                | 50%         | calculatedDamage | currentPower         | increase        | 200%                 | self   | -        | -         | buff     | OnTurnStart  |
| Damage Over Time         |  5 | BlockDebuffs, Cleanse | 100%        | calculatedDamage | currentPower         | decrease        | 25%                  | enemy  | -        | -         | debuff   | OnTurnStart  |
| Damage Reduction         |  5 | BlockDebuffs, Cleanse | 100%        | calculatedDamage | currentPower         | decrease        | 50%                  | enemy  | -        | -         | debuff   | OnTurnStart  |
| Pierce                   |  5 | -                     | 100%        | calculatedDamage | currentPower         | bypass          | 50%                  | enemy  | -        | -         | damage   | OnSkillUse   |


**currentHealth**

|       Name      | TP |        IsBlockedBy        | Probability |  affectedStat |   damageBase   | PrimaryFunction | PrimaryFunctionValue | Target | CostType | CostValue | Category | EffectTiming |
|-----------------|----|---------------------------|-------------|---------------|----------------|-----------------|----------------------|--------|----------|-----------|----------|--------------|
| Direct Heal     |  5 | BlockBuffs                | 100%        | currentHealth | currentPower   | increase        | 50%                  | self   | -        | -         | buff     | OnSkillUse   |
| Life Leech      |  5 | BlockBuffs, ReduceHealing | 100%        | currentHealth | damagedealt    | increase        | 50%                  | self   | -        | -         | buff     | OnSkillUse   |
| Heal Over Time  |  5 | BlockBuffs, ReduceHealing | 100%        | currentHealth | currentPower   | increase        | 10%                  | self   | -        | -         | buff     | OnTurnStart  |
| Grievous Wounds |  5 | BlockDebuffs, Cleanse     | 100%        | currentHealth | currentPower   | decrease        | 50%                  | enemy  | -        | -         | debuff   | OnTurnStart  |
| Reflect Damage  |  5 | Dispel                    | 100%        | currentHealth | damagereceived | decrease        | 50%                  | enemy  | -        | -         | buff     | OnTurnStart  |

**activeEffect**

|      Name      | TP | IsBlockedBy  | Probability | affectedStat |  damageBase  | PrimaryFunction  | PrimaryFunctionValue | Target | CostType | CostValue | Category | EffectTiming |
|----------------|----|--------------|-------------|--------------|--------------|------------------|----------------------|--------|----------|-----------|----------|--------------|
| Cleanse        |  5 | -            | 100%        | activeEffect | -            | decrease(remove) | 100%                 | self   | -        | -         | buff     | OnSkillUse   |
| Dispel         |  5 | BlockDebuffs | 100%        | activeEffect | -            | decrease(remove) | 100%                 | enemy  | -        | -         | debuff   | OnSkillUse   |
| Extend Buffs   |  5 | -            | 100%        | activeEffect | currentPower | increase         | 10%                  | self   | -        | -         | buff     | OnSkillUse   |
| Extend Debuffs |  5 | BlockDebuffs | 100%        | activeEffect | currentPower | increase         | 10%                  | enemy  | -        | -         | debuff   | OnSkillUse   |
| Reduce Debuffs |  5 | -            | 100%        | activeEffect | currentPower | decrease         | 10%                  | self   | -        | -         | buff     | OnSkillUse   |
| Reduce Buffs   |  5 | BlockDebuffs | 100%        | activeEffect | currentPower | decrease         | 10%                  | enemy  | -        | -         | debuff   | OnSkillUse   |


**newActiveEffect**

|      Name     | TP |      IsBlockedBy      | Probability |   affectedStat  | damageBase | PrimaryFunction | PrimaryFunctionValue | Target | CostType | CostValue | Category | EffectTiming |
|---------------|----|-----------------------|-------------|-----------------|------------|-----------------|----------------------|--------|----------|-----------|----------|--------------|
| Block Debuffs |  5 | Dispel                | 100%        | newActiveEffect | -          | block           | 100%                 | self   | -        | -         | buff     | OnTurnStart  |
| Block Buffs   |  5 | BlockDebuffs, Cleanse | 100%        | newActiveEffect | -          | block           | 100%                 | enemy  | -        | -         | debuff   | OnTurnStart  |

**skill**

| Name | TP |      IsBlockedBy      | Probability | affectedStat | damageBase | PrimaryFunction | PrimaryFunctionValue | Target | CostType | CostValue | Category | EffectTiming |
|------|----|-----------------------|-------------|--------------|------------|-----------------|----------------------|--------|----------|-----------|----------|--------------|
| Stun |  5 | BlockDebuffs, Cleanse | 100%        | skill        | -          | block           | 100%                 | enemy  | -        | -         | debuff   | OnTurnStart  |

**template**

| Name | TP | IsBlockedBy | Probability | affectedStat | damageBase | PrimaryFunction | PrimaryFunctionValue | Target | CostType | CostValue | Category | EffectTiming |
|------|----|-------------|-------------|--------------|------------|-----------------|----------------------|--------|----------|-----------|----------|--------------|
|      |    |             |             |              |            |                 |                      |        |          |           |          |              |


\pagebreak

# Effects old grouping

- Immediate
  - Increase power/damage
    - Execution: `Immediately kill the enemy while below 10% health`
    - BuffTurnBonusDamage: `Bonus Damage to amount of all remaining Buff Turns`
    - Finisher: `adds bonus damage if enemy is low`
    - DebuffTurnBonusDamage: `Bonus Damage to Amount of all remaining Debuff Turns of Enemy`
  - Increase Current Health
    - Heal: `Restores health`
    - BuffHeal: `Stronger direct heal but remove one random buff`
    - BuffHeal2: `Heal based on the number of Buffs`
  - change Duration
    - ExtendBuffs: `+1 Turn for each 10 Total Power`
    - ExtendDebuffs: `+1 Turn for each 10 Total Power`
    - ReduceBuffs: `-1 Turn for each 10 Total Power`
    - ReduceDebuffs: `-1 Turn for each 10 Total Power`
  - remove effect
    - DebuffCleanse: `Remove the same number of Debuffs as the enemy currently has`
    - HealthCleanse: `Pay 10% of your current health to remove a random Debuff`
    - Cleanse: `Removes all negative effects from the target`
    - Dispel: `Removes all positive effects from the target`
  - special
    - Pierce: `ingore shield effects`
- Duration
  - Deal Damage
    - DOT: `Applies a damaging effect that deals damage at the start of each turn`
    - reflect damage: `Reflect 50% of the damage you Receive`
  - Increase power/damage
    - critical strike: `50% Chance to double the damage`
    - IncPower: `Increases damage output by 50%`
    - Fury: `each attack increases damage of next attack`
    - Vulnerability: `Receive 50% more damage`
  - Increase current health
    - Heal Over Time (HOT): `Restores health at the start of each turn`
    - Lifeleech: `Gain 50% of your Damage in Health`
  - mitigate damage
    - evasion: `50% Chance to dont get damage`
    - Immunity: `Receive 50% less damage`
    - Adapation: `Receive 10% less Damage from repeated sources`
    - Shield: `Create a shield that absorbs damage equal to 25% of max health`
    - Weaken: `Reduces target's damage output by 50%`
    - Confusion: `50% Chance to miss the skill`
  - mitigate effect
    - Block Debuffs: `Prevents new debuffs from being applied while active`
    - Resistance: `50% Chance to block an incoming Debuff`
    - BlockBuffs: `Prevents the target from receiving buffs`
    - ReduceHealing: `Reduces all healing received by 50%`
  - special
    - Distraction: `50% Chance to attack itself`

