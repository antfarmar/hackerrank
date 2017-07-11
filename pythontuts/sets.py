#import sys
#file = sys.stdin
file = open( r'.\data\sets.txt' )

data = []
for line in file:
	data.append(line.strip().split())

print(data)
set1 = set(data[1])
set2 = set(data[3])
print(set1, set2)
symdif = set1.symmetric_difference(set2)
symdif = list(map(int, symdif))
symdif.sort()
[print(n) for n in sorted(symdif)]
