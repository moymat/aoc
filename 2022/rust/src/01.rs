mod inputs;

use inputs::get_input;

fn get_packs() -> Vec<Vec<i32>> {
    let input = get_input("01");

    let mut packs: Vec<Vec<i32>> = vec![];
    for line in input.lines() {
        if line.trim() == "" {
            packs.push(vec![]);
        } else {
            let calories = line.trim().parse().unwrap();
            match packs.last_mut() {
                Some(v) => v.push(calories),
                None => packs.push(vec![calories]),
            }
        }
    }

    return packs;
}

struct Max {
    calories: i32,
    idx: usize,
}

fn get_max_calories(packs: &Vec<Vec<i32>>) -> Max {
    let mut max = Max {
        calories: 0,
        idx: 0,
    };

    for (i, pack) in packs.iter().enumerate() {
        let mut total = 0;
        for calories in pack {
            total += calories;
        }
        if total > max.calories {
            max.calories = total;
            max.idx = i;
        }
    }

    return max;
}

fn part1() {
    let max = get_max_calories(&get_packs());

    println!("part1: {}", max.calories)
}

fn part2() {
    let mut packs = get_packs();

    let max1 = get_max_calories(&packs);
    packs.remove(max1.idx);
    let max2 = get_max_calories(&packs);
    packs.remove(max2.idx);
    let max3 = get_max_calories(&packs);

    let total = max1.calories + max2.calories + max3.calories;

    println!("part2: {}", total);
}

fn main() {
    part1();
    part2();
}
