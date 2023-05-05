#!/bin/python3
#https://www.hackerrank.com/challenges/extra-long-factorials/problem

import math
import os
import random
import re
import sys

# Complete the extraLongFactorials function below.
def extraLongFactorials(n):
    res = facturail(n)
    print(res)

def facturail(n):
    if n == 0 or n == 1:
        return 1
    elif n == 2:
        return 2
    else:
        fact = 2
        for i in range(3, n + 1):
            fact *= i
        return fact
        


if __name__ == '__main__':
    n = int(input())
    extraLongFactorials(n)
