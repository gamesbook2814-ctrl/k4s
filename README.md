# k4s

A lightweight Terminal UI for K3s/Kubernetes cluster management, built on [Charm](https://charm.sh/).

![k4s demo](https://img.shields.io/badge/version-0.3.0-blue)
![Go Version](https://img.shields.io/badge/go-1.21+-00ADD8?logo=go)
![License](https://img.shields.io/badge/license-MIT-green)

## Screenshots

<p align="center">
  <img src="screenshots/kubeconfig.svg" alt="Kubeconfig selector" width="880"/>
</p>
<p align="center"><em>Kubeconfig selector — choose between multiple clusters</em></p>

<p align="center">
  <img src="screenshots/namespaces.svg" alt="Namespaces view" width="880"/>
</p>
<p align="center"><em>Namespaces view with sidebar navigation</em></p>

<p align="center">
  <img src="screenshots/pods.svg" alt="Pods view" width="880"/>
</p>
<p align="center"><em>Pods view with real-time CPU/Memory metrics</em></p>

<p align="center">
  <img src="screenshots/logs.svg" alt="Log viewer" width="880"/>
</p>
<p align="center"><em>Streaming log viewer with timestamps and search</em></p>

<p align="center">
  <img src="screenshots/events.svg" alt="Events view" width="880"/>
</p>
<p align="center"><em>Cluster events with follow mode and warning filtering</em></p>

## Features

- **Real-time Monitoring** - Live pods, deployments, services, events with auto-refresh
- **Resource Metrics** - CPU/Memory usage (requires metrics-server)
- **Multi-Pod Log Tailing** - Stream logs from multiple pods simultaneously with `Shift+L`
- **Streaming Logs** - Follow logs with search & highlighting
- **Crush-Inspired UI** - Purple-accented theme with sidebar layout and transparent overlays
- **SSH Integration** - Connect to nodes and inspect containers via crictl
- **Keyboard-driven** - Vim-style navigation

## Quick Start

```bash
# From source
git clone https://github.com/LywwKkA-aD/k4s.git
cd k4s && make install

# Or download from releases
# https://github.com/LywwKkA-aD/k4s/releases
```

Configuration is stored at `~/.k4s/config.yaml` (auto-created on first run).

## Keybindings

| Key | Action |
|-----|--------|
| `?` | Help |
| `1-5` | Switch views (Namespaces/Pods/Deployments/Services/Events) |
| `9` | SSH Hosts |
| `j/k` | Navigate |
| `Enter` | Select |
| `l` | Logs |
| `L` | Multi-pod logs |
| `q` | Quit |

See full documentation in [docs/](docs/).

## Contributing

Contributions welcome! Please check our [issues](https://github.com/LywwKkA-aD/k4s/issues) or open a new one.

## License

MIT License - see [LICENSE](LICENSE)

---

[![Stargazers repo roster for @LywwKkA-aD/k4s](https://reporoster.com/stars/LywwKkA-aD/k4s)](https://github.com/LywwKkA-aD/k4s/stargazers)
