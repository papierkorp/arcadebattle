# Effect Usage

**Turn Process for use skill**

1. loop effects for turn start
- entity: `both`
- usageTiming: `etiOnTurnStart`
2. `use skill` command
3. loop if an effect is blocked / target is changed
- entity: `both`
- usageTiming: `etiOnSkillStart`
4. loop effects to create `fullSkillPower` (`currentPower * skillMulti * each effectMulti`)
- usageTiming: `etiOnSkillCalculation`
5. use fullSkillPower to do the damage
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
  ecaRemoveEffect
  ecaChangeEffectTurn
)

type EffectTiming int
const (
  etiOnTurnStart EffectTiming = iota
  etiOnSkillStart
  etiOnSkillCalculation
  etiOnTurnEnd
  etiOnEffectRemoval
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

# Calculations

example fullSkillPower:

```
// new player name difficulty health power speed
// new skill <name> <dmgmulti> <duration> [effect effect effect...]

new player John normal 125 15 8
new skill Fireball 1.5 4

1.5    * 15      = 23
23     * 1,46    = 34
```

|       Category      |      trigger       |      Description      |             Calc             | Value |
|---------------------|--------------------|-----------------------|------------------------------|-------|
| SkillDamage         | -                  | 1 damage per 1 power  | fullSkillPower               | 23    |
| Heal                | onTurnEnd          | 2 Health per 3 power  | fullSkillPower / 3 * 2       | 16    |
| DoDamage            | onSkillCalculation | -                     | -                            | -     |
| IncreasePower       | onSkillCalculation | 0.1 multi per 5 power | fullSkillPower / 5 * 0,1 + 1 | 1,46  |
| IncreaseDamageDone  | onSkillCalculation | 0.1 multi per 5 power | fullSkillPower / 5 * 0,1 + 1 | 1,46  |
| DecreaseDamageTaken | onSkillCalculation | 0.1 multi per 5 power | fullSkillPower / 5 * 0,1 + 1 | 1,46  |
| IncreaseHealing     | onSkillCalculation | 0.1 multi per 3 power | fullSkillPower / 3 * 0,1 + 1 | 1,76  |
| BlockDebuffs        | onSkillStart       | -                     | -                            | -     |
| BlockDamage         | onSkillStart       | -                     | -                            | -     |
| StopSkill           | onSkillStart       | -                     | -                            | -     |
| ChangeTarget        | onSkillStart       | -                     | -                            | -     |
| RemoveEffect        | onSkillStart       | 1 effect per 10 power | fullSkillPower / 10          | 2     |
| DecreasePower       | onSkillCalculation | 0.1 multi per 5 power | fullSkillPower / 5 * 0,1 + 1 | 1,46  |
| BlockHealing        | onSkillStart       | -                     | -                            | -     |
| TakeDamage          | onTurnStart        | 2 Damage per 3 power  | fullSkillPower / 3 * 2       | 16    |
| DecreaseDamageDone  | onSkillCalculation | 0.1 multi per 5 power | fullSkillPower / 5 * 0,1 + 1 | 1,46  |
| IncreaseDamageTaken | onSkillCalculation | 0.1 multi per 5 power | fullSkillPower / 5 * 0,1 + 1 | 1,46  |
| DecreaseHealing     | onSkillCalculation | 0.1 multi per 5 power | fullSkillPower / 5 * 0,1 + 1 | 1,46  |
| BlockBuffs          | onSkillStart       | -                     | -                            | -     |
| ChangeEffectTurn    | onSkillStart       | 1 turn per 10 power   | fullSkillPower / 10          | 2     |

|      trigger       |      description      |
|--------------------|-----------------------|
| onSkillCalculation | 0.1 multi per 5 power |
| onTurnEnd          | 2 y per 3 power       |
| onTurnStart        | 2 y per 3 power       |
| onSkillStart       | 1 y per 10 power      |

unless otherwise defined


# Game Effects Table


## Buffs (Self)

|         Name         |      Trigger       |       Category      |                                   Description                                   |
|----------------------|--------------------|---------------------|---------------------------------------------------------------------------------|
| heal1                | onTurnEnd          | Heal                | Gain 50% of your Damage in Health                                               |
| heal2                | onTurnEnd          | Heal                | Heal 50% of the Damage you take                                                 |
| heal3                | onTurnEnd          | Heal                | Restores health at the start of each turn                                       |
| doDamage1            | onTurnEnd          | DoDamage            | Reflect 50% of the damage you Receive                                           |
| increasePower1       | onSkillCalculation | IncreasePower       | Increase power                                                                  |
| increasePower2       | onSkillCalculation | IncreasePower       | Power increases by 10% each turn, up to 50%                                     |
| increasePower3       | onSkillCalculation | IncreasePower       | Power increases by 10% for each 10% of your missing health                      |
| increasePower4       | onSkillCalculation | IncreasePower       | Each consecutive use of the same skill increases power by 15%                   |
| increaseDamageDone1  | onSkillCalculation | IncreaseDamageDone  | Adds bonus damage if enemy is low                                               |
| increaseDamageDone2  | onSkillCalculation | IncreaseDamageDone  | 50% Chance to double the damage                                                 |
| increaseDamageDone3  | onSkillCalculation | IncreaseDamageDone  | Each attack increases damage of next attack                                     |
| decreaseDamageTaken1 | onSkillCalculation | DecreaseDamageTaken | Receive 50% less damage                                                         |
| decreaseDamageTaken2 | onSkillCalculation | DecreaseDamageTaken | Receive 10% less Damage from repeated sources                                   |
| decreaseDamageTaken3 | onSkillCalculation | DecreaseDamageTaken | Convert 20% of damage taken into a DoT on yourself that deals less total damage |
| increaseHealing1     | onSkillCalculation | IncreaseHealing     | Healing effects become 30% stronger for each debuff you have                    |
| increaseHealing2     | onSkillCalculation | IncreaseHealing     | Healing effects are 100% stronger when below 30% health                         |
| blockDebuffs1        | onSkillStart       | BlockDebuffs        | Prevents new debuffs from being applied while active                            |
| blockDebuffs2        | onSkillStart       | BlockDebuffs        | 50% Chance to block an incoming Debuff                                          |
| blockDamage1         | onSkillStart       | BlockDamage         | 50% Chance to dont get damage                                                   |
| stopSkill1           | onSkillStart       | StopSkill           | 50% Chance to block an incoming skill                                           |
| changeTarget1        | onSkillStart       | ChangeTarget        | 50% Chance the attack is mirrored                                               |
| removeEffect1        | onSkillStart       | RemoveEffect        | remove a random debuff                                                          |
| removeEffect2        | onSkillStart       | RemoveEffect        | remove a random buff of the enemy when attacked                                 |
| removeEffect3        | onSkillStart       | RemoveEffect        | remove a random debuff when attacked                                            |

## Debuffs (Enemy)

|         Name         |      Trigger       |       Category      |                              Description                              |
|----------------------|--------------------|---------------------|-----------------------------------------------------------------------|
| DecreasePower1       | onSkillCalculation | DecreasePower       | Reduce targets Power                                                  |
| BlockHealing1        | onSkillStart       | BlockHealing        | 50% Chance to block heals                                             |
| TakeDamage1          | onTurnStart        | TakeDamage          | Immediately kill the enemy while below 10% health                     |
| TakeDamage2          | onTurnStart        | TakeDamage          | Applies a damaging effect that deals damage at the start of each turn |
| TakeDamage3          | onTurnStart        | TakeDamage          | When cleansed, explodes                                               |
| TakeDamage4          | onTurnStart        | TakeDamage          | Deals damage based on maximum health                                  |
| DecreaseDamageDone1  | onSkillCalculation | DecreaseDamageDone  | Reduces target's damage output by 50%                                 |
| IncreaseDamageTaken1 | onSkillCalculation | IncreaseDamageTaken | Receive 50% more damage                                               |
| DecreaseHealing1     | onSkillCalculation | DecreaseHealing     | Reduces all healing received by 50%                                   |
| BlockBuffs1          | onSkillStart       | BlockBuffs          | Prevents the target from receiving buffs and healing effects          |
| StopSkill1           | onSkillStart       | StopSkill           | 50% Chance to miss the skill                                          |
| StopSkill2           | onSkillStart       | StopSkill           | Cannot use the same skill twice in a row                              |
| ChangeTarget1        | onSkillStart       | ChangeTarget        | 50% Chance to attack itself                                           |
| RemoveEffect1        | onTurnStart        | RemoveEffect        | remove a random Buff                                                  |
| RemoveEffect2        | onSkillStart       | RemoveEffect        | remove a random Buff when attacked                                    |
| RemoveEffect3        | onSkillStart       | RemoveEffect        | remove a random debuff of the enemy when attacked                     |


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
