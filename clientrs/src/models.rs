use chrono::NaiveDateTime;

use crate::schema::clientviews;

#[derive(Serialize, Deserialize, Queryable)]
#[table_name = "test"]
pub struct TestView {
    pub id: i64,
    pub name: String, 
    pub view_time: NaiveDateTime,
    pub email: String
}
