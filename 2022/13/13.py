import ast
from functools import cmp_to_key

def compare_lists(l1: list, l2: list):
    i = 0
    while l1[i:] and l2[i:] and (l := l1[i], r := l2[i]):
        i+=1
        if type(l) == int and type(r) == int:
            if l < r: return 1
            elif l > r: return -1
            else: continue
        if type(l) == list and not type(r) == list:
            r = [r]
        elif not type(l) == list and type(r) == list:
            l = [l]
        res = compare_lists(l, r)
        if res != 0: return res

    if not l1[i:] and l2[i:]:
        return 1
    if l1[i:] and not l2[i:]:
        return -1
    return 0

inp = [[ast.literal_eval(a), ast.literal_eval(b)] for line in 
    open('onAy.txt').read().split('\n\n') for a, b in [line.split()]]

indices = [i+1 for i, (l, r) in enumerate(inp) if compare_lists(l, r) > 0]
print(sum(indices))

f_inp = [p for pairs in inp for p in pairs]
f_inp.extend([[[2]], [[6]]])
f_inp.sort(key=cmp_to_key(compare_lists), reverse=True)
print((f_inp.index([[2]])+1) * (f_inp.index([[6]])+1))
