use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;

fn main() {
    let lines = read_lines("./input.txt").expect("Should have been able to read the file");

    let games: Vec<Game> = lines
        .map(|line| {
            let unwrapped = line.expect("Should have been able to get line");
            let values: Vec<&str> = unwrapped.splitn(2, ' ').collect();
            return Game {
                opponent: Weapon::from_letter(values[0]),
                mine: Weapon::from_letter(values[1]),
                desired_outcome: Outcome::from_letter(values[1]),
            };
        })
        .collect();

    println!("Sum for all games v1 is {}", games.iter().map(|game| game.calculate_score_v1()).sum::<i32>());
    println!("Sum for all games v2 is {}", games.iter().map(|game| game.calculate_score_v2()).sum::<i32>());
}

#[derive(Copy, Clone)]
enum Weapon {
    Unknown = 0,
    Rock = 1,
    Paper = 2,
    Scissor = 3,
}

#[derive(Copy, Clone)]
struct Game {
    mine: Weapon,
    opponent: Weapon,
    desired_outcome: Outcome,
}

#[derive(Copy, Clone)]
enum Outcome {
    Lose,
    Draw,
    Win,
}

impl Game {
    fn calculate_score_v1(self) -> i32 {
        return Game::calculate_score(self.mine, self.opponent);
    }

    fn calculate_score_v2(self) -> i32 {
        let mine: Weapon;
        match self.desired_outcome {
            Outcome::Draw => mine = self.opponent,
            Outcome::Win => {
                match self.opponent {
                    Weapon::Rock => mine = Weapon::Paper,
                    Weapon::Paper => mine = Weapon::Scissor,
                    _ => mine = Weapon::Rock
                }
            }
            _ => {
                match self.opponent {
                    Weapon::Rock => mine = Weapon::Scissor,
                    Weapon::Paper => mine = Weapon::Rock,
                    _ => mine = Weapon::Paper,
                }
            }
        }
        return Game::calculate_score(mine, self.opponent);
    }

    fn calculate_score(mine: Weapon, opponent: Weapon) -> i32 {
        let score = mine.score();
        let diff = score - opponent.score();
        if diff == 0 {
            return score + 3;
        }
        if diff == 1 || diff == -2 {
            return score + 6;
        }
        return score;
    }
}

impl Weapon {
    fn score(self) -> i32 {
        return self as i32;
    }

    fn from_letter(letter: &str) -> Weapon {
        if letter == "A" || letter == "X" {
            return Weapon::Rock;
        }
        if letter == "B" || letter == "Y" {
            return Weapon::Paper;
        }
        if letter == "C" || letter == "Z" {
            return Weapon::Scissor;
        }
        return Weapon::Unknown;
    }
}

impl Outcome {
    fn from_letter(letter: &str) -> Outcome {
        if letter == "Y" {
            return Outcome::Draw;
        }
        if letter == "Z" {
            return Outcome::Win;
        }
        return Outcome::Lose;
    }
}

// The output is wrapped in a Result to allow matching on errors
// Returns an Iterator to the Reader of the lines of the file.
fn read_lines<P>(filename: P) -> io::Result<io::Lines<io::BufReader<File>>>
    where P: AsRef<Path>, {
    let file = File::open(filename)?;
    Ok(io::BufReader::new(file).lines())
}