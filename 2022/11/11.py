import math
import copy

monkeys = []
for m in open('inp').read().split('\n\n'):
    parts = m.splitlines()
    parts = [p.split(':')[1].strip() for p in parts]
    items = [int(it) for it in parts[1].split(', ')]
    mult = parts[2].split()[-2] == '*'
    num = parts[2].split()[-1]
    num = int(num) if num.isdigit() else 0 
    test = int(parts[3].split()[-1])
    test_res = [int(parts[4].split()[-1]), int(parts[5].split()[-1])] # [true, false]
    monkeys.append({'items':items, 'mult':mult, 'num':num, 'test':test, 'test_res':test_res})

prod = math.prod([m['test'] for m in monkeys])
def solve(monkeys, n):
    t = [0 for _ in range(len(monkeys))]
    for i in range(n):
        for i, monkey in enumerate(monkeys):
            while len(monkey['items']) and (item := monkey['items'].pop(0)):
                num = monkey['num'] if monkey['num'] else item
                item = item * num if monkey['mult'] else item + num
                item = item // 3 if n==20 else item%prod
                if item%monkey['test'] == 0:
                    monkeys[monkey['test_res'][0]]['items'].append(item)
                else:
                    monkeys[monkey['test_res'][1]]['items'].append(item)
                t[i]+=1
    return math.prod(sorted(t)[-2:])

print("part 1:", solve(copy.deepcopy(monkeys), 20))
print("part 2:", solve(copy.deepcopy(monkeys), 10000))
