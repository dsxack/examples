use std::thread;

#[no_mangle]
pub extern fn process(number: i32, threads: i32) {
    let handles: Vec<_> = (0..threads).map(|_| {
        thread::spawn(move || {
            let mut x = 0;
            for _ in 0..number {
                x += 1
            }
        })
    }).collect();

    for h in handles {
//        println!("Поток завершился со счетом={}",
//                 h.join().map_err(|_| "Не удалось соединиться с потоком!").unwrap());
        h.join().map_err(|_| "Не удалось соединиться с потоком!").unwrap()
    }
}
