def heightChecker(heights: list[int]) -> int:
    expected = sorted(heights)
    noOfIncrctPos = 0
    rowLength = len(heights)

    for i in range(rowLength):
        if heights[i] != expected[i]:
            noOfIncrctPos += 1

    return noOfIncrctPos


print(heightChecker([1, 1, 2, 1, 3, 4, 5]))
