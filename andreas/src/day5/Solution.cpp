//
// Created by Andreas on 05.12.2023.
//

#include <iostream>
#include <string>
#include <vector>

std::string getNextToken(std::string& s) {
    size_t pos = 0;
    std::string token;
    if ((pos = s.find(" ")) == std::string::npos) {
        if (s.length() > 0) {
            std::string z = s;
            s = "";
            return z;
        }
        return "";
    }
    token = s.substr(0, pos);
    s.erase(0, pos + 1);
    while (s[0] == ' ') {
        s.erase(0, 1);
    }
    return token;
}

// you might need a little bit of patience if you want to run this :)
void solve2() {
    long long seeds[1000];
    for (int i = 0; i < 1000; i++) {
        seeds[i] = -1;
    }
    std::string s, z;
    std::getline(std::cin, s);
    getNextToken(s);
    int c = 0;
    while (!(z = getNextToken(s)).empty()) {
        seeds[c++] = (std::stoll(z));
    }
    std::vector<std::vector<long long *>> maps;
    std::getline(std::cin, s);
    for (int i = 0; i < 7; i++) {
        std::getline(std::cin, s);
        std::vector<long long *> map;
        while (true) {
            std::getline(std::cin, s);
            if (s.empty()) {
                break;
            }
            auto* line = (long long *)malloc(3 * sizeof(long));
            for (int k = 0; k < 3; k++) {
                z = getNextToken(s);
                line[k] = std::stoll(z);
            }
            map.push_back(line);
        }
        maps.push_back(map);
    }

    long long min = LONG_LONG_MAX;
    int d = 0;
    for (int i = 0; i < 1000; i++) {
        if (seeds[i] == -1) {
            break;
        }
        for (long long u = seeds[i]; u < seeds[i] + seeds[i + 1]; u++) {
            long long seed = u;
            for (const auto &map: maps) {
                for (long long *line: map) {
                    if (seed - line[1] < line[2] && seed >= line[1]) {
                        // correct range
                        seed = seed - line[1] + line[0];
                        break;
                    }
                }
            }
            if (u % 1000000 == 0) {
                std::cout << "feasability check " << d++ << std::endl;
            }
            min = std::min(min, seed);
        }
        std::cout << "done with iteration " << i << std::endl;
        i++;
    }

    for (auto map: maps) {
        for (long long *line: map) {
            free(line);
        }
    }

    std::cout << min << std::endl;
}


void solve1() {
    std::vector<long long> seeds;
    std::string s, z;
    std::getline(std::cin, s);
    getNextToken(s);
    while (!(z = getNextToken(s)).empty()) {
        seeds.push_back(std::stoll(z));
    }
    std::vector<std::vector<long long *>> maps;
    std::getline(std::cin, s);
    for (int i = 0; i < 7; i++) {
        std::getline(std::cin, s);
        std::vector<long long *> map;
        while (true) {
            std::getline(std::cin, s);
            if (s.empty()) {
                break;
            }
            auto* line = (long long *)malloc(3 * sizeof(long));
            for (int k = 0; k < 3; k++) {
                z = getNextToken(s);
                line[k] = std::stoll(z);
            }
            map.push_back(line);
        }
        maps.push_back(map);
    }

    long long min = LONG_LONG_MAX;
    for (long long u: seeds) {
        long long seed = u;
        for (const auto& map: maps) {
            for (long long *line: map) {
                if (seed - line[1] < line[2] && seed >= line[1]) {
                    // correct range
                    seed = seed - line[1] + line[0];
                    break;
                }
            }
        }

        min = std::min(min, seed);
    }

    for (auto map: maps) {
        for (long long *line: map) {
            free(line);
        }
    }

    std::cout << min << std::endl;
}

int main() {
    solve2();
}