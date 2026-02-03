package tui

import (
	"context"
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/LywwKkA-aD/k4s/internal/adapter/k8s"
	"github.com/LywwKkA-aD/k4s/internal/adapter/ssh"
	"github.com/LywwKkA-aD/k4s/internal/domain"
	"github.com/LywwKkA-aD/k4s/internal/logger"
)

const podRefreshInterval = 5 * time.Second

// ViewState represents the current view
type ViewState int

const (
	ViewKubeConfigSelect ViewState = iota
	ViewConnecting
	ViewNamespaces
	ViewPods
	ViewPodDetails
	ViewLogs
	ViewMain
	ViewSSHHosts
	ViewSSHConnecting
	ViewCrictlContainers
	ViewCrictlLogs
	ViewNodeInfo
)

// Messages for async operations
type connectResultMsg struct {
	client      *k8s.Client
	clusterInfo *domain.ClusterInfo
	err         error
}

type namespacesResultMsg struct {
	namespaces []domain.Namespace
	err        error
}

type podsResultMsg struct {
	pods []domain.Pod
	err  error
}

type podDetailsResultMsg struct {
	pod    *domain.Pod
	events []domain.PodEvent
	err    error
}

type podRefreshTickMsg struct{}

type podDeleteResultMsg struct {
	podName string
	err     error
}

type podRestartResultMsg struct {
	podName string
	err     error
}

type logsResultMsg struct {
	logs string
	err  error
}

type logLineMsg struct {
	line string
}

type logStreamEndedMsg struct {
	err error
}

type containersResultMsg struct {
	containers []string
	err        error
}

// SSH-related messages
type sshConnectResultMsg struct {
	err error
}

type sshCrictlContainersMsg struct {
	containers []ssh.CrictlContainer
	err        error
}

type sshNodeInfoMsg struct {
	info *domain.NodeInfo
	err  error
}

type sshCrictlLogsMsg struct {
	logs string
	err  error
}

type sshCrictlLogLineMsg struct {
	line string
}

type sshCrictlLogStreamEndedMsg struct {
	err error
}

// App is the main TUI application model
type App struct {
	styles             Styles
	width              int
	height             int
	ready              bool
	config             *domain.Config
	selectedConfig     *domain.KubeConfig
	viewState          ViewState
	kubeConfigList     list.Model
	namespaceList      list.Model
	podList            list.Model
	podDetails         PodDetailsModel
	selectedPodName    string
	k8sClient          *k8s.Client
	clusterInfo        *domain.ClusterInfo
	connectionStatus   domain.ConnectionStatus
	spinner            spinner.Model
	err                error
	loading            bool
	podCount           int
	namespaceCount     int
	confirmDialog      ConfirmDialog
	notification       Notification
	logViewer          LogViewer
	containerSelector  ContainerSelector
	logStreamCancel    context.CancelFunc
	logStreamActive    bool
	logLineChan        <-chan string

	// SSH-related fields
	sshHostList             list.Model
	sshClient               *ssh.Client
	selectedSSHHost         *domain.SSHHost
	crictlContainers        []ssh.CrictlContainer
	crictlContainerList     list.Model
	nodeInfo                *domain.NodeInfo
	passphraseInput         PassphraseInput
	crictlLogViewer         CrictlLogViewer
	selectedCrictlContainer *ssh.CrictlContainer
	crictlLogStreamCancel   context.CancelFunc
	crictlLogStreamActive   bool
	crictlLogLineChan       <-chan string

	// Help screen
	helpScreen HelpScreen

	// Search input
	searchInput SearchInput
}

// NewApp creates a new App instance with configuration
func NewApp(cfg *domain.Config) *App {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(colorPrimary)

	app := &App{
		styles:            DefaultStyles(),
		config:            cfg,
		viewState:         ViewKubeConfigSelect,
		connectionStatus:  domain.StatusDisconnected,
		spinner:           s,
		podDetails:        NewPodDetailsModel(DefaultStyles()),
		confirmDialog:     NewConfirmDialog(),
		notification:      NewNotification(),
		logViewer:         NewLogViewer(DefaultStyles()),
		containerSelector: NewContainerSelector(),
		passphraseInput:   NewPassphraseInput(),
		crictlLogViewer:   NewCrictlLogViewer(DefaultStyles()),
		helpScreen:        NewHelpScreen(),
		searchInput:       NewSearchInput(),
	}

	// If only one kubeconfig, auto-select it
	if len(cfg.KubeConfigs) == 1 {
		app.selectedConfig = &cfg.KubeConfigs[0]
	}

	return app
}

// Init implements tea.Model
func (a *App) Init() tea.Cmd {
	cmds := []tea.Cmd{a.spinner.Tick}

	// Auto-connect if only one kubeconfig
	if a.selectedConfig != nil {
		a.viewState = ViewConnecting
		a.connectionStatus = domain.StatusConnecting
		cmds = append(cmds, a.connectToCluster(a.selectedConfig.Path))
	}

	return tea.Batch(cmds...)
}

// connectToCluster returns a command that connects to the cluster
func (a *App) connectToCluster(kubeconfigPath string) tea.Cmd {
	return func() tea.Msg {
		client, err := k8s.NewClient(kubeconfigPath)
		if err != nil {
			return connectResultMsg{err: err}
		}

		ctx := context.Background()
		if err := client.CheckConnection(ctx); err != nil {
			return connectResultMsg{err: err}
		}

		info, err := client.GetClusterInfo(ctx)
		if err != nil {
			return connectResultMsg{client: client, err: err}
		}

		return connectResultMsg{client: client, clusterInfo: info}
	}
}

// fetchNamespaces returns a command that fetches namespaces
func (a *App) fetchNamespaces() tea.Cmd {
	return func() tea.Msg {
		if a.k8sClient == nil {
			return namespacesResultMsg{err: fmt.Errorf("not connected to cluster")}
		}

		ctx := context.Background()
		namespaces, err := a.k8sClient.GetNamespaces(ctx)
		if err != nil {
			return namespacesResultMsg{err: err}
		}

		return namespacesResultMsg{namespaces: namespaces}
	}
}

// fetchPods returns a command that fetches pods
func (a *App) fetchPods() tea.Cmd {
	return func() tea.Msg {
		if a.k8sClient == nil {
			return podsResultMsg{err: fmt.Errorf("not connected to cluster")}
		}

		ctx := context.Background()
		pods, err := a.k8sClient.GetPods(ctx, a.k8sClient.CurrentNamespace())
		if err != nil {
			return podsResultMsg{err: err}
		}

		return podsResultMsg{pods: pods}
	}
}

// fetchPodDetails returns a command that fetches pod details and events
func (a *App) fetchPodDetails(podName string) tea.Cmd {
	return func() tea.Msg {
		if a.k8sClient == nil {
			return podDetailsResultMsg{err: fmt.Errorf("not connected to cluster")}
		}

		ctx := context.Background()
		namespace := a.k8sClient.CurrentNamespace()

		pod, err := a.k8sClient.GetPod(ctx, namespace, podName)
		if err != nil {
			return podDetailsResultMsg{err: err}
		}

		events, err := a.k8sClient.GetPodEvents(ctx, namespace, podName)
		if err != nil {
			// Non-fatal: continue without events
			events = nil
		}

		return podDetailsResultMsg{pod: pod, events: events}
	}
}

// schedulePodRefresh returns a command that triggers a pod refresh after interval
func (a *App) schedulePodRefresh() tea.Cmd {
	return tea.Tick(podRefreshInterval, func(t time.Time) tea.Msg {
		return podRefreshTickMsg{}
	})
}

// deletePod returns a command that deletes a pod
func (a *App) deletePod(podName string) tea.Cmd {
	return func() tea.Msg {
		if a.k8sClient == nil {
			return podDeleteResultMsg{podName: podName, err: fmt.Errorf("not connected to cluster")}
		}

		ctx := context.Background()
		err := a.k8sClient.DeletePod(ctx, a.k8sClient.CurrentNamespace(), podName)
		return podDeleteResultMsg{podName: podName, err: err}
	}
}

// restartPod returns a command that restarts a pod (by deleting it)
func (a *App) restartPod(podName string) tea.Cmd {
	return func() tea.Msg {
		if a.k8sClient == nil {
			return podRestartResultMsg{podName: podName, err: fmt.Errorf("not connected to cluster")}
		}

		ctx := context.Background()
		err := a.k8sClient.DeletePod(ctx, a.k8sClient.CurrentNamespace(), podName)
		return podRestartResultMsg{podName: podName, err: err}
	}
}

// fetchContainers returns a command that fetches container names for a pod
func (a *App) fetchContainers(podName string) tea.Cmd {
	logger.Debug("fetchContainers called for pod: %s", podName)
	return func() tea.Msg {
		if a.k8sClient == nil {
			logger.Error("fetchContainers: k8sClient is nil")
			return containersResultMsg{err: fmt.Errorf("not connected to cluster")}
		}

		ctx := context.Background()
		logger.Debug("Calling GetPodContainers for %s in namespace %s", podName, a.k8sClient.CurrentNamespace())
		containers, err := a.k8sClient.GetPodContainers(ctx, a.k8sClient.CurrentNamespace(), podName)
		if err != nil {
			logger.Errorf(err, "GetPodContainers failed")
		} else {
			logger.Debug("GetPodContainers returned %d containers", len(containers))
		}
		return containersResultMsg{containers: containers, err: err}
	}
}

// fetchLogs returns a command that fetches logs for a pod
func (a *App) fetchLogs(podName, container string, tailLines int64, timestamps bool) tea.Cmd {
	return func() tea.Msg {
		if a.k8sClient == nil {
			return logsResultMsg{err: fmt.Errorf("not connected to cluster")}
		}

		ctx := context.Background()
		opts := k8s.LogOptions{
			Container:  container,
			TailLines:  tailLines,
			Timestamps: timestamps,
			Follow:     false,
		}

		logs, err := a.k8sClient.GetPodLogs(ctx, a.k8sClient.CurrentNamespace(), podName, opts)
		return logsResultMsg{logs: logs, err: err}
	}
}

// stopLogStream stops the current log stream
func (a *App) stopLogStream() {
	if a.logStreamCancel != nil {
		a.logStreamCancel()
		a.logStreamCancel = nil
	}
	a.logStreamActive = false
}

// Update implements tea.Model
func (a *App) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return a.handleKeyMsg(msg)

	case tea.WindowSizeMsg:
		a.width = msg.Width
		a.height = msg.Height
		a.ready = true
		a.kubeConfigList = newKubeConfigList(
			a.config.KubeConfigs,
			a.width-4,
			a.height-10,
			a.styles,
		)
		a.namespaceList = newNamespaceList(
			nil,
			a.width-4,
			a.height-12,
			a.styles,
		)
		a.podList = newPodList(
			nil,
			a.width-4,
			a.height-12,
			a.styles,
		)
		a.podDetails.SetSize(a.width-4, a.height-12)
		a.logViewer.SetSize(a.width-4, a.height-14)
		a.confirmDialog.SetWidth(a.width)
		a.notification.SetWidth(a.width)
		a.containerSelector.SetWidth(a.width)
		// Initialize SSH host list
		a.sshHostList = newSSHHostList(
			a.config.SSHHosts,
			a.width-4,
			a.height-12,
			a.styles,
		)
		a.crictlContainerList = newCrictlContainerList(
			nil,
			a.width-4,
			a.height-12,
			a.styles,
		)
		a.passphraseInput.SetWidth(a.width)
		a.crictlLogViewer.SetSize(a.width-4, a.height-14)
		a.helpScreen.SetSize(a.width, a.height)
		return a, nil

	case connectResultMsg:
		return a.handleConnectResult(msg)

	case namespacesResultMsg:
		return a.handleNamespacesResult(msg)

	case podsResultMsg:
		return a.handlePodsResult(msg)

	case podDetailsResultMsg:
		return a.handlePodDetailsResult(msg)

	case podRefreshTickMsg:
		// Only refresh if we're on the pods view, dialog is not visible, and not filtering
		if a.viewState == ViewPods && a.k8sClient != nil && !a.confirmDialog.IsVisible() && !a.podList.SettingFilter() && a.podList.FilterState() == list.Unfiltered {
			return a, tea.Batch(a.fetchPods(), a.schedulePodRefresh())
		}
		// Reschedule even if we skip refresh (to check again later)
		if a.viewState == ViewPods && a.k8sClient != nil {
			return a, a.schedulePodRefresh()
		}
		return a, nil

	case podDeleteResultMsg:
		return a.handlePodDeleteResult(msg)

	case podRestartResultMsg:
		return a.handlePodRestartResult(msg)

	case notificationExpiredMsg:
		a.notification.Hide()
		return a, nil

	case containersResultMsg:
		return a.handleContainersResult(msg)

	case logsResultMsg:
		return a.handleLogsResult(msg)

	case logLineMsg:
		return a.handleLogLine(msg)

	case logStreamEndedMsg:
		return a.handleLogStreamEnded(msg)

	case sshConnectResultMsg:
		return a.handleSSHConnectResult(msg)

	case sshCrictlContainersMsg:
		return a.handleSSHCrictlContainers(msg)

	case sshNodeInfoMsg:
		return a.handleSSHNodeInfo(msg)

	case sshCrictlLogsMsg:
		return a.handleCrictlLogsResult(msg)

	case sshCrictlLogLineMsg:
		return a.handleCrictlLogLine(msg)

	case sshCrictlLogStreamEndedMsg:
		return a.handleCrictlLogStreamEnded(msg)

	case spinner.TickMsg:
		var cmd tea.Cmd
		a.spinner, cmd = a.spinner.Update(msg)
		return a, cmd
	}

	// Update child components based on view state
	switch a.viewState {
	case ViewKubeConfigSelect:
		var cmd tea.Cmd
		a.kubeConfigList, cmd = a.kubeConfigList.Update(msg)
		return a, cmd
	case ViewNamespaces:
		var cmd tea.Cmd
		a.namespaceList, cmd = a.namespaceList.Update(msg)
		return a, cmd
	case ViewPods:
		var cmd tea.Cmd
		a.podList, cmd = a.podList.Update(msg)
		return a, cmd
	case ViewPodDetails:
		var cmd tea.Cmd
		a.podDetails, cmd = a.podDetails.Update(msg)
		return a, cmd
	case ViewLogs:
		var cmd tea.Cmd
		a.logViewer, cmd = a.logViewer.Update(msg)
		return a, cmd
	case ViewSSHHosts:
		var cmd tea.Cmd
		a.sshHostList, cmd = a.sshHostList.Update(msg)
		return a, cmd
	case ViewCrictlContainers:
		var cmd tea.Cmd
		a.crictlContainerList, cmd = a.crictlContainerList.Update(msg)
		return a, cmd
	case ViewCrictlLogs:
		var cmd tea.Cmd
		a.crictlLogViewer, cmd = a.crictlLogViewer.Update(msg)
		return a, cmd
	}

	return a, nil
}

func (a *App) handleConnectResult(msg connectResultMsg) (tea.Model, tea.Cmd) {
	if msg.err != nil {
		a.err = msg.err
		a.connectionStatus = domain.StatusError
		a.viewState = ViewMain
		return a, nil
	}

	a.k8sClient = msg.client
	a.clusterInfo = msg.clusterInfo
	a.connectionStatus = domain.StatusConnected
	a.viewState = ViewNamespaces
	a.err = nil
	a.loading = true

	// Fetch namespaces after successful connection
	return a, a.fetchNamespaces()
}

func (a *App) handleNamespacesResult(msg namespacesResultMsg) (tea.Model, tea.Cmd) {
	a.loading = false

	if msg.err != nil {
		a.err = msg.err
		return a, nil
	}

	a.namespaceCount = len(msg.namespaces)
	updateNamespaceList(&a.namespaceList, msg.namespaces)
	a.err = nil
	return a, nil
}

func (a *App) handlePodsResult(msg podsResultMsg) (tea.Model, tea.Cmd) {
	a.loading = false

	if msg.err != nil {
		a.err = msg.err
		return a, nil
	}

	a.podCount = len(msg.pods)

	// Don't update list while filtering - it breaks the filter state
	if !a.podList.SettingFilter() && a.podList.FilterState() == list.Unfiltered {
		updatePodList(&a.podList, msg.pods)
	}
	a.err = nil
	return a, nil
}

func (a *App) handlePodDetailsResult(msg podDetailsResultMsg) (tea.Model, tea.Cmd) {
	a.loading = false

	if msg.err != nil {
		a.err = msg.err
		return a, nil
	}

	a.podDetails.SetPod(msg.pod, msg.events)
	a.err = nil
	return a, nil
}

func (a *App) handlePodDeleteResult(msg podDeleteResultMsg) (tea.Model, tea.Cmd) {
	if msg.err != nil {
		notifCmd := a.notification.Show(
			fmt.Sprintf("Failed to delete pod: %v", msg.err),
			NotificationError,
		)
		return a, notifCmd
	}

	notifCmd := a.notification.Show(
		fmt.Sprintf("Pod '%s' deleted successfully", msg.podName),
		NotificationSuccess,
	)

	// Go back to pods view and refresh
	a.viewState = ViewPods
	a.selectedPodName = ""
	return a, tea.Batch(notifCmd, a.fetchPods(), a.schedulePodRefresh())
}

func (a *App) handlePodRestartResult(msg podRestartResultMsg) (tea.Model, tea.Cmd) {
	if msg.err != nil {
		notifCmd := a.notification.Show(
			fmt.Sprintf("Failed to restart pod: %v", msg.err),
			NotificationError,
		)
		return a, notifCmd
	}

	notifCmd := a.notification.Show(
		fmt.Sprintf("Pod '%s' restarting...", msg.podName),
		NotificationSuccess,
	)

	// Go back to pods view and refresh
	a.viewState = ViewPods
	a.selectedPodName = ""
	return a, tea.Batch(notifCmd, a.fetchPods(), a.schedulePodRefresh())
}

func (a *App) handleContainersResult(msg containersResultMsg) (tea.Model, tea.Cmd) {
	a.loading = false

	if msg.err != nil {
		logger.Errorf(msg.err, "Failed to get containers for pod %s", a.selectedPodName)
		a.err = msg.err
		return a, nil
	}

	logger.Debug("Got %d containers for pod %s: %v", len(msg.containers), a.selectedPodName, msg.containers)

	// Set up log viewer with pod and containers
	a.logViewer.SetPod(a.selectedPodName, a.k8sClient.CurrentNamespace(), msg.containers)
	a.viewState = ViewLogs

	logger.Debug("Switching to ViewLogs, fetching logs for container: %s", a.logViewer.Container())

	// Fetch initial logs
	return a, a.fetchLogs(
		a.selectedPodName,
		a.logViewer.Container(),
		a.logViewer.TailLines(),
		a.logViewer.Timestamps(),
	)
}

func (a *App) handleLogsResult(msg logsResultMsg) (tea.Model, tea.Cmd) {
	a.loading = false

	if msg.err != nil {
		logger.Errorf(msg.err, "Failed to get logs")
		a.err = msg.err
		return a, nil
	}

	logLen := len(msg.logs)
	logger.Debug("Received logs: %d bytes", logLen)

	a.logViewer.SetLogs(msg.logs)
	a.err = nil

	// Don't auto-start streaming - user must press 'f' to follow
	return a, nil
}

func (a *App) handleLogLine(msg logLineMsg) (tea.Model, tea.Cmd) {
	if a.viewState != ViewLogs {
		return a, nil
	}

	a.logViewer.AppendLog(msg.line)

	// Continue reading from stream if active
	if a.logStreamActive && a.logLineChan != nil {
		return a, a.waitForLogLine(a.logLineChan)
	}

	return a, nil
}

func (a *App) handleLogStreamEnded(msg logStreamEndedMsg) (tea.Model, tea.Cmd) {
	a.logStreamActive = false

	// If context was cancelled (user stopped follow), just return
	if msg.err == context.Canceled {
		return a, nil
	}

	// If follow mode is still enabled and we're on logs view, auto-restart the stream
	if a.logViewer.IsFollowing() && a.viewState == ViewLogs {
		logger.Debug("Log stream ended, auto-restarting (follow still enabled)")
		return a, a.startLogStreaming()
	}

	// Show notification if there was an error
	if msg.err != nil {
		notifCmd := a.notification.Show(
			fmt.Sprintf("Log stream ended: %v", msg.err),
			NotificationInfo,
		)
		return a, notifCmd
	}

	return a, nil
}

// startLogStreaming starts the log streaming using a subscription pattern
func (a *App) startLogStreaming() tea.Cmd {
	if a.k8sClient == nil {
		return nil
	}

	// Stop any existing stream
	a.stopLogStream()

	// Create a cancellable context
	ctx, cancel := context.WithCancel(context.Background())
	a.logStreamCancel = cancel
	a.logStreamActive = true

	opts := k8s.LogOptions{
		Container:  a.logViewer.Container(),
		TailLines:  0, // Don't re-fetch tail when streaming
		Timestamps: a.logViewer.Timestamps(),
		Follow:     true,
	}

	lineChan := make(chan string, 100)
	a.logLineChan = lineChan

	// Start streaming in a goroutine
	go func() {
		defer close(lineChan)
		_ = a.k8sClient.StreamPodLogs(
			ctx,
			a.k8sClient.CurrentNamespace(),
			a.logViewer.PodName(),
			opts,
			lineChan,
		)
	}()

	// Return a command that reads from the channel
	return a.waitForLogLine(lineChan)
}

// waitForLogLine returns a command that waits for the next log line
func (a *App) waitForLogLine(lineChan <-chan string) tea.Cmd {
	return func() tea.Msg {
		line, ok := <-lineChan
		if !ok {
			return logStreamEndedMsg{}
		}
		return logLineMsg{line: line}
	}
}

// SSH-related commands and handlers

// connectToSSHHost connects to an SSH host and returns a command
func (a *App) connectToSSHHost(host domain.SSHHost) tea.Cmd {
	// Close existing connection first
	a.closeSSHConnection()

	// Create and store client
	a.sshClient = ssh.NewClient(host)
	a.selectedSSHHost = &host

	return a.retrySSHConnection()
}

// retrySSHConnection retries SSH connection using the existing client (useful after setting passphrase)
func (a *App) retrySSHConnection() tea.Cmd {
	return func() tea.Msg {
		if a.sshClient == nil {
			return sshConnectResultMsg{err: fmt.Errorf("no SSH client")}
		}

		ctx := context.Background()

		if err := a.sshClient.Connect(ctx); err != nil {
			return sshConnectResultMsg{err: err}
		}

		if err := a.sshClient.TestConnection(ctx); err != nil {
			a.sshClient.Close()
			return sshConnectResultMsg{err: err}
		}

		return sshConnectResultMsg{err: nil}
	}
}

// fetchCrictlContainers returns a command that fetches crictl containers
func (a *App) fetchCrictlContainers() tea.Cmd {
	return func() tea.Msg {
		if a.sshClient == nil {
			return sshCrictlContainersMsg{err: fmt.Errorf("not connected to SSH host")}
		}

		ctx := context.Background()
		containers, err := a.sshClient.ListContainers(ctx)
		return sshCrictlContainersMsg{containers: containers, err: err}
	}
}

// fetchNodeInfo returns a command that fetches node information
func (a *App) fetchNodeInfo() tea.Cmd {
	return func() tea.Msg {
		if a.sshClient == nil {
			return sshNodeInfoMsg{err: fmt.Errorf("not connected to SSH host")}
		}

		ctx := context.Background()
		info, err := a.sshClient.GetNodeInfo(ctx)
		return sshNodeInfoMsg{info: info, err: err}
	}
}

func (a *App) handleSSHConnectResult(msg sshConnectResultMsg) (tea.Model, tea.Cmd) {
	a.loading = false

	if msg.err != nil {
		// Check if passphrase is required
		if msg.err == ssh.ErrPassphraseRequired {
			// Show passphrase input dialog
			hostName := ""
			if a.selectedSSHHost != nil {
				hostName = a.selectedSSHHost.Name
			}
			a.passphraseInput.Show(hostName)
			a.viewState = ViewSSHConnecting
			return a, nil
		}
		a.err = msg.err
		a.viewState = ViewSSHHosts
		return a, nil
	}

	// Connection successful - switch to crictl containers view
	a.viewState = ViewCrictlContainers
	a.err = nil
	a.loading = true

	// Fetch containers and node info
	return a, tea.Batch(a.fetchCrictlContainers(), a.fetchNodeInfo())
}

func (a *App) handleSSHCrictlContainers(msg sshCrictlContainersMsg) (tea.Model, tea.Cmd) {
	a.loading = false

	if msg.err != nil {
		a.err = msg.err
		return a, nil
	}

	a.crictlContainers = msg.containers
	updateCrictlContainerList(&a.crictlContainerList, msg.containers)
	a.err = nil
	return a, nil
}

func (a *App) handleSSHNodeInfo(msg sshNodeInfoMsg) (tea.Model, tea.Cmd) {
	if msg.err != nil {
		// Non-fatal, just log
		logger.Debug("Failed to get node info: %v", msg.err)
		return a, nil
	}

	a.nodeInfo = msg.info
	return a, nil
}

// closeSSHConnection closes the current SSH connection
func (a *App) closeSSHConnection() {
	if a.sshClient != nil {
		a.sshClient.Close()
		a.sshClient = nil
	}
	a.selectedSSHHost = nil
	a.nodeInfo = nil
	a.crictlContainers = nil
}

// fetchCrictlLogs returns a command that fetches crictl container logs
func (a *App) fetchCrictlLogs() tea.Cmd {
	return func() tea.Msg {
		if a.sshClient == nil || a.selectedCrictlContainer == nil {
			return sshCrictlLogsMsg{err: fmt.Errorf("not connected or no container selected")}
		}

		ctx := context.Background()
		opts := ssh.CrictlLogOptions{
			TailLines:  a.crictlLogViewer.TailLines(),
			Timestamps: a.crictlLogViewer.Timestamps(),
		}

		logs, err := a.sshClient.ContainerLogs(ctx, a.selectedCrictlContainer.ContainerID, opts)
		return sshCrictlLogsMsg{logs: logs, err: err}
	}
}

// stopCrictlLogStream stops the current crictl log stream
func (a *App) stopCrictlLogStream() {
	if a.crictlLogStreamCancel != nil {
		a.crictlLogStreamCancel()
		a.crictlLogStreamCancel = nil
	}
	a.crictlLogStreamActive = false
}

// startCrictlLogStreaming starts the crictl log streaming
func (a *App) startCrictlLogStreaming() tea.Cmd {
	if a.sshClient == nil || a.selectedCrictlContainer == nil {
		return nil
	}

	// Stop any existing stream
	a.stopCrictlLogStream()

	// Create a cancellable context
	ctx, cancel := context.WithCancel(context.Background())
	a.crictlLogStreamCancel = cancel
	a.crictlLogStreamActive = true

	opts := ssh.CrictlLogOptions{
		TailLines:  0, // Don't re-fetch tail when streaming
		Timestamps: a.crictlLogViewer.Timestamps(),
		Follow:     true,
	}

	lineChan := make(chan string, 100)
	a.crictlLogLineChan = lineChan

	containerID := a.selectedCrictlContainer.ContainerID

	// Start streaming in a goroutine
	go func() {
		defer close(lineChan)
		_ = a.sshClient.StreamContainerLogs(ctx, containerID, opts, lineChan)
	}()

	// Return a command that reads from the channel
	return a.waitForCrictlLogLine(lineChan)
}

// waitForCrictlLogLine returns a command that waits for the next crictl log line
func (a *App) waitForCrictlLogLine(lineChan <-chan string) tea.Cmd {
	return func() tea.Msg {
		line, ok := <-lineChan
		if !ok {
			return sshCrictlLogStreamEndedMsg{}
		}
		return sshCrictlLogLineMsg{line: line}
	}
}

func (a *App) handleCrictlLogsResult(msg sshCrictlLogsMsg) (tea.Model, tea.Cmd) {
	a.loading = false

	if msg.err != nil {
		logger.Errorf(msg.err, "Failed to get crictl logs")
		a.err = msg.err
		return a, nil
	}

	a.crictlLogViewer.SetLogs(msg.logs)
	a.err = nil
	return a, nil
}

func (a *App) handleCrictlLogLine(msg sshCrictlLogLineMsg) (tea.Model, tea.Cmd) {
	if a.viewState != ViewCrictlLogs {
		return a, nil
	}

	a.crictlLogViewer.AppendLog(msg.line)

	// Continue reading from stream if active
	if a.crictlLogStreamActive && a.crictlLogLineChan != nil {
		return a, a.waitForCrictlLogLine(a.crictlLogLineChan)
	}

	return a, nil
}

func (a *App) handleCrictlLogStreamEnded(msg sshCrictlLogStreamEndedMsg) (tea.Model, tea.Cmd) {
	a.crictlLogStreamActive = false

	// If context was cancelled (user stopped follow), just return
	if msg.err == context.Canceled {
		return a, nil
	}

	// If follow mode is still enabled and we're on logs view, auto-restart the stream
	if a.crictlLogViewer.IsFollowing() && a.viewState == ViewCrictlLogs {
		logger.Debug("Crictl log stream ended, auto-restarting (follow still enabled)")
		return a, a.startCrictlLogStreaming()
	}

	// Show notification if there was an error
	if msg.err != nil {
		notifCmd := a.notification.Show(
			fmt.Sprintf("Log stream ended: %v", msg.err),
			NotificationInfo,
		)
		return a, notifCmd
	}

	return a, nil
}

func (a *App) handleKeyMsg(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	key := msg.String()
	logger.Debug("Key pressed: '%s', viewState: %d", key, a.viewState)

	// Don't handle keys during connection
	if a.viewState == ViewConnecting {
		if key == "ctrl+c" {
			return a, tea.Quit
		}
		return a, nil
	}

	// Handle confirmation dialog if visible
	if a.confirmDialog.IsVisible() {
		confirmed, cancelled := a.confirmDialog.Update(msg)
		if confirmed {
			action := a.confirmDialog.Action()
			targetName := a.confirmDialog.TargetName()
			a.confirmDialog.Hide()

			switch action {
			case ConfirmActionDeletePod:
				return a, a.deletePod(targetName)
			case ConfirmActionRestartPod:
				return a, a.restartPod(targetName)
			}
		}
		if cancelled {
			a.confirmDialog.Hide()
		}
		return a, nil
	}

	// Handle help screen if visible
	if a.helpScreen.IsVisible() {
		if key == "?" || key == "esc" {
			a.helpScreen.Hide()
		}
		return a, nil
	}

	// Handle container selector if visible
	if a.containerSelector.IsVisible() {
		selected, cancelled := a.containerSelector.Update(msg)
		if selected {
			container := a.containerSelector.SelectedContainer()
			a.containerSelector.Hide()
			a.logViewer.SetContainer(container)
			a.stopLogStream()
			a.loading = true
			return a, a.fetchLogs(
				a.logViewer.PodName(),
				container,
				a.logViewer.TailLines(),
				a.logViewer.Timestamps(),
			)
		}
		if cancelled {
			a.containerSelector.Hide()
		}
		return a, nil
	}

	// Handle passphrase input if visible
	if a.passphraseInput.IsVisible() {
		passphrase, submitted, cancelled, cmd := a.passphraseInput.Update(msg)
		if submitted {
			a.passphraseInput.Hide()
			if a.sshClient != nil && a.selectedSSHHost != nil {
				// Set passphrase on existing client and retry connection
				a.sshClient.SetPassphrase(passphrase)
				a.loading = true
				return a, a.retrySSHConnection()
			}
		}
		if cancelled {
			a.passphraseInput.Hide()
			a.viewState = ViewSSHHosts
		}
		return a, cmd
	}

	// Handle search input if visible (in log views)
	if a.searchInput.IsVisible() {
		query, submitted, cancelled, cmd := a.searchInput.Update(msg)
		if submitted || cancelled {
			a.searchInput.Hide()
			if cancelled {
				// Clear search
				if a.viewState == ViewLogs {
					a.logViewer.SetSearchQuery("")
				} else if a.viewState == ViewCrictlLogs {
					a.crictlLogViewer.SetSearchQuery("")
				}
			}
		} else {
			// Update search as user types
			if a.viewState == ViewLogs {
				a.logViewer.SetSearchQuery(query)
				a.searchInput.SetMatchCount(a.logViewer.MatchCount())
			} else if a.viewState == ViewCrictlLogs {
				a.crictlLogViewer.SetSearchQuery(query)
				a.searchInput.SetMatchCount(a.crictlLogViewer.MatchCount())
			}
		}
		return a, cmd
	}

	// Handle filter mode for lists
	if a.viewState == ViewKubeConfigSelect && a.kubeConfigList.SettingFilter() {
		var cmd tea.Cmd
		a.kubeConfigList, cmd = a.kubeConfigList.Update(msg)
		return a, cmd
	}
	if a.viewState == ViewNamespaces && a.namespaceList.SettingFilter() {
		var cmd tea.Cmd
		a.namespaceList, cmd = a.namespaceList.Update(msg)
		return a, cmd
	}
	if a.viewState == ViewPods && a.podList.SettingFilter() {
		var cmd tea.Cmd
		a.podList, cmd = a.podList.Update(msg)
		return a, cmd
	}
	if a.viewState == ViewSSHHosts && a.sshHostList.SettingFilter() {
		var cmd tea.Cmd
		a.sshHostList, cmd = a.sshHostList.Update(msg)
		return a, cmd
	}
	if a.viewState == ViewCrictlContainers && a.crictlContainerList.SettingFilter() {
		var cmd tea.Cmd
		a.crictlContainerList, cmd = a.crictlContainerList.Update(msg)
		return a, cmd
	}

	switch msg.String() {
	case "ctrl+c":
		return a, tea.Quit

	case "?":
		// Show help screen (except during connection)
		if a.viewState != ViewConnecting && a.viewState != ViewSSHConnecting {
			a.helpScreen.Toggle()
			return a, nil
		}

	case "q":
		switch a.viewState {
		case ViewMain, ViewNamespaces, ViewPods, ViewPodDetails, ViewLogs, ViewSSHHosts, ViewCrictlContainers, ViewCrictlLogs, ViewNodeInfo:
			a.stopLogStream()        // Clean up any active log stream
			a.stopCrictlLogStream()  // Clean up crictl log stream
			a.closeSSHConnection()   // Clean up SSH connection
			return a, tea.Quit
		case ViewKubeConfigSelect:
			return a, tea.Quit
		}

	case "enter":
		switch a.viewState {
		case ViewKubeConfigSelect:
			if item, ok := a.kubeConfigList.SelectedItem().(kubeConfigItem); ok {
				a.selectedConfig = &item.kubeConfig
				a.viewState = ViewConnecting
				a.connectionStatus = domain.StatusConnecting
				a.err = nil
				return a, a.connectToCluster(item.kubeConfig.Path)
			}
		case ViewNamespaces:
			if item, ok := a.namespaceList.SelectedItem().(namespaceItem); ok {
				a.k8sClient.SetNamespace(item.namespace.Name)
				a.clusterInfo.Namespace = item.namespace.Name
				a.viewState = ViewPods
				a.loading = true
				// Fetch pods and start auto-refresh
				return a, tea.Batch(a.fetchPods(), a.schedulePodRefresh())
			}
		case ViewPods:
			if item, ok := a.podList.SelectedItem().(podItem); ok {
				a.selectedPodName = item.pod.Name
				a.viewState = ViewPodDetails
				a.loading = true
				return a, a.fetchPodDetails(item.pod.Name)
			}
		case ViewSSHHosts:
			if item, ok := a.sshHostList.SelectedItem().(sshHostItem); ok {
				a.viewState = ViewSSHConnecting
				a.loading = true
				a.err = nil
				return a, a.connectToSSHHost(item.host)
			}
		case ViewCrictlContainers:
			if item, ok := a.crictlContainerList.SelectedItem().(crictlContainerItem); ok {
				container := item.container
				a.selectedCrictlContainer = &container
				nodeName := ""
				if a.selectedSSHHost != nil {
					nodeName = a.selectedSSHHost.Name
				}
				a.crictlLogViewer.SetContainer(container.ContainerID, container.Name, nodeName)
				a.viewState = ViewCrictlLogs
				a.loading = true
				return a, a.fetchCrictlLogs()
			}
		}

	case "r":
		// Refresh
		switch a.viewState {
		case ViewNamespaces:
			if a.k8sClient != nil {
				a.loading = true
				return a, a.fetchNamespaces()
			}
		case ViewPods:
			if a.k8sClient != nil {
				a.loading = true
				return a, a.fetchPods()
			}
		case ViewPodDetails:
			if a.k8sClient != nil && a.selectedPodName != "" {
				a.loading = true
				return a, a.fetchPodDetails(a.selectedPodName)
			}
		case ViewLogs:
			if a.k8sClient != nil {
				a.stopLogStream()
				a.loading = true
				return a, a.fetchLogs(
					a.logViewer.PodName(),
					a.logViewer.Container(),
					a.logViewer.TailLines(),
					a.logViewer.Timestamps(),
				)
			}
		case ViewMain:
			if a.selectedConfig != nil && a.connectionStatus != domain.StatusConnected {
				a.viewState = ViewConnecting
				a.connectionStatus = domain.StatusConnecting
				a.err = nil
				return a, a.connectToCluster(a.selectedConfig.Path)
			}
		case ViewCrictlContainers:
			if a.sshClient != nil {
				a.loading = true
				return a, a.fetchCrictlContainers()
			}
		case ViewCrictlLogs:
			if a.sshClient != nil && a.selectedCrictlContainer != nil {
				a.stopCrictlLogStream()
				a.loading = true
				return a, a.fetchCrictlLogs()
			}
		}

	case "l":
		// Go to logs
		if a.viewState == ViewPodDetails && a.selectedPodName != "" {
			// From pod details - pod already selected
			logger.Debug("Opening logs from pod details for: %s", a.selectedPodName)
			a.loading = true
			return a, a.fetchContainers(a.selectedPodName)
		}
		if a.viewState == ViewPods {
			// From pods list - get selected pod
			if item, ok := a.podList.SelectedItem().(podItem); ok {
				a.selectedPodName = item.pod.Name
				logger.Debug("Opening logs from pods list for: %s", a.selectedPodName)
				a.loading = true
				return a, a.fetchContainers(item.pod.Name)
			}
			logger.Warn("No pod selected in pods list")
		}

	case "d":
		// Delete pod
		if a.viewState == ViewPodDetails && a.selectedPodName != "" {
			a.confirmDialog.Show(ConfirmActionDeletePod, a.selectedPodName)
			return a, nil
		}
		// Also allow deletion from pods list view
		if a.viewState == ViewPods {
			if item, ok := a.podList.SelectedItem().(podItem); ok {
				a.confirmDialog.Show(ConfirmActionDeletePod, item.pod.Name)
				return a, nil
			}
		}

	case "R":
		// Restart pod (Shift+R)
		if a.viewState == ViewPodDetails && a.selectedPodName != "" {
			a.confirmDialog.Show(ConfirmActionRestartPod, a.selectedPodName)
			return a, nil
		}
		// Also allow restart from pods list view
		if a.viewState == ViewPods {
			if item, ok := a.podList.SelectedItem().(podItem); ok {
				a.confirmDialog.Show(ConfirmActionRestartPod, item.pod.Name)
				return a, nil
			}
		}

	case "f":
		// Toggle follow mode in log viewer
		if a.viewState == ViewLogs {
			following := a.logViewer.ToggleFollowing()
			if following {
				// Start streaming
				return a, a.startLogStreaming()
			} else {
				// Stop streaming
				a.stopLogStream()
			}
			return a, nil
		}
		// Toggle follow mode in crictl log viewer
		if a.viewState == ViewCrictlLogs {
			following := a.crictlLogViewer.ToggleFollowing()
			if following {
				return a, a.startCrictlLogStreaming()
			} else {
				a.stopCrictlLogStream()
			}
			return a, nil
		}

	case "t":
		// Toggle timestamps in log viewer
		if a.viewState == ViewLogs {
			a.logViewer.ToggleTimestamps()
			a.stopLogStream()
			a.loading = true
			return a, a.fetchLogs(
				a.logViewer.PodName(),
				a.logViewer.Container(),
				a.logViewer.TailLines(),
				a.logViewer.Timestamps(),
			)
		}
		// Toggle timestamps in crictl log viewer
		if a.viewState == ViewCrictlLogs && a.selectedCrictlContainer != nil {
			a.crictlLogViewer.ToggleTimestamps()
			a.stopCrictlLogStream()
			a.loading = true
			return a, a.fetchCrictlLogs()
		}

	case "c":
		// Change container in log viewer
		if a.viewState == ViewLogs && len(a.logViewer.Containers()) > 1 {
			a.containerSelector.Show(a.logViewer.Containers(), a.logViewer.Container())
			return a, nil
		}

	case "/":
		// Start search in log views
		if a.viewState == ViewLogs || a.viewState == ViewCrictlLogs {
			a.searchInput.Show()
			return a, nil
		}

	case "n":
		// Next search match
		if a.viewState == ViewLogs && a.logViewer.SearchQuery() != "" {
			a.searchInput.NextMatch()
			return a, nil
		}
		if a.viewState == ViewCrictlLogs && a.crictlLogViewer.searchQuery != "" {
			a.searchInput.NextMatch()
			return a, nil
		}

	case "N":
		// Previous search match
		if a.viewState == ViewLogs && a.logViewer.SearchQuery() != "" {
			a.searchInput.PrevMatch()
			return a, nil
		}
		if a.viewState == ViewCrictlLogs && a.crictlLogViewer.searchQuery != "" {
			a.searchInput.PrevMatch()
			return a, nil
		}

	case "0":
		// Go to namespaces view
		if (a.viewState == ViewMain || a.viewState == ViewPods || a.viewState == ViewPodDetails) && a.connectionStatus == domain.StatusConnected {
			a.viewState = ViewNamespaces
			a.loading = true
			return a, a.fetchNamespaces()
		}

	case "1":
		// Go to pods view
		if (a.viewState == ViewMain || a.viewState == ViewNamespaces || a.viewState == ViewPodDetails) && a.connectionStatus == domain.StatusConnected {
			a.viewState = ViewPods
			a.loading = true
			return a, tea.Batch(a.fetchPods(), a.schedulePodRefresh())
		}

	case "9":
		// Go to SSH hosts view
		if len(a.config.SSHHosts) > 0 {
			a.viewState = ViewSSHHosts
			a.err = nil
			return a, nil
		}

	case "esc":
		switch a.viewState {
		case ViewMain:
			// Go back to pods
			if a.connectionStatus == domain.StatusConnected {
				a.viewState = ViewPods
				return a, tea.Batch(a.fetchPods(), a.schedulePodRefresh())
			}
			// Or go back to kubeconfig selection
			if len(a.config.KubeConfigs) > 1 {
				a.viewState = ViewKubeConfigSelect
				a.k8sClient = nil
				a.clusterInfo = nil
				a.connectionStatus = domain.StatusDisconnected
				return a, nil
			}
		case ViewLogs:
			// Stop streaming and go back to pod details
			a.stopLogStream()
			a.logViewer.Clear()
			a.viewState = ViewPodDetails
			a.loading = true
			return a, a.fetchPodDetails(a.selectedPodName)
		case ViewPodDetails:
			// Go back to pods
			a.viewState = ViewPods
			a.selectedPodName = ""
			return a, tea.Batch(a.fetchPods(), a.schedulePodRefresh())
		case ViewPods:
			// Go back to namespaces
			a.viewState = ViewNamespaces
			return a, a.fetchNamespaces()
		case ViewNamespaces:
			// Go back to kubeconfig selection if multiple configs
			if len(a.config.KubeConfigs) > 1 {
				a.viewState = ViewKubeConfigSelect
				a.k8sClient = nil
				a.clusterInfo = nil
				a.connectionStatus = domain.StatusDisconnected
				return a, nil
			}
		case ViewSSHHosts:
			// Go back to pods view (or namespaces if not connected)
			if a.connectionStatus == domain.StatusConnected {
				a.viewState = ViewPods
				return a, tea.Batch(a.fetchPods(), a.schedulePodRefresh())
			}
			a.viewState = ViewNamespaces
			return a, a.fetchNamespaces()
		case ViewCrictlContainers, ViewNodeInfo:
			// Go back to SSH hosts
			a.closeSSHConnection()
			a.viewState = ViewSSHHosts
			return a, nil
		case ViewCrictlLogs:
			// Stop streaming and go back to containers
			a.stopCrictlLogStream()
			a.crictlLogViewer.Clear()
			a.selectedCrictlContainer = nil
			a.viewState = ViewCrictlContainers
			return a, nil
		}
	}

	// Pass key to child components
	switch a.viewState {
	case ViewKubeConfigSelect:
		var cmd tea.Cmd
		a.kubeConfigList, cmd = a.kubeConfigList.Update(msg)
		return a, cmd
	case ViewNamespaces:
		var cmd tea.Cmd
		a.namespaceList, cmd = a.namespaceList.Update(msg)
		return a, cmd
	case ViewPods:
		var cmd tea.Cmd
		a.podList, cmd = a.podList.Update(msg)
		return a, cmd
	case ViewPodDetails:
		var cmd tea.Cmd
		a.podDetails, cmd = a.podDetails.Update(msg)
		return a, cmd
	case ViewLogs:
		var cmd tea.Cmd
		a.logViewer, cmd = a.logViewer.Update(msg)
		return a, cmd
	case ViewSSHHosts:
		var cmd tea.Cmd
		a.sshHostList, cmd = a.sshHostList.Update(msg)
		return a, cmd
	case ViewCrictlContainers:
		var cmd tea.Cmd
		a.crictlContainerList, cmd = a.crictlContainerList.Update(msg)
		return a, cmd
	case ViewCrictlLogs:
		var cmd tea.Cmd
		a.crictlLogViewer, cmd = a.crictlLogViewer.Update(msg)
		return a, cmd
	}

	return a, nil
}

// View implements tea.Model
func (a *App) View() string {
	if !a.ready {
		return "Loading..."
	}

	var view string
	switch a.viewState {
	case ViewKubeConfigSelect:
		view = a.renderKubeConfigSelect()
	case ViewConnecting:
		view = a.renderConnecting()
	case ViewNamespaces:
		view = a.renderNamespacesView()
	case ViewPods:
		view = a.renderPodsView()
	case ViewPodDetails:
		view = a.renderPodDetailsView()
	case ViewLogs:
		view = a.renderLogsView()
	case ViewMain:
		view = a.renderMainView()
	case ViewSSHHosts:
		view = a.renderSSHHostsView()
	case ViewSSHConnecting:
		view = a.renderSSHConnecting()
	case ViewCrictlContainers:
		view = a.renderCrictlContainersView()
	case ViewCrictlLogs:
		view = a.renderCrictlLogsView()
	default:
		view = ""
	}

	// Overlay help screen if visible
	if a.helpScreen.IsVisible() {
		return a.overlayHelpScreen(view)
	}

	return view
}

// overlayHelpScreen overlays the help screen centered on the view
func (a *App) overlayHelpScreen(view string) string {
	if !a.helpScreen.IsVisible() {
		return view
	}

	helpView := a.helpScreen.View()
	return lipgloss.Place(
		a.width,
		a.height,
		lipgloss.Center,
		lipgloss.Center,
		helpView,
		lipgloss.WithWhitespaceChars(" "),
		lipgloss.WithWhitespaceForeground(lipgloss.Color("#1a1a1a")),
	)
}

func (a *App) renderKubeConfigSelect() string {
	header := a.renderHeader()
	content := a.styles.Content.
		Width(a.width - 4).
		Height(a.height - 10).
		Render(a.kubeConfigList.View())
	footer := a.renderFooter()

	return lipgloss.JoinVertical(lipgloss.Left, header, content, footer)
}

func (a *App) renderConnecting() string {
	header := a.renderHeader()

	spinnerView := a.spinner.View()
	connectingText := fmt.Sprintf("%s Connecting to cluster...", spinnerView)

	if a.selectedConfig != nil {
		connectingText = fmt.Sprintf("%s Connecting to %s...", spinnerView, a.selectedConfig.Name)
	}

	content := a.styles.Content.
		Width(a.width - 4).
		Height(a.height - 10).
		Render(connectingText)

	footer := a.renderFooter()

	return lipgloss.JoinVertical(lipgloss.Left, header, content, footer)
}

func (a *App) renderNamespacesView() string {
	header := a.renderHeader()

	var contentStr string
	if a.loading && a.namespaceCount == 0 {
		contentStr = fmt.Sprintf("%s Loading namespaces...", a.spinner.View())
	} else if a.err != nil {
		contentStr = a.renderError()
	} else {
		// Title with namespace count
		title := fmt.Sprintf("Namespaces (%d)", a.namespaceCount)
		titleLine := lipgloss.NewStyle().
			Bold(true).
			Foreground(colorPrimary).
			Render(title)

		// Column headers (widths must match delegate: 40 + 12 + age)
		headerLine := lipgloss.NewStyle().
			Bold(true).
			Foreground(colorMuted).
			Render(fmt.Sprintf("  %-45s %-12s %s", "NAME", "STATUS", "AGE"))

		contentStr = titleLine + "\n" + headerLine + "\n" + a.namespaceList.View()
	}

	content := a.styles.Content.
		Width(a.width - 4).
		Height(a.height - 12).
		Render(contentStr)

	footer := a.renderFooter()

	return lipgloss.JoinVertical(lipgloss.Left, header, content, footer)
}

func (a *App) renderPodsView() string {
	header := a.renderHeader()

	var contentStr string
	if a.loading && a.podCount == 0 {
		contentStr = fmt.Sprintf("%s Loading pods...", a.spinner.View())
	} else if a.err != nil {
		contentStr = a.renderError()
	} else {
		// Title with pod count
		title := fmt.Sprintf("Pods (%d)", a.podCount)
		titleLine := lipgloss.NewStyle().
			Bold(true).
			Foreground(colorPrimary).
			Render(title)

		// Column headers (widths must match delegate: 45 + 7 + 12 + 8 + age)
		headerLine := lipgloss.NewStyle().
			Bold(true).
			Foreground(colorMuted).
			Render(fmt.Sprintf("  %-45s %-7s %-12s %-8s %s", "NAME", "READY", "STATUS", "RESTARTS", "AGE"))

		contentStr = titleLine + "\n" + headerLine + "\n" + a.podList.View()
	}

	content := a.styles.Content.
		Width(a.width - 4).
		Height(a.height - 10).
		Render(contentStr)

	footer := a.renderFooter()

	view := lipgloss.JoinVertical(lipgloss.Left, header, content, footer)

	// Overlay confirmation dialog if visible
	if a.confirmDialog.IsVisible() {
		view = a.overlayDialog(view)
	}

	return view
}

func (a *App) renderPodDetailsView() string {
	header := a.renderHeader()

	var contentStr string
	if a.loading {
		contentStr = fmt.Sprintf("%s Loading pod details...", a.spinner.View())
	} else if a.err != nil {
		contentStr = a.renderError()
	} else {
		// Title
		title := fmt.Sprintf("Pod: %s", a.selectedPodName)
		titleLine := lipgloss.NewStyle().
			Bold(true).
			Foreground(colorPrimary).
			Render(title)

		scrollInfo := lipgloss.NewStyle().
			Foreground(colorMuted).
			Render(fmt.Sprintf(" (%.0f%%)", a.podDetails.ScrollPercent()*100))

		contentStr = titleLine + scrollInfo + "\n" + a.podDetails.View()
	}

	content := a.styles.Content.
		Width(a.width - 4).
		Height(a.height - 10).
		Render(contentStr)

	footer := a.renderFooter()

	view := lipgloss.JoinVertical(lipgloss.Left, header, content, footer)

	// Overlay confirmation dialog if visible
	if a.confirmDialog.IsVisible() {
		view = a.overlayDialog(view)
	}

	return view
}

func (a *App) renderLogsView() string {
	header := a.renderHeader()

	var contentStr string
	if a.loading {
		contentStr = fmt.Sprintf("%s Loading logs...", a.spinner.View())
	} else if a.err != nil {
		contentStr = a.renderError()
	} else {
		// Log viewer header
		logHeader := a.logViewer.RenderHeader()
		// Add search input if visible
		if a.searchInput.IsVisible() {
			logHeader += "\n" + a.searchInput.View()
		}
		contentStr = logHeader + "\n" + a.logViewer.View()
	}

	content := a.styles.Content.
		Width(a.width - 4).
		Height(a.height - 10).
		Render(contentStr)

	footer := a.renderFooter()

	view := lipgloss.JoinVertical(lipgloss.Left, header, content, footer)

	// Overlay container selector if visible
	if a.containerSelector.IsVisible() {
		view = a.overlayContainerSelector(view)
	}

	return view
}

func (a *App) renderMainView() string {
	header := a.renderHeader()
	content := a.renderContent()
	footer := a.renderFooter()

	return lipgloss.JoinVertical(lipgloss.Left, header, content, footer)
}

func (a *App) renderHeader() string {
	title := a.styles.Title.Render("k4s")
	subtitle := a.styles.Subtitle.Render("Kubernetes TUI for K3s")

	headerContent := fmt.Sprintf("%s  %s", title, subtitle)

	if a.selectedConfig != nil {
		configBadge := a.styles.StatusBar.Render(a.selectedConfig.Name)
		headerContent = fmt.Sprintf("%s  %s  %s", title, subtitle, configBadge)
	}

	return a.styles.Header.
		Width(a.width - 4).
		Render(headerContent)
}

func (a *App) renderContent() string {
	if a.err != nil {
		return a.renderError()
	}

	if a.selectedConfig == nil {
		return a.styles.Content.Render("No kubeconfig selected")
	}

	var content string
	if a.clusterInfo != nil && a.connectionStatus == domain.StatusConnected {
		content = fmt.Sprintf(`
  Cluster:    %s
  Context:    %s
  Namespace:  %s
  Status:     Connected

Navigation:
  0 - Namespaces view
  1 - Pods view

Upcoming features:
  • Pod operations (Step 7)
  • Real-time streaming logs (Step 8)
`, a.clusterInfo.Name, a.clusterInfo.Context, a.clusterInfo.Namespace)
	} else {
		content = fmt.Sprintf(`
  Kubeconfig: %s
  Path:       %s
  Status:     %s

Press 'r' to retry connection
`, a.selectedConfig.Name, a.selectedConfig.Path, a.connectionStatus)
	}

	return a.styles.Content.
		Width(a.width - 4).
		Height(a.height - 10).
		Render(content)
}

func (a *App) renderFooter() string {
	var helpText string
	switch a.viewState {
	case ViewKubeConfigSelect:
		parts := []string{"↑/↓: navigate", "enter: select", "/: filter"}
		if len(a.config.SSHHosts) > 0 {
			parts = append(parts, "9: ssh")
		}
		parts = append(parts, "q: quit")
		helpText = joinStrings(parts, " • ")
	case ViewConnecting:
		helpText = "ctrl+c: cancel"
	case ViewNamespaces:
		parts := []string{"↑/↓: navigate", "enter: select", "/: filter", "1: pods", "r: refresh"}
		if len(a.config.SSHHosts) > 0 {
			parts = append(parts, "9: ssh")
		}
		if len(a.config.KubeConfigs) > 1 {
			parts = append(parts, "esc: back")
		}
		parts = append(parts, "q: quit")
		helpText = joinStrings(parts, " • ")
	case ViewPods:
		parts := []string{"↑/↓: navigate", "enter: details", "l: logs", "d: delete", "R: restart", "/: filter", "0: ns", "r: refresh"}
		if len(a.config.SSHHosts) > 0 {
			parts = append(parts, "9: ssh")
		}
		parts = append(parts, "esc: back", "q: quit")
		helpText = joinStrings(parts, " • ")
	case ViewPodDetails:
		parts := []string{"↑/↓: scroll", "l: logs", "d: delete", "R: restart", "0: ns", "1: pods", "r: refresh"}
		if len(a.config.SSHHosts) > 0 {
			parts = append(parts, "9: ssh")
		}
		parts = append(parts, "esc: back", "q: quit")
		helpText = joinStrings(parts, " • ")
	case ViewLogs:
		parts := []string{"↑/↓/g/G: scroll", "f: follow", "t: timestamps", "r: refresh"}
		if len(a.logViewer.Containers()) > 1 {
			parts = append(parts, "c: container")
		}
		parts = append(parts, "esc: back", "q: quit")
		helpText = joinStrings(parts, " • ")
	case ViewMain:
		parts := []string{"0: namespaces", "1: pods"}
		if len(a.config.KubeConfigs) > 1 || a.connectionStatus == domain.StatusConnected {
			parts = append(parts, "esc: back")
		}
		if a.connectionStatus != domain.StatusConnected {
			parts = append(parts, "r: retry")
		}
		parts = append(parts, "q: quit")
		helpText = joinStrings(parts, " • ")
	case ViewSSHHosts:
		helpText = "↑/↓: navigate • enter: connect • /: filter • esc: back • q: quit"
	case ViewSSHConnecting:
		helpText = "ctrl+c: cancel"
	case ViewCrictlContainers:
		helpText = "↑/↓: navigate • enter: logs • /: filter • r: refresh • esc: back • q: quit"
	case ViewCrictlLogs:
		helpText = "↑/↓/g/G: scroll • f: follow • t: timestamps • r: refresh • esc: back • q: quit"
	}

	// Show notification if visible
	if a.notification.IsVisible() {
		return a.styles.Footer.
			Width(a.width - 4).
			Render(a.notification.View())
	}

	help := a.styles.Help.Render(helpText)

	// Status badge with color
	var statusStyle lipgloss.Style
	switch a.connectionStatus {
	case domain.StatusConnected:
		statusStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFFFF")).
			Background(colorSuccess).
			Padding(0, 1)
	case domain.StatusConnecting:
		statusStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFFFF")).
			Background(colorWarning).
			Padding(0, 1)
	case domain.StatusError:
		statusStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFFFF")).
			Background(colorError).
			Padding(0, 1)
	default:
		statusStyle = a.styles.StatusBar
	}

	statusText := a.connectionStatus.String()
	if a.clusterInfo != nil && a.connectionStatus == domain.StatusConnected {
		statusText = fmt.Sprintf("%s • %s", a.clusterInfo.Context, a.clusterInfo.Namespace)
	}
	statusBadge := statusStyle.Render(statusText)

	return a.styles.Footer.
		Width(a.width - 4).
		Render(fmt.Sprintf("%s  %s", help, statusBadge))
}

func (a *App) renderError() string {
	return a.styles.Content.
		Width(a.width - 4).
		Height(a.height - 10).
		Render(renderErrorBox(a.err, a.width))
}

func joinStrings(parts []string, sep string) string {
	if len(parts) == 0 {
		return ""
	}
	result := parts[0]
	for i := 1; i < len(parts); i++ {
		result += sep + parts[i]
	}
	return result
}

// overlayContainerSelector overlays the container selector centered on screen
func (a *App) overlayContainerSelector(view string) string {
	selector := a.containerSelector.View()
	if selector == "" {
		return view
	}

	// Use lipgloss.Place to center the dialog on a full-screen background
	return lipgloss.Place(
		a.width,
		a.height,
		lipgloss.Center,
		lipgloss.Center,
		selector,
		lipgloss.WithWhitespaceChars(" "),
		lipgloss.WithWhitespaceForeground(lipgloss.Color("#1a1a1a")),
	)
}

// overlayDialog overlays the confirmation dialog centered on screen
func (a *App) overlayDialog(view string) string {
	dialog := a.confirmDialog.View()
	if dialog == "" {
		return view
	}

	// Use lipgloss.Place to center the dialog on a full-screen background
	return lipgloss.Place(
		a.width,
		a.height,
		lipgloss.Center,
		lipgloss.Center,
		dialog,
		lipgloss.WithWhitespaceChars(" "),
		lipgloss.WithWhitespaceForeground(lipgloss.Color("#1a1a1a")),
	)
}

// SSH-related render functions

func (a *App) renderSSHHostsView() string {
	header := a.renderHeader()

	var contentStr string
	if len(a.config.SSHHosts) == 0 {
		contentStr = "No SSH hosts configured.\n\nAdd hosts to ~/.k4s/config.yaml:\n\nssh_hosts:\n  - name: \"my-node\"\n    host: \"192.168.1.100\"\n    user: \"admin\"\n    key_path: \"~/.ssh/id_rsa\"\n    port: 22"
	} else if a.err != nil {
		contentStr = a.renderError()
	} else {
		// Title
		title := fmt.Sprintf("SSH Hosts (%d)", len(a.config.SSHHosts))
		titleLine := lipgloss.NewStyle().
			Bold(true).
			Foreground(colorPrimary).
			Render(title)

		contentStr = titleLine + "\n\n" + a.sshHostList.View()
	}

	content := a.styles.Content.
		Width(a.width - 4).
		Height(a.height - 12).
		Render(contentStr)

	footer := a.renderFooter()

	return lipgloss.JoinVertical(lipgloss.Left, header, content, footer)
}

func (a *App) renderSSHConnecting() string {
	header := a.renderHeader()

	spinnerView := a.spinner.View()
	var connectingText string
	if a.selectedSSHHost != nil {
		connectingText = fmt.Sprintf("%s Connecting to %s@%s...", spinnerView, a.selectedSSHHost.User, a.selectedSSHHost.Host)
	} else {
		connectingText = fmt.Sprintf("%s Connecting via SSH...", spinnerView)
	}

	content := a.styles.Content.
		Width(a.width - 4).
		Height(a.height - 12).
		Render(connectingText)

	footer := a.renderFooter()

	view := lipgloss.JoinVertical(lipgloss.Left, header, content, footer)

	// Overlay passphrase input if visible
	if a.passphraseInput.IsVisible() {
		dialog := a.passphraseInput.View()
		return lipgloss.Place(
			a.width,
			a.height,
			lipgloss.Center,
			lipgloss.Center,
			dialog,
			lipgloss.WithWhitespaceChars(" "),
			lipgloss.WithWhitespaceForeground(lipgloss.Color("#1a1a1a")),
		)
	}

	return view
}

func (a *App) renderCrictlContainersView() string {
	header := a.renderHeader()

	var contentStr string
	if a.loading && len(a.crictlContainers) == 0 {
		contentStr = fmt.Sprintf("%s Loading containers...", a.spinner.View())
	} else if a.err != nil {
		contentStr = a.renderError()
	} else {
		// Title with node info
		var titleParts []string
		if a.selectedSSHHost != nil {
			titleParts = append(titleParts, fmt.Sprintf("Node: %s", a.selectedSSHHost.Name))
		}
		titleParts = append(titleParts, fmt.Sprintf("Containers: %d", len(a.crictlContainers)))

		titleLine := lipgloss.NewStyle().
			Bold(true).
			Foreground(colorPrimary).
			Render(joinStrings(titleParts, " • "))

		// Node info line
		var nodeInfoLine string
		if a.nodeInfo != nil {
			infoStyle := lipgloss.NewStyle().Foreground(colorMuted)
			infoParts := []string{}
			if a.nodeInfo.Hostname != "" {
				infoParts = append(infoParts, a.nodeInfo.Hostname)
			}
			if a.nodeInfo.OS != "" {
				infoParts = append(infoParts, a.nodeInfo.OS)
			}
			if a.nodeInfo.Memory != "" {
				infoParts = append(infoParts, fmt.Sprintf("Mem: %s", a.nodeInfo.Memory))
			}
			if a.nodeInfo.LoadAvg != "" {
				infoParts = append(infoParts, fmt.Sprintf("Load: %s", a.nodeInfo.LoadAvg))
			}
			nodeInfoLine = infoStyle.Render(joinStrings(infoParts, " | "))
		}

		// Column headers: NAME(25) POD(30) NS(15) STATE(10) AGE(6)
		headerLine := lipgloss.NewStyle().
			Bold(true).
			Foreground(colorMuted).
			Render(fmt.Sprintf("  %-25s %-30s %-15s %-10s %s", "NAME", "POD", "NAMESPACE", "STATE", "AGE"))

		if nodeInfoLine != "" {
			contentStr = titleLine + "\n" + nodeInfoLine + "\n" + headerLine + "\n" + a.crictlContainerList.View()
		} else {
			contentStr = titleLine + "\n" + headerLine + "\n" + a.crictlContainerList.View()
		}
	}

	content := a.styles.Content.
		Width(a.width - 4).
		Height(a.height - 12).
		Render(contentStr)

	footer := a.renderFooter()

	return lipgloss.JoinVertical(lipgloss.Left, header, content, footer)
}

func (a *App) renderCrictlLogsView() string {
	header := a.renderHeader()

	var contentStr string
	if a.loading {
		contentStr = fmt.Sprintf("%s Loading logs...", a.spinner.View())
	} else if a.err != nil {
		contentStr = a.renderError()
	} else {
		// Log viewer header
		logHeader := a.crictlLogViewer.RenderHeader()
		// Add search input if visible
		if a.searchInput.IsVisible() {
			logHeader += "\n" + a.searchInput.View()
		}
		contentStr = logHeader + "\n" + a.crictlLogViewer.View()
	}

	content := a.styles.Content.
		Width(a.width - 4).
		Height(a.height - 10).
		Render(contentStr)

	footer := a.renderFooter()

	return lipgloss.JoinVertical(lipgloss.Left, header, content, footer)
}
