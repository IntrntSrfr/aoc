inp = [x.strip().split(':')[1] for x in open('inp').readlines()]

def solve():
    p1_score = 0
    counts = [1] * len(inp)
    parts = [[set(part.split()) for part in card.split('|')] for card in inp]

    for i, (p1,p2) in enumerate(parts):
        matches = len(p1&p2)
        if matches > 0:
            p1_score += 2**(matches - 1)
        for j in range(matches):
            counts[i+j+1] += counts[i]
    return p1_score, sum(counts)


print(solve())
