use std::io;

fn check(parens: String) -> bool {
    let mut count = 0;
    for c in parens.chars() {
        if c == '(' {
            count += 1;
        } else {
            if count <= 0 {
                return false;
            }
            count -= 1;
        }
    }
    count == 0
}

fn main() {
    io::stdin().lines().for_each(|l| {
        if check(l.unwrap()) {
            println!("yes")
        } else {
            println!("no")
        }
    });
}
