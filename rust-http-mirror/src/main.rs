extern crate hyper;

use std::io;
use hyper::server::{Server, Request, Response};
use hyper::header::ContentLength;

fn main() {
    Server::http("0.0.0.0:8000").unwrap().handle(|mut req: Request, mut res: Response| {
        if let Some(len) = req.headers.get::<ContentLength>() {
            res.headers_mut().set(len.clone());
        }

        io::copy(&mut req, &mut res.start().unwrap()).unwrap();
    }).unwrap();
}