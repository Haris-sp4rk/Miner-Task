import math 

def trignometricidentity():
    try:
        x =int(input("Enter a value to calculate trignometric identitities:"))
    except:
        print("An exception occurred")
    print("Value of sin is::",math.sin(x))
    print("Value of cos is::",math.cos(x))
    print("Value of tan is::",math.tan(x))


    