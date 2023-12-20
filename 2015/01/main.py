inp = [x for x in open('inp').read()]
count = 0
for i, step in enumerate( inp):
    count += 1 if step == '(' else -1
    if count == -1:
        print(i+1)
print(count)