package kata.its_almost_compression.java;

import java.util.Scanner;

class Solution {
    public static String compress(String str) {
        if (str.length() == 0) {
            return str;
        }

        char curr = str.charAt(0);
        int count = 1;
        String compressed = "";
        for (int i = 1; i < str.length(); i++) {
            if (curr != str.charAt(i)) {
                if (count > 2) {
                    compressed += Integer.toString(count) + curr;
                } else {
                    for (int j = 0; j < count; j++) {
                        compressed += curr;
                    }
                }
                curr = str.charAt(i);
                count = 1;
            } else {
                count++;
            }
        }

        if (count > 2) {
            compressed += Integer.toString(count) + curr;
        } else {
            for (int j = 0; j < count; j++) {
                compressed += curr;
            }
        }

        return compressed;
    }

    public static void main(String args[]) {
        Scanner scanner = new Scanner(System.in);
        while (scanner.hasNextLine()) {
            String line = scanner.nextLine().strip();
            System.out.println(compress(line));
        }
        scanner.close();
    }
}