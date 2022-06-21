def taylorExpansion(x, numOfTerms):

    fact = 1
    sum = 0.0

    for i in range(1, numOfTerms):
        fact= fact * i
        sum = sum+(pow(x, i)/fact)

    sum = sum + 1
    return sum


# print(taylorExpansion(2, 10))