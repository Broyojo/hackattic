package kata.open_parens.java;

import java.util.Scanner;

class Solution {
    public static void main(String args[]) {
        Scanner scanner = new Scanner(System.in);
        while (scanner.hasNextLine()) {
            String line = scanner.nextLine();
            if (check(line)) {
                System.out.println("yes");
            } else {
                System.out.println("no");
            }
        }
        scanner.close();
    }

    public static boolean check(String parens) {
        int count = 0;
        for (int i = 0; i < parens.length(); i++) {
            char c = parens.charAt(i);
            switch (c) {
                case '(':
                    count++;
                    break;
                case ')':
                    if (count <= 0) {
                        return false;
                    }
                    count--;
                    break;
                default:
                    System.err.println("unknown character: " + c);
                    System.exit(1);
            }
        }
        return count == 0;
    }
}