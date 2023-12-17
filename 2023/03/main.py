from itertools import permutations

inp = [line.strip() for line in open('inp').readlines()]

blocks = [] # ([(i, j1), (i, j2), (i, j3)], number)
for i, line in enumerate(inp):
    current_int_str = ''
    current_int_positions = []
    for j, char in enumerate(line):
        if char.isdigit():
            current_int_str += char
            current_int_positions.append((i, j))
        else:
            if current_int_str == '':
                continue
            blocks.append((current_int_positions, int(current_int_str)))
            current_int_str = ''
            current_int_positions = []

p1_score = 0
for block in blocks:
    is_part = False
    for coords in block[0]:
        for i in range(-1, 2):
            for j in range(-1, 2):
                if j == 0 and i == 0:
                    continue
                check_y = coords[0]+i
                check_x = coords[1]+j
                if ((check_y, check_x) in coords or 
                    check_y < 0 or check_y > len(inp)-1 or 
                    check_x < 0 or check_x > len(inp[0])-1): 
                    continue

                char = inp[check_y][check_x]
                if char in '.1234567890' :
                    continue
                is_part = True
    if is_part:
        p1_score += block[1]
print(p1_score)
