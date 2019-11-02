class Client < ActiveRecord::Base 
	self.table_name = "cli"
#ApplicationRecord
	validates :name, presence: true
	validates :birth_dateemail, presence: true
end