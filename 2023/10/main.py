inp = [list(row.strip()) for row in open("inp").readlines()]
directions = [(-1, 0), (1, 0), (0, -1), (0, 1)]

representations = {
    "|": [".#.", ".#.", ".#."],
    "-": ["...", "###", "..."],
    "L": [".#.", ".##", "..."],
    "J": [".#.", "##.", "..."],
    "7": ["...", "##.", ".#."],
    "F": ["...", ".##", ".#."],
    ".": ["...", "...", "..."],
}

def scale_up_grid(original_grid, fallback="F"):
    scaled_grid = []
    src = (-1, -1)
    for y, row in enumerate(original_grid):
        new_rows = ["", "", ""]
        for x, tile in enumerate(row):
            if tile == "S":
                src = (y * 3 + 1, x * 3 + 1)
                tile = fallback
            representation = representations[tile]
            for i in range(3):
                new_rows[i] += representation[i]
        scaled_grid.extend(new_rows)
    return scaled_grid, src


def bfs(grid, start):
    def is_valid(x, y):
        return 0 <= x < len(grid) and 0 <= y < len(grid[0]) and grid[x][y] == "#"

    queue, visited, max_n = [(start, 0)], set([start]), -1
    while queue:
        (y, x), n = queue.pop(0)
        max_n = max(max_n, n)
        for dx, dy in directions:
            n_y, n_x = y + dy, x + dx
            if not is_valid(n_y, n_x) or (n_y, n_x) in visited:
                continue
            queue.append(((n_y, n_x), n + 1))
            visited.add((n_y, n_x))
    return max_n

s_grid, src = scale_up_grid(inp)
#for line in s_grid:
#    print(line)
p1_score = bfs(s_grid, (src))//3
print(p1_score)
