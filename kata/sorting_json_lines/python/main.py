import fileinput
import json
import re

# This is not needed since the example code is wrong

# def extract_balance(line):
#     pattern = r'"balance":\s*\d+'

#     def get_num(field):
#         return int(field.split(":")[-1].strip())

#     matches = re.findall(pattern, line)
#     extra_pos = line.find("extra")
#     match_positions = [(line.find(match), match) for match in matches]

#     if extra_pos == -1:
#         return get_num(matches[0])

#     # find the match position which is right after the extra position
#     for pos, match in match_positions:
#         if pos > extra_pos:
#             return get_num(match)

#     raise ValueError("no extra pos")


# def main():
#     lines = [
#         (line.strip(), extract_balance(line.strip())) for line in fileinput.input()
#     ]

#     lines.sort(key=lambda l: l[1])

#     def read_name(line):
#         i = 2
#         name = ""
#         while line[i] != '"':
#             name += line[i]
#             i += 1
#         return name

#     for line, bal in lines:
#         name = read_name(line)
#         print(f"{name}: {bal:,}")


def main():
    lines = [json.loads(line) for line in fileinput.input()]
    augmented_lines = []
    for line in lines:
        balance = (
            line["extra"]["balance"]
            if "extra" in line
            else line[([k for k in line.keys() if k != "extra"][0])]["balance"]
        )
        augmented_lines.append((line, balance))

    augmented_lines.sort(key=lambda l: l[1])

    for line, balance in augmented_lines:
        name = [k for k in line.keys() if k != "extra"][0]
        print(f"{name}: {balance:,}")


if __name__ == "__main__":
    main()
