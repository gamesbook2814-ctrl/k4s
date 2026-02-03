# k4s - Kubernetes TUI Project Plan

## Project Overview
Building a Terminal User Interface (TUI) for K3s cluster management, similar to k9s but named k4s. The tool will provide interactive debugging and monitoring capabilities for Kubernetes clusters with additional SSH/crictl integration.

## Technology Stack

### Core Libraries
- **[bubbletea](https://github.com/charmbracelet/bubbletea)** - Main TUI framework (event-driven model)
- **[bubbles](https://github.com/charmbracelet/bubbles)** - Pre-built UI components (lists, tables, viewports, spinners)
- **[lipgloss](https://github.com/charmbracelet/lipgloss)** - Styling and layout system

### Kubernetes Integration
- **[client-go](https://github.com/kubernetes/client-go)** - Official Kubernetes Go client
- **[k8s.io/api](https://pkg.go.dev/k8s.io/api)** - Kubernetes API types
- **[k8s.io/apimachinery](https://pkg.go.dev/k8s.io/apimachinery)** - Kubernetes API machinery

### SSH & Remote Execution
- **[golang.org/x/crypto/ssh](https://pkg.go.dev/golang.org/x/crypto/ssh)** - SSH client library

### Configuration
- **[viper](https://github.com/spf13/viper)** - Configuration file management (YAML)
- **[gopkg.in/yaml.v3](https://pkg.go.dev/gopkg.in/yaml.v3)** - YAML parsing

### Utilities
- **[color](https://github.com/fatih/color)** - Optional: color support for non-lipgloss areas

## MVP Development Plan

### Step 1: Project Setup & Basic Structure
**Goal:** Initialize project with dependencies and basic app skeleton

- [ ] Initialize Go module (`go mod init github.com/yourusername/k4s`)
- [ ] Install core dependencies (bubbletea, bubbles, lipgloss)
- [ ] Create basic bubbletea application that renders "Hello k4s"
- [ ] Set up main.go with proper error handling
- [ ] Create basic lipgloss theme/styles

**Deliverable:** Running TUI app that displays welcome screen

---

### Step 2: Configuration System
**Goal:** Load and manage kubeconfig files from ~/.k4s/config.yaml

- [ ] Create `~/.k4s/` directory structure on first run
- [ ] Define config.yaml structure (list of kubeconfig paths)
- [ ] Implement config loading with viper
- [ ] Create default config if none exists
- [ ] Display available kubeconfigs in TUI (list component)
- [ ] Implement kubeconfig selection/switching

**Config Format:**
```yaml
kubeconfigs:
  - name: "local-k3s"
    path: "/home/user/.kube/config"
    default: true
  - name: "production"
    path: "/home/user/.kube/prod-config"
```

**Deliverable:** Can select active kubeconfig from TUI

---

### Step 3: Kubernetes Client Integration
**Goal:** Connect to K8s cluster and retrieve basic data

- [ ] Install client-go dependency
- [ ] Create K8s client wrapper from selected kubeconfig
- [ ] Implement connection health check
- [ ] Display connection status in TUI header
- [ ] Handle connection errors gracefully
- [ ] Get current context and namespace

**Deliverable:** TUI shows "Connected to: cluster-name / namespace"

---

### Step 4: Namespace View
**Goal:** List and switch between namespaces

- [ ] Fetch all namespaces from cluster
- [ ] Display namespaces in a list (using bubbles list component)
- [ ] Show namespace status (Active/Terminating)
- [ ] Implement namespace selection (Enter key)
- [ ] Add namespace filtering/search (type to search)
- [ ] Display selected namespace in header

**Key Bindings:**
- Arrow keys / j,k - Navigate
- Enter - Select namespace
- / - Search mode
- Esc - Back

**Deliverable:** Can browse and select namespaces

---

### Step 5: Pods View (Core Feature)
**Goal:** Display pods in selected namespace with essential information

- [ ] Fetch pods for current namespace
- [ ] Create table view with columns:
  - Name
  - Ready (X/Y containers)
  - Status (Running/Pending/Failed/etc.)
  - Restarts
  - Age
  - Node (optional)
- [ ] Implement auto-refresh (every 5 seconds)
- [ ] Add status color coding (green=running, yellow=pending, red=failed)
- [ ] Show pod count in view header
- [ ] Handle empty namespace (no pods)

**Key Bindings:**
- Arrow keys / j,k - Navigate
- r - Refresh now
- Enter - View pod details

**Deliverable:** Live-updating pods list view

---

### Step 6: Pod Details View
**Goal:** Show detailed information about selected pod

- [ ] Display pod metadata (name, namespace, labels, annotations)
- [ ] Show pod conditions
- [ ] List containers with their status
- [ ] Display resource requests/limits
- [ ] Show pod events (last 10)
- [ ] Add navigation back to pods list

**Key Bindings:**
- Esc / q - Back to pods list
- l - View logs (next step)
- d - Delete pod (with confirmation)

**Deliverable:** Detailed pod information screen

---

### Step 7: Pod Operations
**Goal:** Perform basic pod operations

- [ ] **Delete Pod:**
  - Add confirmation dialog
  - Execute pod deletion via client-go
  - Show success/error message
  - Auto-refresh pods list
- [ ] **Restart Pod:**
  - Delete pod (K8s will recreate via deployment)
  - Confirmation required
- [ ] Add operation status notifications (bottom bar)

**Key Bindings:**
- d - Delete (with Y/N confirmation)
- Shift+R - Restart (with Y/N confirmation)

**Deliverable:** Can delete and restart pods from TUI

---

### Step 8: Log Viewer with Streaming
**Goal:** View and follow pod logs in real-time (CRITICAL for debugging)

- [ ] Implement log fetching for selected container
- [ ] Create scrollable log viewport (bubbles viewport component)
- [ ] **Add log streaming/following:**
  - Use K8s watch API for log stream
  - Auto-scroll to bottom on new logs
  - Toggle follow mode on/off
- [ ] If pod has multiple containers:
  - Show container selection menu
  - Display current container in header
- [ ] Add log filtering (search within logs)
- [ ] Implement log tail options (last 100/500/1000 lines)
- [ ] Add timestamps toggle
- [ ] Show container status (if crashed, show why)

**Key Bindings:**
- Esc / q - Back to pod details
- f - Toggle follow mode (like `kubectl logs -f`)
- / - Search in logs
- g - Jump to top
- G - Jump to bottom
- t - Toggle timestamps
- c - Change container (if multi-container pod)

**Deliverable:** Real-time log streaming viewer

---

### Step 9: Navigation & View Switching
**Goal:** Implement main navigation between different views

- [ ] Create main menu/navigation bar
- [ ] Implement view types:
  - Pods (default)
  - Namespaces
  - (Future: Deployments, Services, etc.)
- [ ] Add view switching shortcuts
- [ ] Show current view in header
- [ ] Maintain state when switching views (selected namespace, etc.)

**Key Bindings:**
- Tab / Shift+Tab - Cycle views
- 0 - Namespaces
- 1 - Pods
- ? - Help/keyboard shortcuts

**Deliverable:** Can navigate between Namespaces and Pods views

---

### Step 10: SSH & crictl Integration (Foundation)
**Goal:** Add SSH configuration and basic crictl execution

- [ ] Extend config.yaml with SSH hosts:
  ```yaml
  ssh_hosts:
    - name: "k3s-node-1"
      host: "192.168.1.100"
      user: "admin"
      key_path: "~/.ssh/id_rsa"
      port: 22
  ```
- [ ] Create SSH client wrapper
- [ ] Implement SSH connection test
- [ ] Add SSH hosts view (list configured hosts)
- [ ] Execute basic crictl command (e.g., `crictl ps`)
- [ ] Display crictl output in TUI
- [ ] Add crictl view for container inspection

**Key Bindings:**
- 9 - SSH/crictl view
- s - Select host
- Enter - Execute default command

**Deliverable:** Can connect via SSH and run crictl commands

---

### Step 11: Polish & Error Handling
**Goal:** Make the app production-ready

- [ ] Add comprehensive error handling:
  - Kubeconfig not found
  - Cluster unreachable
  - SSH connection failed
  - Invalid config.yaml
- [ ] Implement loading spinners for slow operations
- [ ] Add help screen (? key) with all keyboard shortcuts
- [ ] Create status bar showing:
  - Current cluster
  - Current namespace
  - Connection status
  - Last refresh time
- [ ] Add graceful shutdown (Ctrl+C)
- [ ] Write user documentation (README.md)

**Deliverable:** Stable, user-friendly MVP

---

### Step 12: Build & Distribution
**Goal:** Package for distribution

- [ ] Create build script for multiple platforms (Linux, macOS)
- [ ] Add version flag (`k4s --version`)
- [ ] Create installation instructions
- [ ] Set up GitHub repository
- [ ] Tag v0.1.0 release

**Deliverable:** Distributable k4s binary

---

## MVP Feature Checklist

### Must-Have (MVP Core)
- [x] Config management (~/.k4s/config.yaml)
- [x] Multiple kubeconfig support
- [x] Namespace browsing and switching
- [x] Pods list view with live updates
- [x] Pod details view
- [x] **Log viewing with streaming/follow mode** (essential for debugging)
- [x] Pod deletion
- [x] Pod restart
- [x] Basic SSH + crictl integration

### Nice-to-Have (Post-MVP)
- [ ] Deployments view
- [ ] Services view
- [ ] ConfigMaps/Secrets view
- [ ] Events view
- [ ] Port forwarding
- [ ] Exec into pod (shell)
- [ ] YAML view/edit
- [ ] Resource metrics (CPU/Memory)
- [ ] Multiple theme support

## Development Timeline Estimate

- **Week 1:** Steps 1-4 (Setup, Config, K8s Client, Namespaces)
- **Week 2:** Steps 5-7 (Pods View, Details, Operations)
- **Week 3:** Step 8 (Log Viewer with Streaming) - Critical feature
- **Week 4:** Steps 9-10 (Navigation, SSH/crictl)
- **Week 5:** Steps 11-12 (Polish, Build, Release)

**Total: 5 weeks for MVP**

## Success Criteria

The MVP is complete when:
1. Can manage multiple kubeconfigs from config file
2. Can browse namespaces and pods
3. Can view real-time streaming logs (like `kubectl logs -f`)
4. Can delete/restart pods
5. Can connect to nodes via SSH and run crictl
6. Handles errors gracefully
7. Has responsive, intuitive UI

## Next Steps

1. Start with Step 1: Initialize project
2. Get basic bubbletea app running
3. Iterate through steps 2-12
4. Test on real K3s cluster throughout development
5. Gather feedback and iterate

---

**Let's build k4s! 🚀**
