from sys import stdin
#file = stdin
file = open( r".\data\fillingjars.txt" ) #stdin
data = file.read().strip().split()
file.close()

numJars = int(data[0])
numOps  = int(data[1])
theOps  = list(map(int, data[2:]))

totalCandies = 0

for i in range(0, len(theOps), 3):
	a,b,k = theOps[i:i+3]
	totalCandies += (b - a + 1) * k
	print(a,b,k)

print(theOps)
print(totalCandies//numJars)
