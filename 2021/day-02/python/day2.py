import sys

f = open(sys.argv[1], "r")
game = int(sys.argv[2])
horizontal = 0
depth = 0
aim = 0

for x in f.read().split("\n"):
    ins = x.split(" ")
    if ins[0] == "forward":
        horizontal += int(ins[1])
        if game == 2:
            depth += aim * int(ins[1])
    elif ins[0] == "up":
        if game == 1:
            depth -= int(ins[1])
        else:
            aim -= int(ins[1])
    elif ins[0] == "down":
        if game == 1:
            depth += int(ins[1])
        else:
            aim += int(ins[1])

print(depth * horizontal)
