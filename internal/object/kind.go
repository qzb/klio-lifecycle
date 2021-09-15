package object

type Kind string

const (
	BuilderKind     Kind = "Builder"
	DeployerKind    Kind = "Deployer"
	EnvironmentKind Kind = "Environment"
	ProjectKind     Kind = "Project"
	PusherKind      Kind = "Pusher"
	ServiceKind     Kind = "Service"
	TaggerKind      Kind = "Tagger"
)
