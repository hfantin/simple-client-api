use chrono::NaiveDateTime;

use crate::schema::clients;

#[derive(Serialize, Deserialize, Queryable)]
#[table_name = "test"]
pub struct ClientView {
    pub id: i64,
    pub name: String, 
    pub view_time: NaiveDateTime,
    pub email: String
}
