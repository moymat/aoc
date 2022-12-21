mod inputs;

use inputs::get_input;

fn main() {
    part1();
}

#[derive(Debug)]
struct Move {
    amount: u32,
    start_idx: usize,
    end_idx: usize,
}

fn get_crates(part: &str) -> Vec<Vec<Option<char>>> {
    let mut crates = part
        .split("\n")
        .map(|line| {
            line.chars()
                .enumerate()
                .filter_map(|(idx, char)| {
                    if idx > 0 && (idx - 1) % 4 == 0 {
                        if char.is_whitespace() {
                            Some(None)
                        } else {
                            Some(Some(char))
                        }
                    } else {
                        None
                    }
                })
                .collect()
        })
        .collect::<Vec<Vec<Option<char>>>>();

    crates.truncate(crates.len() - 1);

    crates
}

fn get_moves(part: &str) -> Vec<Move> {
    part.lines()
        .map(|line| {
            let parts = line.split(" ").collect::<Vec<&str>>();
            Move {
                amount: parts[1].parse::<u32>().unwrap(),
                start_idx: parts[3].parse::<usize>().unwrap() - 1,
                end_idx: parts[5].parse::<usize>().unwrap() - 1,
            }
        })
        .collect::<Vec<Move>>()
}

fn part1() {
    let input = get_input("05");
    let parts = input.split("\n\n").collect::<Vec<&str>>();

    let mut crates = get_crates(parts[0]);
    let moves = get_moves(parts[1]);

    for crane_move in moves {
        for _ in 0..crane_move.amount {
            let to_move = crates
                .iter_mut()
                .find_map(|line| match line[crane_move.start_idx] {
                    Some(char) => {
                        line[crane_move.start_idx] = None;
                        Some(char)
                    }
                    None => None,
                })
                .unwrap();

            let mut line_to_move_to: Option<usize> = None;
            for (idx, line) in crates.iter().enumerate() {
                if let None = line[crane_move.end_idx] {
                    line_to_move_to = Some(idx);
                }
            }

            match line_to_move_to {
                Some(idx) => {
                    crates[idx][crane_move.end_idx] = Some(to_move);
                }
                None => {
                    let mut new_line = vec![None; crates[0].len()];
                    new_line[crane_move.end_idx] = Some(to_move);
                    crates.insert(0, new_line);
                }
            }
        }
    }

    let mut all_first = String::new();
    for stack_idx in 0..crates[0].len() {
        for line in crates.clone() {
            if let Some(char) = line[stack_idx] {
                all_first.push(char);
                break;
            }
        }
    }

    println!("part1: {}", all_first);
}
