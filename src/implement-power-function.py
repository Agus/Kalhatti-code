import math

class CustomMath:
    def pow(self,x,n,d):
        return abs(math.pow(x,n) % d);

print(pow(2,3,3))
