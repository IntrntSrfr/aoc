inp = [(d, int(n)) for line in open('inp') for d,n in [line.split()]]
print(inp)

dirs = {'R': (1, 0), 'L': (-1, 0), 'U': (0, 1), 'D': (0, -1)}
visited = {(0, 0)}

H, T = [0, 0], [0, 0]
for d, n in inp:
  dx, dy = dirs[d]
  for _ in range(n):
    H[0], H[1] = H[0]+dx, H[1]+dy
    if max(abs(H[0]-T[0]), abs(H[1] - T[1])) > 1:
      T[0], T[1] = H[0] - dx, H[1] - dy
    visited.update({(T[0], T[1])})
print(len(visited))