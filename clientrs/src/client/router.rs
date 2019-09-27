use crate::connection;
use crate::client;
use rocket;

pub fn create_routes() {
    rocket::ignite()
        .manage(connection::init_pool())
        .mount(
            "/v1/clients",
            routes![
                client::handler::all,
                client::handler::get
            ],
        )
        .launch();
}
