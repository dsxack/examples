use std::time::{Duration, Instant};
use std::thread::sleep;

fn main() {
    let now = Instant::now();
    sleep(Duration::new(2, 0));
    assert_eq!(2, now.elapsed().as_secs());

    let duration = Duration::from_millis(2569);
    assert_eq!(2, duration.as_secs());
    assert_eq!(569000000, duration.subsec_nanos());
}
