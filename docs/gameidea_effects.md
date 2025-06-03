# Effect Usage

**Turn Process for playerTurn**

1. loop effects for turn start
- loop on entity: `self`
  - usageTiming: `etiOnTurnStart`
  - stats loop
2. `use skill` command
- loop on entity: `enemy`
  - usageTiming: `etiOnSkillStart`
  - block (`ecaBlockDamage`, `ecaStopSkill`, `ecaBlockDebuffs`, )
  - change target (`ecaChangeTarget`)
- loop on entity: `self`
  - usageTiming: `etiOnSkillCalculation`
  - block (`ecaBlockBuffs`)
  - decrease/increase power (`ecaIncreasePower`, `ecaDecreasePower`)
  - decrease/increase damage (`ecaIncreaseOutgoingDamage`, , `ecaDecreaseOutgoingDamage`)
- deal damage (entity function)
  - loop on entity: `enemy`
    - useageTiming: `etiOnIncmoingDamage`, `etiOnActualDamage`
      - decrease/increase damage (`ecaDecreaseIncomingDamage`, `ecaIncreaseIncomingDamage`)
- apply effects
  - loop on entity: `enemy`
    - useageTiming: `etiOnIncomingEffect`
3. loop effects for turn end
- loop on entity: `self`
  - usageTiming: `etiOnTurnEnd`, `etiOnEffectRemoval`
  - stats loop

- stats loop
  - block (`ecaBlockHealing`)
  - decrease/increase healing (`ecaIncreaseHealing`, `ecaDecreaseHealing`)
  - decrease/increase health (`ecaHeal`, `ecaDoDamage`, `ecaTakeDamage`)
  - decrease/increase strength (`ecaIncreaseStrength`)
  - remove/add effects (`ecaRemoveEffect`, `ecaChangeEffectTurn`)

# Effect necessities

```go
type SkillEffect struct {}

type EffectCategory int
const (
  ecaHeal effectCategory = iota
  ecaDoDamage
  ecaIncreasePower
  ecaIncreaseStrength
  ecaIncreaseOutgoingDamage
  ecaDecreaseIncomingDamage
  ecaIncreaseHealing
  ecaAddEffect
  ecaBlockDebuffs
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


type EffectType int
const (
  etyBuff EffectType = iota
  etyDebuff
)
```

# Calculations

- rawStrength => strength from stats
- modifiedStrength => currentStrength +/- buff/debuff effects - source Entity

- rawSkillDamage => modifiedStrength * skillDmgmulti (1 damage per 1 rawSkillDamage)
- modifiedSkillDamage => rawSkillDamage +/- buff/debuff - source Entity

- rawEffectValue => based on effect (e.g. increase by 10% = modifiedStrength x 0,1 / 1 damage/heal per 1 value)
- modifiedEffectDamage => rawEffectValue +/- buff/debuff - source Entity (e.g. ???)
- modifiedEffectHeal => rawEffectValue +/- buff/debuff - source Entity (e.g. increaseHealing)

- actualSkillDamageTaken => modifiedSkillDamage
- actualEffectDamageTaken => modifiedEffectDamage
- actualEffectHealTaken => modifiedEffectHeal

```
e.g. Gain 50% of your Damage in Health + Healing effects are 10% stronger

actualSkillDamageTaken = 50
rawEffectPower = 50% of actualSkillDamageTaken (50) => 25
modifiedEffectPower = 10% of rawEffectPower (25) => 2.5 = 3 + rawEffectPower (25) => 28 (25*1,1)
```

# Game Effects Table

**affected Value**

- currentHealth
- modifiedEffectDamage
- actualSkillDamageTaken
- actualEffectDamageTaken

- buff
  - currentHealth
  - rawStrength
  - rawSkillDamage
  - rawEffectValue
  - rawEffectValue
  - randomBuffTurns
  - allBuffsTurns
  - randomDebuffTurns
  - allDebuffsTurns
- debuff
  - currentHealth
  - rawStrength
  - rawSkillDamage
  - rawEffectValue
  - rawEffectValue
  - randomBuffTurns
  - allBuffsTurns
  - randomDebuffTurns
  - allDebuffsTurns

**Trigger**

- onDealDamage
- onDamageTaken
- onTurnStart
- onTurnEnd
- onEffectRemoved


**Condition**



**effectValue**

- x% of modifiedSkillDamage
- x% of actualSkillDamageTaken/actualEffectDamageTaken
- x per modifiedStrength



## Buffs (Self)

|         Name         |        Trigger        |          Category         |     outputBase    |                          Description                           |
|----------------------|-----------------------|---------------------------|-------------------|----------------------------------------------------------------|
| heal1                | etiOnActualDamage     | ecaHeal                   | OutgoingDamage    | Gain 50% of your Damage in Health                              |
| heal2                | etiOnActualDamage     | ecaHeal                   | ActualDamageTaken | Heal 50% of the Damage you take                                |
| heal3                | etiOnTurnStart        | ecaHeal                   | effectMulti       | Restores health at the start of each turn                      |
| heal4                | etiOnTurnEnd          | ecaHeal                   | effectMulti       | Restores health at the end of each turn                        |
| heal5                | etiOnEffectRemoval    | ecaHeal                   | effectMulti       | gain a massive heal when this effect expires or is removed     |
| doDamage1            | etiOnActualDamage     | ecaDoDamage               | ActualDamageTaken | deal 50% of the damage you Receive                             |
| increaseStrength1    | etiOnTurnStart        | ecaIncreaseStrength       | effectMulti       | Increase strength                                              |
| increaseStrength2    | etiOnTurnStart        | ecaIncreaseStrength       | effectMulti       | strength increases by 10% for each 10% of your missing health  |
| increaseStrength3    | etiOnTurnStart        | ecaIncreaseStrength       | effectMulti       | strength increases by 10% for each active buff you have        |
| increaseStrength4    | etiOnTurnStart        | ecaIncreaseStrength       | effectMulti       | strength increases by 20% when at full health                  |
| increaseStrength5    | etiOnTurnStart        | ecaIncreaseStrength       | effectMulti       | strength increases by 10% for each defeated boss               |
| increasePower1       | etiOnSkillCalculation | ecaIncreasePower          | effectMulti       | increase power                                                 |
| increaseDamageDone1  | etiOnSkillCalculation | ecaIncreaseOutgoingDamage | effectMulti       | Adds bonus damage if enemy is low                              |
| increaseDamageDone2  | etiOnSkillCalculation | ecaIncreaseOutgoingDamage | effectMulti       | 50% Chance to double the damage                                |
| increaseDamageDone3  | etiOnSkillCalculation | ecaIncreaseOutgoingDamage | effectMulti       | Deal double the damage if the same skill is used repeatedly    |
| decreaseDamageTaken1 | etiOnSkillCalculation | ecaDecreaseIncomingDamage | effectMulti       | Receive 50% less damage                                        |
| decreaseDamageTaken2 | etiOnSkillCalculation | ecaDecreaseIncomingDamage | effectMulti       | Receive 10% less Damage from repeated sources                  |
| decreaseDamageTaken3 | etiOnSkillCalculation | ecaDecreaseIncomingDamage |                   |                                                                |
| increaseHealing1     | etiOnSkillCalculation | ecaIncreaseHealing        | effectMulti       | Healing effects are 10% stronger for each debuff you have      |
| increaseHealing2     | etiOnSkillCalculation | ecaIncreaseHealing        | effectMulti       | Healing effects are 100% stronger when below 30% health        |
| increaseHealing3     | etiOnSkillCalculation | ecaIncreaseHealing        | effectMulti       | Healing effects are 10% stronger for each active buff          |
| blockDebuffs1        | etiOnSkillStart       | ecaBlockDebuffs           | -                 | Prevents new debuffs from being applied                        |
| blockDebuffs2        | etiOnSkillStart       | ecaBlockDebuffs           | -                 | 50% Chance to block an incoming Debuff                         |
| blockDebuffs3        | etiOnSkillStart       | ecaBlockDebuffs           | -                 | Block the next 3 debuffs                                       |
| blockDebuffs4        | etiOnSkillStart       | ecaBlockDebuffs           | -                 | Block debuffs that reduce power/damage                         |
| blockDebuffs5        | etiOnSkillStart       | ecaBlockDebuffs           | -                 | Block debuffs that cause damage                                |
| blockDebuffs6        | etiOnSkillStart       | ecaBlockDebuffs           | -                 | Block debuffs that prevent healing                             |
| blockDebuffs7        | etiOnSkillStart       | ecaBlockDebuffs           | -                 | Block debuffs that stop skills/change targets                  |
| blockDebuffs7        | etiOnSkillStart       | ecaBlockDebuffs           | -                 | Block buff prevention and effect manipulation                  |
| blockDamage1         | etiOnSkillStart       | ecaBlockDamage            | effectMulti       | 50% Chance to dont get damage                                  |
| blockDamage2         | etiOnSkillStart       | ecaBlockDamage            | effectMulti       | block damage if below 5% health                                |
| blockDamage3         | etiOnSkillStart       | ecaBlockDamage            | effectMulti       | block damage if the enemy used the same skill as a turn before |
| changeTarget1        | etiOnSkillStart       | ecaChangeTarget           | IncomingDamage    | 50% Chance to reflect the damage                               |
| changeTarget2        | etiOnSkillStart       | ecaChangeTarget           | -                 | 50% Chance to reflect the effects                              |
| changeTarget3        | etiOnSkillStart       | ecaChangeTarget           | IncomingDamage    | reflect the damage if below 5% health                          |
| removeEffect1        | etiOnTurnStart        | ecaRemoveEffect           | -                 | remove a random debuff                                         |
| removeEffect2        | etiOnSkillStart       | ecaRemoveEffect           | -                 | remove a random buff of the enemy when attacked                |
| removeEffect3        | etiOnSkillStart       | ecaRemoveEffect           | -                 | remove a random debuff when attacked                           |
| removeEffect4        | etiOnTurnEnd          | ecaRemoveEffect           | -                 | remove a random debuff that reduce power/damage                |
| removeEffect5        | etiOnTurnEnd          | ecaRemoveEffect           | -                 | remove a random debuff that cause damage                       |
| removeEffect6        | etiOnTurnEnd          | ecaRemoveEffect           | -                 | remove a random debuff that that prevent healing               |
| removeEffect7        | etiOnTurnEnd          | ecaRemoveEffect           | -                 | remove a random debuff that that stop skills/change targets    |
| removeEffect8        | etiOnTurnStart        | ecaRemoveEffect           | -                 | remove the oldest debuff                                       |
| removeEffect9        | etiOnTurnStart        | ecaRemoveEffect           | -                 | remove the newest debuff                                       |
| removeEffect10       | etiOnTurnStart        | ecaRemoveEffect           | -                 | remove the debuff with the most remaining turns                |
| removeEffect11       | etiOnTurnStart        | ecaRemoveEffect           | -                 | remove the debuff with the least remaining turns               |
| addEffect1           | etiOnTurnStart        | ecaAddEffect              | -                 | add a random buff                                              |
| addEffect2           | etiOnTurnEnd          | ecaAddEffect              | -                 | add a random buff                                              |


## Debuffs (Enemy)

|         Name         |        Trigger        |          Category         |   outputBase   |                              Description                              |
|----------------------|-----------------------|---------------------------|----------------|-----------------------------------------------------------------------|
| DecreasePower1       | etiOnSkillCalculation | ecaDecreasePower          | effectMulti    | Reduce targets Power                                                  |
| DecreasePower2       | etiOnSkillCalculation | ecaDecreasePower          |                |                                                                       |
| DecreasePower3       | etiOnSkillCalculation | ecaDecreasePower          |                |                                                                       |
| DecreasePower4       | etiOnSkillCalculation | ecaDecreasePower          |                |                                                                       |
| DecreasePower5       | etiOnSkillCalculation | ecaDecreasePower          |                |                                                                       |
| BlockHealing1        | etiOnSkillStart       | ecaBlockHealing           | effectMulti    | 50% Chance to block heals                                             |
| BlockHealing2        | etiOnSkillStart       | ecaBlockHealing           |                |                                                                       |
| BlockHealing3        | etiOnSkillStart       | ecaBlockHealing           |                |                                                                       |
| TakeDamage1          | etiOnTurnStart        | ecaTakeDamage             | -              | Immediately kill the enemy while below 10% health                     |
| TakeDamage2          | etiOnTurnStart        | ecaTakeDamage             | effectMulti    | Applies a damaging effect that deals damage at the start of each turn |
| TakeDamage3          | etiOnEffectRemoval    | ecaTakeDamage             | effectMulti    | When cleansed, explodes                                               |
| TakeDamage4          | etiOnTurnStart        | ecaTakeDamage             | effectMulti    | Deals damage based on maximum health                                  |
| TakeDamage5          | etiOnTurnStart        | ecaTakeDamage             |                |                                                                       |
| TakeDamage6          | etiOnTurnStart        | ecaTakeDamage             |                |                                                                       |
| DecreaseDamageDone1  | etiOnSkillCalculation | ecaDecreaseOutgoingDamage | effectMulti    | Reduces target's damage output by 50%                                 |
| DecreaseDamageDone2  | etiOnSkillCalculation | ecaDecreaseOutgoingDamage |                |                                                                       |
| DecreaseDamageDone3  | etiOnSkillCalculation | ecaDecreaseOutgoingDamage |                |                                                                       |
| DecreaseDamageDone4  | etiOnSkillCalculation | ecaDecreaseOutgoingDamage |                |                                                                       |
| DecreaseDamageDone5  | etiOnSkillCalculation | ecaDecreaseOutgoingDamage |                |                                                                       |
| DecreaseDamageDone6  | etiOnSkillCalculation | ecaDecreaseOutgoingDamage |                |                                                                       |
| IncreaseDamageTaken1 | etiOnSkillCalculation | ecaIncreaseIncomingDamage | effectMulti    | Receive 50% more damage                                               |
| IncreaseDamageTaken2 | etiOnSkillCalculation | ecaIncreaseIncomingDamage |                |                                                                       |
| IncreaseDamageTaken3 | etiOnSkillCalculation | ecaIncreaseIncomingDamage |                |                                                                       |
| IncreaseDamageTaken4 | etiOnSkillCalculation | ecaIncreaseIncomingDamage |                |                                                                       |
| IncreaseDamageTaken5 | etiOnSkillCalculation | ecaIncreaseIncomingDamage |                |                                                                       |
| IncreaseDamageTaken6 | etiOnSkillCalculation | ecaIncreaseIncomingDamage |                |                                                                       |
| DecreaseHealing1     | etiOnSkillCalculation | ecaDecreaseHealing        | effectMulti    | Reduces all healing received by 50%                                   |
| DecreaseHealing2     | etiOnSkillCalculation | ecaDecreaseHealing        |                |                                                                       |
| DecreaseHealing3     | etiOnSkillCalculation | ecaDecreaseHealing        |                |                                                                       |
| BlockBuffs1          | etiOnSkillStart       | ecaBlockBuffs             | -              | Prevents the target from receiving buffs                              |
| BlockBuffs2          | etiOnSkillStart       | ecaBlockBuffs             | effectMulti    | Prevents the target from receiving healing effects                    |
| BlockBuffs3          | etiOnSkillStart       | ecaBlockBuffs             | effectMulti    | Prevents the target from receiving ecaIncreaseOutgoingDamage effects  |
| ChangeTarget1        | etiOnSkillStart       | ecaChangeTarget           | IncomingDamage | 50% Chance to attack itself                                           |
| RemoveEffect1        | etiOnTurnStart        | ecaRemoveEffect           | -              | remove a random Buff                                                  |
| RemoveEffect2        | etiOnSkillStart       | ecaRemoveEffect           | -              | remove a random Buff when attacked                                    |
| RemoveEffect3        | etiOnSkillStart       | ecaRemoveEffect           | -              | remove a random debuff of the enemy when attacked                     |
| addEffect1           | etiOnTurnStart        | ecaAddEffect              | -              | add a random debuff                                                   |
| addEffect2           | etiOnTurnEnd          | ecaAddEffect              | -              | add a random debuff                                                   |


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
