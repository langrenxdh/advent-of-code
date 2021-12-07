import sys
import numpy as np

positions = open("../puzzle.dat", "r").read().split(",")
positions = [int(i) for i in positions]
positions.sort()
mark=0
for i in range(positions[-1]):
    sum = i
    for ii in positions:
        sum += 2*(ii-i)
    if sum <= 0:
        break
    mark=i

avg = np.full(len(positions), mark, dtype=int)
dist = abs(np.subtract(positions, avg))
ans=0
for i in dist:
    ans+=(i*(i+1))/2
print(ans)