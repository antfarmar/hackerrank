from sys import stdin
#file = stdin
file = open( r".\data\solvemefirst.txt" ) #stdin

def solveMeFirst(a,b):
  return a+b

a = int(file.readline().strip())
b = int(file.readline().strip())

print(a,b)
print(solveMeFirst(a,b))

file.close()
