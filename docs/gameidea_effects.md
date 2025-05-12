# Effect Usage

**Turn Process for use skill**

1. check effects for turn start - effectgroups: 
- take damage
2. `use skill` command
3. loop if an effect is blocked (effect.type block) - effectgroups: 
- block debuffs
- block buffs
- block healing
4. loop for special effects - effectgroups: 
- block damage
- stop skill
- change target
5. determine `basicSkillPower` (`currentPower * skillMulti`)
6. loop effects of both entities to create `fullSkillPower` (`currentPower * skillMulti * effectMulti`) (can be more effect multis, e.g. 0,5 for damage reduction or 2,0 for crit) with `effect.multi` - effectgroups: 
- increase damage done
- decrease damage taken 
- decrease damage done 
- increase power 
- increase damage taken 
- decrease healing 
- increase healing
7. use fullSkillPower to do the damage
8. loop effects for turn end and reduction of turns of effects and reactions - effectgroups: 
- heal 
- do damage

**Turn Process for use talisman**

1. check effects for turn start
- execution
2. `use talisman` command
...


# Effectlist (grouping)

https://github.com/papierkorp/arcadebattle/blob/1c62a25f4e3e0624f7643347efaf1ef25b9b21d0/docs/gameidea.md

**buff**

- increase power
  - IncPower: `Increases damage output by 50%`
  - Adrenaline Rush: `Power increases by 10% each turn, up to 50%`
  - Vengeance: `Power increases by 10% for each 10% of your missing health`
  - Combo Mastery: `Each consecutive use of the same skill increases power by 15%`
- increase damage done
  - Finisher: `adds bonus damage if enemy is low`
  - critical strike: `50% Chance to double the damage`
  - Fury: `each attack increases damage of next attack`
- decrease damage taken
  - Immunity: `Receive 50% less damage`
  - Adapation: `Receive 10% less Damage from repeated sources`
  - Damage Dispersal: `Convert 20% of damage taken into a DoT on yourself that deals less total damage`
- increase healing
  - Regenerative Core: `Healing effects become 30% stronger for each debuff you have`
  - Crisis Recovery: `Healing effects are 100% stronger when below 30% health`
- block debuffs
  - Block Debuffs: `Prevents new debuffs from being applied while active`
  - Resistance: `50% Chance to block an incoming Debuff`
- block damage
  - evasion: `50% Chance to dont get damage`
- stop skill
  - luckyTiming: `50% Chance to block an incoming skill`
- heal
  - Lifeleech: `Gain 50% of your Damage in Health`
  - lastStand: `Heal 50% of the Damage you take`
  - Heal Over Time (HOT): `Restores health at the start of each turn`
- do damage
  - reflect damage: `Reflect 50% of the damage you Receive`
- change target
  - Mirror: `50% Chance the attack is mirrored`

**debuff**

- decrease power
- decrease damage done
  - Weaken: `Reduces target's damage output by 50%`
- increase damage taken
  - Vulnerability: `Receive 50% more damage`
- decrease healing
  - ReduceHealing: `Reduces all healing received by 50%`
- block buffs
  - BlockBuffs: `Prevents the target from receiving buffs and healing effects`
- block healing
- stop skill
  - Confusion: `50% Chance to miss the skill`
  - Mental Block: `cannot use the same skill twice in a row`
- take damage
  - Execution: `Immediately kill the enemy while below 10% health`
  - Bleeding: `Applies a damaging effect that deals damage at the start of each turn`
  - Unstable Affliction: `When cleansed, explodes`
  - Soul Burn: `Deals damage based on maximum health`
- change target
  - Distraction: `50% Chance to attack itself`

**more ideas**

- get a random buff/debuff when attacking/attacked
- remove a random buff/debuff when attacking/attacked
- status swap (swap all buffs/debuffs)
- steal enemy power
- sacrifice health to deal damage/increase damage
- effect only does damage when it naturally runs out of turns (no removal) but increase the damage for every turn it passed
- Health Siphon: Each turn, steal a small amount of health from the enemy

# Talismanlist (one time use items)

- do damage
  - BuffTurnBonusDamage: `Bonus Damage to amount of all remaining Buff Turns`
  - DebuffTurnBonusDamage: `Bonus Damage to Amount of all remaining Debuff Turns of Enemy`
- heal
  - Heal: `Immediately restores health based on power stat`
  - BuffHeal: `Stronger direct heal but remove one random buff`
  - HealthCleanse: `Pay 10% of your current health to remove a random Debuff`
- remove effects
  - Cleanse: `Removes all negative effects from the target`
  - Dispel: `Removes all positive effects from the target`
- change effect turns
  - ExtendBuffs: `+1 Turn for each 10 Total Power`
  - ExtendDebuffs: `+1 Turn for each 10 Total Power`
  - ReduceBuffs: `-1 Turn for each 10 Total Power`
  - ReduceDebuffs: `-1 Turn for each 10 Total Power` 