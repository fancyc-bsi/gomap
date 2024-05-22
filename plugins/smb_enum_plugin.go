package plugins

func init() {
	smbEnumPlugin := &BasePlugin{
		Name:        "SMBEnum",
		Command:     "smbclient",
		Args:        []string{"-L", "//{ip}", "-N"},
		IsHostBased: false,
	}
	RegisterPlugin(smbEnumPlugin, []string{"139", "445"}, []string{"microsoft-ds"})
}
