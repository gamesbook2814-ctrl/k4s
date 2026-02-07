# Views

## Namespaces View (`1`)

Browse and select Kubernetes namespaces.

**Columns:**
- Namespace name
- Status
- Age

## Pods View (`2`)

List all pods in the selected namespace.

**Columns:**
- Pod name
- Ready containers (X/Y)
- Status (color-coded)
- Restart count
- Age
- CPU/Memory (toggle with `m`, requires metrics-server)

**Features:**
- Auto-refreshes every 5 seconds
- Color-coded status (Running=green, Pending=yellow, Failed=red)

**Actions:** `l` logs, `L` multi-pod logs, `d` delete, `R` restart, `m` metrics

## Pod Details (Enter on pod)

Detailed information about a pod.

**Sections:**
- Metadata (labels, annotations)
- Container information
- Resource requests/limits
- Recent events

## Deployments View (`3`)

List all deployments in the selected namespace.

**Columns:**
- Deployment name
- Ready/Desired replicas
- Up-to-date count
- Available count
- Age

**Actions:** `s` scale, `d` delete, `R` restart

## Services View (`4`)

List all services in the selected namespace.

**Columns:**
- Service name
- Type (ClusterIP, NodePort, LoadBalancer)
- Cluster IP
- External IP
- Ports
- Age

## Events View (`5`)

Cluster-wide events in log-style format.

**Features:**
- Auto-follow mode (`f`)
- Filter warnings only (`w`)
- Filter by resource kind (`k`)
- Color-coded by event type (Normal=muted, Warning=red)

## Log Viewer (`l`)

View pod logs with streaming support.

**Features:**
- Follow mode for real-time logs
- Multi-container support (press `c` to switch)
- Timestamp toggle (`t`)
- Search with highlighting (`/`, `n`, `N`)
- ANSI color preservation

## Multi-Pod Log Viewer (`Shift+L`)

View streaming logs from multiple pods simultaneously.

**Features:**
- Select pods via multi-select dialog with `* All Pods` option
- Interleaved log output with `==> pod-name/container <==` headers between sections
- Follow mode with auto-scroll (`f` to toggle)
- Auto-reconnects individual streams on connection drops
- First container per pod auto-selected

**Format:**
```
==> my-app-abc/app <==
2026-02-06 INFO Starting server on port 8080
==> my-app-def/app <==
2026-02-06 INFO Starting server on port 8080
==> my-app-abc/app <==
2026-02-06 INFO Request received GET /health
```

## UI Layout (v0.3.0)

### Sidebar
All connected views display a right sidebar panel (~25% width) containing:
- ASCII art K4S logo with diagonal stripe accents
- Cluster connection status with `◉`/`○` indicators
- Active cluster name, namespace, and kubeconfig path
- Navigation menu with `▸` active view indicator
- Auto-hides on narrow terminals (<90 columns)

### Crush-Inspired Theme
Color palette based on Crush design language:
- Primary purple (`#7D56F4`) for active elements and accents
- Accent purple (`#AD58F7`) for logo stripes
- Selection highlight: `▌` thick left bar with dark purple background (`#2a2550`)
- Text hierarchy: bright (`#E0E0E0`), muted (`#626262`), subtle (`#444444`), dim (`#333333`)

### Transparent Dialogs
All dialogs (confirm, scale, container select, pod multi-select, help) float over visible content instead of blanking the background.

### Footer
Crush-style format: `key action · key action` with thin `─` separator line above.

## SSH Hosts View (`9`)

Connect to K3s nodes via SSH.

**Features:**
- View containers on the node via crictl
- Inspect container logs
- See node system information

See [SSH Integration](ssh.md) for setup details.
