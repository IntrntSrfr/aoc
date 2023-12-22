from functools import cache

def get_solutions(row: str, b):
    def is_valid(s, b):
        sizes = [len(block) for block in "".join(s).split(".") if block]
        return sizes == b

    @cache
    def backtrack(i, cur: tuple, b):
        if i == len(row):
            if is_valid(cur, b):
                s.append("".join(cur))
            return
        cur_copy = list(cur)
        if row[i] == "?":
            cur_copy[i] = "#"
            backtrack(i + 1, tuple(cur_copy), b)
            cur_copy[i] = "."
            backtrack(i + 1, tuple(cur_copy), b)
        else:
            cur_copy[i] = row[i]
            backtrack(i + 1, tuple(cur_copy), b)

    s = []
    backtrack(0, tuple(row), b)
    return len(s)

inp = [x.strip() for x in open("inp").readlines()]

score = 0
for line in inp:
    pattern, nums_str = line.split()
    pattern = "#".join([pattern] * 3)
    nums_str = ",".join([nums_str] * 3)
    nums = [int(x) for x in nums_str.split(",")]
    # print(pattern, nums)
    # print('#'.join([pattern]*5))
    s = get_solutions(pattern, tuple(nums))
    # print(s)
    score += s
print(score)
