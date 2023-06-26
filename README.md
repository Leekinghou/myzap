# 😎 myzap

![License](https://img.shields.io/badge/license-MIT-blue.svg)  

**[中文版](README-CN.md)**

MyZap is a powerful logging library built on top of Zap. It aims to provide a better logging experience by offering the following features:

## Features

1. 🎨 Customizable Log Format: You can easily customize the log format according to your preferences. With MyZap, you have full control over how your log messages are structured.

2. ⚙️ Logging Options: MyZap supports various logging options that allow you to configure the behavior of your logs. You can specify the log level, output location, and more.

3. 📦 Multiple Output Locations: MyZap enables you to direct logs of different levels to different output locations. For example, you can send error logs to a file and info logs to the console simultaneously.

4. 🔄 Log Rotation: MyZap supports log rotation based on both size and time. You can configure the maximum log file size or set a specific time interval for log rotation.

## Installation

To use MyZap in your project, simply import it using Go modules:

```shell
go get github.com/Leekinghou/myzap
```

## Usage

Here's a simple example demonstrating how to use MyZap in your Go application:

```go
package main

import (
	"github.com/Leekinghou/myzap"
)

func main() {
	// Create a new logger
	logger := myzap.NewLogger()

	// Set the log format
	logger.SetLogFormat("[${LEVEL}] ${MESSAGE}")

	// Set different output locations for different log levels
	logger.SetOutputLocation(myzap.LogLevelError, "/path/to/error.log")
	logger.SetOutputLocation(myzap.LogLevelInfo, "/path/to/info.log")

	// Enable log rotation
	logger.EnableLogRotation(10, myzap.LogRotationBySize) // Rotate logs when they reach 10 MB
	logger.EnableLogRotation(24, myzap.LogRotationByTime) // Rotate logs every 24 hours

	// Start logging
	logger.Debug("This is a debug log")
	logger.Info("This is an info log")
	logger.Warn("This is a warning log")
	logger.Error("This is an error log")
}
```

For more details on how to customize the logger and use advanced features, please refer to the [documentation](https://github.com/Leekinghou/myzap/docs).

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvement, feel free to open an issue or submit a pull request.

## License

MyZap is open-source software licensed under the [MIT license](LICENSE).