from math import prod


inp_t, inp_d = [list(map(int, line.split()[1:])) for line in open('inp')]

f = lambda x, l: x * (l - x)

def solve(inp_t, inp_d, part):
    if part == 2:
        inp_t = [int(''.join(map(str, inp_t)))]
        inp_d = [int(''.join(map(str, inp_d)))]
    
    scores = []
    for t, d in zip(inp_t, inp_d):
        scores.append(sum(f(x, t) > d for x in range(t+1)))
    return prod(scores)

print(solve(inp_t, inp_d, 1))
print(solve(inp_t, inp_d, 2))