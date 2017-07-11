from sys import stdin
#file = stdin
file = open( r".\data\chocolatefeast.txt" ) #stdin
data = list(map(int, file.read().strip().split()))[1:]
file.close()

def numChocs(cash, cost, disc):
	total = cash // cost
	wrappers = total

	while wrappers >= disc:
		wrappers += 1 - disc
		total += 1

	return total

for i in range(0, len(data), 3):
	print(numChocs(*data[i:i+3]))

