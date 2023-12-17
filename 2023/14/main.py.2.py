inp = [list(x.strip()) for x in open('inp').readlines()]

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

def p2(inp):
    seen = []
    cur_inp = inp

    tgt = 1_000_000_000
    while len(seen) != tgt:
        cur_inp_str = ''.join([x for y in cur_inp for x in y])
        if cur_inp_str in seen:
            found_idx = seen.index(cur_inp_str)
            cycle_length = len(seen)-found_idx
            tgt = ((tgt - found_idx) % cycle_length) + found_idx + cycle_length

        cur_inp = cycle(cur_inp)
        seen.append(cur_inp_str)
    return cur_inp

p1_out = roll(inp)
scores = sum([len(p1_out)-y for y, y_val in enumerate(p1_out) for x in y_val if 'O' in x])
print(scores)

p2_out = p2(inp)
scores = sum([len(p2_out)-y for y, y_val in enumerate(p2_out) for x in y_val if 'O' in x])
print(scores)


