use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;

fn main() {
    challenge_1();
    challenge_2();
}

fn challenge_1() {
    let lines = read_lines("./input.txt").expect("Should have been able to read the file");
    let priorities: String = String::from(" abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ");

    let sum = lines.map(|line_result| {
        let line = line_result.expect("should have been able to read line");
        let (compartment_1, compartment_2) = line.split_at(line.len() / 2);
        let letters = compartment_1.chars().filter(|c| compartment_2.contains(*c)).collect::<Vec<char>>();//.map(|c| priorities.find(c)).collect();
        return priorities.find(letters[0]).expect("should have been able to find priority");
    }).sum::<usize>();

    println!("sum of priorities: {}", sum);
}

fn challenge_2() {
    let lines = read_lines("./input.txt").expect("Should have been able to read the file");
    let priorities: String = String::from(" abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ");

    let mut line_1: String = String::new();
    let mut line_2: String = String::new();
    let badges_sum = lines.map(|line_result| {
        if line_1 == "" {
            line_1 = line_result.expect("should have been able to read line 1");
            return 0;
        }
        if line_2 == "" {
            line_2 = line_result.expect("should have been able to read line 2");
            return 0;
        }
        let letters = line_result
            .expect("should have been able to read line 3")
            .chars()
            .filter(|c| line_1.contains(*c) && line_2.contains(*c))
            .collect::<Vec<char>>();
        line_1 = String::new();
        line_2 = String::new();
        return priorities.find(letters[0]).expect("should have been able to find badge priority");
    }).sum::<usize>();
    println!("badges sum: {}", badges_sum);
}

// The output is wrapped in a Result to allow matching on errors
// Returns an Iterator to the Reader of the lines of the file.
fn read_lines<P>(filename: P) -> io::Result<io::Lines<io::BufReader<File>>>
    where P: AsRef<Path>, {
    let file = File::open(filename)?;
    Ok(io::BufReader::new(file).lines())
}