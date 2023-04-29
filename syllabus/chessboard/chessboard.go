package chessboard

// Declare a type named Rank which stores if a square is occupied by a piece - this will be a slice of bools
type Rank []bool

// Declare a type named Chessboard which contains a map of eight Ranks, accessed with keys from "A" to "H"
type Chessboard map[string]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank string) int {

	rankRow := cb[rank]

	count := 0

	for _, value := range rankRow {
		if value {
			count++
		}
	}

	return count

}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) int {

	count := 0
	if file < 1 && file > 8 {
		return count
	}

	// iterating through a map is random
	for _, cbValue := range cb {

		for fileIdex, fileValue := range cbValue {

			if fileIdex == file-1 && fileValue {
				count++

			}
		}
	}

	return count

}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) int {

	count := 0
	for _, cbValue := range cb {

		for range cbValue {
			count++
		}
	}

	return count
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) int {

	count := 0
	for _, cbValue := range cb {

		for _, fileValue := range cbValue {

			if fileValue {
				count++
			}

		}
	}

	return count
}
