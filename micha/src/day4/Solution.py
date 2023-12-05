file = open('input.txt')
games = file.readlines()

# part one
points = 0
for game in games:
    p = 0.5
    winning = game.split(': ')[1].split('| ')[0].rstrip().split(' ')
    effective = game.split(': ')[1].split('| ')[1].rstrip().split(' ')
    for n in effective:
        if n in [x for x in winning if x]:
            p = 2 * p
    if p != 0.5:
        points += p
print(points)

# part two
copyMap = None
for i in range(len(games)):
    if copyMap is None:
        copyMap = {i: 1}
    else:
        copyMap.update({i: 1})
for i in range(len(games)):
    matching = 0
    winning = games[i].split(': ')[1].split('| ')[0].rstrip().split(' ')
    effective = games[i].split(': ')[1].split('| ')[1].rstrip().split(' ')
    for n in effective:
        if n in [x for x in winning if x]:
            matching += 1
    instances = copyMap.get(i)
    for m in range(matching):
        copyMap.update({(i + m + 1): copyMap.get(i + m + 1) + instances})
SumCards = 0
for el in copyMap.values():
    SumCards += el
print(SumCards)
