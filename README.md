## Before you run the application

You need to create a self-signed TLS certificate inside the folder ./tls
To do this, run the following commands (starting from the project's root directory)

```
mkdir tls
cd tls
go run /path/to/your/go/standard/library/src/crypto/tls/generate_cert.go --rsa-bits=2048 --host=localhost
```

Note that you path to your Go standard library may differ depending on which system you are using and how you installed Go. For macOS, if you installed Go using Homebrew, it will probably be something like `/usr/local/Cellar/go/<version>/libexec/src/crypto/tls/generate_cert.go`. For Linux installations, if you've downloaded the Go binaries and followed the recommended method of installing Go, it will probably be `/usr/local/go/src/crypto/tls/generate_cert.go`. If installed via package manager (like `apt` for Ubuntu based distros), it will probably be located in `/usr/share/go-<version>/src/crypto/tls/generate_cert.go`.