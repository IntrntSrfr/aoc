from heapq import heappop, heappush

def solve(grid, min_steps, max_steps):
    def in_bounds(grid, y, x):
        return 0 <= y < len(grid) and 0 <= x < len(grid[0])

    dirs = [(0, 1), (1, 0), (0, -1), (-1, 0)]
    # (heat, y, x, steps, dir_idx)
    pq = [(0, 0, 0, 0, 0)]
    seen = set()

    while pq:
        heat, y, x, steps, dir_idx = heappop(pq)
        if (y, x) == (len(grid) - 1, len(grid[0]) - 1) and steps >= min_steps:
            return heat
        if (y, x, steps, dir_idx) in seen:
            continue
        seen.add((y, x, steps, dir_idx))

        for i, (dy, dx) in enumerate(dirs):
            ny, nx = y + dy, x + dx

            if (
               dir_idx == (i + 2) % len(dirs) or 
               not in_bounds(grid, ny, nx) or
               steps < min_steps and i != dir_idx or
               steps > max_steps
            ):
                continue
            
            nsteps = steps + 1 if i == dir_idx else 1
            ndist = heat + grid[ny][nx]
            heappush(pq, (ndist, ny, nx, nsteps, i))
    return -1

inp = [list(map(int, x.strip())) for x in open("inp").readlines()]
print(solve(inp, 0, 3), solve(inp, 4, 10))
# 936 1157
