## Log Prettifier – Command Line Log Formatter

> Note: Currently in Beta, work in progress.

### Installation
To install the latest version of Log Prettifier, run the following command:
```
go install github.com/olbrichattila/logpretty/tree/main/cmd@latest
```

### Overview

Log Prettifier is a powerful command-line utility designed to format and improve the readability of your log files. Ideal for real-time debugging, it allows you to pipe your log data into the tool for a cleaner, more structured output.

Log Prettifier supports a wide range of log formats, with automatic log type detection for ease of use.

### Usage

You can use Log Prettifier with standard input, making it highly versatile for various logging needs. Below are some examples of how to use it:

* Format a log file:
```
cat ./mylog.log | logpretty
```

* Monitor and format logs in real time:
```
tail -f ./mylog.log | logpretty
```

### Supported Log Formats
Currently, Log Prettifier supports the following log formats:

- PHP Logs
- Apache Logs
- Laravel Logs
- Yii1 Logs
- Yii2 Logs
- Yii2 Logs (JSON format)
- Generic JSON Logs
- Any logs with space or bracket-separated entries
- More log formats coming soon!

### Upcoming Features
In future updates, Log Prettifier will include:
- Additional log format support
- A command-line parameter to disable auto-detection and specify a log type manually
- Custom log format configuration via environment variables or .env files (to be determined)

This project is licensed under the MIT License – see the LICENSE file for details.