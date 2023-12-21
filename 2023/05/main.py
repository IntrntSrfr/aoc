inp = open('inp').read().split('\n\n')
#print(inp)

seeds = list(map(int, inp.pop(0).split(' ')[1:]))
#print(seeds)

maps = []
for m in inp:
    parts = m.split()[2:]
    maps.append([list(map(int, parts[i:i+3])) for i in range(0,len(parts),3)])

for i, map in enumerate(maps):
    next_seeds = []
    #print(seeds)
    for seed in seeds:
        mapped = False
        for part in map:
            if seed in range(part[1], part[1]+part[2]):
                #print(seed, part)
                next_seeds.append(seed+(part[0]-part[1]))
                mapped = True
                continue
        if not mapped:
            next_seeds.append(seed)
    seeds = next_seeds
print(min(seeds))
            