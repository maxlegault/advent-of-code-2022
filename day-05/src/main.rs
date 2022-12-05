use std::cell::RefCell;
use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;

fn main() {
    let deck = Deck::from_input("./initial-layout.txt");
    deck.process_instructions("./instructions.txt");
    println!("top crates: {}", deck.get_top_crates());
}

struct Deck {
    stacks: RefCell<Vec<Vec<char>>>,
}

impl Deck {
    fn from_input(filename: &str) -> Deck {
        let lines = read_lines(filename).expect("Should have been able to read initial layout file");
        let mut stacks: Vec<Vec<char>> = Vec::new();
        for _ in 1..10 {
            stacks.push(Vec::new());
        }
        let valid_chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ";
        for line_result in lines {
            if let Ok(line) = line_result {
                let mut index = 0;
                if line.contains('[') {
                    for i in (1..line.len()).step_by(4) {
                        let char = line.chars().nth(i).expect("should have read a char");
                        if valid_chars.contains(char) {
                            let stack = &mut stacks[index];
                            stack.insert(0, char);
                        }
                        index += 1;
                    }
                }
            }
        }
        return Deck { stacks: RefCell::new(stacks) };
    }

    fn process_instructions(&self, filename: &str) {
        let lines = read_lines(filename).expect("Should have been able to read instructions file");
        let mut stacks = self.stacks.borrow_mut();
        for line_result in lines {
            if let Ok(line) = line_result {
                let parts = line.splitn(6, ' ').collect::<Vec<&str>>();
                let amount: i32 = String::from(parts[1]).parse::<i32>().expect("should have been able to parse amount");
                let from_index: usize = String::from(parts[3]).parse::<usize>().expect("should have been able to parse from index") - 1;
                let to_index: usize = String::from(parts[5]).parse::<usize>().expect("should have been able to parse to index") - 1;
                for _ in 0..amount {
                    let char = stacks.get_mut(from_index).unwrap().pop().unwrap();
                    stacks.get_mut(to_index).unwrap().push(char);
                }
            }
        }
    }

    fn get_top_crates(&self) -> String {
        let mut output = String::new();
        self.stacks.borrow_mut().iter().for_each(|s| output.push(*s.last().expect("should have a char")));
        return output;
    }
}

// The output is wrapped in a Result to allow matching on errors
// Returns an Iterator to the Reader of the lines of the file.
fn read_lines<P>(filename: P) -> io::Result<io::Lines<io::BufReader<File>>>
    where P: AsRef<Path>, {
    let file = File::open(filename)?;
    Ok(io::BufReader::new(file).lines())
}