import copy

inp = [sec for sec in open('bigboy.txt').read().split('\n\n')]

start = inp[0].splitlines()
stacks = [[] for _ in start[-1].split()]
inst = [[int(a) for a in line.split() if a.isnumeric()] for line in inp[1].splitlines()]

for z in reversed(start[:-1]):
  z_crates = [c[1] for c in [z[i:i+4] for i in range(0, len(z), 4)]]
  [stacks[i].append(c) for i,c in enumerate(z_crates) if c != " "]

def p1(stacks):
  for n, f, t in inst:
    [stacks[t-1].append(stacks[f-1].pop()) for _ in range(n)]
  return ''.join([s[-1] for s in stacks])

def p2(stacks):
  for n, f, t in inst:
    stacks[t-1].extend(stacks[f-1][-n:])
    stacks[f-1] = stacks[f-1][:-n]
  return ''.join([s[-1] for s in stacks])

def main(p):
  res = p1(copy.deepcopy(stacks)) if p == 1 else p2(copy.deepcopy((stacks)))
  print("part {}: {}".format(p, res))

main(1)
main(2)
