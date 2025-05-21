# TCP Chat – NetCat Clone in Go

## Description

This project is a simplified clone of the NetCat (`nc`) utility, built in Go. It enables multiple clients to connect to a server via TCP and chat in real-time. The server listens on a specified port and handles multiple client connections simultaneously.

## Features

- TCP server with support for up to 10 clients
- Each client must enter a non-empty username
- Clients receive full chat history upon joining
- Messages are timestamped and identified by the sender
- Join and leave notifications are broadcast to all clients
- Empty messages are ignored
- Default port is 8989
- Graceful handling of client disconnections

## Usage

### Run the Server

```bash
go run .             # Runs on default port 8989
go run . 2525        # Runs on specified port 2525
