package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	var totalBirdCount int
	for i := 0; i < len(birdsPerDay); i++ {
		totalBirdCount += birdsPerDay[i]
	}
	return totalBirdCount
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	birdsPerWeek := 0
	for i := 7 * (week - 1); (i < week*7) && (i < len(birdsPerDay)); i++ {
		birdsPerWeek += birdsPerDay[i]
	}
	return birdsPerWeek
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i += 2 {
		birdsPerDay[i] += 1
	}
	return birdsPerDay
}
