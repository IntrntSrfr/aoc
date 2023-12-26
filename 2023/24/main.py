from itertools import combinations
import numpy as np

def boom(p1, p2):
    x1, y1, _ = p1[0]
    dx1, dy1, _ = p1[1]
    x2, y2, _ = p2[0]
    dx2, dy2, _ = p2[1]
    coeffs = np.array([[dx1, -dx2], [dy1, -dy2]])
    consts = np.array([x2 - x1, y2 - y1])
    try:
        ts = np.linalg.solve(coeffs, consts)
        ix = x1 + dx1 * ts[0]
        iy = y1 + dy1 * ts[0]
        return ix, iy, ts
    except np.linalg.LinAlgError:
        return None

inp = [
    [list(map(int, y.split(","))) for y in x.strip().split("@")]
    for x in open("inp").readlines()
]

def p1():
    min_val, max_val = 7, 27
    hits = 0
    for p1, p2 in combinations(inp, 2):
        ans = boom(p1, p2)
        if not ans:
            continue
        x, y, t = ans
        if (
            min_val <= x <= max_val
            and min_val <= y <= max_val
            and not sum(n < 0 for n in t)
        ):
            hits += 1
    return hits

print(p1())
