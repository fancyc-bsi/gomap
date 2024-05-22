package plugins

func init() {
	rpcLdapPlugin := &BasePlugin{
		Name:        "RPCLDAP",
		Command:     "enum4linux-ng",
		Args:        []string{"-A", "--rpc", "{ip}"},
		IsHostBased: false,
	}
	RegisterPlugin(rpcLdapPlugin, []string{"135"}, []string{"msrpc", "ldap"})
}
