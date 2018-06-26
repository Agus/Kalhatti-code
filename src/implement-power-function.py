import math

class CustomMath:
    def pow(self,x,n,d):
        return abs(math.pow(x,n) % d);
