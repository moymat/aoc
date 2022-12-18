mod inputs;

use inputs::get_input;
use std::str::FromStr;

fn main() {
    part1();
    part2();
}

fn get_lines(input: &String) -> Vec<&str> {
    return input
        .trim()
        .lines()
        .map(|l| l.trim())
        .filter(|l| l.len() > 0)
        .collect();
}

fn convert_to_i32(part: &str) -> Vec<i32> {
    part.split("-")
        .map(|s| match FromStr::from_str(s) {
            Ok(num) => num,
            Err(_) => panic!("string can't be converted to i32"),
        })
        .collect()
}

fn part1() {
    let input = get_input("04");
    let lines = get_lines(&input);

    let mut total_overlap = 0;
    for pair in lines {
        let splited: Vec<&str> = pair.split(",").collect();
        let first = convert_to_i32(splited[0]);
        let second = convert_to_i32(splited[1]);

        if (first[0] <= second[0] && first[1] >= second[1])
            || (first[0] >= second[0] && first[1] <= second[1])
        {
            total_overlap += 1;
        }
    }

    println!("part1: {}", total_overlap);
}

fn part2() {
    let input = get_input("04");
    let lines = get_lines(&input);

    let mut partial_overlaps = 0;

    for pair in lines {
        let splited: Vec<&str> = pair.split(",").collect();
        let first = convert_to_i32(splited[0]);
        let second = convert_to_i32(splited[1]);

        let first_range = first[0]..first[1] + 1;
        let second_range = second[0]..second[1] + 1;

        let mut overlap = false;

        for num in first_range.clone() {
            if second_range.contains(&num) {
                overlap = true;
                break;
            }
        }

        if overlap == false {
            for num in second_range {
                if first_range.contains(&num) {
                    overlap = true;
                    break;
                }
            }
        }

        if overlap == true {
            partial_overlaps += 1;
        }
    }

    println!("part2: {}", partial_overlaps);
}
