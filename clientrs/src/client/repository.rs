#![allow(proc_macro_derive_resolution_fallback)]

use crate::schema::cli::table;
use crate::client::Cli;
use crate::diesel::prelude::*;

pub fn all(connection: &PgConnection) -> QueryResult<Vec<Cli>> {
    table.load::<Cli>(&*connection)
}

pub fn get(id: i32, connection: &PgConnection) -> QueryResult<Cli> {
    table.find(id).get_result::<Cli>(connection)
}