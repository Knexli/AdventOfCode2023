//
// Created by Andreas on 04.12.2023.
//

#include <iostream>
#include <string>
#include <unordered_set>

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

void solve1() {
    int res = 0;
    while (true) {
        std::string s, z;
        std::getline(std::cin, s);
        if (s.empty()) {
            break;
        }
        getNextToken(s);
        getNextToken(s);
        std::unordered_set<int> winningNums;
        int c = 0;
        bool inWinningNums = true;
        while (!(z = getNextToken(s)).empty()) {
            if (z == "|") {
                inWinningNums = false;
                continue;
            }
            if (inWinningNums) {
                winningNums.insert(std::stoi(z));
            } else {
                if (winningNums.find(std::stoi(z)) != winningNums.end()) {
                    c = c == 0 ? 1 : c * 2;
                }
            }
        }
        res += c;
    }
    std::cout << res << std::endl;
}

void solve2() {
    int res = 0;
    int dp[1000];
    for (int i = 0; i < 1000; ++i) {
        dp[i] = 1;
    }
    int cardCounter = 0;
    while (true) {
        std::string s, z;
        std::getline(std::cin, s);
        if (s.empty()) {
            break;
        }
        getNextToken(s);
        getNextToken(s);
        std::unordered_set<int> winningNums;
        int c = 0;
        bool inWinningNums = true;
        while (!(z = getNextToken(s)).empty()) {
            if (z == "|") {
                inWinningNums = false;
                continue;
            }
            if (inWinningNums) {
                winningNums.insert(std::stoi(z));
            } else {
                if (winningNums.find(std::stoi(z)) != winningNums.end()) {
                    c++;
                }
            }
        }
        for (int i = cardCounter + 1; i <= cardCounter + c; i++) {
            dp[i] += dp[cardCounter];
        }
        res += dp[cardCounter];
        cardCounter++;
    }
    std::cout << res << std::endl;
}

int main() {
//    solve1();
    solve2();
}