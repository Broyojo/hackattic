import fileinput
from datetime import datetime, timedelta


def main():
    for line in fileinput.input():
        date = datetime.fromtimestamp(0) + timedelta(days=int(line), hours=5)
        weekdays = [
            "Monday",
            "Tuesday",
            "Wednesday",
            "Thursday",
            "Friday",
            "Saturday",
            "Sunday",
        ]
        print(weekdays[date.weekday()])


if __name__ == "__main__":
    main()
