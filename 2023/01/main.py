
name_to_int = {
    "one": "1",
    "two": "2",
    "three": "3",
    "four": "4",
    "five": "5",
    "six": "6",
    "seven": "7",
    "eight": "8",
    "nine" : "9"
}

def p1(word: str, rev: bool = False):
    for i in reversed(word) if rev else word:
        if i.isdigit():
            return i
    return ''

def replace_keys_in_string(s, replacements):
    new_string = ""
    i = 0
    while i < len(s):
        replaced = False
        for key, value in replacements.items():
            if s[i:i+len(key)] == key:
                new_string += value
                i += len(key)
                replaced = True
                break
        if not replaced:
            new_string += s[i]
            i += 1
    return new_string

def replace_keys_in_string_backwards(s, replacements):
    new_string = ""
    i = len(s)
    while i > 0:
        replaced = False
        for key, value in replacements.items():
            key_length = len(key)
            if i - key_length >= 0 and s[i-key_length:i] == key:
                new_string = value + new_string
                i -= key_length
                replaced = True
                break
        if not replaced:
            i -= 1
            new_string = s[i] + new_string
    return new_string

def p2(word: str, rev: bool = False) -> str:
    if rev:
        word = replace_keys_in_string_backwards(word, name_to_int)
    else:
        word = replace_keys_in_string(word, name_to_int)
    return p1(word, rev)

inp = [x.strip() for x in open('inp', 'r').readlines()]

lol = [int(''.join([p2(w), p2(w, True)])) for w in inp]
print(sum(lol))
