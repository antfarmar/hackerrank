from sys import stdin
#file = stdin
file = open( r".\data\loveletter.txt" ) #stdin
file.readline() # skip N

cases = []
for line in file:
	cases.append(line.strip())

file.close()

def minOpsPalindrome(s):
	cnt = 0
	for i in range(len(s)//2):
		cnt += abs(ord(s[i]) - ord(s[-(i+1)]))
	return cnt

for string in cases:
	print(minOpsPalindrome(string))


print(*[minOpsPalindrome(s) for s in cases], sep='\n')

