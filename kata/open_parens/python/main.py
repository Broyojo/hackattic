import fileinput


def check(parens):
    count = 0
    for c in parens:
        if c == "(":
            count += 1
        else:
            if count <= 0:
                return False
            count -= 1
    return count == 0


def main():
    for line in fileinput.input():
        if check(line.strip()):
            print("yes")
        else:
            print("no")


if __name__ == "__main__":
    main()
