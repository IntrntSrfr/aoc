inp = [list(map(int, x.strip().split())) for x in open('inp').readlines()]

def solve(inp):
    score = 0
    for line in inp:
        diffs = []
        while line:
            diffs.append(line)
            line = [line[i + 1] - line[i] for i in range(len(line) - 1)]
        for i in range(len(diffs) - 2, -1, -1):
            diffs[i].append(diffs[i][-1] + diffs[i + 1][-1])
        score += diffs[0][-1]
    return score

rev_inp = [list(reversed(x)) for x in inp]
print(solve(inp), solve(rev_inp))