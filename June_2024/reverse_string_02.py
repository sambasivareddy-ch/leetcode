def reverseString(s: list[str]) -> None:
    n = len(s)
    for i in range(n//2):
        s[i], s[n-i-1] = s[n-i-1], s[i]


s = ['H', 'E', 'L', 'L', 'O']
reverseString(s)
print(s)
