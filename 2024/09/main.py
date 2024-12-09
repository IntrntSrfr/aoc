inp = open('inp').read().strip()

lol: list[int] = []
for i, v in enumerate(inp):
    lol.extend([i//2]*int(v) if i%2==0 else [-1]*int(v))
    #id = i//2
    #lol.append f'{(i//2)}'*int(v) if i%2==0 else '*'*int(v)

def p1(inp: list[int]):
    #print(inp)
    inp1 = inp.copy()
    inp2 = []
    try:
        while -1 in inp:
            inp1 = inp.copy()
            #print(inp)
            x = inp.pop()
            while x == -1:
                x = inp.pop()
            # find first -1
            i = inp.index(-1)
            inp[i] = x
            #print(inp)
            inp2 = inp.copy()
        #print(inp)
    except Exception:
        print(inp1)
        print(inp2)
        
    
    print(sum(a*b for a, b in enumerate(inp)))

if __name__ == '__main__':
    p1(lol)
