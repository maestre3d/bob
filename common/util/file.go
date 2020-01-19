package util

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	global "github.com/maestre3d/bob/common/config"
	"github.com/maestre3d/bob/entity"
	"gopkg.in/yaml.v2"
)

var (
	usrDir string
)

func init() {
	usr, err := os.UserHomeDir()
	if err != nil {
		log.Panic(err)
	}
	usrDir = usr

	if _, err := os.Stat(usr + global.UniversalPath); os.IsNotExist(err) {
		// Create App path
		err = os.Mkdir(usr+global.UniversalPath, os.ModePerm)
		if err != nil {
			log.Panic(err)
		}

		if err != nil {
			log.Panic(err)
		}
	}

	if _, err = ioutil.ReadFile(usr + global.SettingFilePath); err != nil {
		err = CreateSettingsFile()
		if err != nil {
			log.Panic(err)
		}
	}
}

// CreateSettingsFile Create the initial App File
func CreateSettingsFile() error {
	// Create new App model
	app := new(entity.AppModel)
	app.Name = global.AppName + " stable - The Microservice builder"
	app.Version = "1.0"
	app.TotalApplications = 0
	app.InstalledAt, app.LastUsedAt = time.Now(), time.Now()

	// Marshal to YAML
	out, err := yaml.Marshal(&app)
	if err != nil {
		return err
	}

	// Create file
	err = ioutil.WriteFile(usrDir+global.SettingFilePath, out, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

// GetSettings Get Settings
func GetSettings() *entity.AppModel {
	in, err := ioutil.ReadFile(usrDir + global.SettingFilePath)
	if err != nil {
		return nil
	}

	app := new(entity.AppModel)

	err = yaml.Unmarshal(in, &app)
	if err != nil {
		return nil
	}

	return app
}

// OverrideSettings Override existing settings
func OverrideSettings(app *entity.AppModel) error {
	app.LastUsedAt = time.Now()
	out, err := yaml.Marshal(app)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(usrDir+global.SettingFilePath, out, os.ModePerm)
}

// GetCurrentWorkspace Get current workspace
func GetCurrentWorkspace() string {
	workString := ""

	// Get current dir
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	currentDir := strings.Split(dir+"/", "/")
	c := 0

	app := GetSettings()
	for _, workspace := range app.Workspaces {
		root := strings.Split(workspace.Path, "/")
		if len(currentDir) < len(root) {
			return ""
		}

		for i, path := range root {
			if path != currentDir[i] {
				return workString
			}

			c++
		}

		if c < len(root) {
			return workString
		}

		workString = workspace.Path
	}

	return workString
}
