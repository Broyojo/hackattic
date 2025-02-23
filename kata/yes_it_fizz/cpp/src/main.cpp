#include <iostream>

void fizzbuzz(int n, int m) {
    for (int i = n; i <= m; i++) {
        if (i % 3 == 0 && i % 5 == 0) {
            std::cout << "FizzBuzz";
        } else if (i % 3 == 0) {
            std::cout << "Fizz";
        } else if (i % 5 == 0) {
            std::cout << "Buzz";
        } else {
            std::cout << i;
        }

        std::cout << std::endl;
    }
}

int main() {
    std::string input;
    std::getline(std::cin, input);

    int n = std::stoi(input.substr(0, input.find(" ")));
    int m = std::stoi(input.substr(input.find(" ") + 1, input.length()));

    fizzbuzz(n, m);
}