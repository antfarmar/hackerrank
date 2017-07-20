from itertools import combinations
from sys import stdin
#file = stdin
file = open( r".\data\acmicpcteam.txt" ) #stdin
data = file.read().strip().split()
file.close()

print(data)

numPeople = int(data[0])
numTopics = int(data[1])
topicMaps   = data[2:]


def bestPair(bitMaps):
    counts = {'max': 0, 'num': 0}
    for pair in combinations(bitMaps, 2):
        topicsKnown = bin(int(pair[0], 2) | int(pair[1], 2)).count('1')
        if topicsKnown > counts['max']:
            counts['num'] = 1
            counts['max'] = topicsKnown
        elif topicsKnown == counts['max']:
            counts['num'] += 1
    return counts

ans = bestPair(topicMaps)
print(ans['max'], ans['num'], sep='\n')
