# Rover

![Gopher on a rover on Mars](./assets/gopher.webp)

Welcome to the Rover project! This repository contains a Go-based application designed to simulate Rover communications.

## Table of Contents

- [Features](#features)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Usage](#usage)
- [Project Structure](#project-structure)
- [Contributing](#contributing)
- [License](#license)

## Features

- Rover movements
- Socket communication

## Getting Started

### Prerequisites

Make sure you have the following installed:

- Go (version 1.23 or later)
- Git

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/MrTanguy/rover.git
   ```
2. Navigate to the project directory:
   ```bash
   cd rover
   ```
3. Install dependencies:
   ```bash
   go mod tidy
   ```

### Usage

Run the application with:
```bash
go run main.go
```

Connect to the socket with:
```bash
telnet [IP] [PORT]
```

## Project Structure

- `.github/workflows/` - CI/CD workflows for GitHub Actions.
- `assets/` - Readme image.
- `builder/` - Rover Builder.
- `communication/` - Socket.
- `rover/` - Rover movements.
- `topologie/` - Planet Builder.
- `utils/` - Usefull functions.
- `main.go` - Entry point of the application.

## Contributing

Contributions are welcome! Please follow these steps:

1. Fork the repository.
2. Create a new branch:
   ```bash
   git checkout -b feature-branch
   ```
3. Commit your changes:
   ```bash
   git commit -m "Add some feature"
   ```
4. Push to the branch:
   ```bash
   git push origin feature-branch
   ```
5. Open a pull request.

## License

This project is licensed under the [LICENSE NAME] License - see the [LICENSE](LICENSE) file for details.

---

For further assistance, feel free to open an issue or contact the repository owner.
