## Prerequisites

- **Go**: Ensure you have [Go](https://golang.org/doc/install) installed (version >=1.18).
- **PostgreSQL**: Install [PostgreSQL](https://www.postgresql.org/download/), and ensure it's running.
- **Git**: Install [Git](https://git-scm.com/).

## Installation

1. **Clone the Repository**

   Clone the repository to your local machine using Git.

   ```bash
   git clone https://github.com/langdag/poe_chat_api.git

2. **Install Docker and Docker Compose
   Ubuntu:

   # Add Docker's official GPG key:
   sudo apt-get update
   sudo apt-get install ca-certificates curl
   sudo install -m 0755 -d /etc/apt/keyrings
   sudo curl -fsSL https://download.docker.com/linux/ubuntu/gpg -o /etc/apt/keyrings/docker.asc
   sudo chmod a+r /etc/apt/keyrings/docker.asc

   # Add the repository to Apt sources:
   echo \
   "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/ubuntu \
   $(. /etc/os-release && echo "$VERSION_CODENAME") stable" | \
   sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
   sudo apt-get update

   # Install Docker and packages:
   sudo apt-get install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin

   # Install Docker Compose
   sudo curl -L "https://github.com/docker/compose/releases/download/1.29.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
   sudo chmod +x /usr/local/bin/docker-compose