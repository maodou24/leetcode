package solution

func lengthOfLongestSubstring(s string) int {
    var ans int
    
    for i := 0; i < len(s); i++ {
        m := make(map[byte]struct{}, 0)

        j := i
        for j < len(s){
            _, ok := m[s[j]]
            if ok {
                break
            }
            m[s[j]] = struct{}{}
            j++
        }
        ans = max(ans, j-i)
    }

    return ans
}