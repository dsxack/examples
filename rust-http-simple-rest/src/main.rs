extern crate iron;
extern crate router;

use iron::prelude::*;
use iron::status;
use router::Router;

fn main() {
    let mut router = Router::new();

    router.get("/", |req: &mut Request| {
        Ok(Response::with((status::Ok, "Hello" + req. + "!")))
    }, "index");

    Iron::new(router).http("localhost:3000").unwrap();
}
