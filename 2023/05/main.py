inp = open("inp").read().split("\n\n")
seeds = list(map(int, inp.pop(0).split(" ")[1:]))
maps = []
for m in inp:
    parts = m.split()[2:]
    maps.append([list(map(int, parts[i : i + 3])) for i in range(0, len(parts), 3)])


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


def split(r1: range, r2: range):
    overlap_start = max(r1.start, r2.start)
    overlap_end = min(r1.stop, r2.stop)

    # Splitting into three parts
    before = range(r1.start, overlap_start) if overlap_start > r1.start else None
    overlap = range(overlap_start, overlap_end) if overlap_start < overlap_end else None
    after = range(overlap_end, r1.stop) if overlap_end < r1.stop else None

    return before, overlap, after


def combine(r1: range, r2: range):
    return (
        (min(r1.start, r2.start), max(r1.stop, r2.stop))
        if r1.start <= r2.stop and r2.start <= r1.stop
        else None
    )


def p2(seeds, maps):
    seeds = [range(seeds[x], seeds[x] + seeds[x + 1]) for x in range(0, len(seeds), 2)]
    
    for map in maps:
        next_seeds = []
        for seed in seeds:
            for dst, src, r in map:
                print(seed, dst, src, r)
                lol = split(range(src, src+r), seed)
                print(lol)
        seeds = next_seeds
    print(seeds)


print(p1(seeds, maps))
p2(seeds, maps)
