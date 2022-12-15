use std::fs;
fn main() {
    let content = fs::read_to_string("./src/input.txt").expect("Failed to read file");
    let elfs: Vec<&str> = content.split("\n").collect();
    let mut count: i32 = 0;
    for pair_of_elfs in elfs {
        let pair_of_elfs_array: Vec<&str> = pair_of_elfs.split(",").collect();
        let first_elf_pages: Vec<&str> = pair_of_elfs_array[0].split("-").collect();
        let second_elf_pages: Vec<&str> = pair_of_elfs_array[1].split("-").collect();
        let first_elf_starts_from: i32 = first_elf_pages[0].parse().unwrap();
        let first_elf_end_with: i32 = first_elf_pages[1].parse().unwrap();
        let second_elf_starts_from: i32 = second_elf_pages[0].parse().unwrap();
        let second_elf_end_with: i32 = second_elf_pages[1].parse().unwrap();
        if first_elf_starts_from <= second_elf_starts_from && ( first_elf_end_with >= second_elf_starts_from || first_elf_end_with >= second_elf_end_with) {
            count += 1;
        } else if second_elf_starts_from <= first_elf_starts_from && ( second_elf_end_with >= first_elf_starts_from || second_elf_end_with >= first_elf_end_with) {
            count += 1;
        }
    }
    println!("{count}")
}
