use std::io;

fn main() {
    io::stdin().lines().for_each(|l| {
        println!(
            "{}",
            u16::from_str_radix(&l.unwrap().replace("#", "1").replace(".", "0"), 2)
                .expect("parsing failed")
        )
    });
}
