import random
numbers = 20
random.seed(2021)
l = [random.choice([0, 1]) for x in range(numbers)]
print(l)
