#!/bin/python3
#https://www.hackerrank.com/challenges/non-divisible-subset/problem

import math
import os
import random
import re
import sys

#
# Complete the 'nonDivisibleSubset' function below.
#
# The function is expected to return an INTEGER.
# The function accepts following parameters:
#  1. INTEGER k
#  2. INTEGER_ARRAY s
#

def nonDivisibleSubset(k, s):
    rem_num = [0] * k

    rems = [None] * len(s)
    for i, v in enumerate(s):
        rem = v % k
        rem_num[rem] += 1
        rems[i] = rem
    
    high = -1
    if (k / 2) % 1 == 0:
        high = k // 2
    else:
        high = (k // 2) + 1
    
    res = 0
    for i in range(1, high):
        if rem_num[i] >= rem_num[k - i]:
           res += rem_num[i] 
        else:
            res += rem_num[k - i] 
    if rem_num[0] > 0:
        res += 1

    if (k / 2) % 1 == 0 and rem_num[k // 2] > 0:
        res += 1

    return res
    
        

    


if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    first_multiple_input = input().rstrip().split()

    n = int(first_multiple_input[0])

    k = int(first_multiple_input[1])

    s = list(map(int, input().rstrip().split()))

    result = nonDivisibleSubset(k, s)

    fptr.write(str(result) + '\n')

    fptr.close()
