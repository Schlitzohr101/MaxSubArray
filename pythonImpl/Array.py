import random
import math
class Array:
    ar = []
    def __init__(self, n):
        newAr(self,n)

    def __init__(self, ar):
        self.ar = ar

    def newAr(self, n):
        for i in range(n):
            self.ar[i] = random.randint(-50,50)

    def solution1(self):
        max_sum = 0
        n = len(self.ar)
        for i in range(n):
            for j in range(i,n):
                this_sum = 0
                for k in range (i,j+1):#needed to be equal to j in the range <=
                    this_sum += self.ar[k]
                if this_sum > max_sum:
                    max_sum = this_sum
        return max_sum

    def solution2(self):
        max_sum = 0
        n = len(self.ar)
        for i in range(n):
            this_sum = 0
            for j in range(i,n):
                this_sum += self.ar[j]
            if this_sum > max_sum:
                max_sum = this_sum
        return max_sum

    def solution3(self, l, r):

        if l == r:
            return self.ar[l]
        if r == l+1:
            return max(self.ar[l],self.ar[r],self.ar[l]+self.ar[r])

        m = l+r/2
        print('m: ',m)

        mss_left = self.solution3(l,m)
        print("left: ",mss_left)
        mss_right = self.solution3(m,r)
        print("right: ", mss_right)
        mss_middle = self.middle3(l,m,r)
        print("middle: " ,mss_middle)
        return max(mss_left,mss_right,mss_middle)

    def middle3(self,l,m,r):
        max_left_sum = float('-inf')
        sum = 0
        for i in range(m,l-1):
            sum += self.ar[i]
            if sum > max_left_sum:
                max_left_sum = sum
        max_right_sum = float('-inf')
        sum = 0
        for i in range(m+1,r+1):
            sum += self.ar[i]
            if sum > max_right_sum:
                max_right_sum = sum
        return max_left_sum + max_right_sum