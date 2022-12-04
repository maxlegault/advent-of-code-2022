use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;

fn main() {
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

// The output is wrapped in a Result to allow matching on errors
// Returns an Iterator to the Reader of the lines of the file.
fn read_lines<P>(filename: P) -> io::Result<io::Lines<io::BufReader<File>>>
    where P: AsRef<Path>, {
    let file = File::open(filename)?;
    Ok(io::BufReader::new(file).lines())
}