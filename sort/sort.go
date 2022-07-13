package sort

func quickSort(nums []int) {
	quickSortHelp(nums, 0, len(nums)-1)
}
// quick sort相当于 二叉树的前序遍历
func quickSortHelp(nums []int, lo, hi int) {
	if (lo >= hi) {
        return;
    }
    // 对 nums[lo..hi] 进行切分
    // 使得 nums[lo..p-1] <= nums[p] < nums[p+1..hi]
    p := partition(nums, lo, hi);
    // 去左右子数组进行切分
    quickSortHelp(nums, lo, p - 1);
    quickSortHelp(nums, p + 1, hi);
}

func partition(nums []int, lo, hi int) int{
	// 选择第一个作为pivot 并将其置空，这里的置空值的是将这里逻辑置空，方便找比pivot小的数搬过来
	pivot := nums[lo]
	for lo < hi {
		for lo < hi && pivot <= nums[hi] {
			hi --
			// 从后往前找，第一个比pivot小的数
		}
		// 将其搬到当前空出来的位置
		nums[lo] = nums[hi]

		for lo < hi && pivot > nums[lo] {
			lo ++
			// 从前往后找第一个大于pivot的数
		}
		// 将其搬到当前空出来的位置，nums[i] 因为已经搬走了，所以逻辑为空
		nums[hi] = nums[lo]
	}
	// 将pivot置于 正确的位置
	nums[lo] = pivot
	return lo
}