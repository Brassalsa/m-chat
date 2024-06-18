# M-Chat

## Description

M-Chat is a real-time chat application built with Go, htmx, Tailwind CSS, and MongoDB. It offers a sleek user interface and efficient communication powered by modern web technologies.

## Installation

To set up M-Chat locally, follow these steps:

1. Clone the repository:
   ```bash
   git clone https://github.com/Brassalsa/m-chat.git
   ```
   Navigate to the project directory:
   cd m-chat

## Install dependencies:

### Go dependencies

```bash
go get .
```

### Tailwind CSS (assuming you have Node.js and npm installed)

Globally install tailwind css

```bash
npm install -g tailwindcss
```

## Usage

**Important note: Install makefile tools on your machine to run make scripts**

### Build Binary

To build binary for your os run:

```bash
make build
```

### Run Server

To run start the Go server:

```bash
make run
```

Access the application at http://localhost:3000.

### Run Dev server

**Important note: Install [air](https://github.com/air-verse/air) on your machine to run dev server**

```bash
make dev
```

Access the application at http://localhost:3000.

## Run Tests

For running tests:

```bash
make test
```

### Tailwind CSS

Build main.css by running:

```bash
make css-dev
```

Build and watch for css changes:

```bash
make css-dev
```
