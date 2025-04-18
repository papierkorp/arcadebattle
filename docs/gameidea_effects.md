# Effects

Primary Function:

- increase
  - power (for fight, max)
  - current_health (for fight, max)
  - damagedealt
  - damagereceived
  - self effect durations
  - enemy effect durations
- decrease
  - power (for fight, max)
  - current_health (for fight, max)
  - damagedealt
  - damagereceived
  - self effect durations
  - enemy effect durations
- block
  - damage
  - application of effect
  - increase of health
  - increase of power
  - skill
- ???
  - Shield + Pierce

Targets:

- self
- enemy

Each Effect can be placed as on of this 3:

- based on power\*Skillmulti
- based on EffectCount (self/enemy)
- based on Turnsleft for all Effects (self/enemy)

Category:

- buff
- debuff

Battle Loop:

- Start of Turn: apply effect
- action: add effect to list
- end of turn:

---

---

---

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
