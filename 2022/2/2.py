inp = [l.strip() for l in open('inp')]

t = {'A X': 4, 'A Y': 8, 'A Z': 3, 
     'B X': 1, 'B Y': 5, 'B Z': 9, 
     'C X': 7, 'C Y': 2, 'C Z': 6}
print("part 1:", sum(t[l] for l in inp))

t2 = {'A X': 3, 'A Y': 4, 'A Z': 8, 
      'B X': 1, 'B Y': 5, 'B Z': 9, 
      'C X': 2, 'C Y': 6, 'C Z': 7}
print("part 2:", sum(t2[l] for l in inp))

