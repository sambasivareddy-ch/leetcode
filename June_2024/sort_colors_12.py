def sortColors(nums: list[int]) -> None:
    index = 0
    length = len(nums)
    currentColor = 0

    while index < length:
        if nums[index] != currentColor:
            for i in range(index+1, length):
                if nums[i] == currentColor:
                    nums[i], nums[index] = nums[index], nums[i]
                    break
        index += 1
        if currentColor not in nums[index:]:
            currentColor += 1


sortColors([2, 0, 2, 1, 1, 0])
