// palindrome-test.go

package reshmask

import "testing"

func TestIsPalindrome(t *testing.T) {
    testCases := []struct {
        input    string
        expected bool
    }{
        {"A man a plan a canal Panama", true},
        {"racecar", true},
        {"hello", false},
        // Add more test cases as needed
    }

    for _, tc := range testCases {
        t.Run(tc.input, func(t *testing.T) {
            actual := IsPalindrome(tc.input)
            if actual != tc.expected {
                t.Errorf("IsPalindrome(%s) = %v; want %v", tc.input, actual, tc.expected)
            }
        })
    }
}
