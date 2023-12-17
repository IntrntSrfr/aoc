inp = [[[int(n) for n in r.split('-')] for r in line.strip().split(',')] for line in open('inp')]

subsets = overlaps = 0
for line in inp:
  s1, s2 = set(range(line[0][0], line[0][1]+1)), set(range(line[1][0], line[1][1]+1))
  overlaps += any(s1 & s2)
  subsets += s1 <= s2 or s2 <= s1

print("part 1:", subsets)
print("part 2:", overlaps)