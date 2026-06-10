# go-bidding

A practical journey to build a realtime bidding system in Go, starting from a single WebSocket server and evolving into a distributed event-driven platform.

This repository contains the source code accompanying the **Go Bidding** article series, where we gradually explore the architectural challenges behind realtime systems such as auctions, live bidding, and event-driven applications.

## ✨ What You'll Learn

Throughout this series, we will build and evolve the system step by step:

* Realtime communication using WebSocket
* Redis Pub/Sub integration
* Bid validation and state management
* Atomic operations and race condition handling
* Persistence and reliability strategies
* Horizontal scaling techniques
* Distributed workers and event pipelines
* Redis Streams consumer coordination
* Snapshot and recovery mechanisms
* Sharding and partition ownership
* Fault tolerance and failover
* Event Sourcing and CQRS
* Observability in distributed systems
* Transitioning toward an event-driven platform

## 📚 Related Articles

| #  | Article                                                                     |
| -- | --------------------------------------------------------------------------- |
| 1  | Fondasi Realtime Bidding: Menyiapkan Redis untuk Sistem WebSocket di Golang |
| 2  | Menghidupkan Realtime: Implementasi WebSocket di Golang                     |
| 3  | Menghubungkan Realtime dengan Redis: Pub/Sub di Golang                      |
| 4  | Membangun Logic Bidding: Validasi dan State di Golang                       |
| 5  | Menjaga Konsistensi: Atomic Bid dan Race Condition di Golang                |
| 6  | Dari Event ke Persistence: Menjaga Reliability Sistem Realtime              |
| 7  | Scaling Realtime Bidding: Multi Room dan Horizontal Scaling di Golang       |
| 8  | Distributed Worker dan Event Processing Pipeline di Golang                  |
| 9  | Consumer Group dan Worker Coordination di Redis Streams                     |
| 10 | Snapshot dan Fast Recover                                                   |
| 11 | Sharding dan Partition Ownership                                            |
| 12 | Fault Tolerance dan Failover                                                |
| 13 | Event Sourcing dan CQRS                                                     |
| 14 | Observability di Distributed System                                         |
| 15 | Dari Realtime Server ke Event-Driven Platform                               |

Read the complete series on Medium:

👉 https://andriantriputra.medium.com

## 🛠️ Tech Stack

* Go
* WebSocket
* Redis
* Redis Pub/Sub
* Redis Streams
* Docker
* Event-Driven Architecture

## 🎯 Project Goal

The goal of this repository is not merely to build a bidding application, but to understand how modern distributed systems evolve from simple implementations into resilient and scalable architectures.

Each chapter introduces a new challenge and incrementally improves the system design.

## 👤 Author

**Andrian Tri Putra**

* Medium: https://andriantriputra.medium.com
* GitHub: https://github.com/andriantp
* GitHub (alternative): https://github.com/AndrianTriPutra

## 📄 License

Licensed under the Apache License 2.0.
