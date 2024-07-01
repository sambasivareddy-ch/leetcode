def threeConsecutiveOdds(arr):
    n = len(arr)
    hasThreeConsecutiveOdds = False

    if n <= 2:
        return False

    for i in range(n-2):
        if arr[i]%2 and arr[i+1]%2 and arr[i+2]%2:
            hasThreeConsecutiveOdds = True 
            break 
    
    return hasThreeConsecutiveOdds

print(threeConsecutiveOdds([1,2,5,6]))