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

    private static int solve2(Scanner sc) {
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
        List<Node> currNodes = adj.entrySet().stream().filter(e -> e.getKey().endsWith("A"))
                .map(Map.Entry::getValue)
                .collect(Collectors.toList());
        System.out.println("amount starting nodes: " + currNodes.size() +
                ", " + currNodes.stream().map(e -> e.name).toList());
        Node[] nodes = new Node[currNodes.size()];
        for (int j = 0; j < currNodes.size(); j++) {
            nodes[j] = currNodes.get(j);
        }
        Set<String> visited = new HashSet<>();
        while (!isDestination(nodes)) {
            for (int j = 0; j < nodes.length; j++) {
                if (lr[i % n] == 'L') {
                    nodes[j] = adj.get(nodes[j].l);
                } else {
                    nodes[j] = adj.get(nodes[j].r);
                }
            }
            String s = Arrays.stream(nodes).map(e -> e.name).reduce(String::concat).orElseThrow();
            if (visited.contains(s)) {
                System.out.println("Ran in cycle after iteration " + i);
                break;
            }
            visited.add(s);
            if (i % 10_000_000 == 0) {
                System.out.println("Reached iteration " + i);
            }
            i++;
        }
        return i;
    }

    private static boolean isDestination(Node[] nodes) {
        for (Node n : nodes) {
            if (!n.name.endsWith("Z")) {
                return false;
            }
        }
        return true;
    }
}
