#from sys import stdin

file = open( r".\data\manasastones.txt" )
data = list(map(int, file.read().strip().split()))[1:]
file.close()

def lastStone(numStones, diffA, diffB):
	diffs = set()
	for i in range(numStones):
		diffs.add( (numStones-i-1)*diffA + i*diffB)
	return diffs

# For each test case triple...
for i in range(0, len(data), 3):
	print( *sorted( lastStone(*data[i:i+3]) ) )

