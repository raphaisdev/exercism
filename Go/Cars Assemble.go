package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
    return (float64(productionRate) / float64(100)) * successRate
	panic("CalculateWorkingCarsPerHour not implemented")
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
    return int(CalculateWorkingCarsPerHour(productionRate, successRate)) / 60
	panic("CalculateWorkingCarsPerMinute not implemented")
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
    var moduleTen int = int(carsCount / 10) * 95000
    var restTen int = int(carsCount % 10) * 10000
    return uint(moduleTen + restTen)
	panic("CalculateCost not implemented")
}
