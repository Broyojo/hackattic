package kata.almost_binary.java;

import java.util.Scanner;

class Solution {
    public static void main(String args[]) {
        Scanner scanner = new Scanner(System.in);
        while (scanner.hasNextLine()) {
            String line = scanner.nextLine();
            System.out.println(Integer.parseUnsignedInt(line.replace("#", "1").replace(".", "0"), 2));
        }
        scanner.close();
    }
}