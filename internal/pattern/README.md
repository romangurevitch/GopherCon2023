# Understanding Concurrent Design Patterns in Golang

Concurrent design patterns in Go enable the development of applications with parallel processing, where multiple
computations are executed simultaneously. This document provides an insight into different concurrent design patterns,
comparing them based on certain attributes to aid in understanding and choosing the right pattern based on the
requirements.

## Table of Contents

1. [Introduction to Concurrent Design Patterns](#introduction)
2. [Comparison Descriptions](#comparison-descriptions)
3. [Comparison Table](#comparison-table)

## Introduction

Concurrent design patterns are crucial in building efficient and scalable applications, especially in a language like Go
which has built-in support for concurrent programming. Understanding the differences between these patterns, their use
cases, and how they handle data can lead to better design decisions when developing Go applications.

## Comparison Descriptions

- **Input Type**:
    - **Single**: The input data size is fixed and singular.
    - **Bounded**: The size of the input data is known and finite.
    - **Unbounded**: The size of the input data is unknown or potentially infinite.

- **Process Duration**:
    - **Short**: Suitable for operations that complete quickly.
    - **Long**: Suitable for operations that take a significant amount of time or run indefinitely.
    - **Short to Long**: Adaptable to both short and long running processes.

- **Synchronization**: The mechanism through which concurrent tasks coordinate with each other.

- **Latency**: The time it takes to process one unit of work.

- **Throughput**: The amount of work done or the number of tasks completed per unit of time.

- **Data Flow**: The path along which data travels through the system or process.

## Comparison Table

| Design Pattern                   | Input Type | Process Duration | Synchronization                   | Latency                                                    | Throughput                                            | Data Flow                        | Use Case and Application Examples                                     |
|----------------------------------|------------|------------------|-----------------------------------|------------------------------------------------------------|-------------------------------------------------------|----------------------------------|-----------------------------------------------------------------------|
| Future                           | Single     | Short            | Blocking until result ready       | <span style="color:red">Potential Increased Latency</span> | Standard Throughput                                   | Request -> Computation -> Result | Async Computations, Async API Calls                                   |
| Pipeline                         | Unbounded  | Long             | Sequential Execution              | <span style="color:red">Sequential Latency</span>          | <span style="color:red">Sequential Throughput</span>  | Stage-wise Processing            | Stream Processing, Data Transformation Pipelines                      |
| Fan-out Fan-in                   | Bounded    | Short to Long    | Goroutine synchronization         | <span style="color:green">Reduced Latency</span>           | <span style="color:green">Increased Throughput</span> | Task -> Worker -> Aggregator     | CPU bound parallel tasks, Data Processing, Image Processing           |
| Worker Pool                      | Unbounded  | Short to Long    | Worker Coordination               | <span style="color:green">Reduced Latency</span>           | <span style="color:green">Increased Throughput</span> | Task -> Worker -> Result         | I/O or CPU Bound Tasks, Task Processing Systems                       |
| Dynamic Rate-Limited Worker Pool | Unbounded  | Long             | Rate Limiter, Worker Coordination | Controlled Latency                                         | Controlled Throughput                                 | Task -> Worker -> Result         | External Rate Limits, Resource Management, API Clients, Microservices |
| Pub-Sub                          | Unbounded  | Long             | Topic-based Subscription          | Event Delivery Latency                                     | Varied Based on Subscribers                           | Event Broadcast                  | Event Broadcasting, Event Notification Systems                        |

This table encapsulates a comparative overview of various concurrent design patterns in Go, delineating their key
attributes, typical use cases, and application examples.
