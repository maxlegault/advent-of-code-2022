use std::fs;

fn main() {
    let data = fs::read("./input.txt").expect("should have read file to string");
    let mut found = false;
    let mut index = 0;
    while !found {
        let marker = &data[index..index+4];
        if is_marker_valid(marker) {
            found = true;
        } else {
            index += 1;
        }
    }
    println!("first marker after character {}", index + 4);
}

fn is_marker_valid(marker: &[u8]) -> bool {
    for x in 0..4 {
        for y in 0..4 {
            if x != y && marker[x] == marker[y] {
                return false;
            }
        }
    }
    return true;
}
