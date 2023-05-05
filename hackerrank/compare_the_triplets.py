#!/bin/python3
#https://www.hackerrank.com/challenges/compare-the-triplets/problem

import math
import os
import random
import re
import sys

from functools import reduce
import operator

# Complete the compareTriplets function below.
def compareTriplets(a, b):
    alice_point = reduce(operator.add, map(lambda x, y: x > y, a, b))
    bob_point = reduce(operator.add, map(lambda x, y: x > y, b, a))

    return [alice_point, bob_point]

if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    a = list(map(int, input().rstrip().split()))

    b = list(map(int, input().rstrip().split()))

    result = compareTriplets(a, b)

    fptr.write(' '.join(map(str, result)))
    fptr.write('\n')

    fptr.close()
