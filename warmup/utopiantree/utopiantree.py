from sys import stdin
#file = stdin
file = open( r".\data\utopiantree.txt" ) #stdin
file.readline() # skip N

def growth(cycles):
	height = 1
	for n in range(1,cycles+1):
		if n % 2 == 0:
			height += 1
		else:
			height *= 2
	return height

for line in file:
	n = int(line.strip())
	print(growth(n))

file.close()
