use std::fs;

fn main() {
    let data = fs::read("./input.txt").expect("should have read file to string");
    find_marker(&data[0..data.len()], 4);
    find_marker(&data[0..data.len()], 14);
}

fn find_marker(data: &[u8], len: usize) {
    let mut found = false;
    let mut index: usize = 0;
    while !found {
        let marker = &data[index..index+len];
        if is_marker_valid(marker, len) {
            found = true;
        } else {
            index += 1;
        }
    }
    println!("first marker of length {} after character {}", len, index + len);
}

fn is_marker_valid(marker: &[u8], len: usize) -> bool {
    for x in 0..len {
        for y in 0..len {
            if x != y && marker[x] == marker[y] {
                return false;
            }
        }
    }
    return true;
}
