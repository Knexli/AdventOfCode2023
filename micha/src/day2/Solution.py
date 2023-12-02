import re

import numpy as np

file = open('input.txt')
games = file.readlines()

sumId = 0
colourMap = {
    "red": "12",
    "green": "13",
    "blue": "15"
}

# part one
for game in games:
    definition = game.split(': ')[0]
    gameId = int(definition.split(' ')[1])
    breach = False
    draws = game.split(': ')[1].split('; ')
    for draw in draws:
        cube = draw.split(', ')
        for colourTuple in cube:
            amt = colourTuple.split(' ')[0]
            colour = colourTuple.split(' ')[1]
            if int(colourMap.get(colour.strip())) < int(amt):
                breach = True
                break
    if not breach:
        sumId += gameId
print(sumId)

# part two
powerSum = 0
for game in games:
    definition = game.split(': ')[0]
    draws = game.split(': ')[1].split('; ')
    A = None
    for draw in draws:
        v = np.zeros(3)
        cube = draw.split(', ')
        for colourTuple in cube:
            amt = colourTuple.split(' ')[0]
            colour = colourTuple.split(' ')[1].strip()
            if colour == 'red':
                v[0] = amt
            elif colour == 'blue':
                v[1] = amt
            elif colour == 'green':
                v[2] = amt
        if A is None:
            A = np.array(v)
        else:
            A = np.vstack((A, v))
    power = 1
    for column in A.T:
        power *= max(column)
    powerSum += power
print(powerSum)


