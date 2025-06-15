# Effect Usage

**Turn Process for playerTurn**

1. loop effects for turn start
- loop on entity: `self`
  - usageTiming: `etiOnTurnStart`
  - stats loop
2. `use skill` command
- loop on entity: `enemy`
  - usageTiming: `etiOnSkillStart`
  - block
  - change target
- loop on entity: `self`
  - usageTiming: `etiOnSkillCalculation`
  - block
  - decrease/increase Strength 
  - decrease/increase damage
- deal damage (entity function)
  - loop on entity: `enemy`
    - useageTiming: `etiOnIncmoingDamage`, `etiOnActualDamage`
      - decrease/increase damage
- apply effects
  - loop on entity: `enemy`
    - useageTiming: `etiOnIncomingEffect`
3. loop effects for turn end
- loop on entity: `self`
  - usageTiming: `etiOnTurnEnd`, `etiOnEffectRemoval`
  - stats loop

- stats loop
  - block
  - decrease/increase healing
  - decrease/increase health
  - decrease/increase strength
  - remove/add effects



# Calculations

- rawStrength => strength from stats
- modifiedStrength => currentStrength(rawStrength) +/- buff/debuff effects - source Entity

- rawSkillDamage => modifiedStrength * skillDmgmulti (1 damage per 1 rawSkillDamage)
- rawEffectDamage => based on effect or 0
- rawEffectHeal => based on effect or 0
- rawTalismanDamage => based on talisman or 0
- rawTalismanHeal => based on talisman or 0

- rawTurnDamage => rawSkillDamage + rawEffectDamage + rawTalismanDamage
- rawTurnHeal => rawEffectHeal + rawTalismanHeal

- modifiedTurnDamage => rawTurnDamage +/- buff/debuff effects - source Entity
- modifiedTurnHeal => rawTurnHeal +/- buff/debuff effects - source Entity

- actualDamageTaken => modifiedTurnDamage +/- buff/debuff - target Entity
- actualHealTaken => modifiedTurnHeal +/- buff/debuff - target Entity

-----


calculationFunction

```go
// turn is handled with the call of the function

type CalculationValues struct {
  modifiedTurnDamage int
}

func calculate() {
  var source Entity
  var target Entity

  switch skillSource {
  case "player":
    source = GetEntity(PlayerEntity)
    target = GetEntity(BossEntity)
  case "boss":
    source = GetEntity(BossEntity)
    target = GetEntity(PlayerEntity)
  default:
    invalidEntityMsg := GetGameTextError("invalidentity")
    internalErrorMsg := GetGameTextError("internal")
    return fmt.Errorf("%s - %s", internalErrorMsg, invalidEntityMsg)
  }

  modifiedTurnDamage := 0

  if skill {
    modifiedTurnDamage = rawSkillDamage
  }
  

  for _, effect := range s.effectList {

    if effect.trigger == example {
      if effect.condition {
        val := effect.getValue()
        modifiedTurnDamage = modifiedTurnDamage + val

        effect.action(modifiedTurnDamage)
      }
    }

    if effect.trigger == onSkillUsed {}
    if effect.trigger == onTurnStart {}
    if effect.trigger == onTurnEnd {}
    if effect.trigger == onCalculation {}
    if effect.trigger == onApplyDamage {}
    if effect.trigger == onApplyHealing {}
    if effect.trigger == onRemoveActiveEffect {}
    if effect.trigger == onAddActiveEffect {}
  }

}

func getValue(val int) int {
 if val {
  return val
 } else {
  return val
 }
}

func action(val int) {

}


```




# Game Effects Table


**target**

- self
- enemy



**turn**
- own
- enemy



**Trigger**


- onSkillUsed (Skill.Use)
- onTurnStart (target.Turn)
- onTurnEnd (target.Turn)
- onCalculation (Skill.Use)
- onApplyDamage (target.ApplyDamage)
- onApplyHealing (target.ApplyHealing)
- onRemoveActiveEffect (target.RemoveActiveEffect)
- onAddActiveEffect (target.AddActiveEffect)


**Condition**

- whenAtFullHealth
- whenEnemyAtLowHealth
- whenBelow(x%)Health
- x(%) chance
- whenSameSkillWasUsedRoundBefore



**Action**

- setRawEffectDamage
- setRawEffectHeal



```
- increase/decrease modifiedEffectDamage
- increase/decrease modifiedEffectHeal
- increase/decrease modifiedSkillDamage
- apply modifiedEffectDamage currentHealth
- apply modifiedEffectHeal to currentHealth
- reflect modifiedSkillDamage/modifiedEffectDamage
```


**affected Value + effectValue**

rawEffectDamage/rawEffectHeal
- x% of fullTurnDamage




```
modifiedStrength
- x% of rawStrength
- x(%) for each 10% of missing Health
- x(%) for each active Buff
- x(%) for each defeated Boss

rawEffectDamage/rawEffectHeal
- x% for each active debuff
- x% for each active buff
- x2 modifier

- x% of actualDamageTaken
- x per modifiedStrength
- x per maxHealth

modifiedEffectDamage/modifiedSkillDamage
- x%
- x% for each active debuff
- x% for each active buff
- x2 modifier
- x% of modifiedSkillDamage
- x% of actualDamageTaken
- x per modifiedStrength
- set to 0
- x per maxHealth

modifiedEffectHeal
- set to 0
- x%
- x% for each active debuff
- x% for each active buff
- x2 modifier

activeEffectsList
- block next x debuffs/buffs
- block specific debuff/buffs
  - reduce Strength/damage
  - cause damage
  - prevent healing
  - stop skills/change targets
  - block buff prevention/effect manipulation
- add a random debuff/buff
- remove a random debuff/buff
- remove oldest debuff/buff
- remove neweset debuff/buff
- remove debuff/buff with most remaining turns
- remove debuff/buff with least remaining turns
```


## Buffs (Self)

|         Name         |        Trigger        |       Category      |     outputBase    |                          Description                           |
|----------------------|-----------------------|---------------------|-------------------|----------------------------------------------------------------|
| heal1                | etiOnActualDamage     | ecaHeal             | ?                 | Gain 50% of your Damage in Health                              |
| heal2                | etiOnActualDamage     | ecaHeal             | ActualDamageTaken | Heal 50% of the Damage you take                                |
| heal3                | etiOnTurnStart        | ecaHeal             | effectMulti       | Restores health at the start of each turn                      |
| heal4                | etiOnTurnEnd          | ecaHeal             | effectMulti       | Restores health at the end of each turn                        |
| heal5                | etiOnEffectRemoval    | ecaHeal             | effectMulti       | gain a massive heal when this effect expires or is removed     |
| doDamage1            | etiOnActualDamage     | ecaDoDamage         | ActualDamageTaken | deal 50% of the damage you Receive                             |
| increaseStrength1    | etiOnTurnStart        | ecaIncreaseStrength | effectMulti       | Increase strength                                              |
| increaseStrength2    | etiOnTurnStart        | ecaIncreaseStrength | effectMulti       | strength increases by 10% for each 10% of your missing health  |
| increaseStrength3    | etiOnTurnStart        | ecaIncreaseStrength | effectMulti       | strength increases by 10% for each active buff you have        |
| increaseStrength4    | etiOnTurnStart        | ecaIncreaseStrength | effectMulti       | strength increases by 20% when at full health                  |
| increaseStrength5    | etiOnTurnStart        | ecaIncreaseStrength | effectMulti       | strength increases by 10% for each defeated boss               |
| increaseStrength1    | etiOnSkillCalculation | ecaIncreaseStrength | effectMulti       | increase Strength                                              |
| increaseDamageDone1  | etiOnSkillCalculation | ?                   | effectMulti       | Adds bonus damage if enemy is low                              |
| increaseDamageDone2  | etiOnSkillCalculation | ?                   | effectMulti       | 50% Chance to double the damage                                |
| increaseDamageDone3  | etiOnSkillCalculation | ?                   | effectMulti       | Deal double the damage if the same skill is used repeatedly    |
| decreaseDamageTaken1 | etiOnSkillCalculation | ?                   | effectMulti       | Receive 50% less damage                                        |
| decreaseDamageTaken2 | etiOnSkillCalculation | ?                   | effectMulti       | Receive 10% less Damage from repeated sources                  |
| decreaseDamageTaken3 | etiOnSkillCalculation | ?                   |                   |                                                                |
| increaseHealing1     | etiOnSkillCalculation | ecaIncreaseHealing  | effectMulti       | Healing effects are 10% stronger for each debuff you have      |
| increaseHealing2     | etiOnSkillCalculation | ecaIncreaseHealing  | effectMulti       | Healing effects are 100% stronger when below 30% health        |
| increaseHealing3     | etiOnSkillCalculation | ecaIncreaseHealing  | effectMulti       | Healing effects are 10% stronger for each active buff          |
| blockDebuffs1        | etiOnSkillStart       | ecaBlockDebuffs     | -                 | Prevents new debuffs from being applied                        |
| blockDebuffs2        | etiOnSkillStart       | ecaBlockDebuffs     | -                 | 50% Chance to block an incoming Debuff                         |
| blockDebuffs3        | etiOnSkillStart       | ecaBlockDebuffs     | -                 | Block the next 3 debuffs                                       |
| blockDebuffs4        | etiOnSkillStart       | ecaBlockDebuffs     | -                 | Block debuffs that reduce Strength/damage                      |
| blockDebuffs5        | etiOnSkillStart       | ecaBlockDebuffs     | -                 | Block debuffs that cause damage                                |
| blockDebuffs6        | etiOnSkillStart       | ecaBlockDebuffs     | -                 | Block debuffs that prevent healing                             |
| blockDebuffs7        | etiOnSkillStart       | ecaBlockDebuffs     | -                 | Block debuffs that stop skills/change targets                  |
| blockDebuffs7        | etiOnSkillStart       | ecaBlockDebuffs     | -                 | Block buff prevention and effect manipulation                  |
| blockDamage1         | etiOnSkillStart       | ecaBlockDamage      | effectMulti       | 50% Chance to dont get damage                                  |
| blockDamage2         | etiOnSkillStart       | ecaBlockDamage      | effectMulti       | block damage if below 5% health                                |
| blockDamage3         | etiOnSkillStart       | ecaBlockDamage      | effectMulti       | block damage if the enemy used the same skill as a turn before |
| changeTarget1        | etiOnSkillStart       | ecaChangeTarget     | ?                 | 50% Chance to reflect the damage                               |
| changeTarget2        | etiOnSkillStart       | ecaChangeTarget     | -                 | 50% Chance to reflect the effects                              |
| changeTarget3        | etiOnSkillStart       | ecaChangeTarget     | ?                 | reflect the damage if below 5% health                          |
| removeEffect1        | etiOnTurnStart        | ecaRemoveEffect     | -                 | remove a random debuff                                         |
| removeEffect2        | etiOnSkillStart       | ecaRemoveEffect     | -                 | remove a random buff of the enemy when attacked                |
| removeEffect3        | etiOnSkillStart       | ecaRemoveEffect     | -                 | remove a random debuff when attacked                           |
| removeEffect4        | etiOnTurnEnd          | ecaRemoveEffect     | -                 | remove a random debuff that reduce Strength/damage             |
| removeEffect5        | etiOnTurnEnd          | ecaRemoveEffect     | -                 | remove a random debuff that cause damage                       |
| removeEffect6        | etiOnTurnEnd          | ecaRemoveEffect     | -                 | remove a random debuff that that prevent healing               |
| removeEffect7        | etiOnTurnEnd          | ecaRemoveEffect     | -                 | remove a random debuff that that stop skills/change targets    |
| removeEffect8        | etiOnTurnStart        | ecaRemoveEffect     | -                 | remove the oldest debuff                                       |
| removeEffect9        | etiOnTurnStart        | ecaRemoveEffect     | -                 | remove the newest debuff                                       |
| removeEffect10       | etiOnTurnStart        | ecaRemoveEffect     | -                 | remove the debuff with the most remaining turns                |
| removeEffect11       | etiOnTurnStart        | ecaRemoveEffect     | -                 | remove the debuff with the least remaining turns               |
| addEffect1           | etiOnTurnStart        | ecaAddEffect        | -                 | add a random buff                                              |
| addEffect2           | etiOnTurnEnd          | ecaAddEffect        | -                 | add a random buff                                              |


## Debuffs (Enemy)

|         Name         |        Trigger        |       Category      |  outputBase |                              Description                              |
|----------------------|-----------------------|---------------------|-------------|-----------------------------------------------------------------------|
| DecreaseStrength1    | etiOnSkillCalculation | ecaDecreaseStrength | effectMulti | Reduce targets Strength                                               |
| DecreaseStrength2    | etiOnSkillCalculation | ecaDecreaseStrength |             |                                                                       |
| DecreaseStrength3    | etiOnSkillCalculation | ecaDecreaseStrength |             |                                                                       |
| DecreaseStrength4    | etiOnSkillCalculation | ecaDecreaseStrength |             |                                                                       |
| DecreaseStrength5    | etiOnSkillCalculation | ecaDecreaseStrength |             |                                                                       |
| BlockHealing1        | etiOnSkillStart       | ecaBlockHealing     | effectMulti | 50% Chance to block heals                                             |
| BlockHealing2        | etiOnSkillStart       | ecaBlockHealing     |             |                                                                       |
| BlockHealing3        | etiOnSkillStart       | ecaBlockHealing     |             |                                                                       |
| TakeDamage1          | etiOnTurnStart        | ecaTakeDamage       | -           | Immediately kill the enemy while below 10% health                     |
| TakeDamage2          | etiOnTurnStart        | ecaTakeDamage       | effectMulti | Applies a damaging effect that deals damage at the start of each turn |
| TakeDamage3          | etiOnEffectRemoval    | ecaTakeDamage       | effectMulti | When cleansed, explodes                                               |
| TakeDamage4          | etiOnTurnStart        | ecaTakeDamage       | effectMulti | Deals damage based on maximum health                                  |
| TakeDamage5          | etiOnTurnStart        | ecaTakeDamage       |             |                                                                       |
| TakeDamage6          | etiOnTurnStart        | ecaTakeDamage       |             |                                                                       |
| DecreaseDamageDone1  | etiOnSkillCalculation | ?                   | effectMulti | Reduces target's damage output by 50%                                 |
| DecreaseDamageDone2  | etiOnSkillCalculation | ?                   |             |                                                                       |
| DecreaseDamageDone3  | etiOnSkillCalculation | ?                   |             |                                                                       |
| DecreaseDamageDone4  | etiOnSkillCalculation | ?                   |             |                                                                       |
| DecreaseDamageDone5  | etiOnSkillCalculation | ?                   |             |                                                                       |
| DecreaseDamageDone6  | etiOnSkillCalculation | ?                   |             |                                                                       |
| IncreaseDamageTaken1 | etiOnSkillCalculation | ecaIncrease?        | effectMulti | Receive 50% more damage                                               |
| IncreaseDamageTaken2 | etiOnSkillCalculation | ecaIncrease?        |             |                                                                       |
| IncreaseDamageTaken3 | etiOnSkillCalculation | ecaIncrease?        |             |                                                                       |
| IncreaseDamageTaken4 | etiOnSkillCalculation | ecaIncrease?        |             |                                                                       |
| IncreaseDamageTaken5 | etiOnSkillCalculation | ecaIncrease?        |             |                                                                       |
| IncreaseDamageTaken6 | etiOnSkillCalculation | ecaIncrease?        |             |                                                                       |
| DecreaseHealing1     | etiOnSkillCalculation | ecaDecreaseHealing  | effectMulti | Reduces all healing received by 50%                                   |
| DecreaseHealing2     | etiOnSkillCalculation | ecaDecreaseHealing  |             |                                                                       |
| DecreaseHealing3     | etiOnSkillCalculation | ecaDecreaseHealing  |             |                                                                       |
| BlockBuffs1          | etiOnSkillStart       | ecaBlockBuffs       | -           | Prevents the target from receiving buffs                              |
| BlockBuffs2          | etiOnSkillStart       | ecaBlockBuffs       | effectMulti | Prevents the target from receiving healing effects                    |
| BlockBuffs3          | etiOnSkillStart       | ecaBlockBuffs       | effectMulti | Prevents the target from receiving ? effects                          |
| ChangeTarget1        | etiOnSkillStart       | ecaChangeTarget     | ?           | 50% Chance to attack itself                                           |
| RemoveEffect1        | etiOnTurnStart        | ecaRemoveEffect     | -           | remove a random Buff                                                  |
| RemoveEffect2        | etiOnSkillStart       | ecaRemoveEffect     | -           | remove a random Buff when attacked                                    |
| RemoveEffect3        | etiOnSkillStart       | ecaRemoveEffect     | -           | remove a random debuff of the enemy when attacked                     |
| addEffect1           | etiOnTurnStart        | ecaAddEffect        | -           | add a random debuff                                                   |
| addEffect2           | etiOnTurnEnd          | ecaAddEffect        | -           | add a random debuff                                                   |


## Talismans (One-time Use Items)

|        Name       |     Category     |                          Description                          |
|-------------------|------------------|---------------------------------------------------------------|
| DoDamage2         | DoDamage         | Bonus Damage to amount of all remaining Buff Turns            |
| DoDamage3         | DoDamage         | Bonus Damage to Amount of all remaining Debuff Turns of Enemy |
| Heal4             | Heal             | Immediately restores health based on Strength stat               |
| Heal5             | Heal             | Stronger direct heal but remove one random buff               |
| Heal6             | Heal             | Pay 10% of your current health to remove a random Debuff      |
| RemoveEffect4     | RemoveEffect     | Removes all negative effects from the target                  |
| RemoveEffect5     | RemoveEffect     | Removes all positive effects from the target                  |
| ChangeEffectTurn1 | ChangeEffectTurn | +1 Turn for each 10 Total Strength                               |
| ChangeEffectTurn2 | ChangeEffectTurn | +1 Turn for each 10 Total Strength                               |
| ChangeEffectTurn3 | ChangeEffectTurn | -1 Turn for each 10 Total Strength                               |
| ChangeEffectTurn4 | ChangeEffectTurn | -1 Turn for each 10 Total Strength                               |

## Additional Effect Ideas

|        Name        | Type |                                                        Description                                                        |
|--------------------|------|---------------------------------------------------------------------------------------------------------------------------|
| Random Effect      | misc | Get a random buff/debuff when attacking/attacked from a list of random effects                                            |
| Effect Removal     | misc | Remove a random buff/debuff when attacking/attacked                                                                       |
| Status Swap        | misc | Swap all buffs/debuffs                                                                                                    |
| Strength Theft        | misc | Steal enemy Strength                                                                                                         |
| Health Sacrifice   | misc | Sacrifice health to deal damage/increase damage                                                                           |
| Time Bomb          | misc | Effect only does damage when it naturally runs out of turns (no removal) but increase the damage for every turn it passed |
| Health Siphon      | misc | Each turn, steal a small amount of health from the enemy                                                                  |
| influence Talisman | misc | block talisman..                                                                                                          |


# tmp


                          Description                           |
----------------------------------------------------------------|
 Adds bonus damage if enemy is low                              |
 50% Chance to double the damage                                |
 Deal double the damage if the same skill is used repeatedly    |
 Receive 50% less damage                                        |
 Receive 10% less Damage from repeated sources                  |
 Healing effects are 10% stronger for each debuff you have      |
 Healing effects are 100% stronger when below 30% health        |
 Healing effects are 10% stronger for each active buff          |
 Prevents new debuffs from being applied                        |
 50% Chance to block an incoming Debuff                         |
 Block the next 3 debuffs                                       |
 Block debuffs that reduce Strength/damage                      |
 Block debuffs that cause damage                                |
 Block debuffs that prevent healing                             |
 Block debuffs that stop skills/change targets                  |
 Block buff prevention and effect manipulation                  |
 50% Chance to dont get damage                                  |
 block damage if below 5% health                                |
 block damage if the enemy used the same skill as a turn before |
 50% Chance to reflect the damage                               |
 50% Chance to reflect the effects                              |
 reflect the damage if below 5% health                          |
 remove a random debuff                                         |
 remove a random buff of the enemy when attacked                |
 remove a random debuff when attacked                           |
 remove a random debuff that reduce Strength/damage             |
 remove a random debuff that cause damage                       |
 remove a random debuff that that prevent healing               |
 remove a random debuff that that stop skills/change targets    |
 remove the oldest debuff                                       |
 remove the newest debuff                                       |
 remove the debuff with the most remaining turns                |
 remove the debuff with the least remaining turns               |
 add a random buff                                              |
 add a random buff                                              |


|                              Description                              |
|-----------------------------------------------------------------------|
| Reduce targets Strength                                               |
| 50% Chance to block heals                                             |
| Immediately kill the enemy while below 10% health                     |
| Applies a damaging effect that deals damage at the start of each turn |
| When cleansed, explodes                                               |
| Deals damage based on maximum health                                  |
| Reduces target's damage output by 50%                                 |
| Receive 50% more damage                                               |
| Reduces all healing received by 50%                                   |
| Prevents the target from receiving buffs                              |
| Prevents the target from receiving healing effects                    |
| Prevents the target from receiving ? effects  |
| 50% Chance to attack itself                                           |
| remove a random Buff                                                  |
| remove a random Buff when attacked                                    |
| remove a random debuff of the enemy when attacked                     |
| add a random debuff                                                   |
| add a random debuff                                                   |
