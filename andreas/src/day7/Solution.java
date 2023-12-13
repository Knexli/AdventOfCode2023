package day7;

import java.util.*;

public class Solution {

    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        System.out.println(solve(sc));

//        Map<Character, Integer> testHand = new HashMap<>();
//        for (char c : "JJJJJ".toCharArray()) {
//            testHand.putIfAbsent(c, 0);
//            testHand.put(c, testHand.get(c) + 1);
//        }
//        System.out.println(Card.computeHandValue2(testHand));
    }

    private static long solve(Scanner sc) {
        List<Card> cards = new ArrayList<>();
        while (sc.hasNextLine()) {
            String[] s = sc.nextLine().split(" ");
            if (s.length != 2) {
                break;
            }
            cards.add(new Card(s[0].toCharArray(), Integer.parseInt(s[1])));
        }
        Collections.sort(cards);
        long res = 0;
        for (int i = 1; i <= cards.size(); i++) {
            res += (long)i * cards.get(i - 1).bid;
        }
        return res;
    }

    private static class Card implements Comparable<Card> {
        public char[] hand;
        public int bid;

        // 1 for task 1, 2 for task 2
        public static final int COMPARE_MODE = 2;

        public Card(char[] hand, int bid) {
            this.hand = hand;
            this.bid = bid;
        }

        @Override
        public int compareTo(Card other) {
            Map<Character, Integer> thisMap = new HashMap<>();
            Map<Character, Integer> otherMap = new HashMap<>();
            for (char c : hand) {
                thisMap.putIfAbsent(c, 0);
                thisMap.put(c, thisMap.get(c) + 1);
            }
            for (char c : other.hand) {
                otherMap.putIfAbsent(c, 0);
                otherMap.put(c, otherMap.get(c) + 1);
            }
            if (COMPARE_MODE == 1) {
                int vthis = computeHandValue(thisMap);
                int vother = computeHandValue(otherMap);
                if (vthis != vother) {
                    // cards are not the same constellation
                    return vthis - vother;
                }
                return computeTieValue(hand) - computeTieValue(other.hand);
            } else {
                int vthis = computeHandValue2(thisMap);
                int vother = computeHandValue2(otherMap);
                if (vthis != vother) {
                    // cards are not the same constellation
                    return vthis - vother;
                }
                return computeTieValue2(hand) - computeTieValue2(other.hand);
            }
        }

        private static int computeTieValue(char[] hand) {
            String s = "";
            for (char c : hand) {
                s += map(c);
            }
            return Integer.parseInt(s, 16);
        }

        private static int computeTieValue2(char[] hand) {
            String s = "";
            for (char c : hand) {
                s += map2(c);
            }
            return Integer.parseInt(s, 16);
        }

        private static char map(char c) {
            if (c == 'A') {
                return 'e';
            }
            if (c == 'K') {
                return 'd';
            }
            if (c == 'Q') {
                return 'c';
            }
            if (c == 'J') {
                return 'b';
            }
            if (c == 'T') {
                return 'a';
            }
            return c;
        }

        private static char map2(char c) {
            if (c == 'A') {
                return 'e';
            }
            if (c == 'K') {
                return 'd';
            }
            if (c == 'Q') {
                return 'c';
            }
            if (c == 'J') {
                return '1';
            }
            if (c == 'T') {
                return 'a';
            }
            return c;
        }

        private static int computeHandValue(Map<Character, Integer> map) {
            // Five of a kind
            if (map.keySet().size() == 1) {
                return 6;
            }

            // Four of a kind
            if (map.containsValue(4)) {
                return 5;
            }

            if (map.containsValue(3)) {
                if (map.containsValue(2)) {
                    // Full house
                    return 4;
                }
                // Three of a kind
                return 3;
            }

            if (map.values().stream().filter(e -> e == 2).count() == 2) {
                // two pair
                return 2;
            }

            if (map.containsValue(2)) {
                // one pair
                return 1;
            }
            // high card
            return 0;
        }

        public static int computeHandValue2(Map<Character, Integer> map) {
            int jVal = map.getOrDefault('J', 0);
            Map<Character, Integer> m = new HashMap<>();
            for (Map.Entry<Character, Integer> e : map.entrySet()) {
                if (e.getKey() == 'J') {
                    continue;
                }
                m.put(e.getKey(), e.getValue());
            }
            // Five of a kind
            if (m.values().stream().anyMatch(e -> e + jVal >= 5) || jVal == 5) {
                return 6;
            }

            // Four of a kind
            if (m.values().stream().anyMatch(e -> e + jVal >= 4)) {
                return 5;
            }

            // Full house
            if (jVal >= 3) {
                return 4;
            }
            if (m.keySet().size() == 2) {
                return 4;
            }

            // Three of a kind
            if (m.values().stream().anyMatch(e -> e + jVal >= 3)) {
                return 3;
            }

            // two pair
            if (jVal == 2) {
                return 2;
            }
            if (jVal == 1 && m.containsValue(2)) {
                // two pair
                return 2;
            }
            if (map.values().stream().filter(e -> e == 2).count() == 2) {
                return 2;
            }

            // one pair
            if (m.values().stream().anyMatch(e -> e + jVal >= 2)) {
                return 1;
            }

            // high card
            return 0;
        }
    }
}
