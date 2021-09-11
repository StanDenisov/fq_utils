package confstruct

//ConfStruct - structer for configurating microservices
type ConfStruct struct {
	AppMode       string `json:"app_mode,omitempty"`
	AppName       string `json:"app_name,omitempty"`
	PgDBPort      string `json:"pg_db_port,omitempty"`
	PgDBPassword  string `json:"pg_db_password,omitempty"`
	PgDBUsernname string `json:"pg_db_usernname,omitempty"`
	PgDBName      string `json:"pg_db_name,omitempty"`
}
