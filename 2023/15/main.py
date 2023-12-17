def HASH(cur: int, inp: str):
    for c in inp:
        cur += ord(c)
        cur *= 17
        cur %= 256
    return cur

def parse_step(step):
    if '=' in step:
        label, focal = step.split('=')
        focal = int(focal)
        op = '='
    else:
        label = step[:-1]
        focal = -1
        op = '-'
    return label, op, focal

def p1(inp):
    return sum([HASH(0, step) for step in inp])

def p2(inp):
    boxes = [[] for _ in range(256)]
    for step in inp:
        label, op, focal = parse_step(step)
        box_idx = HASH(0, label)
        box = boxes[box_idx]
        found_idx = next((i for i, v in enumerate(box) if v[0] == label), -1)
        if op == '=':
            if found_idx != -1:
                box[found_idx][1] = focal
            else:
                box.append([label, focal])
        else:
            if found_idx != -1:
                del box[found_idx]
    return sum(
        (i + 1) * (j + 1) * lens[1] 
        for i, box in enumerate(boxes) 
        for j, lens in enumerate(box)
    )

inp = [x for x in open('inp').read().split(',')]
print(p1(inp), p2(inp))
