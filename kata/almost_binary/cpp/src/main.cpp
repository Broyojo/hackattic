#include <iostream>

int main() {
    for (std::string line; std::getline(std::cin, line);) {
        uint16_t num = 0;
        for (std::string::iterator c = line.begin(); c != line.end(); c++) {
            switch (*c) {
                case '#':
                    num = num << 1 | 1;
                    break;
                case '.':
                    num = num << 1;
                    break;
                default:
                    std::cerr << "unknown character: " << *c << std::endl;
                    exit(1);
            }
        }
        std::cout << num << std::endl;
    }
}
