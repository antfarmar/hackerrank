from sys import stdin
import re

file = stdin
file = open(r'.\data\decorators.txt', 'rt' )
data = file.read().strip().split()[1:]

print(data)

pat = r'^[\+091]*([0-9]{5})([0-9]{5})$'
exp = re.compile(pat)


def sortInc(lst):
	return sorted(lst)


def stdPhone(phoneList):
	for i in range(len(phoneList)):
		pnum = phoneList[i]
		group = exp.findall(pnum)[0]
		phoneList[i] = "+91 " + group[0] + " " + group[1]
	return phoneList


def sortAndFormatDecorator(sortFunc, fmtFunc):
	def process():
		return sortFunc(fmtFunc(data))
	return process


f = sortAndFormatDecorator(sortInc, stdPhone)
print('\n'.join(f()))

