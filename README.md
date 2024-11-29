## Log prettier

!Note, this is currently in beta, work in progress

Install:
```
Todo
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

