CREATE TABLE bank (
	oid uuid DEFAULT uuid_generate_v4 (),
	bankCode VARCHAR NOT NULL,
	bankName VARCHAR NOT NULL,
	insertedDate TIMESTAMP NOT NULL,
		lastUpdate TIMESTAMP ,
		userInserted VARCHAR ,
		lastUserId VARCHAR ,
	PRIMARY KEY (oid)
);


 
CREATE TABLE permissionpolicyuser (
	oid uuid DEFAULT uuid_generate_v4 (),
	emailName VARCHAR NOT NULL,
	password VARCHAR NOT NULL,
	insertedDate TIMESTAMP NOT NULL,
		lastUpdate TIMESTAMP ,
		userInserted VARCHAR ,
		lastUserId VARCHAR ,
	PRIMARY KEY (oid)
);