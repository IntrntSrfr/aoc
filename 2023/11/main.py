from itertools import combinations

inp = [list(x.strip()) for x in open('inp').readlines()]

def get_empty_lines(input: list[list[str]]):
    def transpose(input: list[list[str]]):
        return list(map(list, zip(*input)))
    
    def get_empty(input: list[list[str]]):
        return [x for x in range(0, len(input)-1) if '#' not in input[x]]
    
    rows = get_empty(input)
    cols = get_empty(transpose(input))
    return rows, cols

points = [(x, y) for y, inner in enumerate(inp) for x, val in enumerate(inner) if val == '#']
ers, ecs = get_empty_lines(inp)

def solve(delta):
    n = 0
    for (sx, sy), (tx, ty) in combinations(points, 2):
        start_row, stop_row = sorted([sy, ty])
        start_col, stop_col = sorted([sx, tx])
        nn = abs(sx - tx) + abs(sy - ty)
        for r in ers:
            if start_row < r < stop_row:
                nn += delta-1
        for c in ecs:
            if start_col < c < stop_col:
                nn += delta-1
        n += nn
    return n

print(solve(2), solve(1000000))
