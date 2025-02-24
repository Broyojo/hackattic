import fileinput


def main():
    for line in fileinput.input():
        parts = line.strip().split()
        decoded = []
        for part in parts:
            if len(part) > 2 and part[0] == "0" and part[1] == "x":
                decoded.append(int(part, 16))
            elif len(part) > 2 and part[0] == "0" and part[1] == "b":
                decoded.append(int(part, 2))
            elif len(part) > 2 and part[0] == "0" and part[1] == "o":
                decoded.append(int(part, 8))
            elif part.isdecimal():
                decoded.append(int(part))
            elif part.isascii():
                decoded.append(ord(part))
        print(decoded)
        print(sum(decoded))


if __name__ == "__main__":
    main()
