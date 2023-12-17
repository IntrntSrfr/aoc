inp = [list(row.strip()) for row in open('inp').readlines()]
#print(inp)

dirs = {
    '|': [[0, -1], [0, 1]],
    '-': [[-1, 0], [1, 0]],
    'L': [[0, -1], [1, 0]],
    'J': [[0, -1], [-1, 0]],
    '7': [[0, 1], [-1, 0]],
    'F': [[0, 1], [1, 0]],
}

def find_start(inp: list[list[str]]) -> tuple[int, int]:
    for y, row in enumerate(inp):
        if 'S' in row:
            return (row.index('S'), y)
    return (-1, -1)

#def find_start_tile(inp: list[list[str]]):

def bfs(start: tuple[int, int], start_tile = 'F'):
    visited = [start]
    q = [start]
    inp[start[1]][start[0]] = start_tile
    nums = [[0 for _ in range(len(inp[0]))] for _ in range(len(inp))]
    max_num = 0
    while q:
        cur = q.pop(0)
        cur_dirs = dirs[inp[cur[1]][cur[0]]]
        for dir in cur_dirs:
            n_x, n_y = cur[0] + dir[0], cur[1] + dir[1]
            if (n_x, n_y) not in visited:
                q.append((n_x, n_y))
                nums[n_y][n_x] = nums[cur[1]][cur[0]] + 1
                if nums[n_y][n_x] > max_num:
                    max_num = nums[n_y][n_x]
                visited.append((n_x, n_y))
    return max_num, visited

max_num, visited = bfs(find_start(inp), 'F')
print(visited)


score = 0 
for y in range(len(inp)):
    for x in range(len(inp[y])):
        if inp[y][x] != '.':
            continue
        wn = 0
        for cx in range(x):
            pass
        if wn % 2 == 1:
            score += 1

print(score)
