#include <iostream>

int main() {
    for (std::string line; std::getline(std::cin, line);) {
        int offset = std::stoi(line);
        int weekday = (offset + 4) % 7;
        if (weekday < 0) {
            weekday += 7;
        }
        std::string weekdays[] = {
            "Sunday",
            "Monday",
            "Tuesday",
            "Wednesday",
            "Thursday",
            "Friday",
            "Saturday",
        };
        std::cout << weekdays[weekday] << std::endl;
    }
}
