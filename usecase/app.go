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

// CreateApp Create new app
func CreateApp(appName, path, description string) error {
	// Get current dir
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	if path != "" && path != "." {
		if sPath := strings.Split(path, "/"); sPath[len(sPath)-1] != "/" {
			path = path + "/" + appName + "/"
		} else {
			path = path + appName + "/"
		}
	}

	// Set default path if not dir specified
	if path == "" || path == "." {
		path = dir + "/" + appName + "/"
	}

	// Sanitize app name
	appName = strings.ToLower(appName)

	// Verify paths and names to avoid duplicates
	app := util.GetSettings()
	for _, workspace := range app.Workspaces {
		if workspace.Name == appName || workspace.Path == path {
			return errors.New("Application already exists")
		}
	}

	// Create a new workspace and write
	workspace := &entity.Workspace{
		Name:        appName,
		Description: description,
		Path:        path,
		CreatedAt:   time.Now(),
	}

	app.Workspaces = append(app.Workspaces, workspace)
	app.TotalApplications++
	log.Printf("APPLICATION: Created new %s application\n", appName)

	return util.OverrideSettings(app)
}

// GetAllAppInfo Get all application(s) information
func GetAllAppInfo() error {
	w := new(tabwriter.Writer)
	app := util.GetSettings()
	isFound := false

	w.Init(os.Stdout, 0, 8, 2, '\t', tabwriter.Debug|tabwriter.AlignRight)
	fmt.Fprintln(w, "Application\tPath\tCreated At")

	for _, workspace := range app.Workspaces {
		fmt.Fprintf(w, "%s\t%s\t%s\n", workspace.Name, workspace.Path, workspace.CreatedAt)
		isFound = true
	}
	fmt.Fprintln(w)
	w.Flush()

	if isFound {
		fmt.Println("Use -> nucleon service {app_name} <- to see all services")
	} else {
		fmt.Println("Application(s) not found")
	}

	return nil
}

// GetAppInfo Get application information
func GetAppInfo(name string) error {
	w := new(tabwriter.Writer)
	app := util.GetSettings()

	w.Init(os.Stdout, 0, 8, 2, '\t', tabwriter.Debug|tabwriter.AlignRight)
	fmt.Fprintln(w, "Application\tPath\tCreated At")
	isFound := false
	for _, workspace := range app.Workspaces {
		if workspace.Name == strings.ToLower(name) {
			fmt.Fprintf(w, "%s\t%s\t%s\n", workspace.Name, workspace.Path, workspace.CreatedAt)
			isFound = true
		}
	}
	fmt.Fprintln(w)
	w.Flush()

	if isFound {
		fmt.Printf("Use -> nucleon service %s <- to see all %s services\n", name, name)
	} else {
		fmt.Println("Application not found")
	}

	return nil
}

// RemoveApp Remove an existing app
func RemoveApp(name string) error {
	app := util.GetSettings()
	deleted := false
	for i, workspace := range app.Workspaces {
		if workspace.Name == strings.ToLower(name) {
			app.Workspaces = removeApp(app.Workspaces, i)
			app.TotalApplications--
			deleted = true
			log.Printf("APPLICATION: Removed %s application\n", name)
		}
	}
	if !deleted {
		fmt.Print("Application not found")
	}

	return util.OverrideSettings(app)
}

func removeApp(s []*entity.Workspace, i int) []*entity.Workspace {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}
