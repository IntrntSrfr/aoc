inp = [list(map(int, x.strip())) for x in open('inp').readlines()]
dirs = [(0, 1), (1, 0), (0, -1), (-1, 0)]

def get_zeroes(inp: list[list[int]]) -> list[tuple[int, int]]:
    return [(x, y) for y, row in enumerate(inp) for x, val in enumerate(row) if val == 0]

def find_paths(inp, start):
    stack, path, distPath = [start], set(), [] 
    while stack:
        x, y = stack.pop()
        if inp[y][x] == 9:
            path.add((x, y))
            distPath.append((x, y))
        for dx, dy in dirs:
            nx, ny = x + dx, y + dy
            if 0 <= nx < len(inp[0]) and 0 <= ny < len(inp) and inp[ny][nx] == inp[y][x] + 1:
                stack.append((nx, ny))
    return len(path), len(distPath)

def solve(inp: list[list[int]]):
    print(*(sum(find_paths(inp, p)[i] for p in get_zeroes(inp)) for i in (0, 1)))

if __name__ == '__main__':
    solve(inp)

