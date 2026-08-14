// Package weather provides a forecasting functionality that describes the
// current weather condition of a specific city.
package weather

// CurrentCondition represents the current weather condition of a city.
var CurrentCondition string

// CurrentLocation represents a city who's current weather is being forecasted.
var CurrentLocation string

// Forecast returns a string value that represents the current weather condition of a specific city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
