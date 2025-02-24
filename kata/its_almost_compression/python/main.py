import fileinput


def main():
    for line in fileinput.input():
        current_char = ""
        current_count = 0
        compressed = ""

        def add(compressed, current_count, current_char):
            if current_count > 2:
                return compressed + f"{current_count}{current_char}"
            else:
                return compressed + current_char * current_count

        for c in line:
            if c != current_char:
                if current_char != "":
                    compressed = add(compressed, current_count, current_char)
                current_char = c
                current_count = 1
            else:
                current_count += 1
        if current_char != "\n":
            compressed = add(compressed, current_count, current_char)
        print(compressed)


if __name__ == "__main__":
    main()
