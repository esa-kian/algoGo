/*
Given a string s, return the longest 
palindromic
 
substring
 in s.

 

Example 1:

Input: s = "babad"
Output: "bab"
Explanation: "aba" is also a valid answer.
Example 2:

Input: s = "cbbd"
Output: "bb"
*/

func longestPalindrome(s string) string {
	length := len(s)
	if length == 0 {
		return ""
	}

	var start, end int
	for i := 0; i < length; {
		if length-i <= (end-start)/2 {
			break
		}
		l, r := i, i
		for r < length-1 && s[r] == s[r+1] {
			r++
		}
		i = r + 1
		for l > 0 && r < length-1 && s[l-1] == s[r+1] {
			l--
			r++
		}
		if r-l > end-start {
			start, end = l, r
		}
	}

	return s[start : end+1]
}
