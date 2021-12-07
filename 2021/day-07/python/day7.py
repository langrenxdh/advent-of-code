import sys

f = open(sys.argv[1], "r")
game = sys.argv[2]
positions = f.read().split(",")
positions = [int(i) for i in positions]
positions.sort()

fuels = []
for p1 in range(positions[-1]):
    fuel = 0
    for p in positions:
        if game == "1":
            fuel += abs(p - p1)
        else:
            n = abs(p - p1)
            fuel += (n * (n + 1)) / 2
    fuels.append(fuel)
fuels.sort()
print(fuels[0])
