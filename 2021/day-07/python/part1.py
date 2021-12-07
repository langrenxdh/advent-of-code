import sys
import numpy as np

positions = open("../puzzle.dat", "r").read().split(",")
positions = [int(i) for i in positions]
positions.sort()
avg = np.full(len(positions), positions[len(positions)/2], dtype=int)
print(sum(abs(np.subtract(positions, avg))))