from math import lcm

instr, nodes = [x.strip() for x in open('inp').read().split('\n\n')]
nodes = nodes.split('\n')

node_map = {}
for node in nodes:
    src, dst = node.split(' = ')
    node_map[src] = dst[1:-1].split(', ')

def get_steps(pos: str, tgt_pos: list[str]) -> int:
    idx, steps = 0, 0
    while pos not in tgt_pos:
        pos = node_map[pos][0 if instr[idx] == 'L' else 1]
        idx = (idx + 1) % len(instr)
        steps += 1
    return steps

def p1() -> int:
    return get_steps('AAA', ['ZZZ'])

def p2() -> int:
    factors: list[int] = []
    p2_src = list(filter(lambda x: x.endswith('A'), node_map.keys()))
    p2_tgt = list(filter(lambda x: x.endswith('Z'), node_map.keys()))

    factors = [get_steps(src, p2_tgt) for src in p2_src]
    return lcm(*factors)

print(p1(), p2())
