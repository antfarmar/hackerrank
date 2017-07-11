#from sys import stdin
#file = stdin
file = open(r'.\data\maplambda.txt', 'rt' )
data = int(file.read().strip())

def fib(n):
	result = []
	a, b = 0, 1
	while n > 0:
		result.append(a)
		a, b = b, a+b
		n -= 1
	return result

print(list(map(lambda x: x**3, fib(data))))
#print(*map(lambda x: x**3, fib(data)), sep="\n")
