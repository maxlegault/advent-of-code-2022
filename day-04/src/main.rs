use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;

fn main() {
    let lines = read_lines("./input.txt").expect("Should have been able to read the file");
    let pairs = lines
        .map(|l| Pair::from_string(l.expect("should have been able to read line")))
        .collect::<Vec<Pair>>();
    println!("complete overlap sum: {}", pairs.iter().map(|p| p.calculate_total_overlap()).sum::<i32>());
}

struct Pair {
    first: Assignment,
    second: Assignment,
}

struct Assignment {
    start: i32,
    end: i32,
}

impl Pair {
    fn from_string(s: String) -> Pair {
        let parts = s.splitn(2, ",").collect::<Vec<&str>>();
        Pair {
            first: Assignment::from_string(String::from(parts[0])),
            second: Assignment::from_string(String::from(parts[1])),
        }
    }

    fn calculate_total_overlap(&self) -> i32 {
        if self.first.start <= self.second.start && self.first.end >= self.second.end {
            return 1;
        }
        if self.second.start <= self.first.start && self.second.end >= self.first.end {
            return 1;
        }
        return 0;
    }
}

impl Assignment {
    fn from_string(s: String) -> Assignment {
        let parts = s.splitn(2, "-").collect::<Vec<&str>>();
        Assignment {
            start: String::from(parts[0]).parse().expect("should have parsed start"),
            end: String::from(parts[1]).parse().expect("should have parsed end"),
        }
    }
}


// The output is wrapped in a Result to allow matching on errors
// Returns an Iterator to the Reader of the lines of the file.
fn read_lines<P>(filename: P) -> io::Result<io::Lines<io::BufReader<File>>>
    where P: AsRef<Path>, {
    let file = File::open(filename)?;
    Ok(io::BufReader::new(file).lines())
}