package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	successRatePercentage := successRate / float64(100)
	workingCarsPerHour := float64(productionRate) * successRatePercentage

	return workingCarsPerHour
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	workingCarsPerHour := CalculateWorkingCarsPerHour(productionRate, successRate)
	minutesInOneHour := 60
	workingCarsPerMinute := int(workingCarsPerHour) / minutesInOneHour

	return workingCarsPerMinute
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	costPerCar := 10_000
	costPerBatchOfTen := 95_000

	carsPerGroup := 10

	numFullBatches := carsCount / carsPerGroup
	remainingCars := carsCount % carsPerGroup

	totalBatchCost := numFullBatches * costPerBatchOfTen
	totalRemainingCarsCost := remainingCars * costPerCar

	totalCost := totalBatchCost + totalRemainingCarsCost

	return uint(totalCost)
}
