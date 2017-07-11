from sys import stdin
from re import compile
from re import match

file = stdin
file = open(r'.\data\validateroman.txt', 'rt' )
data = file.read().strip().split()

pat = r'^M{0,3}(CM|CD|D?C{0,3})(XC|XL|L?X{0,3})(IX|IV|V?I{0,3})$'
exp = compile(pat)

print(*[exp.match(num) != None for num in data], sep='\n')
