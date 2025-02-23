package kata.debasing_64.java;

import java.util.Base64;
import java.util.Scanner;

class Solution {
    public static void main(String args[]) {
        Scanner scanner = new Scanner(System.in);
        while (scanner.hasNextLine()) {
            String line = scanner.nextLine();
            String decoded = new String(Base64.getDecoder().decode(line));
            System.out.println(decoded);
        }
        scanner.close();
    }
}