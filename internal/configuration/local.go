package configuration

var local Configuration = Configuration{
	Version:     "0.1",
	RunningMode: "DEBUG",
	Database: Database{
		Host:         "localhost",
		Port:         "3306",
		UserName:     "root",
		UserPassword: "UnderAll4",
		Name:         "duval",
	},
	RunHash: "",
}
