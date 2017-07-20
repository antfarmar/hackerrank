from sys import stdin
#file = stdin
file = open( r".\data\alternatingchars.txt" ) #stdin
file.readline() # skip N

cases = []
for line in file:
	cases.append(line.strip())

file.close()

for string in cases:
	cnt = 0
	for i in range(len(string)-1):
		if string[i] == string[i+1]:
			cnt += 1
	print(cnt)

# print(len(S) - len(list(itertools.groupby(S))))
