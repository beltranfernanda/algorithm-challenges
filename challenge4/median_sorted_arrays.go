package challenge4

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	a := mergeSort(nums1, nums2)
	var median int
	var number float64
	mod := len(a) % 2
	if mod == 0 {
		median = len(a) / 2
		number = (float64(a[median]) + float64(a[median-1])) / 2

	} else {
		median = len(a) / 2
		number = float64(a[median])
	}
	return number

}

func mergeSort(nums1 []int, nums2 []int) []int {
	i := 0
	j := 0
	length := len(nums1) + len(nums2)
	sortedArray := make([]int, length)
	for i < len(nums1) || j < len(nums2) {
		if i == len(nums1) && j < len(nums2) {
			sortedArray[i+j] = nums2[j]
			j++
			continue
		}
		if j == len(nums2) && i < len(nums1) {
			sortedArray[i+j] = nums1[i]
			i++
			continue
		}
		if nums1[i] > nums2[j] {
			sortedArray[i+j] = nums2[j]
			j++
		} else {
			sortedArray[i+j] = nums1[i]
			i++
		}
	}
	return sortedArray
}

// Most efficient solution
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1, l2 := len(nums1), len(nums2)
	stop := (l1 + l2) / 2
	var nums []int

	if l1 == 0 {
		nums = nums2
	} else if l2 == 0 {
		nums = nums1
	} else {
		nums = make([]int, stop+1)
		for i, j, k := 0, 0, 0; i <= stop; i++ {
			if j >= l1 {
				nums[i] = nums2[k]
				k++
			} else if k >= l2 {
				nums[i] = nums1[j]
				j++
			} else if nums1[j] <= nums2[k] {
				nums[i] = nums1[j]
				j++
			} else {
				nums[i] = nums2[k]
				k++
			}
		}
	}

	if (l1+l2)%2 == 1 {
		return float64(nums[stop])
	} else {
		return (float64(nums[stop]) + float64(nums[stop-1])) * 0.5
	}
}
