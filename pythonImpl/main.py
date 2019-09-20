from Menu import Menu
from Array import Array
import re


def main():
    print("Hello world")
    M = Menu()
    M.printMenu()
    r = input()
    method = choice_selector(r)
    while r != 4:
        method()
        M.printMenu()
        r = input()
        method = choice_selector(r)

def choice_selector(c):
    def choices():
        return {choice1() : 1,
                choice2() : 2,
                choice3() : 3, 
        }.get(c, 4)

    return choices

def choice1():
    given = input("Input elements of your array followed by ',' i.e. : 1,2,...\n:")
    # print(given)
    # print(re.split('|,',given))
    A = Array(given)
    # print(A.ar)
    print("Solution 1: ", A.solution1())
    print("Solution 2: ", A.solution2())
    l = 0
    r = len(A.ar)-1
    print("Solution 3: ",A.solution3(l,r))

    



main()

print("Have a good Day!")