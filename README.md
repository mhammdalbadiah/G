# G

### Simple Git Task Runner (Built with Go)

**G** is a lightweight command-line tool written in **Go** that simplifies  
common Git commands into short, easy-to-use actions.

---

## Features

- One command for **add + commit + push**
- Faster Git workflow
- Cross-platform (macOS, Linux, Windows)
- Single binary, no dependencies

---

## Installation

```bash
git clone https://github.com/mhammdalbadiah/G.git
cd G
go build -o G
sudo mv G /usr/local/bin/



G help
G status
G pull
G push "Commit message"
G clone <repo_url>
