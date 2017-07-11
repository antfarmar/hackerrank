from sys import stdin
from re import compile
from re import match

file = stdin
file = open(r'.\data\validatephone.txt', 'rt' )
data = file.read().strip().split()[1:]

pat = r'^[789][0-9]{9}$'
exp = compile(pat)

def match(phonum):
	if exp.match(phonum):
		return "YES"
	else:
		return "NO"

print(*map(match, data), sep="\n")

#print("YES" if exp.match(input()) else "NO")
