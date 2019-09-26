use diesel::{self, prelude::*};

use rocket_contrib::json::Json;

use crate::models::{Cli};
use crate::schema;
use crate::DbConn;

#[get("/")]
pub fn index() -> &'static str {
    "Application successfully started!"
}


#[get("/v1/clients")]
pub fn list_cli(conn: DbConn) -> Result<Json<Vec<Cli>>, String> {
    use crate::schema::cli::dsl::*;

    cli.load(&conn.0).map_err(|err| -> String {
        println!("Error querying client views: {:?}", err);
        "Error querying client views from the database".into()
    }).map(Json)
}