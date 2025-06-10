package mergesortedarray

func merge(nums1 []int, m int, nums2 []int, n int)  {
    idx1, idx2 := m-1, n-1

    k := m + n -1
    for idx1 >= 0 && idx2 >= 0 {
        if nums1[idx1] >= nums2[idx2] {
            nums1[k] = nums1[idx1]
            idx1--
        } else {
            nums1[k] = nums2[idx2]
            idx2--
        }
        k--
    }

    if idx2 >= 0 {
        copy(nums1[:k+1], nums2[:idx2+1])
    }
} 
