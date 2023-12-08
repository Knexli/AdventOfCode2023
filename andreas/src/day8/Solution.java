package day8;

import java.util.*;
import java.util.stream.Collectors;

public class Solution {

    private record Node(String name, String l, String r) {
    }

    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);
        System.out.println(solve2(sc));
    }

    private static int solve(Scanner sc) {
        char[] lr = sc.nextLine().toCharArray();
        sc.nextLine();
        Map<String, Node> adj = new HashMap<>();
        while (sc.hasNextLine()) {
            String line = sc.nextLine();
            if (line.isBlank()) {
                break;
            }
            String[] split = line.split(" ");
            Node n = new Node(split[0],
                    split[2].substring(1, split[2].length() - 1),
                    split[3].substring(0, split[3].length() - 1));
            adj.put(n.name, n);
        }

        int i = 0, n = lr.length;
        Node currNode = adj.get("AAA");
        while (!currNode.name.equals("ZZZ")) {
            if (lr[i % n] == 'L') {
                currNode = adj.get(currNode.l);
            } else {
                currNode = adj.get(currNode.r);
            }
            i++;
        }
        return i;
    }

    private static long solve2(Scanner sc) {
        char[] lr = sc.nextLine().toCharArray();
        sc.nextLine();
        Map<String, Node> adj = new HashMap<>();
        while (sc.hasNextLine()) {
            String line = sc.nextLine();
            if (line.isBlank()) {
                break;
            }
            String[] split = line.split(" ");
            Node n = new Node(split[0],
                    split[2].substring(1, split[2].length() - 1),
                    split[3].substring(0, split[3].length() - 1));
            adj.put(n.name, n);
        }

        int i, n = lr.length;
        List<Node> currNodes = adj.entrySet().stream().filter(e -> e.getKey().endsWith("A"))
                .map(Map.Entry::getValue)
                .collect(Collectors.toList());
        System.out.println("amount starting nodes: " + currNodes.size() +
                ", " + currNodes.stream().map(e -> e.name).toList());


        // cycleTimes[i][j] = if i == 0: amount of steps after which first node with suffix Z with j-th start is reached
        //                    if i == 1: cycle time to reach that same node again
        int[][] cycleTimes = new int[2][currNodes.size()];
        for (int k = 0; k < currNodes.size(); k++) {
            Node start = currNodes.get(k);
            Node currNode = start;
            int c = 0;
            i = 0;
            while(true) {
                if (lr[i % n] == 'L') {
                    currNode = adj.get(currNode.l);
                } else {
                    currNode = adj.get(currNode.r);
                }
                i++;
                if (currNode.name.endsWith("Z")) {
                    if (c == 0) {
                        cycleTimes[0][k] = i;
                        c = i;
                    } else {
                        cycleTimes[1][k] = i - c;
                        break;
                    }
                }
            }
        }

        long z = 0;
        outer: while (true) {
            // assume possible with z loops of the first startpos
            long totalsteps = cycleTimes[0][0] + z * cycleTimes[1][0];
            for (int k = 1; k < currNodes.size(); k++) {
                if ((totalsteps - cycleTimes[0][k]) % cycleTimes[1][k] != 0) {
                    z++;
                    continue outer;
                }
            }
            return totalsteps;
        }
    }
}
