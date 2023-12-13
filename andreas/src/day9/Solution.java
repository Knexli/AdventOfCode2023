package day9;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.Scanner;

public class Solution {
    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        System.out.println(solve(sc));
    }

    private static long solve(Scanner sc) {
        long res = 0L;
        while (true) {
            String line = sc.nextLine();
            if (line.isBlank()) {
                break;
            }
            List<Integer> l = new ArrayList<>();
            Arrays.stream(line.split(" ")).map(Integer::parseInt).forEach(l::add);
            long[] asArr = new long[l.size()];
            for (int i = 0; i < l.size(); i++) {
                asArr[i] = l.get(i);
            }
            // change to res += getNext(asArr); for task 1
            res += getNextLeft(asArr);
        }
        return res;
    }

    private static long getNext(long[] arr) {
        if (Arrays.stream(arr).allMatch(e -> e == 0)) {
            return 0L;
        }
        long[] next = new long[arr.length - 1];
        for (int i = 0; i < arr.length - 1; i++) {
            next[i] = arr[i + 1] - arr[i];
        }
        long z = getNext(next);
        return arr[arr.length - 1] + z;
    }

    private static long getNextLeft(long[] arr) {
        if (Arrays.stream(arr).allMatch(e -> e == 0)) {
            return 0L;
        }
        long[] next = new long[arr.length - 1];
        for (int i = 0; i < arr.length - 1; i++) {
            next[i] = arr[i + 1] - arr[i];
        }
        long z = getNextLeft(next);
        return arr[0] - z;
    }


}
