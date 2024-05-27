package plugins

func init() {
	smbEnumPlugin := &BasePlugin{
		Name:        "SMBEnum",
		Command:     "enum4linux-ng",
		Args:        []string{"-A", "{ip}"},
		IsHostBased: false,
	}
	RegisterPlugin(smbEnumPlugin, []string{"139", "445"}, []string{"microsoft-ds"})
}
