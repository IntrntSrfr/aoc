from math import sqrt

inp = [line.strip().split() for line in open("inp").readlines()]
dirs = {"R": (0, 1), "D": (1, 0), "L": (0, -1), "U": (-1, 0)}

def parse_inp(inp, part=1):
    if part == 1:
        inp = list(map(lambda x: (x[0], int(x[1])), inp))
    else:
        inp = list(map(lambda x: (list(dirs)[int(x[2][-2])], int(x[2][2:-2], 16)), inp))
    vertices = [(0, 0)]
    for dir, n in inp[:-1]:
        y, x = vertices[-1]
        vertices.append((y + dirs[dir][0] * n, x + dirs[dir][1] * n))
    return vertices

def solve(vertices):
    area = 0
    for i in range(len(vertices)):
        y1, x1 = vertices[i]
        y2, x2 = vertices[(i + 1) % len(vertices)]
        area += (x1 * y2) - (x2 * y1)
        area += sqrt((x2 - x1) ** 2 + (y2 - y1) ** 2)
    return int(abs(area // 2) + 1)

print(solve(parse_inp(inp, 1)), solve(parse_inp(inp, 2)))
