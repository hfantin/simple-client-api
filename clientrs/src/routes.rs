use diesel::{self, prelude::*};

use rocket_contrib::json::Json;

use crate::models::{ClientView};
use crate::schema;
use crate::DbConn;

#[get("/")]
pub fn index() -> &'static str {
    "Application successfully started!"
}


#[get("/v1/clients")]
pub fn list_clients_views(conn: DbConn) -> Result<Json<Vec<ClientView>>, String> {
    use crate::schema::clients::dsl::*;

    clients.load(&conn.0).map_err(|err| -> String {
        println!("Error querying client views: {:?}", err);
        "Error querying client views from the database".into()
    }).map(Json)
}