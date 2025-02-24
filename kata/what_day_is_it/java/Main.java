package kata.what_day_is_it.java;

import java.util.Scanner;

class Solution {
    public static void main(String args[]) {
        Scanner scanner = new Scanner(System.in);
        while (scanner.hasNextLine()) {
            String line = scanner.nextLine();
            int offset = Integer.parseInt(line);
            int weekday = (offset + 4) % 7;
            if (weekday < 0) {
                weekday += 7;
            }
            String[] weekdays = {
                    "Sunday",
                    "Monday",
                    "Tuesday",
                    "Wednesday",
                    "Thursday",
                    "Friday",
                    "Saturday",
            };
            System.out.println(weekdays[weekday]);
        }
        scanner.close();
    }
}