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
