## Log Prettifier – Command Line Log Formatter

> Note: Currently in Beta, work in progress.

### Installation
To install the latest version of Log Prettifier, run the following command:
```
go install github.com/olbrichattila/logpretty/cmd/logpretty@latest
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


### Manual Log Parser Selection:
By default, Log Prettier automatically detects the log format. However, you can override this behavior and force a specific parser by using one of the following parameters:

- php: For PHP logs
- apache: For Apache access logs
- laravel: For Laravel logs
- yii2: For Yii2 logs
- yii1: For Yii1 logs
- json: For JSON logs

### Example:
```
tail -f ./mylog.log | logpretty laravel
```

### Upcoming Features
In future updates, Log Prettifier will include:
- Additional log format support
- Custom log format configuration via environment variables or .env files (to be determined)

This project is licensed under the MIT License – see the LICENSE file for details.