inp = open('inp').read().strip()

lol: list[int] = []
for i, v in enumerate(inp):
    lol.extend([i//2]*int(v) if i%2==0 else [-1]*int(v))
    #id = i//2
    #lol.append f'{(i//2)}'*int(v) if i%2==0 else '*'*int(v)

def p1(inp: list[int]):
    while -1 in inp:
        x = inp.pop()
        i = inp.index(-1)
        inp[i] = x
        while inp[-1] == -1:
            inp.pop()
    print(sum(a*b for a, b in enumerate(inp)))

if __name__ == '__main__':
    p1(lol)
