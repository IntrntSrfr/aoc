from functools import cmp_to_key
from collections import Counter

cards = "23456789TJQKA"

def hand_type(card):
    counts = sorted(Counter(card).values())
    if 5 in counts:
        return 6
    elif 4 in counts:
        return 5
    elif counts == [2, 3]:
        return 4
    elif 3 in counts:
        return 3
    elif counts.count(2) == 2:
        return 2
    elif 2 in counts:
        return 1
    else:
        return 0

def hand_type2(card):
    jokers = card.count("J")
    card = [c for c in card if c != "J"]
    counts = sorted(Counter(card).values(), reverse=True)
    if not counts:
        counts = [0]

    if counts[0] + jokers == 5:
        return 6
    elif counts[0] + jokers == 4:
        return 5
    elif counts[0] + jokers == 3 and counts[1] == 2:
        return 4
    elif counts[0] + jokers == 3:
        return 3
    elif counts[0] == 2 and (jokers or counts[1] == 2):
        return 2
    elif counts[0] == 2 or jokers:
        return 1
    else:
        return 0


def comp_cards(c1, c2):
    for i in range(5):
        if c1[0][i] == c2[0][i]:
            continue
        return -1 if cards.index(c1[0][i]) < cards.index(c2[0][i]) else 1
    return 0

def comp_func(part):
    def comp_hands(c1, c2) -> int:
        ht1, ht2 = (
            (hand_type(c1[0]), hand_type(c2[0]))
            if part == 1
            else (hand_type2(c1[0]), hand_type2(c2[0]))
        )
        if ht1 == ht2:
            return comp_cards(c1, c2)
        return -1 if ht1 < ht2 else 1

    return comp_hands

inp = [(x.split()[0], int(x.split()[1])) for x in open("inp").readlines()]
s1 = sorted(inp, key=cmp_to_key(comp_func(1)))
a1 = sum([(i + 1) * x[1] for i, x in enumerate(s1)])

cards = "J23456789TQKA"
s2 = sorted(inp, key=cmp_to_key(comp_func(2)))
a2 = sum([(i + 1) * x[1] for i, x in enumerate(s2)])
print(a1, a2)
# 253313241 253362743
