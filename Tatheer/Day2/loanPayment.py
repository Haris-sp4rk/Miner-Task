import math
def loanPayment():
    p=int(input('Enter your loan amount : '))
    r=int(input('Enter interest rate per period : '))
    n=int(input('Enter total number of period: '))
    X=(1+r)**(n-1)
    Y=(1+r)**n
    a=p*(r*(Y))/X
    print("Payment amount per period : " ,a)
loanPayment()