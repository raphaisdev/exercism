package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
    totalBirds := 0
    for _, birdCount := range birdsPerDay {
        totalBirds += birdCount
    }
    return totalBirds
	panic("Please implement the TotalBirdCount() function")
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
    return TotalBirdCount(birdsPerDay[((week-1)*7):(week*7)])
	panic("Please implement the BirdsInWeek() function")
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
    for k, v := range birdsPerDay {
        if k%2 != 0 {
            continue
        }
        birdsPerDay[k] = v+1
    }
    return birdsPerDay
	panic("Please implement the FixBirdCountLog() function")
}
