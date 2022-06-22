def binomial():
    from math import factorial as fac

    print("\nBinomial Coefficient Calculator. \n\nSample Values: 8,3 = 56\n")
    
    try:
        x = int(input("Enter First Integer: "))
        y = int(input("Enter Second Integer: "))
    except:
        print("Input type is not an integer or a -ve value")
        
    print('\nBinomial Coefficient:', fac(x) // fac(y) // fac(x - y))



