use std::io;

fn fizzbuzz(n: i32, m: i32) {
    for i in n..=m {
        match (i % 3, i % 5) {
            (0, 0) => {
                println!("FizzBuzz")
            }
            (0, _) => {
                println!("Fizz")
            }
            (_, 0) => {
                println!("Buzz")
            }
            (_, _) => {
                println!("{}", i);
            }
        }
    }
}

fn main() {
    io::stdin().lines().take(1).for_each(|l| {
        let parts = l
            .unwrap()
            .trim()
            .split(" ")
            .map(|num| num.parse().expect("parsing failed"))
            .collect::<Vec<_>>();
        fizzbuzz(parts[0], parts[1]);
    });
}
