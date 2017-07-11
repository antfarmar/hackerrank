from sys import stdin
import re

file = stdin
file = open(r'.\data\validateemail.txt', 'rt' )
data = file.read().strip().split()[1:]

print(data)

import re

pat = r'^[a-zA-Z0-9_-]+@[a-zA-Z0-9]+\.[a-zA-Z0-9]{1,3}$'
exp = re.compile(pat)

def ematch(email):
	if exp.match(email) != None:
		return True
	return False

print(sorted(list(filter(ematch, data))))


#from re import match
#l = sorted([input() for _ in range(int(input()))])
#print(list(filter(lambda i: match(r'^[a-zA-Z0-9_-]+\@[a-zA-Z0-9]+\.[a-zA-Z]{1,3}$', i), data)))
