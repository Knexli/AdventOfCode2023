package day1;

import java.io.File;
import java.io.FileNotFoundException; 
import java.util.Scanner;

public class Solution {
    public static void main(String[] args) {
        try {
            File inp = new File(".\\AdventOfCode2023\\simeon\\src\\day1\\input.txt");
            File test = new File(".\\AdventOfCode2023\\simeon\\src\\day1\\test.txt");
            Scanner myReader = new Scanner(inp);
            int out = 0;
            while (myReader.hasNextLine()) {
                String data = myReader.nextLine();
                //System.out.println(sol(data));
                out += sol2(data);
            }
            System.out.println(out);
            myReader.close();
        } catch (FileNotFoundException e) {
            System.out.println("An error occurred.");
            e.printStackTrace();
        }

    }

    public static int sol(String s){
        int a = 0;
        int b = 0;
        for (char c : s.toCharArray()) {
            if(Character.isDigit(c)){
                int i = Character.getNumericValue(c);
                if(a == 0){
                    a = i;
                    b = i;
                }else{
                    b = i;
                }
            }
        }
        return a*10+b;
    }

    public static int sol2(String s){
        s = s
            .replace("one", "o1e")
            .replace("two", "t2o")
            .replace("three", "t3e")
            .replace("four", "f4r")
            .replace("five", "f5e")
            .replace("six", "s6x")
            .replace("seven", "s7n")
            .replace("eight", "e8t")
            .replace("nine", "n9e");
        int a = 0;
        int b = 0;
        for (char c : s.toCharArray()) {
            if(Character.isDigit(c)){
                int i = Character.getNumericValue(c);
                if(a == 0){
                    a = i;
                    b = i;
                }else{
                    b = i;
                }
            }
        }
        System.out.println(a*10+b);
        return a*10+b;
    }


}
