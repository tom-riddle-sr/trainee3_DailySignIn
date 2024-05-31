package weighted_random

import (
	"math/rand"
)

type (
	WeightedRandom struct {
		Object interface{}
		Weight int32
	}

	WeightedRandomList []WeightedRandom
)

// 透過 WeightedRandomList 產生隨機數
func (wrl WeightedRandomList) Gen() int32 {
	var totalSlice []int32
	for index, item := range wrl {
		var numberSlice []int32

		for i := int32(0); i < item.Weight; i++ {
			numberSlice = append(numberSlice, int32(index))
		}
		// ...將 numberSlice 的元素展開後加入 totalSlice
		totalSlice = append(totalSlice, numberSlice...)
	}
	// rand.Intn 會回傳一個介於 0 到 len(totalSlice) 之間的隨機數
	// 如果要不是0開始，可以使用 rand.Intn(len(totalSlice) - 1) + 1
	randomIndex := rand.Intn(len(totalSlice))
	randomElement := totalSlice[randomIndex]

	return randomElement

}
