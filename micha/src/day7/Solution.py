import numpy as np

file = open('input.txt')
lines = file.readlines()

handMap = {
    'fiveOfKind': 1,
    'fourOfKind': 2,
    'fullHouse': 3,
    'threeOfKind': 4,
    'twoPairs': 5,
    'onePair': 6,
    'none': 7,
}

cardMapPartOne = {
    'A': 1,
    'K': 2,
    'Q': 3,
    'J': 4,
    'T': 5,
    '9': 6,
    '8': 7,
    '7': 8,
    '6': 9,
    '5': 10,
    '4': 11,
    '3': 12,
    '2': 13
}

cardMapPartTwo = {
    'A': 1,
    'K': 2,
    'Q': 3,
    'T': 5,
    '9': 6,
    '8': 7,
    '7': 8,
    '6': 9,
    '5': 10,
    '4': 11,
    '3': 12,
    '2': 13,
    'J': 14,
}


def isFullHouse(s):
    return (s[0] == s[1] == s[2] and s[3] == s[4]) or (s[0] == s[1] and s[2] == s[3] == s[4])


def hasFiveOfKind(s):
    return s[0] == s[1] == s[2] == s[3] == s[4]


def hasFourOfKind(s):
    return (s[0] == s[1] == s[2] == s[3]) or (s[1] == s[2] == s[3] == s[4])


def hasThreeOfKind(s):
    return (s[0] == s[1] == s[2]) or (s[1] == s[2] == s[3]) or (s[2] == s[3] == s[4])


def detHandPartOne(s):
    s = ''.join(sorted(s))

    pairAmt = 0
    for idx, c in enumerate(s):
        if idx > 0 and c == s[idx - 1]:
            pairAmt += 1

    if hasFiveOfKind(s):
        return handMap.get('fiveOfKind')

    if hasFourOfKind(s):
        return handMap.get('fourOfKind')

    if isFullHouse(s):
        return handMap.get('fullHouse')

    if hasThreeOfKind(s):
        return handMap.get('threeOfKind')

    if pairAmt == 2:
        return handMap.get('twoPairs')

    if pairAmt == 1:
        return handMap.get('onePair')

    return np.inf


def detHandPartTwo(s):
    s = ''.join(sorted(s))

    jokerAmt = 0
    for c in s:
        if c == 'J':
            jokerAmt += 1

    pairAmt = 0
    for idx, c in enumerate(s):
        if (not hasFiveOfKind(s) and not hasFourOfKind(s) and not isFullHouse(s) and not hasThreeOfKind(s)
                and idx > 0 and c == s[idx - 1]):
            pairAmt += 1

    if hasFiveOfKind(s) or (hasFourOfKind(s) and jokerAmt > 0) or (hasThreeOfKind(s) and jokerAmt == 2) or (
            isFullHouse(s) and jokerAmt == 3):
        return handMap.get('fiveOfKind')

    if hasFourOfKind(s) or (hasThreeOfKind(s) and jokerAmt == 1) or (pairAmt == 2 and jokerAmt == 2) or jokerAmt == 3:
        return handMap.get('fourOfKind')

    if isFullHouse(s) or (pairAmt == 2 and jokerAmt > 0):
        return handMap.get('fullHouse')

    if hasThreeOfKind(s) or (pairAmt == 1 and jokerAmt > 0) or jokerAmt > 1:
        return handMap.get('threeOfKind')

    if pairAmt == 2:
        return handMap.get('twoPairs')

    if pairAmt == 1 or s.__contains__('J'):
        return handMap.get('onePair')

    return handMap.get('none')


def insert(ranking, hand, multiplier, cardMap):
    for idx in range(len(ranking)):
        if detHandPartTwo(hand) > detHandPartTwo(ranking[idx][0]):
        # if detHandPartOne(hand) > detHandPartOne(ranking[idx][0]):
            ranking.insert(idx, (hand, multiplier))
            return ranking
        elif detHandPartTwo(hand) == detHandPartTwo(ranking[idx][0]):
        # elif detHandPartOne(hand) == detHandPartOne(ranking[idx][0]):
            elChars = list(ranking[idx][0])
            for charIdx, c in enumerate(list(hand)):
                if cardMap.get(c) > cardMap.get(elChars[charIdx]):
                    ranking.insert(idx, (hand, multiplier))
                    return ranking
                if cardMap.get(c) < cardMap.get(elChars[charIdx]):
                    break
    ranking.append((hand, multiplier))
    return ranking


ranking = None
for line in lines:
    sLine = line.rstrip().split(' ')
    hand = sLine[0]
    if ranking is None:
        ranking = [(hand, int(sLine[1]))]
    else:
        ranking = insert(ranking, hand, int(sLine[1]), cardMapPartTwo)
        # ranking = insert(ranking, hand, int(sLine[1]), cardMapPartOne)

Sum = 0
for idx, ranking in enumerate(ranking):
    Sum += (idx + 1) * ranking[1]
print(Sum)
