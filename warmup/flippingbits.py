from sys import stdin
#import ctypes
#file = stdin
file = open( r".\data\flippingbits.txt" ) #stdin
file.readline() # skip N

# 2's-complement trick
def flipbits(n):
   return  (1<<32) + ~n

for line in file:
	a = int(line.strip())
	#print(bin(a))
	print(flipbits(a))

#print(2<<132)
file.close()
