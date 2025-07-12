package solution

func isSubsequence(s string, t string) bool {
    var idx int

    last := len(s)-1
    for i := range t {
        if idx > last {
            break
        }
        if t[i] == s[idx] {
            idx++
        }
    }
    return idx > last
}