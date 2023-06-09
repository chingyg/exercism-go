package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {

	totalBirds := 0

	for index := range birdsPerDay {
		totalBirds += birdsPerDay[index]
	}

	return totalBirds

}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {

	startIndex := (week - 1) * 7

	return TotalBirdCount(birdsPerDay[startIndex : week*7])

}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {

	// Increment the value of every even index
	for index := range birdsPerDay {
		if index%2 == 0 {
			birdsPerDay[index] += 1
		}
	}

	return birdsPerDay

}
