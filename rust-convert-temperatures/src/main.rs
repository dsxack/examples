use std::io;

fn main() {
    println!("Print temperature in celsius");

    let mut celsius = String::new();

    io::stdin().read_line(&mut celsius)
        .expect("Failed to read line");

    let celsius: f64 = celsius.trim().parse()
        .expect("Please type a number!");

    println!("{} celsius = {} fahrenheit", celsius, convert_celsius_to_fahrenheit(celsius))
}

fn convert_celsius_to_fahrenheit(celsius: f64) -> f64 {
    celsius * 1.8 + 32.0
}
