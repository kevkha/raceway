package config

import (
	"encoding/json"
	"os"
)

// Raceway holds the base settings for the app
type Raceway struct {
	AppPort  string `json:"app_port"`
	AppDebug bool   `json:"app_debug"`
}

// Database
type Database struct {
	DBHost  string `json:"db_host"`
	DBPort  int    `json:"db_port"`
	DBUser  string `json:"db_user"`
	DBPass  string `json:"db_pass"`
	DBName  string `json:"db_name"`
	DBDebug bool   `json:"db_debug"`
}

// Scenarios
type Scenarios struct {
	ScenarioDir    string `json:"scenario_dir"`
	ScenarioFormat string `json:"scenario_format"`
}

// Graphite
type Graphite struct {
	GraphiteHost string `json:"graphite_host"`
}

// Scheduler
type Scheduler struct {
	SchedulerDBBucket  string `json:"scheduler_db_bucket"`
	SchedulerDBName    string `json:"scheduler_db_name"`
	SchedulerDBTimeout int    `json:"scheduler_db_timeout"`
}

// Config contains the Raceway configuration
type Config struct {
	Raceway       Raceway   `json:"raceway"`
	Database      Database  `json:"database"`
	Scenarios     Scenarios `json:"scenarios"`
	Graphite      Graphite  `json:"graphite"`
	Scheduler     Scheduler `json:"scheduler"`
	HTMLOutputDir string    `json:"html_output_dir"`
}

// Load builds a config obj
func Load(cf string) (*Config, error) {
	confFile, err := os.Open(cf)
	if err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(confFile)
	var conf Config
	if err = decoder.Decode(&conf); err != nil {
		return nil, err
	}
	return &conf, nil
}
