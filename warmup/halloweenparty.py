from sys  import stdin

file = open( r".\data\halloweenparty.txt" )
data = list(map(int, file.read().strip().split()))[1:]
file.close()

print(*[k*k//4 for k in data], sep="\n")

