package global

const (
	AppName         string = "Nucleon"
	ShortName       string = "nucleon"
	GenericError    string = AppName + " Error: %s\n"
	ServiceError    string = "Missing service name\nTry with the following command:\n" + ShortName + " generate service {name}"
	PathError       string = "Missing path name\nTry with the following command:\n" + ShortName + " generate service {name} {path}"
	UniversalPath   string = "/." + ShortName + "/"
	SettingFilePath string = UniversalPath + ShortName + ".yaml"
)
