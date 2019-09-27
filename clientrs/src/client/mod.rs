#![allow(proc_macro_derive_resolution_fallback)]
use super::schema::cli;


pub mod handler;
pub mod repository;
pub mod router;

#[derive(Queryable, AsChangeset, Serialize, Deserialize)]
#[table_name = "cli"]
pub struct Cli {
    pub id: i32,
    pub name: String,
    pub birth_date: String,
    pub email: String
}

