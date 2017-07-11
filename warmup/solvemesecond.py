from sys import stdin
#file = stdin
file = open( r".\data\solvemesecond.txt" ) #stdin
file.readline() # skip N

def solveMeSecond(a,b):
   return a+b

for line in file:
	a,b = map(int, line.strip().split())
	print(a,b)
	print(solveMeSecond(a,b))

file.close()
