# GoLatencyTester
SingleSideLatency - A UDP-based Network Delay Measurement Tool

SingleSideLatency is a simple yet powerful tool for measuring network delay using UDP protocol. It enables accurate one-way latency measurement between a sender and receiver. This Go language project provides high-precision results by using precise timestamps to calculate delays.

Key Features:
- Measure one-way network latency using UDP packets.
- High precision and accuracy in latency measurements.
- Easy-to-use command-line interface for both sender and receiver.
- Suitable for benchmarking and performance testing of network applications.

Usage:
To measure network latency, run the sender and receiver with appropriate IP address and port settings. The sender will send UDP packets with timestamps, and the receiver will calculate the one-way delay and display the results in milliseconds.

Dependencies:
This project is implemented in Go language and relies on the net package for UDP communication.

Get Started:
1. Clone the repository.
2. Run 'go run main.go -mode send -address [receiver IP] -port [receiver port]' to start the sender.
3. Run 'go run main.go -mode recv -port [receiver port]' to start the receiver.

Note: Ensure that the receiver is running before starting the sender.

Enjoy precise network delay measurements with SingleSideLatency!
