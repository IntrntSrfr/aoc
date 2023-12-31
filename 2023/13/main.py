def transpose(l):
    return ["".join(r) for r in zip(*l)]

def find_mirror_idx(mirr):
    mirror_idxs = []
    for i in range(len(mirr) - 1):
        if mirr[i] == mirr[i + 1]:
            mirror_idxs.append(i)
    rets = []
    for idx in mirror_idxs:
        for i in range(0, idx + 1):
            if (
                (idx - i == 0 or idx + i + 1 == len(mirr) - 1) and 
                mirr[idx - i] == mirr[idx + i + 1]
            ):
                rets.append(idx + 1)
                break
            if mirr[idx - i] != mirr[idx + i + 1]:
                break
    return rets


def p2_mirror_score(mirr):
    init_idxs = set(find_mirror_idx(mirr))
    init_idxs_t = set(find_mirror_idx(transpose(mirr)))

    for y in range(len(mirr)):
        for x in range(len(mirr[y])):
            mirr[y][x] = p2_map[mirr[y][x]]
            idx = set(find_mirror_idx(mirr))
            idx_t = set(find_mirror_idx(transpose(mirr)))

            for idx in idx - init_idxs:
                return idx * 100
            for idx_t in idx_t - init_idxs_t:
                return idx_t
            mirr[y][x] = p2_map[mirr[y][x]]
    return 0

def p1(inp):
    score = 0
    for mirr in inp:
        for v in find_mirror_idx(mirr):
            score += 100 * v
        for v in find_mirror_idx(transpose(mirr)):
            score += v
    return score

def p2(inp):
    score = 0
    for mirr in inp:
        score += p2_mirror_score(mirr)
    return score

inp = [
    [[char for char in row] for row in pattern.split("\n")]
    for pattern in open("inp").read().split("\n\n")
]
p2_map = {"#": ".", ".": "#"}

print(p1(inp), p2(inp))
# 33520 34824
