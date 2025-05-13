# Effect Usage

**Turn Process for use skill**

1. loop effects for turn start
- entity: `both`
- usageTiming: `etiOnTurnStart`
2. `use skill` command
3. loop if an effect is blocked / target is changed
- entity: `both`
- usageTiming: `etiOnSkillStart`
5. determine `basicSkillPower` (`currentPower * skillMulti`)
6. loop effects to create `fullSkillPower` (`currentPower * skillMulti * each effectMulti`)
- usageTiming: `etiOnSkillCalculation`
7. use fullSkillPower to do the damage
8. loop effects for turn end and reduction of turns of effects and reactions - trigger: `onTurnEnd`
- usageTiming: `etiOnTurnEnd`

**Turn Process for use talisman**

1. loop effects for turn start
- entity: `both`
- usageTiming: `etiOnTurnStart`
2. `use talisman` command
...


# Effect necessities

```go
type SkillEffect struct {
  name                 string
  description          string
  talentpointCosts     int
  probability          float32
  type                 EffectType
  category             EffectCategory
  execute              func()
  checkCondition       func() bool
  usageTiming          EffectTiming
  multi                float32
}

type EffectCategory int
const (
  ecaHeal EffectCategory = iota
  ecaDoDamage
  ecaIncreasePower
  ecaIncreaseDamageDone
  ecaDecreaseDamageTaken
  ecaIncreaseHealing
  ecaaBlockDebuffs
  ecaBlockDamage
  ecaStopSkill
  ecaChangeTarget
  ecaTakeDamage
  ecaDecreasePower
  ecaDecreaseDamageDone
  ecaIncreaseDamageTaken
  ecaDecreaseHealing
  ecaBlockBuffs
  ecaBlockHealing
)

type EffectTiming int
const (
  etiOnTurnStart EffectTiming = iota
  etiOnSkillStart
  etiOnSkillCalculation
  etiOnTurnEnd
)

type EffectType int
const (
  etyBuff EffectType = iota
  etyDebuff
)

type BattleState struct {
  currentHealth        int
  currentPower         int
  totalBuffTurnsCount  int
  totalBuffCount       int
  totalDebuffTurnCount int
  totalDebuffCount     int
  activeEffectsList    []ActiveEffect
  currentBattlePhase   BattlePhase
}

type ActiveEffect struct {
  skillEffect SkillEffect
  totalPower  int
  turnsLeft   int
  source      Entity
  target      Entity
}
```

# Game Effects Table

## Buffs (Self)

|         Name         |      Trigger       |        Category       |                                   Description                                   |
|----------------------|--------------------|-----------------------|---------------------------------------------------------------------------------|
| Lifeleech            | onTurnEnd          | heal                  | Gain 50% of your Damage in Health                                               |
| Last Stand           | onTurnEnd          | heal                  | Heal 50% of the Damage you take                                                 |
| Heal Over Time (HOT) | onTurnEnd          | heal                  | Restores health at the start of each turn                                       |
| Reflect Damage       | onTurnEnd          | do damage             | Reflect 50% of the damage you Receive                                           |
| Increase Power       | onSkillCalculation | increase power        | Increases damage output by 50%                                                  |
| Adrenaline Rush      | onSkillCalculation | increase power        | Power increases by 10% each turn, up to 50%                                     |
| Vengeance            | onSkillCalculation | increase power        | Power increases by 10% for each 10% of your missing health                      |
| Combo Mastery        | onSkillCalculation | increase power        | Each consecutive use of the same skill increases power by 15%                   |
| Finisher             | onSkillCalculation | increase damage done  | Adds bonus damage if enemy is low                                               |
| Critical Strike      | onSkillCalculation | increase damage done  | 50% Chance to double the damage                                                 |
| Fury                 | onSkillCalculation | increase damage done  | Each attack increases damage of next attack                                     |
| Immunity             | onSkillCalculation | decrease damage taken | Receive 50% less damage                                                         |
| Adaptation           | onSkillCalculation | decrease damage taken | Receive 10% less Damage from repeated sources                                   |
| Damage Dispersal     | onSkillCalculation | decrease damage taken | Convert 20% of damage taken into a DoT on yourself that deals less total damage |
| Regenerative Core    | onSkillCalculation | increase healing      | Healing effects become 30% stronger for each debuff you have                    |
| Crisis Recovery      | onSkillCalculation | increase healing      | Healing effects are 100% stronger when below 30% health                         |
| Block Debuffs        | onSkillStart       | block debuffs         | Prevents new debuffs from being applied while active                            |
| Resistance           | onSkillStart       | block debuffs         | 50% Chance to block an incoming Debuff                                          |
| Evasion              | onSkillStart       | block damage          | 50% Chance to dont get damage                                                   |
| Lucky Timing         | onSkillStart       | stop skill            | 50% Chance to block an incoming skill                                           |
| Mirror               | onSkillStart       | change target         | 50% Chance the attack is mirrored                                               |

## Debuffs (Enemy)

|         Name        |      Trigger       |        Category       |                              Description                              |
|---------------------|--------------------|-----------------------|-----------------------------------------------------------------------|
| Execution           | onTurnStart        | take damage           | Immediately kill the enemy while below 10% health                     |
| Bleeding            | onTurnStart        | take damage           | Applies a damaging effect that deals damage at the start of each turn |
| Unstable Affliction | onTurnStart        | take damage           | When cleansed, explodes                                               |
| Soul Burn           | onTurnStart        | take damage           | Deals damage based on maximum health                                  |
| Weaken              | onSkillCalculation | decrease damage done  | Reduces target's damage output by 50%                                 |
| Vulnerability       | onSkillCalculation | increase damage taken | Receive 50% more damage                                               |
| Reduce Healing      | onSkillCalculation | decrease healing      | Reduces all healing received by 50%                                   |
| Block Buffs         | onSkillStart       | block buffs           | Prevents the target from receiving buffs and healing effects          |
| Confusion           | onSkillStart       | stop skill            | 50% Chance to miss the skill                                          |
| Mental Block        | onSkillStart       | stop skill            | Cannot use the same skill twice in a row                              |
| Distraction         | onSkillStart       | change target         | 50% Chance to attack itself                                           |

## Talismans (One-time Use Items)

|           Name           |       Category      |                          Description                          |
|--------------------------|---------------------|---------------------------------------------------------------|
| Buff Turn Bonus Damage   | do damage           | Bonus Damage to amount of all remaining Buff Turns            |
| Debuff Turn Bonus Damage | do damage           | Bonus Damage to Amount of all remaining Debuff Turns of Enemy |
| Heal                     | heal                | Immediately restores health based on power stat               |
| Buff Heal                | heal                | Stronger direct heal but remove one random buff               |
| Health Cleanse           | heal                | Pay 10% of your current health to remove a random Debuff      |
| Cleanse                  | remove effects      | Removes all negative effects from the target                  |
| Dispel                   | remove effects      | Removes all positive effects from the target                  |
| Extend Buffs             | change effect turns | +1 Turn for each 10 Total Power                               |
| Extend Debuffs           | change effect turns | +1 Turn for each 10 Total Power                               |
| Reduce Buffs             | change effect turns | -1 Turn for each 10 Total Power                               |
| Reduce Debuffs           | change effect turns | -1 Turn for each 10 Total Power                               |

## Additional Effect Ideas

|        Name        | Type |                                                        Description                                                        |
|--------------------|------|---------------------------------------------------------------------------------------------------------------------------|
| Random Effect      | misc | Get a random buff/debuff when attacking/attacked                                                                          |
| Effect Removal     | misc | Remove a random buff/debuff when attacking/attacked                                                                       |
| Status Swap        | misc | Swap all buffs/debuffs                                                                                                    |
| Power Theft        | misc | Steal enemy power                                                                                                         |
| Health Sacrifice   | misc | Sacrifice health to deal damage/increase damage                                                                           |
| Time Bomb          | misc | Effect only does damage when it naturally runs out of turns (no removal) but increase the damage for every turn it passed |
| Health Siphon      | misc | Each turn, steal a small amount of health from the enemy                                                                  |
| influence Talisman | misc | block talisman..                                                                                                          |
