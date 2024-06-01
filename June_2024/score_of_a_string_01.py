def scoreOfString(s: str) -> int:
    score = 0  # initialize to 0
    strLength = len(s)

    for i in range(0, strLength-1):
        score += abs(ord(s[i]) - ord(s[i+1]))

    return score


s = input("Enter the string:")
print("Score:", scoreOfString(s))
