/*
Given a non-negative integer x, return the square root of x rounded down to the nearest integer. The returned integer should be non-negative as well.

You must not use any built-in exponent function or operator.

For example, do not use pow(x, 0.5) in c++ or x ** 0.5 in python.
 

Example 1:

Input: x = 4
Output: 2
Explanation: The square root of 4 is 2, so we return 2.
Example 2:

Input: x = 8
Output: 2
Explanation: The square root of 8 is 2.82842..., and since we round it down to the nearest integer, 2 is returned.
 

*/

func mySqrt(x int) int {
    min := 0
    max := x + 1

    for min < max {
        avg := min + (max - min) / 2

        if avg * avg > x {
            max = avg
        } else {
            min = avg + 1
        }
    }

    return min - 1
}
