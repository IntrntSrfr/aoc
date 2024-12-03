inp = [list(map(int, line.split())) for line in open('inp')]

def is_safe(report: list[int]) -> bool:
    diffs = [j-i for i,j in zip(report, report[1:])]
    return all(abs(d) <= 3 for d in diffs) and \
        (all(d < 0 for d in diffs) or 
         all(d > 0 for d in diffs))
    
def p1(inp: list[list[int]]):
    print(sum([is_safe(l) for l in inp]))

def p2(inp: list[list[int]]):
    print(sum([is_safe(l) or \
        any([is_safe(l[:i] + l[i+1:]) for i in range(len(l))]) for l in inp]))

if __name__ == '__main__':
    p1(inp)
    p2(inp)
