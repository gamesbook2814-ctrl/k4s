# k4s Features Documentation

Comprehensive documentation of all features in k4s v0.2.0.

## Table of Contents

- [Overview](#overview)
- [Core Features](#core-features)
- [Views](#views)
- [Keyboard Navigation](#keyboard-navigation)
- [Configuration](#configuration)
- [Architecture](#architecture)

---

## Overview

**k4s** is a Terminal User Interface (TUI) for Kubernetes/K3s cluster management, inspired by [k9s](https://k9scli.io/). It provides an interactive, keyboard-driven interface for monitoring and managing Kubernetes clusters.

### Key Capabilities
- Multi-cluster management via kubeconfig files
- Real-time pod monitoring with auto-refresh
- Live log streaming with follow mode
- Pod lifecycle operations (delete, restart)
- Deployment management (scale, restart, delete)
- Services and Events views
- Resource metrics (CPU/Memory) when metrics-server available
- SSH integration for container runtime inspection
- Vim-style keyboard navigation

---

## Core Features

### 1. Multi-Kubeconfig Support

Manage multiple Kubernetes clusters from a single configuration file.

**Capabilities:**
- Load multiple kubeconfig files
- Switch between clusters on the fly
- Set default cluster for auto-selection
- Expand `~` paths automatically

**Configuration:**
```yaml
kubeconfigs:
  - name: "local-k3s"
    path: "~/.kube/config"
    default: true
  - name: "production"
    path: "~/.kube/prod-config"
```

### 2. Real-time Pod Monitoring

Live-updating pod list with essential information at a glance.

**Information Displayed:**
| Column | Description |
|--------|-------------|
| Name | Pod name |
| Ready | Ready containers (X/Y) |
| Status | Pod phase with color coding |
| Restarts | Total container restarts |
| Age | Time since creation |
| CPU | Current CPU usage (toggle with `m`) |
| Memory | Current memory usage (toggle with `m`) |

**Status Color Coding:**
| Status | Color |
|--------|-------|
| Running | Green |
| Pending | Yellow |
| Failed | Red |
| Succeeded | Blue |
| Unknown | Gray |

**Auto-refresh:** Every 5 seconds (preserves filter state)

### 3. Streaming Log Viewer

View and follow pod logs in real-time, similar to `kubectl logs -f`.

**Features:**
- Real-time log streaming with follow mode
- Multi-container pod support with container selector
- Search highlighting (press `/` to search)
- Timestamp toggle (press `t`)
- Configurable tail lines (default: 500)
- Scroll navigation (vim-style and arrow keys)

**Log Viewer Controls:**
| Key | Action |
|-----|--------|
| `f` | Toggle follow mode |
| `t` | Toggle timestamps |
| `/` | Search in logs |
| `n` | Next search match |
| `N` | Previous match |
| `g` | Go to top |
| `G` | Go to bottom |
| `c` | Change container |

### 4. Pod Operations

Perform essential pod operations with safety confirmations.

**Delete Pod:**
- Confirmation dialog required
- Graceful deletion via Kubernetes API
- Success/error notification
- Auto-refresh of pod list

**Restart Pod:**
- Confirmation dialog required
- Deletes pod (Kubernetes recreates via controller)
- Useful for deployment-managed pods

### 5. SSH Integration

Connect to K3s nodes via SSH for direct container runtime access.

**Authentication Methods:**
1. **SSH Agent** (Recommended)
   - Uses existing ssh-agent for key management
   - No passphrase prompts needed

2. **Private Key with Passphrase**
   - Prompts for passphrase when key is protected
   - Secure passphrase input (masked)

**Configuration:**
```yaml
ssh_hosts:
  - name: "k3s-node-1"
    host: "192.168.1.100"
    user: "admin"
    key_path: "~/.ssh/id_rsa"
    port: 22
```

### 6. crictl Integration

Inspect containers at the runtime level via crictl on remote nodes.

**Capabilities:**
- List all containers on a node
- View container logs directly
- See node system information:
  - Hostname
  - Operating system
  - Memory usage
  - Load average

**Use Cases:**
- Debug container issues at runtime level
- Inspect containers not visible in Kubernetes API
- View system-level container information

---

## Views

### Kubeconfig Selector
- Displayed on startup if multiple kubeconfigs configured
- Auto-skipped if only one kubeconfig or default is set
- Navigate and select active cluster

### Namespace Browser (`1`)
- List all namespaces in cluster
- Show namespace status (Active/Terminating)
- Filter namespaces by name
- Select namespace for pod viewing

### Pod List (`2`)
- Display pods in selected namespace
- Real-time status updates
- Color-coded status indicators
- Filter pods by name
- Select pod for details or operations
- Toggle CPU/Memory metrics (`m`)

### Deployments (`3`)
- List all deployments in namespace
- Show ready/desired replica counts
- Color-coded status (green=ready, yellow=partial, red=unavailable)
- Scale deployments (`s`) - interactive replica input
- Restart deployments (`R`) - rolling restart
- Delete deployments (`d`) with confirmation

### Services (`4`)
- List all services in namespace
- Show type, cluster IP, external IP, ports
- View service details
- Color-coded by type (LoadBalancer, NodePort, ClusterIP)

### Events (`5`)
- Cluster-wide events in log-style format
- Auto-follow mode (`f`) - scroll to new events
- Filter warnings only (`w`)
- Filter by resource kind (`k`) - cycle through Pod, Deployment, ReplicaSet, Service, Node
- Color-coded by event type (red=Warning, muted=Normal)

### Pod Details (Enter)
- Comprehensive pod information:
  - Metadata (name, namespace, UID)
  - Labels and annotations
  - Container specifications
  - Resource requests/limits
  - Recent events (last 10)
- Scrollable view for long content

### Log Viewer (`l`)
- Full-screen log display
- Real-time streaming capability
- Search with highlighting
- Container selection for multi-container pods

### SSH Host Selector (`9`)
- List configured SSH hosts
- Connection status indicators
- Select host for crictl access

### crictl Container List
- List containers via `crictl ps`
- Show container ID, name, state
- Select container for log viewing

### crictl Log Viewer
- Stream container logs via SSH
- Similar controls to Kubernetes log viewer

---

## Keyboard Navigation

### Global Shortcuts
| Key | Action |
|-----|--------|
| `?` | Show/hide help |
| `q` | Quit application |
| `Ctrl+C` | Force quit |
| `Esc` | Go back / Cancel |
| `r` | Refresh current view |

### List Navigation
| Key | Action |
|-----|--------|
| `↑` / `k` | Move up |
| `↓` / `j` | Move down |
| `Enter` | Select / Open |
| `/` | Filter / Search |
| `g` | Go to top |
| `G` | Go to bottom |

### View Switching
| Key | View |
|-----|------|
| `1` | Namespaces |
| `2` | Pods |
| `3` | Deployments |
| `4` | Services |
| `5` | Events |
| `9` | SSH Hosts |

### Pod Actions
| Key | Action |
|-----|--------|
| `l` | View logs |
| `d` | Delete pod |
| `R` | Restart pod (Shift+R) |
| `m` | Toggle metrics |
| `Enter` | View details |

### Deployment Actions
| Key | Action |
|-----|--------|
| `s` | Scale deployment |
| `d` | Delete deployment |
| `R` | Restart deployment |
| `Enter` | View details |

### Events Actions
| Key | Action |
|-----|--------|
| `f` | Toggle follow mode |
| `w` | Toggle warnings only |
| `k` | Cycle kind filter |

### Scrolling (Details/Logs)
| Key | Action |
|-----|--------|
| `↑` / `↓` | Scroll line |
| `PgUp` / `PgDn` | Scroll page |
| `Home` / `g` | Go to top |
| `End` / `G` | Go to bottom |

---

## Configuration

### Configuration File Location
```
~/.k4s/config.yaml
```

### Full Configuration Schema

```yaml
# Kubernetes cluster configurations
kubeconfigs:
  - name: string        # Display name for cluster
    path: string        # Path to kubeconfig file (supports ~)
    default: bool       # Auto-select this cluster (optional)

# SSH host configurations for crictl access
ssh_hosts:
  - name: string        # Display name for host
    host: string        # Hostname or IP address
    user: string        # SSH username
    key_path: string    # Path to SSH private key (supports ~)
    port: int           # SSH port (default: 22)
```

### Log File Location
```
~/.k4s/logs/k4s-YYYY-MM-DD.log
```

### Auto-created Files
On first run, k4s creates:
- `~/.k4s/` directory
- `~/.k4s/config.yaml` with default template
- `~/.k4s/logs/` directory for debug logs

---

## Architecture

### Technology Stack

| Component | Library |
|-----------|---------|
| TUI Framework | charmbracelet/bubbletea v1.3.10 |
| UI Components | charmbracelet/bubbles v0.21.0 |
| Styling | charmbracelet/lipgloss v1.1.0 |
| Kubernetes Client | k8s.io/client-go v0.35.0 |
| Kubernetes Metrics | k8s.io/metrics v0.35.0 |
| SSH Client | golang.org/x/crypto v0.47.0 |
| Configuration | spf13/viper v1.21.0 |

### Project Structure

```
k4s/
├── cmd/k4s/main.go           # Entry point
├── internal/
│   ├── domain/               # Pure domain models
│   │   ├── config.go         # Configuration models
│   │   ├── pod.go            # Pod/Container models
│   │   ├── namespace.go      # Namespace model
│   │   ├── cluster.go        # Cluster connection models
│   │   └── errors.go         # Domain errors
│   ├── adapter/
│   │   ├── config/           # Config file loading
│   │   ├── k8s/              # Kubernetes client
│   │   │   ├── client.go     # Client creation/connection
│   │   │   ├── pod.go        # Pod operations
│   │   │   └── logs.go       # Log streaming
│   │   ├── ssh/              # SSH client
│   │   │   ├── client.go     # SSH connections
│   │   │   └── crictl.go     # crictl parsing
│   │   └── tui/              # Terminal UI
│   │       ├── app.go        # Main application
│   │       ├── styles.go     # Lipgloss styles
│   │       └── [views].go    # Individual views
│   └── logger/               # File-based logging
├── Makefile                  # Build automation
└── go.mod                    # Go module
```

### Design Principles

1. **Clean Architecture** - Domain logic separated from adapters
2. **Hexagonal Architecture** - Ports and adapters pattern
3. **Event-driven** - Bubbletea message passing model
4. **Async Operations** - Non-blocking I/O for all external calls
5. **Graceful Degradation** - Non-fatal logging, error recovery

---

## Platform Support

| Platform | Architecture | Status |
|----------|--------------|--------|
| Linux | amd64 | Supported |
| Linux | arm64 | Supported |
| macOS | amd64 | Supported |
| macOS | arm64 (Apple Silicon) | Supported |
| Windows | - | Not supported |

---

## Dependencies

### Runtime Requirements
- Access to a Kubernetes cluster (kubeconfig)
- SSH access to nodes (for crictl features)
- Terminal with 256-color support (recommended)

### Build Requirements
- Go 1.21 or later
- Make (for build automation)
- golangci-lint (for linting)
