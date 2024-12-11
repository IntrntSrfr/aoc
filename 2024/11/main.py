from collections import Counter
from functools import cache

inp = [x for x in open("inp").read().split()]

@cache
def t(n):
    if n == "0":
        return ["1"]
    elif len(n) % 2 == 0:
        return [str(int(n[: len(n) // 2])), str(int(n[len(n) // 2 :]))]
    else:
        return [str(int(n) * 2024)]

c = Counter(inp)
for i in range(75):
    c = Counter({stone: freq for key, freq in c.items() for stone in t(key)})
    if i + 1 in {25, 75}:
        print(sum(c.values()))
