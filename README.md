# Logger Example in Go

This Go program demonstrates how to use a custom logging library (`logger`) to add additional information to log entries and handle concurrent logging within goroutines.

## Features

- **Logger Initialization**: The logger is initialized using the `NewDefaultLogger` function.
- **Adding Additional Info**: Additional key-value pairs can be added to the log entries for context, using the `AddAdditionalInfo` function.
- **Logging in Goroutines**: The program logs messages within multiple goroutines to demonstrate concurrent logging.
- **Deleting Additional Info**: Information can be deleted from the logger context using the `DeleteAdditionalInfo` function.
- **Log Output**: Logs are written with context information, showing how the logger behaves with multiple goroutines and log modifications.

## Code Walkthrough

### 1. Logger Initialization

```go
logger := logger.NewDefaultLogger()
```

The logger is created using the `NewDefaultLogger()` function.

### 2. Adding Additional Information

```go
logger.AddAdditionalInfo("a", "1")
```

This line adds a key-value pair `a:1` to the logger's context, which will be included in any log entries made while this information is still present.

### 3. Logging with Context

```go
logger.Infow("routine1")
```

The logger logs the message "routine1" with the additional context `{"a": "1"}`.

### 4. Concurrent Logging in Goroutines

Two separate goroutines are created to log messages with different context:

```go
go func() {
    logger.Infow("routine2")
}()
go func() {
    logger.AddAdditionalInfo("b", "2")
    logger.Infow("routine3")
}()
```

- The first goroutine logs "routine2" with the context `{"a": "1"}` (from the main routine).
- The second goroutine adds new information (`b:2`) and logs "routine3" with this new context (`{"b": "2"}`).

### 5. Deleting Information

```go
logger.DeleteAdditionalInfo(1)
```

This line deletes the additional context added in the main routine (`{"a": "1"}`). After this, logs from the main routine will no longer contain the `a` key-value pair.

### 6. Final Log in the Main Routine

```go
logger.Infow("routine1")
```

This logs "routine1" again, but without the context `{"a": "1"}` since it was deleted earlier.

### 7. Waiting for Goroutines

```go
time.Sleep(time.Second)
```

This ensures that the main routine waits long enough for all goroutines to complete their logging before the program terminates.

## Log Output Example

The output will look something like this:

```
INFO    test/test.go:11 routine1        {"a": "1"}
INFO    test/test.go:20 routine1
INFO    test/test.go:17 routine3        {"b": "2"}
INFO    test/test.go:13 routine2
```

### Explanation of Output:
1. The first log entry shows "routine1" with the context `{"a": "1"}`.
2. The second log entry is from the main routine, after the additional info has been deleted, so no context is logged.
3. The third entry comes from the goroutine that logs "routine3" with the context `{"b": "2"}`.
4. The fourth entry comes from the goroutine logging "routine2" with the context `{"a": "1"}`.
