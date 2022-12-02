use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;

fn main() {
    let lines = read_lines("./input.txt").expect("Should have been able to read the file");

    let mut elf_calories: Vec<u32> = Vec::new();

    let mut current_count = 0u32;
    for line in lines {
        if let Ok(calories) = line {
            if calories == "" {
                elf_calories.push(current_count);
                current_count = 0;
            } else {
                current_count += calories.trim().parse::<u32>().expect("Should have been able to parse line as u8");
            }
        }
    }
    if current_count > 0 {
        elf_calories.push(current_count);
    }

    println!("Max for elf calories: {}", elf_calories.iter().max().unwrap())
}

// The output is wrapped in a Result to allow matching on errors
// Returns an Iterator to the Reader of the lines of the file.
fn read_lines<P>(filename: P) -> io::Result<io::Lines<io::BufReader<File>>>
    where P: AsRef<Path>, {
    let file = File::open(filename)?;
    Ok(io::BufReader::new(file).lines())
}