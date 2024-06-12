def relativeSortArray(arr1: list[int], arr2: list[int]) -> list[int]:
    counter = {}
    for i in arr1:
        if i in counter:
            counter[i] += 1
        else:
            counter[i] = 1

    resultArr = []
    for i in arr2:
        resultArr.extend([i]*counter[i])
        del counter[i]

    remainArr = []
    for i, j in counter.items():
        remainArr.extend([i]*j)

    remainArr.sort()

    resultArr.extend(remainArr)

    return resultArr


print(relativeSortArray(
    [2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19], [2, 1, 4, 3, 9, 6]))
