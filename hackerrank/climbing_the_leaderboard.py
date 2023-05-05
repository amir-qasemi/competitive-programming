#!/bin/python3
#https://www.hackerrank.com/challenges/climbing-the-leaderboard/problem


import math
import os
import random
import re
import sys

# Complete the climbingLeaderboard function below.
def climbingLeaderboard(scores, alice):
    ranks = [None] * len(scores)
    curr_rank = 1
    ranks[0] = 1
    for i in range(1, len(scores)):
        if scores[i] == scores[i - 1]:
            ranks[i] = curr_rank
        else:
            curr_rank += 1
            ranks[i] = curr_rank

    alice_rank = [None] * len(alice)
    j = len(scores) - 1
    for i in range(0, len(alice)):
        if alice[i] > scores[j]:
            while j >= 0 and alice[i] > scores[j]:
                j -= 1

        if j < 0:
            for k in range(i, len(alice)):
                alice_rank[k] = 1
            break

        if alice[i] == scores[j]:
            alice_rank[i] = ranks[j]
        elif alice[i] < scores[j]:
            alice_rank[i] = ranks[j] + 1            

    return alice_rank
            



if __name__ == '__main__':
    fptr = open(os.environ['OUTPUT_PATH'], 'w')

    scores_count = int(input())

    scores = list(map(int, input().rstrip().split()))

    alice_count = int(input())

    alice = list(map(int, input().rstrip().split()))

    result = climbingLeaderboard(scores, alice)

    fptr.write('\n'.join(map(str, result)))
    fptr.write('\n')

    fptr.close()
