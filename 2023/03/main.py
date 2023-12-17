import re

inp = [line.strip() for line in open('inp').readlines()]
blocks = []
for y, line in enumerate(inp):
    blocks.extend([
        (list((y, x) for x in range(m.start(), m.end())), int(m.group())) 
        for m in re.finditer(r'\d+', line)
    ])

gears = []
for y, line in enumerate(inp):
    gears.extend([(y, m.start()) for m in re.finditer(r'\*', line)])

def p2():
    score = 0
    seen = set()
    for (y, x) in gears:
        parts = []
        for dy in range(-1, 2):
            for dx in range(-1, 2):
                if dy == 0 and dx == 0:
                    continue
                if (y+dy, x+dx) in seen:
                    continue
                for block in blocks:
                    if (y+dy, x+dx) in block[0]:
                        parts.append(block[1])
                        seen.update(block[0])
        if len(parts) == 2:
            score += parts[0]*parts[1]
    return score

def p1():
    score = 0
    for block in blocks:
        coords, num = block
        is_part = False
        for (y, x) in coords:
            for dy in range(-1, 2):
                for dx in range(-1, 2):
                    if dy == 0 and dx == 0:
                        continue
                    if y + dy < 0 or y+dy >= len(inp) or x+dx <0 or x+dx >= len(inp[0]):
                        continue
                    if inp[y+dy][x+dx] in '.1234567890':
                        continue
                    is_part = True
        if is_part:
            score += num
    return score

print(p1(), p2())
