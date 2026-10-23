import (
	"slices"
)

type MedianFinder struct {
	nums []int
}

func Constructor() MedianFinder {
	return MedianFinder{}
}

func (mf *MedianFinder) AddNum(num int) {
	mf.nums = append(mf.nums, num)
	slices.Sort(mf.nums)
}

func (mf MedianFinder) FindMedian() float64 {
	length := len(mf.nums)
	if length == 0 {
		return 0
	}

	if length%2 == 0 {
		middle := float64((length - 1)) / float64(2)
		middleLeftIdx, middleRightIdx := int(math.Floor(middle)), int(math.Ceil(middle))
		middleLeft, middleRight := mf.nums[middleLeftIdx], mf.nums[middleRightIdx]
		return float64(middleLeft+middleRight) / 2.0
	} else {
		return float64(mf.nums[length/2])
	}
}
