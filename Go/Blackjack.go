package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
    switch {
        case card=="ten" || card=="king" || card=="queen" || card=="jack":
        	return 10
        case card=="ace":
        	return 11
        case card=="eight":
        	return 8
        case card=="two":
        	return 2
        case card=="nine":
        	return 9
        case card=="three":
        	return 3
        case card=="four":
        	return 4
        case card=="five":
        	return 5
        case card=="six":
        	return 6
        case card=="seven":
        	return 7
        default:
        	return 0
    }
	panic("Please implement the ParseCard function")
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
    sumCards := ParseCard(card1)+ParseCard(card2)
    
    switch {
        case card1=="ace" && card2=="ace":
        	return "P"
        case sumCards >=21:
        	if ParseCard(dealerCard)>=10 {
                return "S"
            }
        	return "W"
        case sumCards >16 && sumCards<21:
        	return "S"
        case sumCards >11 && sumCards<17:
        	if ParseCard(dealerCard) > 6 {
                return "H"
            }
        	return "S"
        default:
        	return  "H"
    }
	panic("Please implement the FirstTurn function")
}
