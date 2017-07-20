from sys  import stdin
from math import sqrt
from math import ceil

#file = stdin
file = open( r".\data\sherlocksquares.txt" ) #stdin
data = list(map(int, file.read().strip().split()))[1:]
file.close()

def numSquares(mini, maxi):
	cnt = 0
	start = ceil(sqrt(mini))
	while start**2 <= maxi:
		cnt   += 1
		start += 1
	return cnt

# For each test case pair...
for i in range(0, len(data), 2):
	print(numSquares(*data[i:i+2]))

