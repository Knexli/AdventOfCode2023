import numpy as np

file = open('input.txt')
lines = file.readlines()

times = lines[0].split(': ')[1].split(' ')
goals = lines[1].split(': ')[1].split(' ')

# part one
while '' in times:
    times.remove('')
while '' in goals:
    goals.remove('')

A = np.empty((len(times), 2))
for idx in range(len(times)):
    v = np.array([times[idx].rstrip(), goals[idx].rstrip()])
    A[idx] = v

margin = np.empty((len(A), 1))
for idx, v in enumerate(A):
    winCntr = 0
    for acceleration in range(int(v[0]) + 1):
        distance = acceleration * (int(v[0]) - acceleration)
        if distance > int(v[1]):
            winCntr += 1
    margin[idx] = winCntr
print(np.prod(margin))

# part two
mergedTime = ''
mergedGoal = ''
for timeSnippet in times:
    mergedTime += timeSnippet.rstrip()
for goalSnippet in goals:
    mergedGoal += goalSnippet.rstrip()

v = np.array([mergedTime, mergedGoal])

winCntr = 0
for acceleration in range(int(v[0]) + 1):
    distance = acceleration * (int(v[0]) - acceleration)
    if distance > int(v[1]):
        winCntr += 1
print(np.prod(winCntr))