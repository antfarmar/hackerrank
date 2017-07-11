#import sys
#file = sys.stdin
file = open( r".\data\listcomprehensions.txt" )
data = file.read().strip().split()

#x,y,z,n = input(), input(), input(), input()
x,y,z = map(eval, map(''.join, (zip(data[:3], ['+1']*3))))
n = int(data[3])
print(x,y,z,n)

coords = [[a,b,c] for a in range(x) for b in range(y) for c in range(z) if a+b+c != n]
print(coords)



