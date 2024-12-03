import re
inp = open("inp").read().strip()

mul_re = re.compile(r"mul\((\d+),(\d+)\)")
do_re = re.compile(r"do\(\)")
dont_re = re.compile(r"don\'t\(\)")

def p1(inp: str):
    print(sum(int(x) * int(y) for x, y in mul_re.findall(inp)))

def p2(inp: str):
    dos = [l.split("don't()")[0] for l in inp.split("do()")]
    print(sum([int(x) * int(y) for d in dos for x, y in mul_re.findall(d)]))

if __name__ == "__main__":
    p1(inp)
    p2(inp)
