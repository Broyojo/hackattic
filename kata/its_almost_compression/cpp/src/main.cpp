#include <iostream>

std::string compress(std::string str) {
    if (str.length() == 0) {
        return str;
    }

    char curr = str[0];
    int count = 1;
    std::string compressed = "";

    for (int i = 1; i < str.length(); i++) {
        if (str[i] != curr) {
            if (count > 2) {
                compressed += std::to_string(count) + curr;
            } else {
                for (int j = 0; j < count; j++) {
                    compressed += curr;
                }
            }
            curr = str[i];
            count = 1;
        } else {
            count++;
        }
    }

    if (count > 2) {
        compressed += std::to_string(count) + curr;
    } else {
        for (int j = 0; j < count; j++) {
            compressed += curr;
        }
    }

    return compressed;
}

int main() {
    for (std::string line; std::getline(std::cin, line);) {
        std::cout << compress(line) << std::endl;
    }
}
