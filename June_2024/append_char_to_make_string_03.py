def appendCharacters(s: str, t: str) -> int:
    searchIndex = 0
    tLength = len(t)
    sLength = len(s)

    i = 0
    while i < tLength and searchIndex < sLength:
        if t[i] == s[searchIndex]:
            i += 1
        searchIndex += 1

    return tLength-i


print(appendCharacters('Hei', 'Hello'))
