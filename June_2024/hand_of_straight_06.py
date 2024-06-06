def isNStraightHand(hand: list[int], groupSize: int) -> bool:
    handLength = len(hand)

    if handLength % groupSize != 0:
        return False

    counter = {}
    sortedUniqueElementsInHand = list(set(hand))
    sortedUniqueElementsInHand.sort()
    for ele in hand:
        if ele in counter:
            counter[ele] += 1
        else:
            counter[ele] = 1

    for i in sortedUniqueElementsInHand:
        if counter[i] > 0:
            count = counter[i]
            for j in range(groupSize):
                if i+j not in counter or counter[i + j] < count:
                    return False
                counter[i + j] -= 1

    return True


print(isNStraightHand([1, 2, 3, 2, 3, 4, 6, 7, 8], 3))
