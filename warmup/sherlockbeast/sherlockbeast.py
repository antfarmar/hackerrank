from sys  import stdin

#file = stdin
file = open( r".\data\sherlockbeast.txt" ) #stdin
data = list(map(int, file.read().strip().split()))[1:]
file.close()

# Maximize x in k = 3x + 5y
for k in data:
  for x in range(k//3, -1, -1):
    if (k-3 * x) % 5 == 0:
      y = (k-3 * x) // 5
      print('5'*3*x + '3'*5*y)
      break
    elif x==0:
      print(-1)

