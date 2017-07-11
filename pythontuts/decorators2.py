from sys import stdin
from operator import itemgetter
#file = stdin
file = open(r'.\data\decorators2.txt', 'rt' )
#data = file.read().strip().split()[1:]

file.readline() # skip N
data = []
for line in file:
	rec = line.strip().split()
	rec[2] = int(rec[2])
	data.append(rec)


def sortAge(records):
	records.sort(key=itemgetter(2))


def titler(record):
	first, second = record[:2]
	sex = record[3]
	if sex == 'M':
		title = 'Mr.'
	else:
		title = 'Ms.'
	return ' '.join([title,first,second])


def showFmtRecords(fmtFunc, records):
	def process():
		sortAge(records)
		for r in records:
			print(fmtFunc(r))
	return process


showTitle = showFmtRecords(titler, data)
showTitle()


