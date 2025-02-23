use std::io;

// Solution inspired from: https://www.thespatula.io/rust/rust_base64/

const BASE64_CHARSET: &[u8; 64] =
    b"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/";

fn base64_decode(s: &str) -> String {
    let mut decoded = vec![];
    let mut buffer = 0;
    let mut bits_collected = 0;

    for c in s.chars() {
        if c != '=' {
            let position = BASE64_CHARSET.iter().position(|&x| x == c as u8);

            match position {
                Some(pos) => {
                    buffer = buffer << 6 | pos as u32;
                    bits_collected += 6;

                    while bits_collected >= 8 {
                        bits_collected -= 8;
                        let byte = (buffer >> bits_collected) & 0xFF;
                        decoded.push(byte as u8);
                    }
                }
                None => panic!("invalid character: {c}"),
            }
        }
    }

    String::from_utf8(decoded).expect("decoding failed")
}

fn main() {
    io::stdin()
        .lines()
        .for_each(|l| println!("{}", base64_decode(&l.unwrap())));
}

#[cfg(test)]
mod test {
    use crate::base64_decode;

    #[test]
    fn test_base64_decode() {
        let samples = [
            ("bGF0ZS1hdC1uaWdodA==", "late-at-night"),
            ("d2l0aC10aGUtcmlzaW5nLWFwZQ==", "with-the-rising-ape"),
            ("dGhlLXJ1dGhsZXNzLXNldmVu", "the-ruthless-seven"),
        ];

        for (base64, ascii) in samples {
            assert!(
                base64_decode(base64) == ascii,
                "base64_decode(\"{base64}\"): \"{}\" != \"{ascii}\"",
                base64_decode(base64)
            );
        }
    }
}
