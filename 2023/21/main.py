import numpy as np

inp = [list(x.strip()) for x in open("inp").readlines()]

src = [
    (y, x)
    for y, y_val in enumerate(inp)
    for x, x_val in enumerate(y_val)
    if x_val == "S"
][0]
dirs = [(-1, 0), (0, 1), (1, 0), (0, -1)]

steps = set([src])
steps2 = []

for i in range(1, 65 + 131 * 2 + 1):
    new_steps = set()
    for y, x in steps:
        for dy, dx in dirs:
            n_y, n_x = y + dy, x + dx
            if inp[n_y % len(inp)][n_x % len(inp)] != "#":
                new_steps.add((n_y, n_x))
    steps.clear()
    steps.update(new_steps)

    if i == 64:
        print(len(steps))
    if i % 131 == 65:
        steps2.append(len(steps))

print(steps2)
coeffs = np.polyfit([0, 1, 2], steps2, 2)
print(np.polyval(coeffs, 26501365 // 131))
