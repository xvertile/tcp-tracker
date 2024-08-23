# TCP Tracker
### quick description
`tcp-tracker` is a simple Go package that wraps a `net.Conn` and tracks the number of bytes read and written over the connection. It also allows you to set a maximum byte limit, after which the connection will be automatically closed.

## Features

- **Byte Counting**: Track the total number of bytes read and written.
- **Max Byte Limit**: Automatically close the connection if the byte limit is exceeded.
- **Prevent Overspill**: Ensure that the connection does not read or write more bytes than expected.

## Installation

To install the package, use the following command:

```bash
go get -u xvertile/tcp-tracker
```

## Usage

Here is a quick example of how to use the `tcp-tracker` package:

```go
package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"xvertile/tcp-tracker/tracker"
)

func main() {
	conn, err := net.Dial("tcp", "icanhazip.com:80")
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()
	maxBytes := int64(1024)
	trackedConn := tracker.CreateCountingConn(conn, maxBytes)
	request := "GET / HTTP/1.1\r\nHost: icanhazip.com\r\nConnection: close\r\n\r\n"
	_, err = trackedConn.Write([]byte(request))
	if err != nil {
		log.Fatalf("Failed to send request: %v", err)
	}
	response := make([]byte, 4096)
	for {
		n, err := trackedConn.Read(response)
		if err != nil && err != io.EOF {
			log.Fatalf("Failed to read response: %v", err)
		}
		if n == 0 || err == io.EOF {
			break
		}
		fmt.Print(string(response[:n]))
	}
	log.Printf("Total bytes used: %d\n", trackedConn.BytesRead)
}

```

### Explanation

- **CountingConn**: Wraps a `net.Conn` to track and limit the number of bytes read and written.
- **MaxBytes**: Automatically closes the connection when the byte limit is exceeded.

### Benifits
Compared to io.Copy, this package provides the following benefits:
- **Byte Counting**: Track the total number of bytes read and written.
- **Max Byte Limit**: Automatically close the connection if the byte limit is exceeded.
- **Prevent Overspill**: Ensure that the connection does not read or write more bytes than expected.

### What about io.Copy?
if a connection is left open indefinitely, the io.Copy will not prevent overspill from happening as io.Copy does not track every read and write only when a connection is fully closed do you get the byte count back from the io.Copy function.

### Running the Example

1. Save the example code to a file, e.g., `main.go`.
2. Run the file using:

   ```bash
   go run main.go
   ```

3. The program will print the response from `icanhazip.com` and log the total number of bytes used.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contributions

Contributions are welcome! Feel free to submit a pull request or open an issue.
