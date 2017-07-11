from sys import stdin

#file = stdin
file = open( r".\data\isfibo.txt" ) #stdin
data = list(map(int, file.read().strip().split()))[1:]
file.close()

def fibList(high):
	fibs = [0]
	x,y = 0,1
	while fibs[-1] <= high:
		x,y = y, x+y
		fibs.append(x)
		high -=1
	return fibs

fibs = fibList(10**10)

for n in data:
	if n in fibs:
		print("isFibo")
	else:
		print("isNotFibo")



#fibonacci = [0, 1]
#while fibonacci[-1] <= 10**10:
#    fibonacci.append(fibonacci[-1] + fibonacci[-2])
