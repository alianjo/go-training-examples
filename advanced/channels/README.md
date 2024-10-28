# README

## Overview

This Go package demonstrates the use of channels for concurrent data transfer between goroutines. It includes two main functions: `sendData` and `receiveData`, which showcase sending and receiving data through a channel, respectively.

## Functions

### sendData

*   Sends integers from 0 to 4 through a channel with a 500ms delay between each send operation.
*   Closes the channel after sending all data.

### receiveData

*   Continuously receives data from a channel and prints it to the console.
*   Detects when the channel is closed and exits the function.

## Usage

1.  Create a channel using `make(chan int)`.
2.  Start the `sendData` function as a goroutine, passing the channel as an argument.
3.  Start the `receiveData` function as a goroutine, passing the same channel as an argument.
4.  Use `time.Sleep` to allow the goroutines to run for a specified duration.

## Example

```go
func main() {
    ch := make(chan int)
    go sendData(ch)
    go receiveData(ch)
    time.Sleep(3 * time.Second)
}
```

## Notes

*   This package uses the `time` package for delay operations.
*   The `select` statement is used in the `receiveData` function to handle channel closure.
*   The `close` function is used to close the channel after sending all data in the `sendData` function.