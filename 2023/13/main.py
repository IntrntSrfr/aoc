#inp = [y for x in open('inp').read().split('\n\n') for y in x.split()]
inp = [[[char for char in row] for row in layer.split('\n')] for layer in open('inp').read().split('\n\n')]
#print(inp)

def transpose(l: list[str]):
    return [''.join(r) for r in zip(*l)]

def find_mirror_idx(grid):
    mirror_idxs = []
    for i in range(len(grid)-1):
        if grid[i] == grid[i+1]:
            mirror_idxs.append(i)

    ret = 0
    for mirror_idx in mirror_idxs:
        for i in range(0, mirror_idx+1):
            if (mirror_idx-i == 0 or mirror_idx + i + 1 == len(grid)-1) and grid[mirror_idx-i] == grid[mirror_idx+i+1]:
                ret = mirror_idx + 1
                break
            if grid[mirror_idx-i] != grid[mirror_idx+i+1]:
                break
    return ret

p2_map = {'#': '.', '.': '#'}
def p2(inp):
    score = 0
    for grid in inp:
        score += check_many(grid)

    return score

def pp(inp):
    for line in inp:
        print(line)

def check_many(grid):
    init_idx = find_mirror_idx(grid)
    init_idx_t = find_mirror_idx(transpose(grid))


    for y in range(len(grid)):
        for x in range(len(grid[y])):
            grid[y][x] = p2_map[grid[y][x]]
            #pp(grid)
            #print()
            idx = find_mirror_idx(grid)
            idx_t = find_mirror_idx(transpose(grid))

            if idx and idx != init_idx:
                #print("normal", init_idx, init_idx_t, idx, idx_t)
                return idx*100
            if idx_t and idx_t != init_idx_t:
                #print("transpose", init_idx, init_idx_t, idx, idx_t)
                return idx_t
            grid[y][x] = p2_map[grid[y][x]]

            """ 
            idx = find_mirror_idx(grid)
            if idx != init_idx and not t:
                return 100*idx    
            if not idx:
                return find_mirror_idx(transpose(grid))
            """
        
    return 0

def p1(inp):
    score = 0
    for grid in inp:
        cur = 100*find_mirror_idx(grid)
        if not cur:
            cur += find_mirror_idx(transpose(grid))
        score+=cur
    return score

# 33520 34824

print(p1(inp), p2(inp))