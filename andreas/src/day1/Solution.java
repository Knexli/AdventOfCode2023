package day1;

import java.util.Map;
import java.util.Scanner;

public class Solution {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        System.out.println(solve2(sc));
    }

    public static int solve(Scanner sc) {
        int c = 0;
        while (true) {
            String s = sc.nextLine();
            if (s.isBlank()) {
                break;
            }
            int f = Integer.MIN_VALUE, l = 0;
            for (char v : s.toCharArray()) {
                if ('0' <= v && v <= '9') {
                    l = v - '0';
                    if (f == Integer.MIN_VALUE) {
                        f = v - '0';
                    }
                }
            }
            c += Integer.parseInt("" + f + l);
        }
        return c;
    }

    public static int solve2(Scanner sc) {
        int c = 0;
        Map<String, Integer> map = Map.ofEntries(
                Map.entry("one", 1),
                Map.entry("two", 2),
                Map.entry("three", 3),
                Map.entry("four", 4),
                Map.entry("five", 5),
                Map.entry("six", 6),
                Map.entry("seven", 7),
                Map.entry("eight", 8),
                Map.entry("nine", 9)
        );
        while (true) {
            String s = sc.nextLine();
            if (s.isBlank()) {
                break;
            }
            int f = -1, l = 0;
            for (int i = 0; i < s.length(); i++) {
                int v = -1;
                if (s.charAt(i) >= '0' && s.charAt(i) <= '9') {
                    v = s.charAt(i) - '0';
                }
                for (Map.Entry<String, Integer> e : map.entrySet()) {
                    if (s.substring(i).startsWith(e.getKey())) {
                        v = e.getValue();
                    }
                }
                if (v != -1) {
                    l = v;
                    if (f == -1) {
                        f = v;
                    }
                }
            }
            c += Integer.parseInt("" + f + l);
        }
        return c;
    }
}
