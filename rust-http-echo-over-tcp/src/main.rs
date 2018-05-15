use std::net::{TcpListener, TcpStream};
use std::thread;

use std::io::Read;
use std::io::Write;

fn handle_client(mut stream: TcpStream) {
    let mut buf;

    stream.write("HTTP/1.1 200 OK\n".as_bytes()).unwrap();
    stream.write("Transfer-Encoding: chunked\n".as_bytes()).unwrap();
    stream.write("\r\n".as_bytes()).unwrap();

    const BUF_SIZE: usize = 512;
    loop {
        buf = [0; BUF_SIZE];

        let m = match stream.read(&mut buf) {
            Ok(m) => {
                if m == 0 {
                    break
                }
                m
            }
            Err(e) => panic!(e)
        };

        stream.write(format!("{:x}\r\n", m).as_bytes()).unwrap();
        stream.write(&buf[0 .. m]).unwrap();
        stream.write("\r\n".as_bytes()).unwrap();

        if m < BUF_SIZE {
            break;
        }
    }

    stream.write("0\r\n\r\n".as_bytes()).unwrap();
}

fn main() {
    let listener = TcpListener::bind("0.0.0.0:8000").unwrap();

    for stream in listener.incoming() {
        match stream {
            Err(e) => println!("failed: {}", e),
            Ok(stream) => {
                thread::spawn(move || {
                    handle_client(stream)
                });
            }
        }
    }
}
