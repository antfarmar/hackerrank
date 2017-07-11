#from sys import stdin

data = open( r".\data\sherlockgcd.txt" )

def gcd(a, b):
    if b == 0:
        return a
    else:
        return gcd(b, a % b)

numCases = int(data.readline())

for _ in range(numCases):
	data.readline() # skip N
	nums = list(map(int, data.readline().strip().split()))

	g = nums[0]
	for n in nums:
		g = gcd(g, n)

	if g == 1:
		print("YES")
	else:
		print("NO")

data.close()
