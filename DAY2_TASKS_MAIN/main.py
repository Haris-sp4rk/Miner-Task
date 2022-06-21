import sys
sys.path.insert(0, 'C:/Users/SHAKIR/Desktop/team/haris/Day2')
sys.path.insert(0, 'C:/Users/SHAKIR/Desktop/team/kashish/day2')
sys.path.insert(0, 'C:/Users/SHAKIR/Desktop/team/Osama/Day2')
sys.path.insert(0, 'C:/Users/SHAKIR/Desktop/team/Tatheer/Day2')
sys.path.insert(0, 'C:/Users/SHAKIR/Desktop/team/Zohaib/Day 2')

from trignometricidentity import trignometricidentity
from taylor import taylorExpansion
from binomial_coefficient import binomial
from loanPayment import loanPayment
from gaussain import gaussain

print("Select from the options what u want to use:")
print("1.Trignometric Identities ")
print("2.Taylor Expansion")
print("3.Binomial Coefficient")
print("4.Loan Payment")
print("5.Gaussian")

try:
  x=int(input("Select Any:"))
except:
  print("An exception occurred")

if(x==1):
    trignometricidentity()
elif(x==2):
    taylorExpansion()
elif(x==3):
    binomial()
elif(x==4):
    loanPayment()
elif(x==5):
    gaussain()
else:
    print("Wrong Input")

