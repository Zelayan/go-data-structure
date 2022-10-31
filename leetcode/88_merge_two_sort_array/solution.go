package leetcode

func merge(nums1 []int, m int, nums2 []int, n int) []int {
    sort := make([]int, 0, m+n)
    i, j := 0, 0
    for {
        if i == m {
            sort = append(sort, nums2[j:]...)
            break
        }
        if j == n {
            sort = append(sort,nums1[i:]...)
			break
        }
        if nums1[i] < nums2[j] {
            sort = append(sort, nums1[i])
            i++
        } else {
            sort = append(sort, nums2[j])
            j++
        }
    }
	return sort
}