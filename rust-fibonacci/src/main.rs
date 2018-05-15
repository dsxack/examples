use std::io;

fn main() {
    println!("Please enter N position of Fibonacci number");

    let mut n = String::new();

    io::stdin().read_line(&mut n)
        .expect("Failed to read line");

    let n: u32 = n.trim().parse()
        .expect("Please type a number!");

    println!("{} of Fibonacci is {}", n, nth_fibonacci(n))
}

fn nth_fibonacci(n: u32) -> u32 {
    let mut first: u32 = 0;
    let mut second: u32 = 1;

    for _ in 1..n {
        let old_first = first;

        first = second;
        second = old_first + second;
    }

    first + second
}