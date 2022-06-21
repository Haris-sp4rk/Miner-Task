def taylorExpansion():
    x=int(input("Input Value of x:"))
    numOfTerms=int(input("Input number of terms:"))
    fact = 1
    sum = 0.0

    for i in range(1, numOfTerms):
        fact= fact * i
        sum = sum+(pow(x, i)/fact)

    sum = sum + 1
    print("The answer is ",sum)

    try:
        x = int(input("Enter value of x:"))
        numOfTerms = int(input("Enter number of terms: "))



        fact = 1
        sum = 0.0

        for i in range(1, numOfTerms):
            fact= fact * i
            sum = sum+(pow(x, i)/fact)

        sum = sum + 1
        return sum
    except:
        print("Invalid Input. Please try again.")


print(taylorExpansion())