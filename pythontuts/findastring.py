#from sys import stdin
#file = sys.stdin
file = open(r'.\data\findastring.txt', 'rt' )
data = file.read().strip().split()

print(data)
txt = data[0]
sub = data[1]
cnt = 0

rng = len(txt) - len(sub) + 1

for i in range(rng):
	if sub == txt[i : i + len(sub)]:
		cnt += 1

print(cnt)

print(sum(sub == txt[i:i+len(sub)] for i in range(rng)))
