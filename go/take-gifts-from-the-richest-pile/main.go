package take_gifts_from_the_richest_pile

import "math"

func pickGifts(gifts []int, k int) int64 {
	for k != 0 {
		max := gifts[0]
		maxIndex := 0
		for i, gift := range gifts {
			if gift > max {
				max = gift
				maxIndex = i
			}
		}
		gifts[maxIndex] = int(math.Sqrt(float64(max)))
		k -= 1
		if max == 1 {
			return int64(len(gifts))
		}
	}
	var sum int64 = 0
	for _, gift := range gifts {
		sum += int64(gift)
	}
	return sum
}
