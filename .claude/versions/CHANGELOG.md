# k4s Changelog

All notable changes to k4s will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.2.0] - 2026-02-04

### Added

#### New Views
- **Deployments view** (key `3`) - List all deployments with ready/desired replica counts
  - Scale deployments (`s`) with interactive replica input
  - Restart deployments (`R`) - rolling restart
  - Delete deployments (`d`) with confirmation
  - Color-coded status (green=ready, yellow=partial, red=unavailable)
- **Services view** (key `4`) - List all services with type, cluster IP, ports
  - Service details with full port mappings
  - Color-coded by type (LoadBalancer, NodePort, ClusterIP)
- **Events view** (key `5`) - Cluster-wide events in log-style format
  - Auto-follow mode (`f`) - scroll to new events
  - Filter warnings only (`w`)
  - Filter by resource kind (`k`) - cycle through Pod, Deployment, ReplicaSet, Service, Node
  - Color-coded by event type (red=Warning, muted=Normal)

#### Pod Metrics
- **Metrics support** (`m`) - Toggle CPU/Memory columns in pod list
  - Displays current CPU usage (e.g., "15m" = 15 millicores)
  - Displays current memory usage (e.g., "128Mi")
  - Requires metrics-server to be installed in cluster
  - Graceful fallback when metrics unavailable

#### Log Viewer Improvements
- **ANSI color preservation** - Source log colors are preserved
- **Smart line wrapping** - Wraps at terminal width without breaking ANSI sequences
- **Proper back navigation** - Esc returns to correct source view (pods list, not pod details)

### Changed
- **Navigation keys** - Changed from `0-4` to `1-5` for view switching
  - `1` = Namespaces
  - `2` = Pods
  - `3` = Deployments
  - `4` = Services
  - `5` = Events
  - `9` = SSH (unchanged)
- **Help screen** - Redesigned to 3-column layout (wider, not taller)
  - Column 1: Global + Navigation
  - Column 2: Pods + Deployments
  - Column 3: Events + Logs

### Fixed
- Up/down navigation not working in Deployments, Services, Events views
- Log viewer Esc returning to wrong view (pod details instead of pod list)
- Long log lines going off screen instead of wrapping
- Footer shifting up when viewing logs with wrapped lines
- Metrics client not being initialized on connect
- Metrics showing "-" instead of actual values

---

## [0.1.0] - 2025-01-XX

### Initial Release

First public release of k4s - Kubernetes TUI for K3s cluster management.

### Added

#### Core Features
- **Multi-kubeconfig support** - Manage multiple K3s/Kubernetes clusters from `~/.k4s/config.yaml`
- **Real-time pod monitoring** - Live-updating pod list with 5-second auto-refresh
- **Streaming log viewer** - Real-time log streaming with follow mode (like `kubectl logs -f`)
- **Pod operations** - Delete and restart pods with confirmation dialogs
- **SSH integration** - Connect to K3s nodes via SSH for container runtime inspection
- **crictl integration** - View containers, logs, and node info directly via crictl

#### User Interface
- Terminal User Interface built with Bubbletea framework
- Vim-style keyboard navigation (`j`/`k`, `g`/`G`)
- Color-coded pod status indicators
- Search/filter functionality in lists and logs
- Help screen with keyboard shortcuts (`?`)
- Toast-style notifications for operations

#### Views
- **Kubeconfig selector** - Choose active cluster configuration
- **Namespace browser** - Browse and select Kubernetes namespaces
- **Pod list** - View all pods with status, restarts, and age
- **Pod details** - Detailed pod information including events
- **Log viewer** - Scrollable logs with search highlighting
- **SSH host selector** - Choose configured SSH hosts
- **crictl container list** - View containers on remote nodes
- **crictl log viewer** - Stream container logs via SSH

#### Configuration
- Auto-created configuration on first run
- Support for multiple kubeconfig paths
- SSH host configuration with key-based authentication
- SSH agent support for passphrase-protected keys

#### Technical
- Clean hexagonal architecture with domain/adapter separation
- File-based debug logging to `~/.k4s/logs/`
- Graceful error handling with user-friendly messages
- Cross-platform support (Linux amd64/arm64, macOS amd64/arm64)

### Build Information
- **Version**: Set via ldflags at build time
- **Commit**: Git commit hash embedded in binary
- **Build Date**: Timestamp of build
- **Go Version**: Runtime Go version displayed

---

## Version Information

### Current Version
- **Version**: 0.1.0
- **Status**: MVP (Minimum Viable Product)
- **Go Version**: 1.21+
- **License**: MIT

### Version Flags
```bash
k4s --version
k4s -v
```

### Version Output Format
```
k4s - Kubernetes TUI for K3s
  Version:    0.1.0
  Commit:     abc1234
  Built:      2025-01-15T10:30:00Z
  Go version: go1.21.0
  OS/Arch:    linux/amd64
```

### Build Variables (ldflags)
| Variable | Description |
|----------|-------------|
| `Version` | Semantic version string |
| `Commit` | Git commit hash (short) |
| `BuildDate` | ISO 8601 build timestamp |

### Build Command
```bash
go build -ldflags "-X main.Version=0.1.0 -X main.Commit=$(git rev-parse --short HEAD) -X main.BuildDate=$(date -u +%Y-%m-%dT%H:%M:%SZ)" ./cmd/k4s
```

---

## Roadmap

### Planned for Future Releases

#### v0.2.0 (Released 2026-02-04)
- [x] Deployments view
- [x] Services view
- [x] Events view
- [x] Resource metrics (CPU/Memory)

#### v0.3.0 (Planned)
- [ ] ConfigMaps/Secrets view
- [ ] Port forwarding
- [ ] Exec into pod (shell)

#### v1.0.0 (Future)
- [ ] YAML view/edit
- [ ] Multiple theme support
- [ ] Plugin system
- [ ] Custom keybindings

---

## Release Process

1. Update version in Makefile
2. Update CHANGELOG.md
3. Create git tag: `git tag -a v0.1.0 -m "Release v0.1.0"`
4. Build releases: `make release`
5. Push tag: `git push origin v0.1.0`
6. Create GitHub release with artifacts
