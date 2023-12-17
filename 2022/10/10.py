inp = [x.split() for x in open('inp')]
#print(inp)

p1 = 0
r, c = 1, 1
for inst in inp:
    c+=1
    if c == 20 or (c-20)%40==0:
        p1 += c*r
    #    print(c, r, c*r)

    if len(inst) < 2:
        continue


    _, n = inst[0], int(inst[1])
    r, c = r+n, c+1

    if c == 20 or (c-20)%40==0:
        p1 += c*r
        #print(c, r, c*r)

print(p1)
