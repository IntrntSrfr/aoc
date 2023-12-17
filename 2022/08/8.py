inp = [[int(j) for j in i.strip()] for i in open('inp')]
rows, cols = len(inp), len(inp[0])

max_s = 0
visible = 0
for r in range(1, rows-1):
  for c in range(1, cols-1):
    s = 1
    vis = False
    for dr, dc in [(0, 1), (0, -1), (1, 0), (-1, 0)]:
      cr, cc, i = r+dr, c+dc, 0
      while 0 <= cr < rows and 0 <= cc < cols:
        i += 1
        if inp[r][c] <= inp[cr][cc]:
          break
        if cr <= 0 or cr >= rows-1 or cc <= 0 or cc >= cols-1:
          vis = True
        cr, cc = cr+dr, cc+dc
      s *= i
    max_s = max(s, max_s)
    visible += 1 if vis else 0
      

print("part 1:", visible+4*len(inp)-4)
print("part 2:", max_s)