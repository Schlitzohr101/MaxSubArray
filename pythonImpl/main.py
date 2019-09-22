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

    given = input("Enter each Method to test i.e.(1234)\n:")
    for i in range(0,len(given)):
        method = measures(given[i])
        method(A)

        
def choice3():
    print("this is the third\n")

def measures(x):
    return {
            '1': measure1,
            '2': measure2,
            '3': measure3,
            '4': measure4,
        }.get(x)


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
        print("elapsed time: ",c.microseconds,"(ms)\n")


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
        print("elapsed time: ",c.microseconds,"(ms)\n")

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
        print("elapsed time: ",c.microseconds,"(ms)\n")
    
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
        print("elapsed time: ",c.microseconds,"(ms)\n")


main()

print("Have a good Day!")