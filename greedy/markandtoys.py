with open( "data.txt" ) as f:
    data = f.read().split()

##with open( filename, "w") as f:
##    f.write("write something here ")

##import sys
##data = sys.stdin.read()
######################################################

print(data)
#d = list(map(int, data.split()))
data = [int(x) for x in data]
csh = data[1]
cst = [x for x in data[2:] if x<=csh]
#cst.sort()
print(data,csh,cst)

cnt, tot = 0,0
for c in sorted(cst):
    tot += c
    if tot < csh:
        cnt+=1
    else:
        break

print(cnt)


#prices = sorted(filter(lambda x: x <= K, map(int, sys.stdin.readline().split())))
