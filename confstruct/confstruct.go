package confstruct

//ConfStruct - structer for configurating microservices
type ConfStruct struct {
	AppPort      string `json:"app_port,omitempty"`
	AppMode      string `json:"app_mode,omitempty"`
	AppName      string `json:"app_name,omitempty"`
	PgDBPort     string `json:"pg_db_port,omitempty"`
	PgDBPassword string `json:"pg_db_password,omitempty"`
	PgDBUserName string `json:"pg_db_user_name,omitempty"`
	PgDBName     string `json:"pg_db_name,omitempty"`
}
