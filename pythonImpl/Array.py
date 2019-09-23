import random
import math
class Array:
    ar = []
    def __init__(self, ar):
        self.ar = ar

    def newAr(self, n):
        for i in range(n):
            self.ar.append(random.randint(-50,50))

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

        m = int((l+r)/2)

        mss_left = self.solution3(l,m)
        mss_right = self.solution3(m+1,r)
        mss_middle = self.middle3(l,m,r)
        return max(mss_left,mss_right,mss_middle)

    def middle3(self,l,m,r):
        max_left_sum = float('-inf')
        sum = 0
        for i in range(m,l-1,-1):
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

    def solution4(self):
        max_sum = 0
        this_sum = 0

        for i in range(len(self.ar)):
            this_sum += self.ar[i]
            if this_sum > max_sum:
                max_sum = this_sum
            elif this_sum < 0:
                this_sum = 0
        return max_sum

    def compare1(self, C, m, n):
        return C.microseconds *  (math.pow(float(n),3.0)/math.pow(float(m),3.0))
        

    def compare2(self, C, m, n):
        a = C.microseconds / (math.pow(float(m),2.0) + m)
        return (a * math.pow(float(n),2.0)) + n

    def compare3(self, C, m, n):
        a = C.microseconds/ (m * math.log2(float(m)))
        return a * n * math.log2(float(n))  

    def compare4(self, C, m, n):
        slope = C.microseconds/m
        return n * slope
