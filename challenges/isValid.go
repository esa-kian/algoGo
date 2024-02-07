// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

// An input string is valid if:

// Open brackets must be closed by the same type of brackets.
// Open brackets must be closed in the correct order.
// Every close bracket has a corresponding open bracket of the same type.

func isValid(s string) bool {
    // first basic condition is the string should be even
    if len(s)%2 == 1 || len(s) == 0 {
        return false
    }

    parentheses := map[rune]rune{
        '(': ')',
        '[': ']',
        '{': '}',
    }

    stack := []rune{}

    for _, p := range s {
        if _, open := parentheses[p]; open {
            stack = append(stack, p)
        } else if len(stack) == 0 || parentheses[stack[len(stack)-1]] != p {
            return false
        } else {
            stack = stack[:len(stack)-1]
        }
    }
    
    return len(stack) == 0


}
