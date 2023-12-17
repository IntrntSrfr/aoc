inp = open('bigboy.txt').read()

def solve(n):
  for i in range(len(inp)-n-1):
    if len(set(inp[i:i+n])) == n:
      return i+n

print(solve(4))
print(solve(14))