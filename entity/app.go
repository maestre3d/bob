package entity

import "time"

// AppModel App Configuration model
type AppModel struct {
	Name              string    `yaml:"name"`
	Version           string    `yaml:"version"`
	TotalApplications int       `yaml:"total_applications"`
	InstalledAt       time.Time `yaml:"intalled_at"`
	LastUsedAt        time.Time `yaml:"last_used_at"`
	Workspaces        []*Workspace
}

// Workspace App Workspace
type Workspace struct {
	Name        string    `yaml:"name"`
	Description string    `yaml:"description"`
	Path        string    `yaml:"path"`
	CreatedAt   time.Time `yaml:"created_at"`
	Services    []*Service
}

// Service App Microservice
type Service struct {
	Name        string    `yaml:"name"`
	Description string    `yaml:"description"`
	Path        string    `yaml:"path"`
	CreatedAt   time.Time `yaml:"created_at"`
}
