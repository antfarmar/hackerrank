from sys import stdin
#file = stdin
file = open( r".\data\maximizingxor.txt" ) #stdin
L,R = list(map(int, file.read().strip().split()))
file.close()

maxXor = max(A^B for A in range(L,R+1) for B in range(L,R+1))
print(maxXor)


