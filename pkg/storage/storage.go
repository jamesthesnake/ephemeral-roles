package storage

// Client is a storage interface to abstract from the system used to implement
// storage
type Client interface {
	List() ([]string, error)
	Store(config *ServerConfig)
	Retrieve(server string) *ServerConfig
	Delete(server string)
}

// ServerConfig is a custom configuration for a Discord server
type ServerConfig struct {
	Name                  string          `json:"name"`
	RolePrefix            string          `json:"rolePrefix"`
	RoleColor             string          `json:"roleColor"`
	VoiceChannelWhitelist map[string]bool `json:"voiceChannelWhitelist"`
	VoiceChannelBlacklist map[string]bool `json:"voiceChannelBlacklist"`
}
