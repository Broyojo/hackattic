import base64
import fileinput


def main():
    for line in fileinput.input():
        decoded = base64.standard_b64decode(line).decode("utf-8")
        print(decoded)


if __name__ == "__main__":
    main()
