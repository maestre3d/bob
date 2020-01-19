package usecase

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/maestre3d/bob/common/util"
	"github.com/maestre3d/bob/entity"
)

// GenerateService Create a new service
func GenerateService(name, appName, description string) error {
	// Verify if not exists then insert
	app := util.GetSettings()
	serviceID := 0
	currentWorkspace := util.GetCurrentWorkspace()
	if appName == "" && currentWorkspace == "" {
		return errors.New("Failed to create service, must be inside an application folder")
	}

	for i, workspace := range app.Workspaces {
		if workspace.Name == strings.ToLower(appName) || workspace.Path == currentWorkspace {
			serviceID = i
			for _, serviceWork := range workspace.Services {
				if serviceWork.Name == strings.ToLower(name) {
					return errors.New("Service already exists")
				}
			}
		}
	}

	service := new(entity.Service)
	service.Name = name
	service.Description = description
	service.Path = app.Workspaces[serviceID].Path + name + "/"
	service.CreatedAt = time.Now()

	app.Workspaces[serviceID].Services = append(app.Workspaces[serviceID].Services, service)
	log.Printf("SERVICE: Created %s/%s service\n", app.Workspaces[serviceID].Name, name)

	return util.OverrideSettings(app)
}

// GetAllServiceInfo Get all service(s) information
func GetAllServiceInfo(appName string) error {
	w := new(tabwriter.Writer)
	app := util.GetSettings()

	w.Init(os.Stdout, 0, 8, 2, '\t', tabwriter.Debug|tabwriter.AlignRight)
	fmt.Fprintln(w, "Service\tPath\tWorkspace\tCreated At")

	for _, workspace := range app.Workspaces {
		if workspace.Name == strings.ToLower(appName) {
			for _, service := range workspace.Services {
				fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", service.Name, service.Path, workspace.Name, service.CreatedAt)
			}
		}
	}
	fmt.Fprintln(w)
	w.Flush()

	return nil
}

// GetServiceInfo Get service information
func GetServiceInfo(appName, name string) error {
	w := new(tabwriter.Writer)
	app := util.GetSettings()

	w.Init(os.Stdout, 0, 8, 2, '\t', tabwriter.Debug|tabwriter.AlignRight)
	fmt.Fprintln(w, "Service\tPath\tWorkspace\tCreated At")

	for _, workspace := range app.Workspaces {
		if workspace.Name == strings.ToLower(appName) {
			for _, service := range workspace.Services {
				if service.Name == strings.ToLower(name) {
					fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", service.Name, service.Path, workspace.Name, service.CreatedAt)
				}
			}
		}
	}
	fmt.Fprintln(w)
	w.Flush()

	return nil
}

// RemoveService Remove an existing service
func RemoveService(appName, name string) error {
	log.Printf("SERVICE: Removed %s/%s service\n", appName, name)
	return nil
}

func removeService(s []*entity.Workspace, i int) []*entity.Workspace {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
