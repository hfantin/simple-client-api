use crate::schema::*;

#[derive(Serialize, Deserialize, Queryable)]
pub struct Cli {
    pub id: i64,
    pub name: String, 
    pub email: String
}
