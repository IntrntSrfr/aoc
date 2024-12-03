from math import prod
import re
inp = open('inp').read().strip()

mul_re = re.compile(r'mul\((\d+),(\d+)\)')
do_re = re.compile(r'do\(\)')
dont_re = re.compile(r'don\'t\(\)')

def p1(inp: str):
    print(sum(int(x)*int(y) for x,y in mul_re.findall(inp)))

def p2(inp: str):
    dos = [(0, True)] + [(x.start(), True) for x in do_re.finditer(inp)]
    donts = [(x.start(), False) for x in dont_re.finditer(inp)]
    
    l = sorted(dos+donts)
    do, res = l.pop(0)[1], 0
    for i in mul_re.finditer(inp):
        while l and l[0][0] < i.start():
            do = l.pop(0)[1]
        res += prod(int(x) for x in i.groups()) if do else 0
    print(res)

if __name__ == '__main__':
    p1(inp)
    p2(inp)
