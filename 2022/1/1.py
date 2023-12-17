l = sorted([sum([int(l) for l in sec.splitlines()]) for sec in open('./inp').read().split('\n\n')])

print("part 1:", l[-1])
print("part 2:", sum(l[-3:]))
