def transpose(inp):
    return list(zip(*inp))

def rotate(inp):
    return [list(reversed(row)) for row in transpose(inp)]

def roll(inp):
    for y in range(len(inp)):
        for x in range(len(inp[y])):
            if inp[y][x] != 'O':
                continue
            dy = 0
            for i in range(y+1):
                if y-i < 0 or (inp[y-i][x] != '.' and y-i != y):
                    break
                dy = i
            if y-dy < 0 or inp[y-dy][x] != '.':
                continue
            inp[y][x], inp[y-dy][x] = inp[y-dy][x], inp[y][x] 
    return inp

def cycle(inp):
    for _ in range(4):
        inp = roll(inp)
        inp = rotate(inp)
    return inp

def solve(inp):
    seen = []
    cur_inp = inp

    tgt = 1_000_000_000
    for i in range(tgt):
        cur_inp_str = ''.join([x for y in cur_inp for x in y])
        if cur_inp_str in seen:
            found_idx = seen.index(cur_inp_str)
            cycle_length = len(seen)-found_idx
            remaining = (tgt-i) % cycle_length
            for _ in range(remaining):
                cur_inp = cycle(cur_inp)
            return cur_inp

        cur_inp = cycle(cur_inp)
        seen.append(cur_inp_str)
    return cur_inp

inp = [list(x.strip()) for x in open('inp').readlines()]

p1_out = roll(inp)
p2_out = solve(inp)
p1_scores = sum([len(p1_out)-y for y, y_val in enumerate(p1_out) for x in y_val if 'O' in x])
p2_scores = sum([len(p2_out)-y for y, y_val in enumerate(p2_out) for x in y_val if 'O' in x])
print(p1_scores, p2_scores)
