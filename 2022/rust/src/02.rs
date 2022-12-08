mod inputs;

use inputs::get_input;

fn main() {
    part1();
    part2();
}

#[derive(Debug, PartialEq, Eq, Clone)]
enum Play {
    Rock,
    Paper,
    Scissors,
}

impl Play {
    fn from_i32(value: i32) -> Play {
        match value {
            65 | 88 => Play::Rock,
            66 | 89 => Play::Paper,
            67 | 90 => Play::Scissors,
            _ => panic!("Wrong value {}", value),
        }
    }

    fn to_result(&self, opponent: &Play) -> Play {
        match self {
            Play::Rock => opponent.beats(),
            Play::Paper => opponent.clone(),
            Play::Scissors => opponent.get_beaten_by(),
        }
    }

    fn get_beaten_by(&self) -> Play {
        match self {
            Play::Paper => Play::Scissors,
            Play::Rock => Play::Paper,
            Play::Scissors => Play::Rock,
        }
    }

    fn beats(&self) -> Play {
        match self {
            Play::Paper => Play::Rock,
            Play::Rock => Play::Scissors,
            Play::Scissors => Play::Paper,
        }
    }

    fn get_fight_result(&self, opponent: &Play) -> i32 {
        if self == opponent {
            3
        } else if opponent == &self.get_beaten_by() {
            0
        } else {
            6
        }
    }

    fn get_points(&self) -> i32 {
        match self {
            Play::Rock => 1,
            Play::Paper => 2,
            Play::Scissors => 3,
        }
    }
}

fn get_result(play1: &Play, play2: &Play) -> i32 {
    play1.get_points() + play1.get_fight_result(&play2)
}

fn convert_round(round: &str) -> Vec<Play> {
    round
        .trim()
        .chars()
        .filter(|char| (*char as i32) >= 65 && (*char as i32) <= 90)
        .map(|char| Play::from_i32(char as i32))
        .collect::<Vec<Play>>()
}

fn part1() {
    let inputs = get_input("02");
    let rounds = inputs.lines();

    let mut points = 0;

    for round in rounds {
        let plays = convert_round(round);
        points += get_result(&plays[1], &plays[0]);
    }

    println!("part1: {}", points);
}

fn part2() {
    let inputs = get_input("02");
    let rounds = inputs.lines();

    let mut points = 0;

    for round in rounds {
        let plays = convert_round(round);
        let to_play = plays[1].to_result(&plays[0]);
        points += get_result(&to_play, &plays[0])
    }

    println!("part2: {}", points);
}
