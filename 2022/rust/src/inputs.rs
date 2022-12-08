use std::{fs, path::Path};

pub fn get_input(num: &str) -> String {
    let str = format!("./src/inputs/{num}.txt");
    let path = Path::new(str.as_str());
    fs::read_to_string(path).expect("Failed reading file")
}
