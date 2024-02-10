/*
Given an input string s and a pattern p, implement regular expression matching with support for '.' and '*' where:

'.' Matches any single character.​​​​
'*' Matches zero or more of the preceding element.
The matching should cover the entire input string (not partial).
*/

func isMatch(s string, p string) bool {
	slen, plen := len(s), len(p)
	dp := make([][]bool, slen+1)
	for i := range dp {
		dp[i] = make([]bool, plen+1)
	}

	for i := 0; i <= slen; i++ {
		for j := 0; j <= plen; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = true
			} else if j > 1 && p[j-1] == '*' {
				dp[i][j] = dp[i][j-2] || (i > 0 && (p[j-2] == '.' || p[j-2] == s[i-1]) && dp[i-1][j])
			} else if i > 0 && j > 0 && (p[j-1] == '.' || p[j-1] == s[i-1]) {
				dp[i][j] = dp[i-1][j-1]
			}
		}
	}
	return dp[slen][plen]
}
