# RapidIO Concurrent Event Consumer Challenge

## Disclaimer

**Advanced Challenge**: This task requires a strong understanding of Go's concurrency features and is considered
advanced.

It is recommended that participants tackle the simpler challenges provided in this workshop before attempting
this Challenge.

This task is intended to test and improve your concurrency skills in Go and is best approached
with prior experience in concurrent programming.

## Overview

The RapidIO challenge invites participants to design and implement a high-performance, concurrent event consumer to
reduce the handling latency of I/O events. This challenge is ideal for developers looking to sharpen their skills in
concurrent programming and real-time system design in Go.

<img src="https://avatars.githubusercontent.com/u/50307654?s=280&v=4" alt="drawing" height="400"/>

## Challenge Description

A sequential event consumer implementation, part of the RapidIO suite, is provided as a starting point. The current
implementation processes events from multiple channels sequentially, which can lead to increased latency as the volume
of events grows.

Your mission is to create a **concurrent** version of the event consumer that significantly reduces event handling
latency, making full use of the Go programming language's concurrency features.

## Provided Materials

- A **sequential event consumer** serves as the baseline for your improvements.
- A **configurable event simulator** that generates a high volume of timestamped I/O events across multiple channels.
- A **plotter** utility to visualize and measure the latency between event creation and handling.
- A set of **simple tests** to verify the correctness of your event consumer implementation.

## Goals

- Implement the [concurrent version](concurrent.go) of the RapidIO interface
- Minimize the time difference between event emission and processing.
- Ensure that your solution is scalable and efficient under high load.

## Getting Started

1. Familiarize yourself with the provided sequential implementation and understand its limitations.
2. Plan your approach to introducing concurrency into the event handling process.
3. Implement your concurrent consumer, aiming to minimize the latency.
4. Use the provided plotter to visualize the latency improvements.
5. Validate your implementation with the provided tests.

