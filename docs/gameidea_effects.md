# Effect Usage

**Turn Process for use skill**

1. loop effects for turn start
- entity: `both`
- usageTiming: `etiOnTurnStart`
2. `use skill` command
3. loop if an effect is blocked / target is changed
- entity: `both`
- usageTiming: `etiOnSkillStart`
4. loop effects to create `rawSkillPower` (`currentPower * skillMulti * each effectMulti`)
- usageTiming: `etiOnSkillCalculation`
5. use rawSkillPower to do the damage
6. loop effects for turn end and reduction of turns of effects and reactions - trigger: `onTurnEnd`
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
  internalName     string
  displayName      string
  description      string
  talentpointCosts int
  probability      float32
  effectType       effectType
  category         effectCategory
  execute          func()
  checkCondition   func() bool
  usageTiming      effectTiming
  multi            float32
}


type EffectCategory int
const (
  ecaHeal effectCategory = iota
  ecaDoDamage
  ecaIncreasePower
  ecaIncreaseOutgoingDamage
  ecaDecreaseIncomingDamage
  ecaIncreaseHealing
  ecaaBlockDebuffs
  ecaBlockDamage
  ecaStopSkill
  ecaChangeTarget
  ecaTakeDamage
  ecaDecreasePower
  ecaDecreaseOutgoingDamage
  ecaIncreaseIncomingDamage
  ecaDecreaseHealing
  ecaBlockBuffs
  ecaBlockHealing
  ecaRemoveEffect
  ecaChangeEffectTurn
)

type EffectTiming int
const (
  etiOnTurnStart effectTiming = iota
  etiOnSkillStart
  etiOnSkillCalculation
  etiOnTurnEnd
  etiOnEffectRemoval
  etiOnSkillEnd
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
  rawSkillPower  int
  turnsLeft   int
  source      Entity
  target      Entity
}
```

# Calculations

- basePower => currentPower from battleState
- modifiedPower => basePower +/- buffs/debuffs source Entity
- rawSkillPower => modifiedPower * skillDmgmulti
- calculatedDamage => 1 damage per 1 rawSkillPower
- outgoingDamage => calculatedDamage +/- buffs/debuffs source Entity
- incomingDamage => outgoingDamage
- actualDamageTaken => incomingDamage +/- buffs/debuffs target Entity

example rawSkillPower:

```
// new player name difficulty health power speed
// new skill <name> <dmgmulti> <duration> [effect effect effect...]

new player John normal 125 15 8
new skill Fireball 1.5 4

1.5    * 15      = 23
23     * 1,46    = 34
```

|          Category         |        Trigger        |      Description      |             Calc            | Value |
|---------------------------|-----------------------|-----------------------|-----------------------------|-------|
| SkillDamage               | -                     | 1 damage per 1 power  | rawSkillPower               | 23    |
| ecaHeal                   | etiOnTurnEnd          | 2 Health per 3 power  | rawSkillPower / 3 * 2       | 16    |
| ecaDoDamage               | etiOnSkillCalculation | -                     | -                           | -     |
| ecaIncreasePower          | etiOnSkillCalculation | 0.1 multi per 5 power | rawSkillPower / 5 * 0,1 + 1 | 1,46  |
| ecaIncreaseOutgoingDamage | etiOnSkillCalculation | 0.1 multi per 5 power | rawSkillPower / 5 * 0,1 + 1 | 1,46  |
| ecaDecreaseIncomingDamage | etiOnSkillCalculation | 0.1 multi per 5 power | rawSkillPower / 5 * 0,1 + 1 | 1,46  |
| ecaIncreaseHealing        | etiOnSkillCalculation | 0.1 multi per 3 power | rawSkillPower / 3 * 0,1 + 1 | 1,76  |
| ecaaBlockDebuffs          | etiOnSkillStart       | -                     | -                           | -     |
| ecaBlockDamage            | etiOnSkillStart       | -                     | -                           | -     |
| ecaStopSkill              | etiOnSkillStart       | -                     | -                           | -     |
| ecaChangeTarget           | etiOnSkillStart       | -                     | -                           | -     |
| ecaRemoveEffect           | etiOnSkillStart       | 1 effect per 10 power | rawSkillPower / 10          | 2     |
| ecaDecreasePower          | etiOnSkillCalculation | 0.1 multi per 5 power | rawSkillPower / 5 * 0,1 + 1 | 1,46  |
| ecaBlockHealing           | etiOnSkillStart       | -                     | -                           | -     |
| ecaTakeDamage             | etiOnTurnStart        | 2 Damage per 3 power  | rawSkillPower / 3 * 2       | 16    |
| ecaDecreaseOutgoingDamage | etiOnSkillCalculation | 0.1 multi per 5 power | rawSkillPower / 5 * 0,1 + 1 | 1,46  |
| ecaIncreaseIncomingDamage | etiOnSkillCalculation | 0.1 multi per 5 power | rawSkillPower / 5 * 0,1 + 1 | 1,46  |
| ecaDecreaseHealing        | etiOnSkillCalculation | 0.1 multi per 5 power | rawSkillPower / 5 * 0,1 + 1 | 1,46  |
| ecaBlockBuffs             | etiOnSkillStart       | -                     | -                           | -     |
| ecaChangeEffectTurn       | etiOnSkillStart       | 1 turn per 10 power   | rawSkillPower / 10          | 2     |

|        Trigger        |      Description      |
|-----------------------|-----------------------|
| etiOnSkillCalculation | 0.1 multi per 5 power |
| etiOnTurnEnd          | 2 y per 3 power       |
| etiOnTurnStart        | 2 y per 3 power       |
| etiOnSkillStart       | 1 y per 10 power      |

unless otherwise defined


# Game Effects Table


## Buffs (Self)

|         Name         |        Trigger        |          Category         |     outputBase    |                         Description                          |
|----------------------|-----------------------|---------------------------|-------------------|--------------------------------------------------------------|
| heal1                | etiOnSkillEnd         | ecaHeal                   | OutgoingDamage    | Gain 50% of your Damage in Health                            |
| heal2                | etiOnSkillEnd         | ecaHeal                   | ActualDamageTaken | Heal 50% of the Damage you take                              |
| heal3                | etiOnTurnStart        | ecaHeal                   | effectMulti       | Restores health at the start of each turn                    |
| heal4                | etiOnTurnEnd          | ecaHeal                   | effectMulti       | Restores health at the end of each turn                      |
| heal4                | etiOnEffectRemoval    | ecaHeal                   | effectMulti       | gain a massive heal when this effect expires or is removed   |
| doDamage1            | etiOnTurnEnd          | ecaDoDamage               | ActualDamageTaken | Reflect 50% of the damage you Receive                        |
| increasePower1       | etiOnSkillCalculation | ecaIncreasePower          | effectMulti       | Increase power                                               |
| increasePower3       | etiOnSkillCalculation | ecaIncreasePower          | effectMulti       | Power increases by 10% for each 10% of your missing health   |
| increaseDamageDone1  | etiOnSkillCalculation | ecaIncreaseOutgoingDamage | effectMulti       | Adds bonus damage if enemy is low                            |
| increaseDamageDone2  | etiOnSkillCalculation | ecaIncreaseOutgoingDamage | effectMulti       | 50% Chance to double the damage                              |
| increaseDamageDone3  | etiOnSkillCalculation | ecaIncreaseOutgoingDamage | effectMulti       | Deal double the damage if the same skill is used repeatedly  |
| decreaseDamageTaken1 | etiOnSkillCalculation | ecaDecreaseIncomingDamage | effectMulti       | Receive 50% less damage                                      |
| decreaseDamageTaken2 | etiOnSkillCalculation | ecaDecreaseIncomingDamage | effectMulti       | Receive 10% less Damage from repeated sources                |
| increaseHealing1     | etiOnSkillCalculation | ecaIncreaseHealing        | effectMulti       | Healing effects become 30% stronger for each debuff you have |
| increaseHealing2     | etiOnSkillCalculation | ecaIncreaseHealing        | effectMulti       | Healing effects are 100% stronger when below 30% health      |
| blockDebuffs1        | etiOnSkillStart       | ecaBlockDebuffs           | -                 | Prevents new debuffs from being applied while active         |
| blockDebuffs2        | etiOnSkillStart       | ecaBlockDebuffs           | -                 | 50% Chance to block an incoming Debuff                       |
| blockDebuffs3        | etiOnSkillStart       | ecaBlockDebuffs           | -                 | Block debuffs from the ecaHeal Category                      |
| blockDamage1         | etiOnSkillStart       | ecaBlockDamage            | effectMulti       | 50% Chance to dont get damage                                |
| stopSkill1           | etiOnSkillStart       | ecaStopSkill              | -                 | 50% Chance to block an incoming skill                        |
| changeTarget1        | etiOnSkillStart       | ecaChangeTarget           | IncomingDamage    | 50% Chance the attack is mirrored                            |
| removeEffect1        | etiOnSkillStart       | ecaRemoveEffect           | -                 | remove a random debuff                                       |
| removeEffect2        | etiOnSkillStart       | ecaRemoveEffect           | -                 | remove a random buff of the enemy when attacked              |
| removeEffect3        | etiOnSkillStart       | ecaRemoveEffect           | -                 | remove a random debuff when attacked                         |

## Debuffs (Enemy)

|         Name         |        Trigger        |          Category         |   outputBase   |                              Description                              |
|----------------------|-----------------------|---------------------------|----------------|-----------------------------------------------------------------------|
| DecreasePower1       | etiOnSkillCalculation | ecaDecreasePower          | effectMulti    | Reduce targets Power                                                  |
| BlockHealing1        | etiOnSkillStart       | ecaBlockHealing           | effectMulti    | 50% Chance to block heals                                             |
| TakeDamage1          | etiOnTurnStart        | ecaTakeDamage             | -              | Immediately kill the enemy while below 10% health                     |
| TakeDamage2          | etiOnTurnStart        | ecaTakeDamage             | effectMulti    | Applies a damaging effect that deals damage at the start of each turn |
| TakeDamage3          | etiOnEffectRemoval    | ecaTakeDamage             | effectMulti    | When cleansed, explodes                                               |
| TakeDamage4          | etiOnTurnStart        | ecaTakeDamage             | effectMulti    | Deals damage based on maximum health                                  |
| DecreaseDamageDone1  | etiOnSkillCalculation | ecaDecreaseOutgoingDamage | effectMulti    | Reduces target's damage output by 50%                                 |
| IncreaseDamageTaken1 | etiOnSkillCalculation | ecaIncreaseIncomingDamage | effectMulti    | Receive 50% more damage                                               |
| DecreaseHealing1     | etiOnSkillCalculation | ecaDecreaseHealing        | effectMulti    | Reduces all healing received by 50%                                   |
| BlockBuffs1          | etiOnSkillStart       | ecaBlockBuffs             | -              | Prevents the target from receiving buffs                              |
| BlockBuffs2          | etiOnSkillStart       | ecaBlockBuffs             | effectMulti    | Prevents the target from receiving healing effects                    |
| BlockBuffs3          | etiOnSkillStart       | ecaBlockBuffs             | effectMulti    | Prevents the target from receiving ecaIncreaseOutgoingDamage effects  |
| StopSkill1           | etiOnSkillStart       | ecaStopSkill              | -              | 50% Chance to miss the skill                                          |
| StopSkill2           | etiOnSkillStart       | ecaStopSkill              | -              | Cannot use the same skill twice in a row                              |
| ChangeTarget1        | etiOnSkillStart       | ecaChangeTarget           | IncomingDamage | 50% Chance to attack itself                                           |
| RemoveEffect1        | etiOnTurnStart        | ecaRemoveEffect           | -              | remove a random Buff                                                  |
| RemoveEffect2        | etiOnSkillStart       | ecaRemoveEffect           | -              | remove a random Buff when attacked                                    |
| RemoveEffect3        | etiOnSkillStart       | ecaRemoveEffect           | -              | remove a random debuff of the enemy when attacked                     |

## Talismans (One-time Use Items)

|        Name       |     Category     |                          Description                          |
|-------------------|------------------|---------------------------------------------------------------|
| DoDamage2         | DoDamage         | Bonus Damage to amount of all remaining Buff Turns            |
| DoDamage3         | DoDamage         | Bonus Damage to Amount of all remaining Debuff Turns of Enemy |
| Heal4             | Heal             | Immediately restores health based on power stat               |
| Heal5             | Heal             | Stronger direct heal but remove one random buff               |
| Heal6             | Heal             | Pay 10% of your current health to remove a random Debuff      |
| RemoveEffect4     | RemoveEffect     | Removes all negative effects from the target                  |
| RemoveEffect5     | RemoveEffect     | Removes all positive effects from the target                  |
| ChangeEffectTurn1 | ChangeEffectTurn | +1 Turn for each 10 Total Power                               |
| ChangeEffectTurn2 | ChangeEffectTurn | +1 Turn for each 10 Total Power                               |
| ChangeEffectTurn3 | ChangeEffectTurn | -1 Turn for each 10 Total Power                               |
| ChangeEffectTurn4 | ChangeEffectTurn | -1 Turn for each 10 Total Power                               |

## Additional Effect Ideas

|        Name        | Type |                                                        Description                                                        |
|--------------------|------|---------------------------------------------------------------------------------------------------------------------------|
| Random Effect      | misc | Get a random buff/debuff when attacking/attacked from a list of random effects                                            |
| Effect Removal     | misc | Remove a random buff/debuff when attacking/attacked                                                                       |
| Status Swap        | misc | Swap all buffs/debuffs                                                                                                    |
| Power Theft        | misc | Steal enemy power                                                                                                         |
| Health Sacrifice   | misc | Sacrifice health to deal damage/increase damage                                                                           |
| Time Bomb          | misc | Effect only does damage when it naturally runs out of turns (no removal) but increase the damage for every turn it passed |
| Health Siphon      | misc | Each turn, steal a small amount of health from the enemy                                                                  |
| influence Talisman | misc | block talisman..                                                                                                          |
