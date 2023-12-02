//
// Created by Andreas on 02.12.2023.
//

#include <iostream>
#include <string>

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
    return token;
}

bool validate(int& r, int& g, int& b) {
    bool res = r <= 12 && g <= 13 && b <= 14;
    r = 0;
    g = 0;
    b = 0;
    return res;
}

void validate2(int& r, int& g, int& b, int& minR, int& minG, int& minB) {
    minR = std::max(minR, r);
    minG = std::max(minG, g);
    minB = std::max(minB, b);
    r = 0;
    g = 0;
    b = 0;
}

void solve1() {
    std::string s;
    int c = 0;
    while (std::getline(std::cin, s)) {
        if (s.empty()) {
            break;
        }
        std::string z;
        getNextToken(s);
        z = getNextToken(s);
        int gamenum = std::stoi(z.substr(0, z.size() - 1));
        int r = 0, g = 0, b = 0;
        std::string y;
        bool addres = true;
        while (!(z = getNextToken(s)).empty()) {
            y = getNextToken(s);
            if (y.substr(0,3) == "red") {
                r += std::stoi(z);
            } else if (y.substr(0,3) == "blu") {
                b += std::stoi(z);
            } else if (y.substr(0, 3) == "gre") {
                g += std::stoi(z);
            }

            if (y.substr(y.length() - 1) == ";") {
                if (!validate(r, g, b)) {
                    addres = false;
                }
            }

        }
        if (!validate(r, g, b)) {
            addres = false;
        }
        if (addres) {
            c += gamenum;
        }
    }
    std::cout << c << std::endl;
}

void solve2() {
    std::string s;
    int c = 0;
    while (std::getline(std::cin, s)) {
        if (s.empty()) {
            break;
        }
        std::string z;
        getNextToken(s);
        z = getNextToken(s);
        int gamenum = std::stoi(z.substr(0, z.size() - 1));
        int r = 0, g = 0, b = 0, minR = 0, minG = 0, minB = 0;
        std::string y;
        bool addres = true;
        while (!(z = getNextToken(s)).empty()) {
            y = getNextToken(s);
            if (y.substr(0,3) == "red") {
                r += std::stoi(z);
            } else if (y.substr(0,3) == "blu") {
                b += std::stoi(z);
            } else if (y.substr(0, 3) == "gre") {
                g += std::stoi(z);
            }

            if (y.substr(y.length() - 1) == ";") {
                validate2(r, g, b, minR, minG, minB);
            }

        }
        validate2(r, g, b, minR, minG, minB);
        c += minR * minG * minB;
    }
    std::cout << c << std::endl;
}


int main() {
    //solve1();
    solve2();
}