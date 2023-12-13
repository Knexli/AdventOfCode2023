package day6;

import java.util.ArrayList;
import java.util.List;
import java.util.Scanner;

public class Solution {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        System.out.println(solve2(sc));
    }

    public static long solve(Scanner sc) {
        String line1 = sc.nextLine();
        String line2 = sc.nextLine();
        List<Integer> times = getV(line1);
        List<Integer> distances = getV(line2);

        long res = 0;
        for (int i = 0; i < times.size(); i++) {
            int time = times.get(i);
            int distance = distances.get(i);

            int c = 0;
            for (int j = 0; j < time; j++) {
                // check if time j is possible
                if ((time - j) * j > distance) {
                    c++;
                }
            }

            res = res == 0 ? c : res * c;
        }
        return res;
    }

    public static long solve2(Scanner sc) {
        String line1 = sc.nextLine();
        String line2 = sc.nextLine();
        List<Integer> times = getV(line1);
        List<Integer> distances = getV(line2);

        long res = 0;
        long time = Long.parseLong(times.stream().map(e -> e + "").reduce(String::concat).orElse("-1"));
        long distance = Long.parseLong(distances.stream().map(e -> e + "").reduce(String::concat).orElse("-1"));

        for (int j = 0; j < time; j++) {
            // check if time j is possible
            if ((time - j) * j > distance) {
                res++;
            }
        }

        return res;
    }

    public static List<Integer> getV(String line) {
        List<Integer> list = new ArrayList<>();
        Scanner s1 = new Scanner(line);
        s1.next();
        while (s1.hasNextInt()) {
            list.add(s1.nextInt());
        }
        return list;
    }
}
