#![feature(proc_macro_hygiene, decl_macro)]

extern crate chrono;
#[macro_use]
extern crate diesel;
#[macro_use]
extern crate rocket;
#[macro_use]
extern crate rocket_contrib;
extern crate serde;
extern crate serde_json;
#[macro_use]
extern crate serde_derive;

pub mod cors;
pub mod models;
pub mod routes;
pub mod schema;

#[database("test")]
pub struct DbConn(diesel::PgConnection);

fn main() {
      rocket::ignite()
        .mount("/", routes![
            routes::index,
            routes::list_cli,
        ])
        .attach(DbConn::fairing())
        .attach(cors::CorsFairing)
        .launch();
}