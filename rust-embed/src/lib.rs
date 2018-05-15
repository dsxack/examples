use std::thread;

#[no_mangle]
pub extern fn process() {
    let handles: Vec<_> = (0..10).map(|_| {
        thread::spawn(|| {
            let mut x = 0;
            for _ in 0..5_000_000 {
                x += 1
            }
            x
        })
    }).collect();

    for h in handles {
        println!("Поток завершился со счетом={}",
            h.join().map_err(|_| "Не удалось соединиться с потоком!").unwrap());
    }
}
