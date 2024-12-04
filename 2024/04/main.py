inp = [[x for x in line.strip()] for line in open('inp').readlines()]

def p1(inp: list[list[str]]):
    letters = {'X': 'M', 'M': 'A', 'A': 'S'}
    dirs = [(1, 0), (0, 1), (-1, 0), (0, -1), (1, 1), (-1, -1), (1, -1), (-1, 1)]
    count = 0
    
    for y in range(len(inp)):
        for x in range(len(inp[0])):
            if inp[y][x] == 'X':
                for d in dirs:
                    if check(inp, x, y, d, 'X', letters):
                        count += 1
    print(count)

def check(inp: list[list[str]], x: int, y: int, d: tuple, i: str, letters: dict[str, str]):
    if i == 'S':
        return True 
    
    n_x, n_y = x + d[0], y + d[1]
    if n_x < 0 or n_y < 0 or n_x >= len(inp[0]) or n_y >= len(inp):
        return False
    
    if inp[n_y][n_x] == letters[i]:
        return check(inp, n_x, n_y, d, letters[i], letters)
    return False

def p2(inp: list[list[str]]):
    letters = {'X':'','M': 'S', 'A':'', 'S': 'M'}
    count = 0
    for y in range(len(inp)):
        for x in range(len(inp[0])):
            if inp[y][x] == 'A':
                if x + 1 >= len(inp[0]) or x - 1 < 0 or y + 1 >= len(inp) or y - 1 < 0:
                    continue
                top_left = inp[y-1][x-1]
                top_right = inp[y-1][x+1]
                if top_left not in ['M', 'S'] or top_right not in ['M', 'S']:
                    continue
                if top_left == letters[inp[y+1][x+1]] and top_right == letters[inp[y+1][x-1]]:
                    count += 1
                
    print(count)

if __name__ == '__main__':
    p1(inp)
    p2(inp)
