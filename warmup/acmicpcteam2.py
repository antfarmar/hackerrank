from sys import stdin
#file = stdin
file = open( r".\data\acmicpcteam.txt" ) #stdin
data = file.read().strip().split()
file.close()

print(data)

numPeople = int(data[0])
numTopics = int(data[1])
bitMaps   = data[2:]
knowList  = []

for i in range(numPeople-1):
	for j in range(i+1, numPeople):
		knowCnt = sum(int(pair[0]) or int(pair[1]) for pair in zip(bitMaps[i], bitMaps[j]))
		knowList.append(knowCnt)

maxKnown = max(knowList)
print(maxKnown, knowList.count(maxKnown), sep='\n')
