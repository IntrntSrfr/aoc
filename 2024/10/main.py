inp = [list(map(int, x.strip())) for x in open('inp').readlines()]
dirs = [(0, 1), (1, 0), (0, -1), (-1, 0)]

def get_zeroes(inp: list[list[int]]) -> list[tuple[int, int]]:
    return [(x, y) for y, row in enumerate(inp) for x, val in enumerate(row) if val == 0]

def score(inp, start):
    stack, visited, score = [start], set(), 0
    while stack:
        x, y = stack.pop()
        if inp[y][x] == 9:
            score += 1
            continue
        for dx, dy in dirs:
            nx, ny = x + dx, y + dy
            if 0 <= nx < len(inp[0]) and 0 <= ny < len(inp) and (nx, ny) not in visited and inp[ny][nx] == inp[y][x] + 1:
                visited.add((nx, ny))
                stack.append((nx, ny))
    return score

def rating(inp, start):
    stack, paths = [(start, [start])], 0
    while stack:
        (x, y), path = stack.pop()
        if inp[y][x] == 9:
            paths += 1
            continue
        for dx, dy in dirs:
            nx, ny = x + dx, y + dy
            if 0 <= nx < len(inp[0]) and 0 <= ny < len(inp) and (nx, ny) not in path and inp[ny][nx] == inp[y][x] + 1:
                stack.append(((nx, ny), path + [(nx, ny)]))
    return paths

def p1(inp: list[list[int]]):
    starts = get_zeroes(inp)
    print(sum([score(inp, p) for p in starts]))
    
def p2(inp: list[list[int]]):
    starts = get_zeroes(inp)
    print(sum([rating(inp, p) for p in starts]))

if __name__ == '__main__':
    p1(inp)
    p2(inp)
