#include <iostream>

bool check(std::string parens) {
    int count = 0;
    for (std::string::iterator c = parens.begin(); c != parens.end(); c++) {
        switch (*c) {
            case '(':
                count++;
                break;
            case ')':
                if (count <= 0) {
                    return false;
                }
                count--;
                break;
            default:
                std::cerr << "error: unknown character: " << *c << std::endl;
                exit(1);
        }
    }
    return count == 0;
}

int main() {
    for (std::string line; std::getline(std::cin, line);) {
        if (check(line)) {
            std::cout << "yes";
        } else {
            std::cout << "no";
        }
        std::cout << std::endl;
    }
}
