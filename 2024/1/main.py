from collections import Counter as c

l1,l2 = zip(*[map(int, i.split()) for i in open('inp').readlines()])

def p1(l1: list, l2: list):
    print(sum([abs(x - y) for x, y in zip(sorted(l1), sorted(l2))]))

def p2(l1: list, l2: list):
    print(sum([x * c(l2)[x] for x in l1]))
    
if __name__ == '__main__':
    p1(l1, l2)
    p2(l1, l2)
