package util

// ContainsInt64 returns true if contains a needle in a heystack.
func ContainsInt64(h []int64, n int64) bool {
	for _, v := range h {
		if v == n {
			return true
		}
	}
	return false
}
