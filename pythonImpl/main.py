from Menu import Menu
from Array import Array
import re
import datetime


def main():
    M = Menu()
    M.printMenu()
    r = input()
    method = choices(r)
    while method != 4:
        method()
        M.printMenu()
        r = input()
        method = choices(r)

def choices(c):
    return {'1' : choice1,
            '2' : choice2,
            '3' : choice3, 
    }.get(c, 4)


def choice1():
    given = input("Input elements of your array followed by ',' i.e. : 1,2,...\n:")
    g = re.split(',',given)
    a = []
    for string in g:
        a.append(int(string))
    A = Array(a)
    # print(A.ar)
    print("Solution 1: ", A.solution1())
    print("Solution 2: ", A.solution2())
    l = 0
    r = len(A.ar)-1
    print("Solution 3: ",A.solution3(l,r))
    print("Solution 4: ",A.solution4())

def choice2():
    given = input("Enter the size of the array to test\n:")
    A = Array([])
    A.newAr(int(given))
    

    # python3 implimentation
    given = str(input("Enter each Method to test i.e.(1234)\n:"))
    for i in range(0,len(given)):
        method = measures(given[i])
        method(A)

        
def choice3():
    given = input("Enter the Method # to test\n:")
    methodNum = int(given)
    method = measures(given)
    given = input("Enter the size of the array to test\n:")
    A = Array([])
    m = int(given)
    A.newAr(m)
    given = input("Enter the size of the array to estimate\n:")
    n = int(given)
    C = method(A)
    umm = switcher(A, methodNum)
    expected = umm(C, m, n)
    print("The expected time for Method #",methodNum," is ", expected, "(µs)\n")
    A.newAr(n)
    method(A)
    
    

def measures(x):
    return {
            '1': measure1,
            '2': measure2,
            '3': measure3,
            '4': measure4,
        }.get(x)

def switcher(A, x):
    def closure(C, m, n):
        return {
            1: A.compare1(C, m, n),
            2: A.compare2(C, m, n),
            3: A.compare3(C, m, n),
            4: A.compare4(C, m, n),
        }.get(x)
    return closure    

def measure1(A):
    print("Method 1 is being tested\n")
    print("Array to be tested: ",A.ar)
    a = datetime.datetime.now()
    print("MSS found: ",A.solution1())
    b = datetime.datetime.now()
    c = b - a
    if c.seconds != 0:
        print("elapsed time: ",c.seconds,"(s)\n")
    else:
        print("elapsed time: ",c.microseconds,"(µs)\n")
    return c


def measure2(A):
    print("Method 2 is being tested\n")
    print("Array to be tested: ",A.ar)
    a = datetime.datetime.now()
    print("MSS found: ",A.solution2())
    b = datetime.datetime.now()
    c = b - a
    if c.seconds != 0:
        print("elapsed time: ",c.seconds,"(s)\n")
    else:
        print("elapsed time: ",c.microseconds,"(µs)\n")
    return c

def measure3(A):
    l = 0
    r = len(A.ar)-1
    print("Method 3 is being tested\n")
    print("Array to be tested: ",A.ar)
    a = datetime.datetime.now()
    print("MSS found: ",A.solution3(l,r))
    b = datetime.datetime.now()
    c = b - a
    if c.seconds != 0:
        print("elapsed time: ",c.seconds,"(s)\n")
    else:
        print("elapsed time: ",c.microseconds,"(µs)\n")
    return c
    
def measure4(A):
    print("Method 4 is being tested\n")
    print("Array to be tested: ",A.ar)
    a = datetime.datetime.now()
    print("MSS found: ",A.solution4())
    b = datetime.datetime.now()
    c = b - a
    if c.seconds != 0:
        print("elapsed time: ",c.seconds,"(s)\n")
    else:
        print("elapsed time: ",c.microseconds,"(µs)\n")
    return c


main()

print("Have a good Day!")