use crate::schema::clients;

#[derive(Serialize, Deserialize, Queryable)]
#[table_name = "clients"]
pub struct Client {
    pub id: i64,
    pub name: String, 
    pub email: String
}
