from sys import stdin
#file = stdin
file = open( r".\data\angrychildren.txt" ) #stdin
data = list(map(int, file.read().strip().split()))
file.close()

n    = data[0]
k    = data[1]
nums = sorted(data[2:])


print(nums)
print(min(nums[i+k-1] - nums[i] for i in range(0, n-k-1)))

# Timeout > 10s
#print(min( [max(nums[i:i+k]) - min(nums[i:i+k]) for i in range(len(nums)-k)] ))


