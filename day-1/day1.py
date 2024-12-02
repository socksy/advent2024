#!/usr/bin/env python3

# Part one
left = []
right = []
with open("input") as f:
    for line in f:
        l,r = line.split()
        left.append(int(l))
        right.append(int(r))

left.sort()
right.sort()

sum = 0
i = 0
for i in range(len(max(left, right))):
    sum += abs(left[i] - right[i])

print(f"part one sum of differences is: {sum}")

#
