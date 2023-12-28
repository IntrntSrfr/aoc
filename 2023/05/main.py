def p1(seeds, maps):
    for map in maps:
        next_seeds = []
        for seed in seeds:
            mapped = False
            for dst, src, r in map:
                if seed in range(src, src + r):
                    next_seeds.append(seed + (dst - src))
                    mapped = True
                    continue
            if not mapped:
                next_seeds.append(seed)
        seeds = next_seeds
    return min(seeds)

def split(r1, r2):
    a, b = r1
    c, d = r2

    if max(a, c) < min(b, d):
        before = (a, max(a, c))
        overlap = (max(a, c), min(b, d))
        after = (min(b, d), b)
        before = before if before[0] != before[1] else ()
        after = after if after[0] != after[1] else ()
        return [before, overlap, after]
    return [r1]

def p2(seeds, maps):
    seeds = [(seeds[x], seeds[x] + seeds[x + 1]) for x in range(0, len(seeds), 2)]
    for map in maps:
        next_seeds = set()
        for seed in seeds:
            split_seeds = set()
            mapped = False
            for dst, src, r in map:
                splits = split(seed, (src, src+r))
                if len(splits) == 3:
                    mapped = True
                    seeds.extend([x for x  in [splits[0], splits[2]] if len(x)])
                    next_seeds.add((splits[1][0]+(dst - src), splits[1][1]+(dst - src)))
                    continue
                else:
                    split_seeds.update(splits)
            if not mapped:
                next_seeds.update(split_seeds)
        seeds = list(next_seeds)
    return min(seeds)[0]

inp = open("inp").read().split("\n\n")
seeds = list(map(int, inp.pop(0).split(" ")[1:]))
maps = []
for m in inp:
    parts = m.split()[2:]
    maps.append([list(map(int, parts[i : i + 3])) for i in range(0, len(parts), 3)])

print(p1(seeds, maps), p2(seeds, maps))
