use crate::connection::DbConn;
use crate::client::repository;
use crate::client::Cli;
use diesel::result::Error;
use rocket::http::Status;
use rocket_contrib::json::Json;
use std::env;

#[get("/")]
pub fn all(connection: DbConn) -> Result<Json<Vec<Cli>>, Status> {
    repository::all(&connection)
        .map(|client| Json(client))
        .map_err(|error| error_status(error))
}

fn error_status(error: Error) -> Status {
    match error {
        Error::NotFound => Status::NotFound,
        _ => Status::InternalServerError,
    }
}

#[get("/<id>")]
pub fn get(id: i32, connection: DbConn) -> Result<Json<Cli>, Status> {
    repository::get(id, &connection)
        .map(|client| Json(client))
        .map_err(|error| error_status(error))
}

fn host() -> String {
    env::var("ROCKET_ADDRESS").expect("ROCKET_ADDRESS must be set")
}

fn port() -> String {
    env::var("ROCKET_PORT").expect("ROCKET_PORT must be set")
}
