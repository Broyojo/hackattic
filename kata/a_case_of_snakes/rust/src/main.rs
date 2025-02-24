use std::io;

fn main() {
    io::stdin().lines().for_each(|l| {
        let line = l.unwrap();

        let mut snake_case = String::new();

        let mut first_upper_case = false;
        for (i, c) in line.chars().enumerate() {
            if c.is_uppercase() {
                if first_upper_case {
                    snake_case += "_";
                } else if i > 2 {
                    let mut bad = false;
                    for c_ in line[..i].chars() {
                        if c_.is_ascii_digit() {
                            bad = true;
                            break;
                        }
                    }
                    if !bad {
                        snake_case.push_str(&line[..i]);
                        snake_case.push('_');
                    }
                }
                snake_case.push(c.to_ascii_lowercase());
                first_upper_case = true;
            } else if first_upper_case {
                snake_case.push(c)
            }
        }

        println!("{}", snake_case);
    });
}
