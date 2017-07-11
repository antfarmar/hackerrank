import sys
N = None
K = None
prices = []

for line in sys.stdin:
    if N is None:
        N, K = [int(i) for i in line.strip().split()]
        continue
    prices = [int(i) for i in line.strip().split()]

prices.sort()
total_price = 0
count = 0
for price in prices:
    total_price += price
    if total_price >= K:
        print(count)
        break
    else:
        count += 1
