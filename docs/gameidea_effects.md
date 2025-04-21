# Effects new grouping

consists of:

- name
- description
- talentpointCosts
- isBlockedBy (can the effect be blocked by another effect)
- probability (if there is a probability involved else 100%)
- affectedStat (see in PrimaryFunction which stats can be affected)
    - power (for fight, max)
    - current_health (for fight, max)
    - damagedealt
    - damagereceived
    - self effect durations
    - enemy effect durations
    - effect (application)
    - health (increase)
    - power (increase)
    - damage (increase)
    - skill
- damageBase (how is the damage calculated)
  - power
  - RemainingTurnsCount
  - TotalEffectsCount
  - enemyLowHealth
  - ...
- PrimaryFunctionValue (always a % value i think?)
- PrimaryFunction (what is the effect doing)
  - increase
  - decrease
  - block
  - ???
    - Shield + Pierce
- Target (who to target)
  - self
  - enemy
- Cost (does it cost something to activate, can be null)
  - nothing (just cast the effect)
  - remove a buff
  - remove some health
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

|           Name           | TP |        IsBlockedBy        | Probability |      affectedStat      |       damageBase       | PrimaryFunction  | PrimaryFunctionValue | Target | CostType | CostValue | Category | EffectTiming |
|--------------------------|----|---------------------------|-------------|------------------------|------------------------|------------------|----------------------|--------|----------|-----------|----------|--------------|
| Direct Damage            |  5 | -                         | 100%        | damagedealt            | power                  | increase         | 100%                 | enemy  | -        | -         | damage   | OnSkillUse   |
| Pierce                   |  5 | -                         | 100%        | damagedealt            | power                  | bypass           | 50%                  | enemy  | -        | -         | damage   | OnSkillUse   |
| Finisher                 |  5 | -                         | 100%        | damagedealt            | power                  | increase         | 50%                  | enemy  | -        | -         | damage   | OnSkillUse   |
| Buff Turn Bonus Damage   |  5 | -                         | 100%        | damagedealt            | self effect durations  | increase         | 5% per turn          | enemy  | -        | -         | damage   | OnSkillUse   |
| Debuff Turn Bonus Damage |  5 | -                         | 100%        | damagedealt            | enemy effect durations | increase         | 5% per turn          | enemy  | -        | -         | damage   | OnSkillUse   |
| Direct Heal              |  5 | BlockBuffs                | 100%        | current_health         | power                  | increase         | 50%                  | self   | -        | -         | buff     | OnSkillUse   |
| Life Leech               |  5 | BlockBuffs, ReduceHealing | 100%        | current_health         | damagedealt            | increase         | 50%                  | self   | -        | -         | buff     | OnSkillUse   |
| Cleanse                  |  5 | -                         | 100%        | effect                 | -                      | decrease(remove) | 100%                 | self   | -        | -         | buff     | OnSkillUse   |
| Dispel                   |  5 | BlockDebuffs              | 100%        | effect                 | -                      | decrease(remove) | 100%                 | enemy  | -        | -         | debuff   | OnSkillUse   |
| Extend Buffs             |  5 | -                         | 100%        | self effect durations  | power                  | increase         | 10%                  | self   | -        | -         | buff     | OnSkillUse   |
| Extend Debuffs           |  5 | BlockDebuffs              | 100%        | enemy effect durations | power                  | increase         | 10%                  | enemy  | -        | -         | debuff   | OnSkillUse   |
| Reduce Debuffs           |  5 | -                         | 100%        | self effect durations  | power                  | decrease         | 10%                  | self   | -        | -         | buff     | OnSkillUse   |
| Reduce Buffs             |  5 | BlockDebuffs              | 100%        | enemy effect durations | power                  | decrease         | 10%                  | enemy  | -        | -         | debuff   | OnSkillUse   |
| Block Debuffs            |  5 | Dispel                    | 100%        | effect                 | -                      | block            | 100%                 | self   | -        | -         | buff     | OnTurnStart  |
| Heal Over Time           |  5 | BlockBuffs, ReduceHealing | 100%        | current_health         | power                  | increase         | 10%                  | self   | -        | -         | buff     | OnTurnStart  |
| Increase Power           |  5 | Dispel                    | 100%        | power                  | power                  | increase         | 50%                  | self   | -        | -         | buff     | OnTurnStart  |
| Shield                   |  5 | Pierce, Dispel            | 100%        | damagereceived         | power                  | absorb           | 25%                  | self   | -        | -         | buff     | OnTurnStart  |
| Reflect Damage           |  5 | Dispel                    | 100%        | current_health         | damagereceived         | decrease         | 50%                  | enemy  | -        | -         | buff     | OnTurnStart  |
| Evasion                  |  5 | Dispel                    | 50%         | damagereceived         | power                  | decrease         | 100%                 | self   | -        | -         | buff     | OnTurnStart  |
| Critical Rate            |  5 | Dispel                    | 50%         | damagedealt            | power                  | increase         | 200%                 | self   | -        | -         | buff     | OnTurnStart  |
| Damage Over Time         |  5 | BlockDebuffs, Cleanse     | 100%        | damagedealt            | power                  | decrease         | 25%                  | enemy  | -        | -         | debuff   | OnTurnStart  |
| Stun                     |  5 | BlockDebuffs, Cleanse     | 100%        | skill                  | -                      | block            | 100%                 | enemy  | -        | -         | debuff   | OnTurnStart  |
| Damage Reduction         |  5 | BlockDebuffs, Cleanse     | 100%        | damagedealt            | power                  | decrease         | 50%                  | enemy  | -        | -         | debuff   | OnTurnStart  |
| Block Buffs              |  5 | BlockDebuffs, Cleanse     | 100%        | effect                 | -                      | block            | 100%                 | enemy  | -        | -         | debuff   | OnTurnStart  |
| Grievous Wounds          |  5 | BlockDebuffs, Cleanse     | 100%        | current_health         | power                  | decrease         | 50%                  | enemy  | -        | -         | debuff   | OnTurnStart  |


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

