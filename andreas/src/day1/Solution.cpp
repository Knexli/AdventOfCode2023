//
// Created by Andreas on 02.12.2023.
//

#include "Solution.h"

#include <iostream>
#include <string>
#include <map>

void solve1() {
    std::string s;
    int c = 0;
    while (std::getline(std::cin, s)) {
        if (s.empty()) {
            break;
        }
        int f = -1, l = -1;
        for (char &a: s) {
            if ('0' <= a && a <= '9') {
                if (f == -1) {
                    f = a - '0';
                }
                l = a - '0';
            }
        }
        c += 10 * f + l;
    }
    std::cout << c;
}

void solve2() {
    std::map<std::string, int> map;
    map["one"] = 1;
    map["two"] = 2;
    map["three"] = 3;
    map["four"] = 4;
    map["five"] = 5;
    map["six"] = 6;
    map["seven"] = 7;
    map["eight"] = 8;
    map["nine"] = 9;

    std::string s;
    int c = 0;
    while (std::getline(std::cin, s)) {
        if (s.empty()) {
            break;
        }
        int v = -1;
        int f = -1, l = -1;
        for (int i = 0; i < s.size(); i++) {
            if ('0' <= s[i] && s[i] <= '9') {
                v = s[i] - '0';
            }
            for (auto const &x: map) {
                if (i + x.first.size() > s.size()) {
                    continue;
                }
                if (s.find(x.first, i) == i) {
                    v = x.second;
                }
            }
            if (v != -1) {
                if (f == -1) {
                    f = v;
                }
                l = v;
            }
        }
        c += 10 * f + l;
    }

    std::cout << c;
}

int main() {
    //solve1();
    solve2();
}