#from sys import stdin
#file = stdin
file = open(r'.\data\findsecondmax.txt', 'rt' )
data = file.read().strip().split()[1:]

nums = list(map(int, data))
nums.sort(reverse=True)
print(nums)

big = nums[0]

for n in nums[1:]:
	if n != big:
		print(n)
		break


#_ = input()
#nums = set([int(val) for val in input().split()])
#print(sorted(nums)[-2])
