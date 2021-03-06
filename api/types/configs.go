// ./docker/api/types/configs.go  //peter


package types

import "github.com/docker/docker/api/types/container"

// configs holds structs used for internal communication between the
// frontend (such as an http server) and the backend (such as the
// docker daemon).

// ContainerCreateConfig is the parameter set to ContainerCreate()
type ContainerCreateConfig struct {
	Name            string
	Config          *container.Config
	HostConfig      *container.HostConfig
	AdjustCPUShares bool
}

// ContainerRmConfig holds arguments for the container remove
// operation. This struct is used to tell the backend what operations
// to perform.
type ContainerRmConfig struct {
	ForceRemove, RemoveVolume, RemoveLink bool
}

// ContainerCommitConfig contains build configs for commit operation,
// and is used when making a commit with the current state of the container.
type ContainerCommitConfig struct {
	Pause   bool
	Repo    string
	Tag     string
	Author  string
	Comment string
	// merge container config into commit config before commit
	MergeConfigs bool
	Config       *container.Config
}

// CriuConfig holds configuration options passed down to libcontainer and CRIU
type CriuConfig struct {
	ImagesDirectory string
	WorkDirectory   string
	LeaveRunning    bool
        PrevImagesDirectory   string  //peter 
        TrackMemory     bool //peter
        EnablePreDump     bool //peter
        AutoDedup         bool //peter
        PageServer        bool //peter
        Address           string  //peter 
        Port              int32   //peter
}

// RestoreConfig holds the restore command options, which is a superset of the CRIU options
type RestoreConfig struct {
	CriuOpts     CriuConfig
	ForceRestore bool
}

// ExecConfig is a small subset of the Config struct that hold the configuration
// for the exec feature of docker.
type ExecConfig struct {
	User         string   // User that will run the command
	Privileged   bool     // Is the container in privileged mode
	Tty          bool     // Attach standard streams to a tty.
	Container    string   // Name of the container (to execute in)
	AttachStdin  bool     // Attach the standard input, makes possible user interaction
	AttachStderr bool     // Attach the standard output
	AttachStdout bool     // Attach the standard error
	Detach       bool     // Execute in detach mode
	Cmd          []string // Execution commands and args
}
