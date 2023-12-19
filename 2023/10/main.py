inp = [list(row.strip()) for row in open("inp").readlines()]
dirs = [(-1, 0), (1, 0), (0, -1), (0, 1)]

tiles = {
    "|": [".#.", ".#.", ".#."],
    "-": ["...", "###", "..."],
    "L": [".#.", ".##", "..."],
    "J": [".#.", "##.", "..."],
    "7": ["...", "##.", ".#."],
    "F": ["...", ".##", ".#."],
    ".": ["...", "...", "..."],
}

def scale_up_grid(grid, fallback="F"):
    new_grid = []
    src = (-1, -1)
    for y, row in enumerate(grid):
        new_rows = ["", "", ""]
        for x, tile in enumerate(row):
            if tile == "S":
                src = (y * 3 + 1, x * 3 + 1)
                tile = fallback
            large_tile = tiles[tile]
            for i in range(3):
                new_rows[i] += large_tile[i]
        new_grid.extend(new_rows)
    return new_grid, src

def bfs(grid, start):
    def is_valid(x, y):
        return 0 <= x < len(grid) and 0 <= y < len(grid[0]) and grid[x][y] == "#"

    q, seen, max_n = [(start, 0)], set([start]), -1
    while q:
        (y, x), n = q.pop(0)
        max_n = max(max_n, n)
        for dx, dy in dirs:
            n_y, n_x = y + dy, x + dx
            if not is_valid(n_y, n_x) or (n_y, n_x) in seen:
                continue
            q.append(((n_y, n_x), n + 1))  # type: ignore
            seen.add((n_y, n_x))
    return seen, max_n // 3

def floodfill(grid, src, loop):
    def is_valid(x, y):
        return (0 <= x < len(grid) and 0 <= y < len(grid[0])
                and (y, x) not in seen and (y, x) not in loop)

    q, seen, count = [src], set([]), 0
    while q:
        (y, x) = q.pop(0)
        if not is_valid(x, y):
            continue
        if x % 3 == 1 and y % 3 == 1:
            count += 1
        seen.add((y, x))
        q.extend([(y + 1, x), (y - 1, x), (y, x + 1), (y, x - 1)])
    return count

s_grid, src = scale_up_grid(inp)
loop, p1_score = bfs(s_grid, (src))
p2_score = floodfill(s_grid, (src[0] + 1, src[1] + 1), loop)
print(p1_score, p2_score)
