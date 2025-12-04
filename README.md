# Lidless

A macOS menubar app that prevents your Mac from sleeping with a simple toggle interface.

## Description

Lidless adds an eye icon to your macOS menubar that lets you quickly enable or disable system sleep. When sleep is disabled, your Mac stays awake - perfect for long-running tasks, presentations, or downloads. The icon changes to show the current state (sun when awake, moon when sleep is enabled).

Features:

- Simple menubar toggle interface
- Visual feedback with icon states (☀️ awake / 🌙 sleep)
- Syncs with system sleep state on startup
- Uses macOS native `pmset` commands with Touch ID authentication

## Installation

### Download Pre-built Binary (Recommended)

Download the latest release for your Mac architecture:

**Intel Macs (amd64):**

```bash
curl -L https://github.com/robertarles/lidless/releases/latest/download/lidless-darwin-amd64.tar.gz | tar xz
sudo mv lidless-darwin-amd64 /usr/local/bin/lidless
chmod +x /usr/local/bin/lidless
```

**Apple Silicon Macs (arm64):**

```bash
curl -L https://github.com/robertarles/lidless/releases/latest/download/lidless-darwin-arm64.tar.gz | tar xz
sudo mv lidless-darwin-arm64 /usr/local/bin/lidless
chmod +x /usr/local/bin/lidless
```

Or download manually from the [releases page](https://github.com/robertarles/lidless/releases).

### Via Go Install

```bash
go install github.com/robertarles/lidless/cmd/lidless@latest
```

After installation, run `lidless` from your terminal to start the menubar app.

### From Source

```bash
# Clone the repository
git clone https://github.com/robertarles/lidless.git
cd lidless

# Build and install to ~/bin
make full-build-install-user

# Or install system-wide to /usr/local/bin
make full-build-install-system
```

## Usage

1. Launch the app - an eye icon appears in your menubar
2. Click the icon to see the current state
3. Select "Toggle Sleep" to enable/disable system sleep
4. The icon changes to reflect the current state:
   - ☀️ Sun = Sleep disabled (system stays awake)
   - 🌙 Moon = Sleep enabled (normal behavior)

The app will prompt for authentication (Touch ID or password) when changing sleep settings.

## OPTIONAL Configure sudo to use Touch ID

This makes toggling VERY much more convenient if you have touch-id.

Edit /etc/pam.d/sudo:
`sudo nano /etc/pam.d/sudo`
Add this as the first line:
  auth       sufficient     pam_tid.so
  
## Requirements

- macOS (uses `pmset` system command)
- Go 1.25.5 or later (for building from source)
