//
// Created by Andreas on 03.12.2023.
//

#include <iostream>
#include <string>
#include <valarray>

bool isNum(char c) {
    return '0' <= c && c <= '9';
}

bool hasNeighborChar(char **map, int l, int i, int j0, int j1) {
    int kx[] = {0, 1, 0, -1, 1, 1, -1, -1};
    int ky[] = {1, 0, -1, 0, 1, -1, 1, -1};
    for (int j = j0; j <= j1; j++) {
        for (int k = 0; k < 8; k++) {
            int x = kx[k] + j;
            int y = ky[k] + i;
            if (y < 0 || x < 0 || x >= l) {
                continue;
            }
            char s = map[y][x];
            if (s != '.' && !isNum(s)) {
                return true;
            }
        }
    }
    return false;
}

int computeNeighborChar(char **map, int **intmap, int l, int i, int j0, int j1, int res) {
    int kx[] = {0, 1, 0, -1, 1, 1, -1, -1};
    int ky[] = {1, 0, -1, 0, 1, -1, 1, -1};
    for (int j = j0; j <= j1; j++) {
        for (int k = 0; k < 8; k++) {
            int x = kx[k] + j;
            int y = ky[k] + i;
            if (y < 0 || x < 0 || x >= l) {
                continue;
            }
            char s = map[y][x];
            if (s == '*') {
                if (intmap[y][x] == -1) {
                    intmap[y][x] = res;
                    return -1;
                }
                return intmap[y][x];
            }
        }
    }
    return -1;
}

int computenum(char** map, int l, int i, int startJ, int j) {
    if (hasNeighborChar(map, l, i, startJ, j - 1)) {
        int res = 0;
        for (int z = startJ; z < j; z++) {
            int num = (map[i][z] - '0');
            for (int v = z; v < j - 1; v++) {
                num *= 10;
            }
            res += num;
        }
        return res;
    }
    return 0;
}

// this would absolutely break if there were more than 2 numbers adjacent to a *
int computenum2(char** map, int** intmap, int l, int i, int startJ, int j) {
    int res = 0;
    for (int z = startJ; z < j; z++) {
        int num = (map[i][z] - '0');
        for (int v = z; v < j - 1; v++) {
            num *= 10;
        }
        res += num;
    }
    int v = computeNeighborChar(map, intmap, l, i, startJ, j - 1, res);
    if (v != -1) {
        return res * v;
    }
    return 0;
}

void solve1() {
    std::string s;
    std::getline(std::cin, s);
    int l = s.length();
    char **map = (char **)(malloc(1000 * sizeof(char *)));
    for (int i = 0; i < 1000; i++) {
        map[i] = (char *)malloc(l * sizeof(char));
    }
    for (int i = 0; i < 1000; i++) {
        for (int j = 0; j < l; j++) {
            map[i][j] = '.';
        }
    }
    for (int j = 0; j < l; j++) {
        map[0][j] = s[j];
    }
    int k = 1;
    while (std::getline(std::cin, s)) {
        if (s.empty()) {
            break;
        }
        for (int j = 0; j < l; j++) {
            map[k][j] = s[j];
        }
        k++;
    }

    int res = 0;

    for (int i = 0; i < 1000; i++) {
        int startJ = -1;
        for (int j = 0; j < l; j++) {
            if (isNum(map[i][j])) {
                if (startJ == -1) {
                    startJ = j;
                }
            } else {
                if (startJ != -1) {
                    int n = computenum(map, l, i, startJ, j);
                    res += n;
                    startJ = -1;
                }
            }
        }
        if (startJ != -1) {
            int n = computenum(map, l, i, startJ, l);
            res += n;
        }
    }

    for (int i = 0; i < 1000; i++) {
        free(map[i]);
    }
    free(map);
    std::cout << res;
}

void solve2() {
    std::string s;
    std::getline(std::cin, s);
    int l = s.length();
    char **map = (char **)(malloc(1000 * sizeof(char *)));
    int **intmap = (int **)(malloc(1000 * sizeof(int *)));
    for (int i = 0; i < 1000; i++) {
        map[i] = (char *)malloc(l * sizeof(char));
        intmap[i] = (int *)malloc(l * sizeof(int));
    }
    for (int i = 0; i < 1000; i++) {
        for (int j = 0; j < l; j++) {
            map[i][j] = '.';
            intmap[i][j] = -1;
        }
    }
    for (int j = 0; j < l; j++) {
        map[0][j] = s[j];
    }
    int k = 1;
    while (std::getline(std::cin, s)) {
        if (s.empty()) {
            break;
        }
        for (int j = 0; j < l; j++) {
            map[k][j] = s[j];
        }
        k++;
    }

    int res = 0;

    for (int i = 0; i < 1000; i++) {
        int startJ = -1;
        for (int j = 0; j < l; j++) {
            if (isNum(map[i][j])) {
                if (startJ == -1) {
                    startJ = j;
                }
            } else {
                if (startJ != -1) {
                    int n = computenum2(map, intmap, l, i, startJ, j);
                    res += n;
                    startJ = -1;
                }
            }
        }
        if (startJ != -1) {
            int n = computenum2(map, intmap, l, i, startJ, l);
            res += n;
        }
    }

    for (int i = 0; i < 1000; i++) {
        free(map[i]);
    }
    free(map);
    std::cout << res;
}


int main() {
    //solve1();
    solve2();
}