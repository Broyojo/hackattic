import fileinput


def main():
    for line in fileinput.input():
        snake_case = ""
        first_upper_case = False
        for i, c in enumerate(line.strip()):
            if c.isupper():
                if first_upper_case:
                    snake_case += "_"
                if not first_upper_case and i > 3:
                    snake_case += line[:i] + "_"  # add back in prefix
                snake_case += c.lower()
                first_upper_case = True
            elif first_upper_case:
                snake_case += c
        print(snake_case)


if __name__ == "__main__":
    main()
