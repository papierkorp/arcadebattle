# Effects new grouping

- talentpointCosts
- probability (if there is a probability involved else 100%)
- affectedStat (see in PrimaryFunction which stats can be affected)
    - calculatedDamage
    - currentHealth
    - currentPower
    - activeEffect
    - randomActiveEffect
    - newActiveEffect
    - allActiveEffects
    - (randomSkill)
- damageBase (how is the damage calculated)
  - currentPower
  - totalBuffTurnCount 
  - totalBuffCount      
  - totalDebuffTurnCount
  - totalDebuffCount    
  - calculatedDamage
- PrimaryFunctionValue (always a % value)
  - increase/decrease
    - % value
    - integer stat value
  - block
    - increase of affectedStat
    - decrease of affectedStat
  - remove
    - 0
    - 100%
    - activeEffect
    - randomActiveEffect
- PrimaryFunction (not available for each affectedStat)
  - increase
  - decrease
  - block
  - remove
- blockType
  - Increase
  - Decrease
  - Remove
- Target (who to target)
  - self
  - enemy
- Cost (does it cost something to activate, can be null, dynamically available for everything, can be used to reduce talentpointcosts)
  - nothing (default, just cast the effect)
  - remove a buff
  - decrease some health
- CostValue (integer, default 0, can only be used it Cost != nothing)
- Category (to which category does it belong, can be calculated => how? if increase...)
  - buff
  - debuff
  - damage
- EffectTiming (when is the effect applied)
  - OnTurnStart
  - OnSkillUse
  - OnTurnEnd
  - OnDurationEnd


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

| Name | TP | Probability | damageBase | PrimaryFunction | PrimaryFunctionValue | blockType | Target | CostType | CostValue | Category | EffectTiming |
|------|----|-------------|------------|-----------------|----------------------|-----------|--------|----------|-----------|----------|--------------|
|      |    |             |            |                 |                      |           |        |          |           |          |              |

---
---
---

**emsCalculatedDamage**

| Name | TP | Probability |      damageBase      | PrimaryFunction | PrimaryFunctionValue | Target | CostType | CostValue | Category | EffectTiming |
|------|----|-------------|----------------------|-----------------|----------------------|--------|----------|-----------|----------|--------------|
| x    | x  | x           | currentPower         | Increase        | x                    | x      | x        | x         | x        | x            |
| x    | x  | x           | currentPower         | Decrease        | x                    | x      | x        | x         | x        | x            |
| x    | x  | x           | currentPower         |                 |                      |        |          |           |          |              |
| x    | x  | x           | currentPower         |                 |                      |        |          |           |          |              |
| x    | x  | x           | currentPower         |                 |                      |        |          |           |          |              |
| x    | x  | x           | currentPower         |                 |                      |        |          |           |          |              |
| x    | x  | x           | currentPower         |                 |                      |        |          |           |          |              |
| x    | x  | x           | totalBuffTurnCount   |                 |                      |        |          |           |          |              |
| x    | x  | x           | totalBuffTurnCount   |                 |                      |        |          |           |          |              |
| x    | x  | x           | totalBuffTurnCount   |                 |                      |        |          |           |          |              |
| x    | x  | x           | totalBuffTurnCount   |                 |                      |        |          |           |          |              |
| x    | x  | x           | totalBuffTurnCount   |                 |                      |        |          |           |          |              |
| x    | x  | x           | totalBuffTurnCount   |                 |                      |        |          |           |          |              |
| x    | x  | x           | totalBuffCount       |                 |                      |        |          |           |          |              |
| x    | x  | x           | totalBuffCount       |                 |                      |        |          |           |          |              |
| x    | x  | x           | totalBuffCount       |                 |                      |        |          |           |          |              |
| x    | x  | x           | totalBuffCount       |                 |                      |        |          |           |          |              |
| x    | x  | x           | totalBuffCount       |                 |                      |        |          |           |          |              |
| x    | x  | x           | totalBuffCount       |                 |                      |        |          |           |          |              |
| x    | x  | x           | totalDebuffTurnCount |                 |                      |        |          |           |          |              |
| x    | x  | x           | totalDebuffTurnCount |                 |                      |        |          |           |          |              |
| x    | x  | x           | totalDebuffTurnCount |                 |                      |        |          |           |          |              |
| x    | x  | x           | totalDebuffTurnCount |                 |                      |        |          |           |          |              |
| x    | x  | x           | totalDebuffTurnCount |                 |                      |        |          |           |          |              |
| x    | x  | x           | totalDebuffTurnCount |                 |                      |        |          |           |          |              |
| x    | x  | x           | totalDebuffCount     |                 |                      |        |          |           |          |              |
| x    | x  | x           | totalDebuffCount     |                 |                      |        |          |           |          |              |
| x    | x  | x           | totalDebuffCount     |                 |                      |        |          |           |          |              |
| x    | x  | x           | totalDebuffCount     |                 |                      |        |          |           |          |              |
| x    | x  | x           | totalDebuffCount     |                 |                      |        |          |           |          |              |
| x    | x  | x           | totalDebuffCount     |                 |                      |        |          |           |          |              |
| x    | x  | x           | calculatedDamage     |                 |                      |        |          |           |          |              |
| x    | x  | x           | calculatedDamage     |                 |                      |        |          |           |          |              |
| x    | x  | x           | calculatedDamage     |                 |                      |        |          |           |          |              |
| x    | x  | x           | calculatedDamage     |                 |                      |        |          |           |          |              |
| x    | x  | x           | calculatedDamage     |                 |                      |        |          |           |          |              |
| x    | x  | x           | calculatedDamage     |                 |                      |        |          |           |          |              |

**emsCurrentHealth**

| Name | TP | Probability | damageBase | PrimaryFunction | PrimaryFunctionValue | Target | CostType | CostValue | Category | EffectTiming |
|------|----|-------------|------------|-----------------|----------------------|--------|----------|-----------|----------|--------------|
|      |    |             |            |                 |                      |        |          |           |          |              |

**emsCurrentPower**

| Name | TP | Probability | damageBase | PrimaryFunction | PrimaryFunctionValue | Target | CostType | CostValue | Category | EffectTiming |
|------|----|-------------|------------|-----------------|----------------------|--------|----------|-----------|----------|--------------|
|      |    |             |            |                 |                      |        |          |           |          |              |

**emsCurrentActiveEffect**

| Name | TP | Probability | damageBase | PrimaryFunction | PrimaryFunctionValue | Target | CostType | CostValue | Category | EffectTiming |
|------|----|-------------|------------|-----------------|----------------------|--------|----------|-----------|----------|--------------|
|      |    |             |            |                 |                      |        |          |           |          |              |

**emsRandomActiveEffect**

| Name | TP | Probability | damageBase | PrimaryFunction | PrimaryFunctionValue | Target | CostType | CostValue | Category | EffectTiming |
|------|----|-------------|------------|-----------------|----------------------|--------|----------|-----------|----------|--------------|
|      |    |             |            |                 |                      |        |          |           |          |              |

**emsNewActiveEffect**

| Name | TP | Probability | damageBase | PrimaryFunction | PrimaryFunctionValue | Target | CostType | CostValue | Category | EffectTiming |
|------|----|-------------|------------|-----------------|----------------------|--------|----------|-----------|----------|--------------|
|      |    |             |            |                 |                      |        |          |           |          |              |

**emsAllActiveEffects**

| Name | TP | Probability | damageBase | PrimaryFunction | PrimaryFunctionValue | Target | CostType | CostValue | Category | EffectTiming |
|------|----|-------------|------------|-----------------|----------------------|--------|----------|-----------|----------|--------------|
|      |    |             |            |                 |                      |        |          |           |          |              |

\pagebreak

