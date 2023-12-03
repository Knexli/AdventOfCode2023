import numpy as np


def findGear(B, pos):
    for vi in range(B.T[0].size):
        vpos = B[vi][0].split(',')
        if vpos[0] == pos[0] and vpos[1] == pos[1]:
            return vi


file = open('input.txt')
lines = file.readlines()
A = np.asarray([list(line) for line in map(lambda s: s.strip(), lines)])
period = '.'
gear = '*'
Sum = 0
GearSum = 0
n = None
gearPos = None
adj = False
B = None

for i in range(A[0].size):
    for j in range(A.T[0].size):
        if np.char.isnumeric(A[i][j]):
            if n is None:
                n = A[i][j]
            else:
                n = n + A[i][j]
            for deltaI in range(-1, 2):
                for deltaJ in range(-1, 2):
                    if 0 < i + deltaI < A[0].size - 1 and 0 < j + deltaJ < A.T[0].size - 1:
                        if A[i + deltaI][j + deltaJ] != period and not np.char.isnumeric(A[i + deltaI][j + deltaJ]):
                            adj = True
                            if A[i + deltaI][j + deltaJ] == gear:
                                gearPos = (str(i + deltaI), str(j + deltaJ))
                            break
                if adj:
                    break
        else:
            n = None
        if (j < A[0].size - 1 and not np.char.isnumeric(A[i][j + 1]) and adj) or (j == A[0].size - 1 and adj):
            Sum += int(n)
            # part two
            if gearPos is not None:
                if B is None:
                    v = [gearPos[0] + ',' + gearPos[1], n, None]
                    B = np.array([v])
                else:
                    vi = findGear(B, gearPos)
                    if vi is None:
                        v = [gearPos[0] + ',' + gearPos[1], n, None]
                        B = np.vstack((B, v))
                    elif B[vi][2] is None:
                        v = [gearPos[0] + ',' + gearPos[1], B[vi][1], n]
                        B = np.vstack((B, v))
                        B = np.delete(B, vi, 0)
                    else:
                        B = np.delete(B, vi, 0)
            n = None
            adj = False
            gearPos = None

for row in B:
    if row[2] is not None:
        GearSum += int(row[1]) * int(row[2])

print(Sum)
print(GearSum)
