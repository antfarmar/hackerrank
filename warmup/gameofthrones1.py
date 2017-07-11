from sys import stdin
#file = stdin
file = open( r".\data\gameofthrones1.txt" ) #stdin
data = file.read().strip()
file.close()

oddCnt = 0

for char in range(ord("a"),ord("z")+1):
    if (data.count(chr(char)) % 2 != 0):
        oddCnt += 1

if (oddCnt <= 1):
    print("YES")
else:
    print("NO")
