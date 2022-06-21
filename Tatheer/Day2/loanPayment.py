import math
def loanPayment():
    try:
        p=int(input('Enter your loan amount : '))
        r=int(input('Enter interest rate per period : '))
        n=int(input('Enter total number of period: '))    
    except:
        print("An exception has occurred!!")
    X=(1+r)**(n-1)
    Y=(1+r)**n
    a=p*((r*(X))/Y)
    a="{:.2f}".format(a)
    print("Payment amount per period : " ,a)

