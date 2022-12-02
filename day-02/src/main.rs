use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;

fn main() {
    let lines = read_lines("./input.txt").expect("Should have been able to read the file");

    let games: Vec<Game> = lines
        .map(|line| {
            let values = line
                .unwrap()
                .splitn(2, ' ')
                .map(Weapon::from_letter)
                .collect::<Vec<Weapon>>();
            return Game {
                opponent: values[0],
                mine: values[1],
            };
        })
        .collect();

    println!("Sum for all games is {}", games.iter().map(|game| game.calculate_score()).sum::<i32>());
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
}

impl Game {
    fn calculate_score(self) -> i32 {
        let score = self.mine.score();
        let diff = score - self.opponent.score();
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

// The output is wrapped in a Result to allow matching on errors
// Returns an Iterator to the Reader of the lines of the file.
fn read_lines<P>(filename: P) -> io::Result<io::Lines<io::BufReader<File>>>
    where P: AsRef<Path>, {
    let file = File::open(filename)?;
    Ok(io::BufReader::new(file).lines())
}