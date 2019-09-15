USE test;

drop table IF EXISTS CLI;

CREATE TABLE `CLI` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `NAME` varchar(50) NOT NULL,
  `BIRTH_DATE` date NOT NULL,
  `EMAIL` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

INSERT INTO test.CLI (NAME,BIRTH_DATE,EMAIL) VALUES 
('Naoma Gourlay' ,'1980-01-31',null).
('Janene Presnell' ,'1980-01-31',null).
('Clint Aguon' ,'1980-01-31',null).
('Elli Casares' ,'1980-01-31',null).
('Toni Keeler' ,'1980-01-31',null).
('Luvenia Schoepp' ,'1980-01-31',null).
('Giovanna Burchill' ,'1980-01-31',null).
('Marjory Kyler' ,'1980-01-31',null).
('Joshua Villalpando' ,'1980-01-31',null).
('Hue Bazile' ,'1980-01-31',null).
('Lanell Brehmer' ,'1980-01-31',null).
('Fran Friedrich' ,'1980-01-31',null).
('Thomas Turrell' ,'1980-01-31',null).
('Marcell Harrah' ,'1980-01-31',null).
('Dorothea Dhillon' ,'1980-01-31',null).
('Valencia Borland' ,'1980-01-31',null).
('Bibi Kitzman' ,'1980-01-31',null).
('Emmitt Swager' ,'1980-01-31',null).
('Ashlie Mcfaddin' ,'1980-01-31',null).
('Tressa Mass' ,'1980-01-31',null).
('Delana Fillion' ,'1980-01-31',null).
('Tamera Tung' ,'1980-01-31',null).
('Cyril Pawloski' ,'1980-01-31',null).
('Jessie Jiles' ,'1980-01-31',null).
('Nieves Talbott' ,'1980-01-31',null).
('Delsie Ram' ,'1980-01-31',null).
('Mendy Weyer' ,'1980-01-31',null).
('Vernon Resnick' ,'1980-01-31',null).
('Johnna Maner' ,'1980-01-31',null).
('Clement Lyons' ,'1980-01-31',null).
('Irish Kogut' ,'1980-01-31',null).
('Charisse Hollins' ,'1980-01-31',null).
('Romaine Heitman' ,'1980-01-31',null).
('Rachelle Longnecker' ,'1980-01-31',null).
('Gidget Gaitan' ,'1980-01-31',null).
('Donald Leonardi' ,'1980-01-31',null).
('Margo Victor' ,'1980-01-31',null).
('Tracie Durrell' ,'1980-01-31',null).
('Anjelica Hilchey' ,'1980-01-31',null).
('Tarra Tarrance' ,'1980-01-31',null).
('Vicki Hypolite' ,'1980-01-31',null).
('Kai Barto' ,'1980-01-31',null).
('Chung Younker' ,'1980-01-31',null).
('Antonetta Oullette' ,'1980-01-31',null).
('Shalanda Oliver' ,'1980-01-31',null).
('Velvet Coursey' ,'1980-01-31',null).
('Maribeth Montijo' ,'1980-01-31',null).
('Adena Poch' ,'1980-01-31',null).
('Randal Bridge' ,'1980-01-31',null).
('Nilda Chi' ,'1980-01-31',null).
('Rachelle Grooms' ,'1980-01-31',null).
('Taunya Villescas' ,'1980-01-31',null).
('Tamesha Hoerr' ,'1980-01-31',null).
('Retta Deherrera' ,'1980-01-31',null).
('Jimmy Munford' ,'1980-01-31',null).
('Sumiko Sifuentes' ,'1980-01-31',null).
('Adam Gilbertson' ,'1980-01-31',null).
('Neva Garofalo' ,'1980-01-31',null).
('Katia Giannone' ,'1980-01-31',null).
('Lakendra Mccrystal' ,'1980-01-31',null).
('Dean Beckett' ,'1980-01-31',null).
('Angelic Samms' ,'1980-01-31',null).
('Hal Hewitt' ,'1980-01-31',null).
('Boyce Lindon' ,'1980-01-31',null).
('Jina Dupler' ,'1980-01-31',null).
('Tomasa Elsasser' ,'1980-01-31',null).
('Jeff Sholes' ,'1980-01-31',null).
('Alexandria Enright' ,'1980-01-31',null).
('Tia Cangelosi' ,'1980-01-31',null).
('Leonore Dimas' ,'1980-01-31',null).
('Ardis Raftery' ,'1980-01-31',null).
('Millicent Lew' ,'1980-01-31',null).
('Roxanne Sorensen' ,'1980-01-31',null).
('Krissy Balducci' ,'1980-01-31',null).
('Lily Wise' ,'1980-01-31',null).
('Daphne Heckart' ,'1980-01-31',null).
('Adrianna Delillo' ,'1980-01-31',null).
('Zachary Beyer' ,'1980-01-31',null).
('Jose Artiaga' ,'1980-01-31',null).
('An Merriam' ,'1980-01-31',null).
('Lovetta Elwood' ,'1980-01-31',null).
('Naomi Pipkin' ,'1980-01-31',null).
('Joni Etheredge' ,'1980-01-31',null).
('Ty Gartrell' ,'1980-01-31',null).
('Malka Banegas' ,'1980-01-31',null).
('Emelda Moisan' ,'1980-01-31',null).
('Trey Villanova' ,'1980-01-31',null).
('Angila Coplin' ,'1980-01-31',null).
('Hye Maurice' ,'1980-01-31',null).
('Hilda Naughton' ,'1980-01-31',null).
('Regenia Gooding' ,'1980-01-31',null).
('Humberto Easler' ,'1980-01-31',null).
('Leesa Oathout' ,'1980-01-31',null).
('Alonzo Schulte' ,'1980-01-31',null).
('Rebeca Blythe' ,'1980-01-31',null).
('Kasha Rakestraw' ,'1980-01-31',null).
('Shad Hazeltine' ,'1980-01-31',null).
('Carman Stolz' ,'1980-01-31',null).
('Jospeh Massman' ,'1980-01-31',null).
('Ladonna Singleterry' ,'1980-01-31',null)
;
