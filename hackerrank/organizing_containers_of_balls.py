#!/bin/python3
#https://www.hackerrank.com/challenges/organizing-containers-of-balls/problem

import math
import os
import random
import re
import sys

# Complete the organizingContainers function below.
def organizingContainers(container):
    n = len(container)
    
    ball_sums = [0] * n
    containers_cap = [0] * n

    for con_idx, single_con in enumerate(container):
        for ball_idx, ball_num in enumerate(single_con):
             containers_cap[con_idx] = containers_cap[con_idx] + ball_num
             ball_sums[ball_idx] = ball_sums[ball_idx] + ball_num 

    ball_sums.sort()
    containers_cap.sort()

    possible = True
    for ball_num, con_cap in zip(ball_sums, containers_cap):
        if ball_num != con_cap:
            possible = False
            break

    if possible:
        return "Possible"
    else:
        return "Impossible"



if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    q = int(input())

    for q_itr in range(q):
        n = int(input())

        container = []

        for _ in range(n):
            container.append(list(map(int, input().rstrip().split())))

        result = organizingContainers(container)

        fptr.write(result + '\n')

    fptr.close()
