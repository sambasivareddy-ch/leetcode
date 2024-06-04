def longestPalindrome(s: str) -> int:
    def longest(start, end):
        if s[start:end] == s[start:end:-1]:
            return (end-start)
        return longest(start+1, end)
        return longest(start, end-1)

    longestLength = longest(0, len(s))
    return longestLength


print(longestPalindrome('hello'))
