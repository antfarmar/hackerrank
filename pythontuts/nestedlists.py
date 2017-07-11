#import sys
#file = sys.stdin
file = open( r".\data\nestedlists.txt" )
data = file.read().strip().split()[1:]

records = [ [data[i], float(data[i+1])] for i in range(0, len(data), 2) ]
print(records)
low = min([r[1] for r in records])
dif = min([r[1] - low for r in records if r[1] != low])
print(dif)
names = [ r[0] for r in records if r[1]-dif == low]
[print(name) for name in sorted(names)]


#from decimal import Decimal
#from itertools import groupby, islice
#from operator import itemgetter

#a = []
#for i in range(int(input())):
#  x, y = (input(), Decimal(input()))
#  a.append((y, x))
#a.sort()
#for k, v in islice(groupby(a, key=itemgetter(0)), 1, 2):
#  for x in v:
#    print(x[1])
