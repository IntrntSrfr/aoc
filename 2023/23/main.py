inp = [x.strip() for x in open('inp').readlines()]
#print(inp)

def dfs(grid, src, tgt, part1=True):
    def is_valid(y, x):
        return 0 <= y < rows and 0 <= x < cols and grid[y][x] != '#'

    rows, cols = len(grid), len(grid[0])
    paths = []
    dirs = {
        '.': [(-1, 0), (0, 1), (1, 0), (0, -1)],
        '^': [(-1, 0)],
        '>': [(0, 1)],  
        'v': [(1, 0)],
        '<': [(0, -1)]
    }
    q = [(*src, 0, set())]

    while q:
        y, x, steps, seen = q.pop()
        if (y, x) == tgt:
            paths.append(steps)
            continue
        if (y, x) in seen:
            continue
        seen.add((y, x))
        for dy, dx in dirs[grid[y][x]] if part1 else dirs['.']:
            ny, nx = y + dy, x + dx
            if is_valid(ny, nx) and (ny, nx) not in seen:
                q.append((ny, nx, steps + 1, seen.copy()))

    return paths

src = (0,1)
tgt = (len(inp)-1, len(inp[0])-2)

print(max(dfs(inp, src, tgt, True)))

#found, seen = dfs(inp, (0,1), (len(inp)-1, len(inp[0])-2))
#print(found)
#print(seen)

def pp(grid, seen):
    for y in range(len(grid)):
        for x in range(len(grid[0])):
            if (y,x) in seen:
                print('O', end='')
                continue
            print(grid[y][x], end='')
        print()

#pp(inp, [])
#print()
#pp(inp, seen)