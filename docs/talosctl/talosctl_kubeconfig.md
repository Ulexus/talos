<!-- markdownlint-disable -->
## talosctl kubeconfig

Download the admin kubeconfig from the node

### Synopsis

Download the admin kubeconfig from the node.
If merge flag is defined, config will be merged with ~/.kube/config or [local-path] if specified.
Otherwise kubeconfig will be written to PWD or [local-path] if specified.

```
talosctl kubeconfig [local-path] [flags]
```

### Options

```
  -f, --force                       Force overwrite of kubeconfig if already present, force overwrite on kubeconfig merge
      --force-context-name string   Force context name for kubeconfig merge
  -h, --help                        help for kubeconfig
  -m, --merge                       Merge with existing kubeconfig (default true)
```

### Options inherited from parent commands

```
      --context string       Context to be used in command
  -e, --endpoints strings    override default endpoints in Talos configuration
  -n, --nodes strings        target the specified nodes
      --talosconfig string   The path to the Talos configuration file (default "/home/user/.talos/config")
```

### SEE ALSO

* [talosctl](talosctl.md)	 - A CLI for out-of-band management of Kubernetes nodes created by Talos

