from itertools import combinations
import numpy as np
from typing import Tuple
from collections import deque

inp = [list(x.strip()) for x in open('inp').readlines()]

def expand(input: list[list[str]]):
    def transpose(input: list[list[str]]):
        return list(map(list, zip(*input)))
    
    def expand_inner(input: list[list[str]]):
        for i in range(len(input)-1, 0, -1):
            if len(set(input[i])) == 1:
                input.insert(i+1, ['.'] * len(input[i]))
                i -= 2
        return input
    
    input = expand_inner(input)
    input = expand_inner(transpose(input))
    return transpose(input)

def manhattan_distance(src, tgt):
    return abs(src[0] - tgt[0]) + abs(src[1] - tgt[1])

new_inp = expand(inp)
#print(np.array(new_inp))

points = [(x, y) for y, inner in enumerate(new_inp) for x, val in enumerate(inner) if val == '#']
points_combos = list(combinations(points, 2))

print(sum([manhattan_distance(src, tgt) for src, tgt in points_combos]))

