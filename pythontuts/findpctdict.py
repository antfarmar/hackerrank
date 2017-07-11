#import sys
#file = sys.stdin
file = open( r".\data\findpctdict.txt" )

data = []
for line in file:
	data.append(line.strip().split())

query = data[len(data)-1][0]
data  = data[1:len(data)-1]
dct = dict()

for d in data:
	d[1:] = map(float, d[1:])
	dct[d[0]] = d[1:]

print(data, query, dct, sep='\n')

print('{:.2f}'.format(sum(dct[query])/3))



#from sys import stdin, stdout
#N = int(stdin.readline())
#M = {}
#for i in range(N):
#    lst = stdin.readline().split()
#    M[lst[0]] = list(map(float, lst[1:]))
#Q = stdin.readline()
#stdout.write('{0:.2f}'.format(sum(M[Q])/len(M[Q])))
