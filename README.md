# Discord Nitro Free Trial Generator

This Go script generates Discord Nitro free trials and saves the access details in `output.txt`. It automates requests to the Discord API to retrieve trial tokens which are then formatted and stored for easy access.

## Prerequisites

Before you can run this script, ensure you have the following installed:

- [Go (Golang)](https://golang.org/dl/) - The programming language used to write and run the script.
- [Git](https://git-scm.com/downloads) - For cloning the repository.

## Installation

Follow these steps to get the script running on your system:

1. **Clone the Repository**:
    ```bash
    git clone https://github.com/PatFDev/DiscordNitroGen
    cd DiscordNitroGen
    ```

2. **Download Dependencies**:
    The script uses the `github.com/google/uuid` package to generate UUIDs. Install it using:
    ```bash
    go get github.com/google/uuid
    ```

## Running the Script

To run the script, use the following command in the terminal:

```bash
go run .
