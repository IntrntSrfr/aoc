from math import prod

inp = [
    [
        [item.strip() for item in column.split(',')] 
        for column in row.split(';')
    ] 
    for row in (game.split(':')[1] 
    for game in open('inp').readlines())
]

def solve(games):
    p1_score = 0
    p2_score = 0
    for i, game in enumerate(games, 1):
        total_vals = {"red": 0, "green": 0, "blue": 0}
        good = True
        max_vals = {"red": 12, "green": 13, "blue": 14}
        for hand in game:
            for color in hand:
                n, c = color.split()
                total_vals[c] = max(total_vals[c], int(n))
                if int(n) > max_vals[c]:
                    good = False
        if good:
            p1_score += (i)
        p2_score += prod(total_vals.values())
    return p1_score, p2_score

print(solve(inp))
