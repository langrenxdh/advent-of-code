import sys

# first para is: sliding window size; second para is: puzzle data file
step = int(sys.argv[1])
f = open(sys.argv[2], "r")
nums = []

for x in f.read().split("\n"):
    nums.append(int(x))

cnt = 0

for idx in range(1, len(nums)):
    if nums[idx - step] < nums[idx]:
        cnt += 1

print("Sliding window:", step, "^_^ Total increase:", cnt)
