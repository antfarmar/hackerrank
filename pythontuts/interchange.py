#import sys
#file = sys.stdin
file = open( r".\data\interchange.txt" )
data = file.read().strip().split()

#tup = input(), input()
tup = tuple(data)
[print(tup[-i]) for i in range(1,3)]



# Other #######################
#a, b = input(), input()
#a, b = b, a

#for t in a,b:
#    print(t)


