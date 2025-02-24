use std::io;

fn main() {
    io::stdin().lines().for_each(|l| {
        let mut current_char: Option<char> = None;
        let mut current_count = 0;
        let mut compressed = String::new();

        for c in l.unwrap().chars() {
            match current_char {
                Some(curr_char) => {
                    if curr_char != c {
                        if current_count > 2 {
                            compressed.push_str(&current_count.to_string());
                            compressed.push(curr_char);
                        } else {
                            for _ in 0..current_count {
                                compressed.push(curr_char);
                            }
                        }
                        current_char = Some(c);
                        current_count = 1;
                    } else {
                        current_count += 1;
                    }
                }
                None => {
                    current_char = Some(c);
                    current_count = 1;
                }
            }
        }
        if let Some(curr_char) = current_char {
            if curr_char != '\n' && current_count > 2 {
                compressed.push_str(&current_count.to_string());
                compressed.push(curr_char);
            } else {
                for _ in 0..current_count {
                    compressed.push(curr_char);
                }
            }
        }
        println!("{}", compressed);
    });
}
