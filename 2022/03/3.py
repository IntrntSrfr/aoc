import string
inp = [line.strip() for line in open('inp')]

shared = []
for line in inp:
  l, r = set(line[:len(line)//2]), set(line[len(line)//2:])
  shared.append((l & r).pop())
print("part 1:", sum([string.ascii_letters.index(x)+1 for x in shared]))

shared = []
for grp in range(0, len(inp), 3):
  g1, g2, g3 = set(inp[grp]), set(inp[grp+1]), set(inp[grp+2])
  shared.append((g1 & g2 & g3).pop())
print("part 2:", sum([string.ascii_letters.index(x)+1 for x in shared]))