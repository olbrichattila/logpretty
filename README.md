## Log prettier

!Note, this is currently in beta, work in progress

Install:
```
go install github.com/olbrichattila/logpretty/tree/main/cmd@latest
```

This command line utility re formats your logs. This can be used for real time debugging:

It uses standard input to receive the data, therefore it can be piped.

Use cases:

```
cat ./mylog.log | logpretty
```

```
tail -f ./mylog.log | logpretty
```

Currently it tries to automatically identify the log type.
The currently supported types:

- PHP
- Apache
- Laravel
- Yii1
- Yii2
- Yii2 (JSON format)
- Generic JSON format
- Any separated logs with [] or blanks
- Others coming soon.

Coming soon:
- New log types
- Command line parameter to skip auto detect and force a log type
- Custom defined log format (from env or .env file, to be decided)
