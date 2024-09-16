/**
1. Document package weather
Since goblins are not as smart as you are, they forgot what the package should do for them. Please write a comment for package weather that describes its contents. The package comment should introduce the package and provide information relevant to the package as a whole.

2. Document the CurrentCondition and CurrentLocation variables
The president of Goblinocus is a bit paranoid and fears uncommented variables are used to destroy their country. Please clarify the usage of the package variables CurrentCondition and CurrentLocation and put the president's mind at ease. This should tell any user of the package what information the variables store, and what they can do with it.

3. Document the Forecast() function
Goblinocus forecast operators want to know what the Forecast() function does (but do not tell them how it works, since unfortunately, they will get more confused). Please write a comment for this function that describes what the function does, but not how it does it.
**/

// Package weather provides tools to display information about wheater.
package weather

// CurrentCondition represents the current weather condition.
var CurrentCondition string

// CurrentLocation represents the current location.
var CurrentLocation string

// Forecast returns a string with information about whearer in a given city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
