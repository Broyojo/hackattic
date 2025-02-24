use std::io;

fn main() {
    io::stdin().lines().for_each(|l| {
        let offset = l.unwrap().parse::<i32>().expect("parsing failed");
        let mut weekday = (offset + 4) % 7;
        if weekday < 0 {
            weekday += 7;
        }
        let weekdays = [
            "Sunday",
            "Monday",
            "Tuesday",
            "Wednesday",
            "Thursday",
            "Friday",
            "Saturday",
        ];
        println!("{}", weekdays[weekday as usize]);
    });
}
