#include <iostream>
#include <vector>
#include <cstring>

const char *BASE64_CHARSET = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/";

std::string base64_decode(std::string str) {
    std::vector<char> decoded;
    int buffer = 0;
    int bits_collected = 0;

    for (std::string::iterator it = str.begin(); it != str.end(); it++) {
        if (*it != '=') {
            int position = -1;
            for (int i = 0; i < strlen(BASE64_CHARSET); i++) {
                if (*it == BASE64_CHARSET[i]) {
                    position = i;
                    break;
                }
            }
            if (position == -1) {
                std::cerr << "invalid character: '" << *it << "'" << std::endl;
                exit(1);
            }

            buffer = buffer << 6 | position;
            bits_collected += 6;
            while (bits_collected >= 8) {
                bits_collected -= 8;
                char byte = (buffer >> bits_collected) & 0xFF;
                decoded.push_back(byte);
            }
        }
    }

    std::string s(decoded.begin(), decoded.end());
    return s;
}

int main() {
    for (std::string line; std::getline(std::cin, line);) {
        std::string result = base64_decode(line);
        std::cout << result << std::endl;
    }
}
