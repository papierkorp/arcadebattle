// Package internal comment
package internal

import (
	"fmt"
	"strings"
)

//todo display all the costs in help

type costList map[string]int

// ---------------------------------------------------------------
// -------------------------- BALANCING --------------------------
// ---------------------------------------------------------------

func getTalentpoints(d difficulty) int {
	switch d {
	case normal:
		return 1000
	case hard:
		return 80
	case expert:
		return 40
	case master:
		return 20
	case torment:
		return 10
	default:
		return 100
	}
}

func getEffectCost(effectName string) (int, error) {
	effectCosts := costList{
		"heal1": 5,
	}
	cost, exists := effectCosts[effectName]
	if !exists {
		internalErrorMsg := GetGameTextError("internal")
		invalidEffectMsg := GetGameTextError("invalideffect")
		return 0, fmt.Errorf("%s: %s (%s)", internalErrorMsg, invalidEffectMsg, effectName)
	}
	return cost, nil
}

func getSkillCost(skillPart string) (int, costList, error) {
	var skillCostsList = costList{
		"dmgmulti":     1,  // 1 per 0.1 increase
		"duration":     5,  // 5 per 1 Turn
		"passivemulti": 10, // if type == passive, effect cost * 10
	}
	cost, exists := skillCostsList[skillPart]
	if !exists {
		internalErrorMsg := GetGameTextError("internal")
		invalidskillcostMsg := GetGameTextError("invalidskillcost")
		return 0, skillCostsList, fmt.Errorf("%s: %s (%s)", internalErrorMsg, invalidskillcostMsg, skillPart)
	}
	return cost, skillCostsList, nil
}

func getStatsCost(statName string) (int, costList, error) {
	var statCostList = costList{
		"hp_cost":    1, // cost per 5 HP
		"power_cost": 1, // cost per 1 Power
		"speed_cost": 2, // cost per 1 speed
	}
	cost, exists := statCostList[statName]
	if !exists {
		internalErrorMsg := GetGameTextError("internal")
		invalidStatMsg := GetGameTextError("invalidstat")
		return 0, statCostList, fmt.Errorf("%s: %s (%s)", internalErrorMsg, invalidStatMsg, statName)
	}
	return cost, statCostList, nil
}

// ---------------------------------------------------------------
// ------------------------- CALCULATION -------------------------
// ---------------------------------------------------------------

func calculateSkillCost(s *Skill) int {
	// calculationMsg := GetGameTextGameMessage("talentcalculation")
	// totalMsg := GetGameTextGameMessage("total")
	// effectMsg := GetGameTextGameMessage("effect")
	// dmgMultiMsg := GetGameTextGameMessage("damagemultiplier")
	// durationMsg := GetGameTextGameMessage("duration")
	// costper01multiMsg := GetGameTextGameMessage("costper01multi")

	// err, dmgMultiCostPer01Multi, _ := getSkillCost("dmgmulti")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// var dmgCost int = 0

	// totalEffectCost := 0
	// err, durationCostperTurn, _ := getSkillCost("duration")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println("\n" + calculationMsg + "\n")

	// for _, effect := range s.effectList {
	// 	effectTotalCost := effect.cost

	// 	additionalCost := effect.duration * durationCostperTurn
	// 	effectTotalCost += additionalCost

	// 	// if s.skilltype != PassiveEffect

	// 	costperturnMsg := GetGameTextGameMessage("costperturn")

	// 	// Damage Over Time   20 (Duration * Cost per Turn: 3 * 5 + Effect: 5)
	// 	fmt.Printf("  %-18s %d (%s * %s: %d * %d + %s: %d)\n",
	// 		effect.name,
	// 		effectTotalCost,
	// 		durationMsg,
	// 		costperturnMsg,
	// 		effect.duration,
	// 		durationCostperTurn,
	// 		effectMsg,
	// 		effect.cost)

	// 	totalEffectCost += effectTotalCost
	// }

	// totalCost := totalEffectCost + dmgCost

	// fmt.Printf("  %-18s %d (%s x %s: %.2f x %d)\n", dmgMultiMsg, dmgCost, dmgMultiMsg, costper01multiMsg, s.dmgmulti, dmgMultiCostPer01Multi)
	// fmt.Println("  " + strings.Repeat("_", 24))
	// fmt.Printf("  %-18s %d\n\n", totalMsg, totalCost)

	// s.talentpointcoststotal = totalCost

	return 0
}

func calculateStatsCost(stats Stats) int {
	hpStats := stats.health
	hpCost, _, err := getStatsCost("hp_cost")
	if err != nil {
		fmt.Println(err)
	}

	powerStats := stats.power
	powerCost, _, err := getStatsCost("power_cost")
	if err != nil {
		fmt.Println(err)
	}

	speedStats := stats.speed
	speedCost, _, err := getStatsCost("speed_cost")
	if err != nil {
		fmt.Println(err)
	}

	hpTotalCost := hpStats * hpCost
	powerTotalCost := powerStats * powerCost
	speedTotalCost := speedStats * speedCost

	totalCost := hpTotalCost + powerTotalCost + speedTotalCost

	playerStatus := GetGameTextStatusPlayer()
	calculationMsg := GetGameTextGameMessage("talentcalculation")
	totalMsg := GetGameTextGameMessage("total")
	costperstatMsg := GetGameTextGameMessage("costperstat")

	// HP         10 (HP x Cost per Stat: 10 x 1)
	// Power      5 (Power x Cost per Stat: 5 x 1)
	// Speed      14 (Speed x Cost per Stat: 7 x 2)
	// ----------------
	// Total      29

	fmt.Println(calculationMsg)
	fmt.Printf("\n  %-10s %d (%s x %s: %d x %d)\n", playerStatus.Health, hpTotalCost, playerStatus.Health, costperstatMsg, hpStats, hpCost)
	fmt.Printf("  %-10s %d (%s x %s: %d x %d)\n", playerStatus.Power, powerTotalCost, playerStatus.Power, costperstatMsg, powerStats, powerCost)
	fmt.Printf("  %-10s %d (%s x %s: %d x %d)\n", playerStatus.Speed, speedTotalCost, playerStatus.Speed, costperstatMsg, speedStats, speedCost)
	fmt.Println("  " + strings.Repeat("-", 16))
	fmt.Printf("  %-10s %d\n", totalMsg, totalCost)
	return totalCost
}

func changeTalentpoints(cost int) error {
	currentTalentpoints := currentPlayer.talentpointsRemaining

	if cost > currentTalentpoints {
		exceededMsg := GetGameTextError("exceededtalentpoints")
		separatorMsg := GetGameTextGameMessage("separator")
		usedMsg := GetGameTextGameMessage("used")
		availableMsg := GetGameTextGameMessage("available")
		return fmt.Errorf("%s, %s: %d %s %s: %d",
			exceededMsg,         // Used too much Talentpoints
			availableMsg,        // Available:
			currentTalentpoints, // available value
			separatorMsg,        //  |
			usedMsg,             // Used:
			cost)                // used value
	}

	currentPlayer.talentpointsRemaining = currentTalentpoints - cost

	return nil
}
