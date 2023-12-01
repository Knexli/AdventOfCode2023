import numpy as np


def traverseCalibrationPartOne(calibration):
    calibration.rstrip()
    firstDigit = 0
    lastDigit = 0
    for i in calibration:
        if i.isnumeric():
            firstDigit = i
            break
    for i in calibration:
        if i.isnumeric():
            lastDigit = i
    return int(firstDigit + lastDigit)


def traverseCalibrationPartTwo(calibration, numbersMap):
    firstDigit = None
    lastDigit = None
    word = ''

    for i in calibration:
        word = word + i
        number = findNumber(i, numbersMap, word)
        if number is not None:
            firstDigit = number
            break
    for i in reversed(calibration):
        word = i + word
        number = findNumber(i, numbersMap, word)
        if number is not None:
            lastDigit = number
            break

    return int(firstDigit + lastDigit)


def findNumber(i, numbersMap, word):
    if i.isnumeric():
        return i
    else:
        for key in numbersMap:
            if word.__contains__(key):
                return numbersMap.get(key)


file1 = open('input.txt')
calibrations = file1.readlines()

resultPart1 = 0
for row in calibrations:
    resultPart1 += traverseCalibrationPartOne(row)
print(resultPart1)

resultPart2 = 0
numbersMap = {
    "one": "1",
    "two": "2",
    "three": "3",
    "four": "4",
    "five": "5",
    "six": "6",
    "seven": "7",
    "eight": "8",
    "nine": "9"
}
for row in calibrations:
    resultPart2 += traverseCalibrationPartTwo(row, numbersMap)
print(np.sum(resultPart2))
