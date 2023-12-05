import numpy as np
import time

file = open('testInput.txt')
lines = file.readlines()
defiPartOne = lines[0].split(': ')[1].rstrip().split(' ')
defiPartTwo = None

for idx, seed in enumerate(defiPartOne):
    if idx % 2 == 0:
        quant = int(defiPartOne[idx + 1])
        if defiPartTwo is None:
            defiPartTwo = [*range(int(seed), int(seed) + int(defiPartOne[idx + 1]), 1)]
        else:
            defiPartTwo = np.append(defiPartTwo, range(int(seed), int(seed) + int(defiPartOne[idx + 1]), 1))
lines.pop(0)

almanac = None
key = None
A = None

for idx, s in enumerate(lines):
    if s != '\n':
        if s.__contains__('map'):
            key = s.split(' ')[0]
            if almanac is None:
                almanac = {key: None}
            else:
                almanac.update({key: None})
        else:
            r = int(s.split(' ')[2].rstrip()) - 1
            v0 = (int(s.split(' ')[0]), int(s.split(' ')[1]))
            v1 = (v0[0] + r, v0[1] + r)
            if A is None:
                A = np.array(v0)
                A = np.vstack((A, v1))
            else:
                A = np.vstack((A, v0))
                A = np.vstack((A, v1))
            almanac.update({key: A})
    else:
        A = None


def evalLocation(defi):
    locationList = np.empty((len(defi)))
    start = time.time()
    for idx, s in enumerate(defi):
        curr = int(s)
        for src in almanac.keys():
            lower = np.min(almanac.get(src)[:, 0])
            upper = np.max(almanac.get(src)[:, 0])

            ref = -1
            refTuple = None
            if lower <= curr <= upper:
                for val in almanac.get(src):
                    if ref <= val[1] <= curr:
                        ref, refTuple = val[1], val
                curr = curr + refTuple[0] - refTuple[1]
            else:
                curr = curr
        locationList[idx] = curr
        if idx == 1000000:
            end = time.time()
            print('1 mio took', end - start)
    return min(locationList)


print(evalLocation(defiPartOne))
print(evalLocation(defiPartTwo))
