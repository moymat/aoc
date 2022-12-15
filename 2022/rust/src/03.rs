mod inputs;

use inputs::get_input;

fn main() {
    part1();
    part2();
}

fn get_value(char: char) -> i32 {
    let value = char as i32;
    if value < 91 {
        value - 64 + 26
    } else {
        value - 96
    }
}

fn part1() {
    let input = get_input("03");

    let rucksacks = input.trim().lines().map(|l| l.trim());

    let mut sum = 0;

    for rucksack in rucksacks {
        let (compartment1, compartment2) = (
            &rucksack[..rucksack.len() / 2],
            &rucksack[(rucksack.len()) / 2..],
        );
        for item in compartment1.chars() {
            if compartment2.contains(item) {
                sum += get_value(item);
                break;
            }
        }
    }

    println!("part1: {}", sum);
}

fn part2() {
    let input = get_input("03");

    let rucksacks: Vec<&str> = input.trim().lines().map(|l| l.trim()).collect();

    let mut sum = 0;

    let mut i = 0;
    while i < rucksacks.len() {
        let (m1, m2, m3) = (rucksacks[i], rucksacks[i + 1], rucksacks[i + 2]);

        for item in m1.chars() {
            if m2.contains(item) && m3.contains(item) {
                sum += get_value(item);
                break;
            }
        }

        i += 3;
    }

    println!("part2: {}", sum);
}
