package day1;

import java.io.File;
import java.io.FileNotFoundException;
import java.util.ArrayList;
import java.util.List;
import java.util.Scanner;

public class Solution {


    public static void main(String[] args) throws FileNotFoundException {
        Scanner scanner = new Scanner(new File("dominik/src/day1/input.txt"));
        List<String> input = new ArrayList<>();
        while (scanner.hasNextLine()) {
            input.add(scanner.nextLine());
        }
        scanner.close();

        run(input);

    }

    private static void run(List<String> input) {
        List<Integer> solution = new ArrayList<>();
        for (String s : input) {
            // String[] parts = s.split("[1-9]");
            String[] parts = s.split("[a-z]*");
            List<String> number = new ArrayList<>();
            for (String part : parts) {
                if (!part.isEmpty()) {
                    number.add(part);
                }
            }
            solution.add(Integer.valueOf(number.get(0) + number.get(number.size() - 1)));
        }
        int x = 0;
        for (Integer i : solution) {
            x += i;
        }
        System.out.println(x);
    }


}