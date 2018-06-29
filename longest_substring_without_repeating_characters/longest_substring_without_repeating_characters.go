package longest_substring_without_repeating_characters

// Given a string, find the length of the longest
// substring without repeating characters.
func lengthOfLongestSubstring(s string) int {

	// Brute
	n := len(s)
	ans := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			if allUnique(s, i, j) {
				ans = Max(ans, j-i)
			}
		}
	}

	return ans

}

func allUnique(s string, start int, end int) bool {
	m := make(map[byte]int)
	for i := start; i < end; i++ {
		if _, ok := m[s[i]]; ok {
			return false
		}
		m[s[i]] = i

	}

	return true

}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func lengthOfLongestSubstring2(s string) int {
	bs := []byte(s)
	var bsMaxlen, start int
	bucket := map[byte]int{}
	for i := 0; i < len(bs); i++ {
		if _, ok := bucket[bs[i]]; ok && start <= bucket[bs[i]] {
			start = bucket[bs[i]] + 1
		} else {
			if bsMaxlen < i-start+1 {
				bsMaxlen = i - start + 1
			}
		}
		bucket[bs[i]] = i
	}
	return bsMaxlen
}
