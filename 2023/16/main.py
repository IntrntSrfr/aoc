from collections import defaultdict

def get_laser_energy(start_pos):
    lasers = [start_pos]
    seen = defaultdict(set)

    while lasers:
        (x, y), (dx, dy) = lasers.pop()
        while True:
            if y < 0 or y >= len(grid) or x < 0 or x >= len(grid[y]) or (dx, dy) in seen[(x, y)]:
                break
            seen[(x, y)].add((dx, dy))
            tile = grid[y][x]
            if tile == '/':
                dx, dy = -dy, -dx
            elif tile == '\\':
                dx, dy = dy, dx
            elif tile == '-' and dy != 0:
                lasers.extend([[[x+1, y],[1, 0]], [[x-1, y],[-1, 0]]])
                break
            elif tile == '|' and dx != 0:
                lasers.extend([[[x, y+1],[0, 1]], [[x, y-1],[0, -1]]])
                break
            x += dx
            y += dy
    return len(seen)

grid = [list(x.strip()) for x in open('inp')]

def p1():
    return get_laser_energy([[0, 0], [1, 0]])

def p2():
    scores = []
    for x in range(len(grid)):
        scores.append(get_laser_energy([[x, 0], [0, 1]]))
        scores.append(get_laser_energy([[x, len(grid) - 1], [0, -1]])) 

    for y in range(1, len(grid) - 1):
        scores.append(get_laser_energy([[0, y], [1, 0]])) 
        scores.append(get_laser_energy([[len(grid) - 1, y], [-1, 0]])) 
    return max(scores)

print(p1(), p2())