from functools import cmp_to_key

card_values = ['2', '3', '4', '5', '6', '7', '8', '9', 'T', 'J', 'Q', 'K', 'A']

inp = [x.strip().split() for x in open('inp').readlines()]
print(inp)

def comp_cards(c1, c2) -> int:
    return 0

sorted_cards = sorted(inp, key=cmp_to_key(comp_cards))
