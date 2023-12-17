import random

rand = random.Random(0)  # set seed
shit = "NSEWRLF"
for _ in range(10_000_000):
    sh = rand.choice(shit)
    if sh in "NSEWF":
        param = rand.randint(100, 1000)
    else:
        param = 90 * rand.randint(1, 3)
    print(sh + str(param))
