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

# Part two

similarity_score = 0
for l in left:
    similar_counter = 0
    for r in right:
        if l == r:
            similar_counter += l
    similarity_score += similar_counter

print(f"similarity score is: {similarity_score}")
