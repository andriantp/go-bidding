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

- Fondasi Realtime Bidding: Menyiapkan Redis untuk Sistem WebSocket di Golang -> [Chapter-1](https://andriantriputra.medium.com/bidding-1-fondasi-realtime-bidding-menyiapkan-redis-untuk-sistem-websocket-di-golang-14186134ebd0)
- Menghidupkan Realtime: Implementasi WebSocket di Golang -> [Chapter-2](https://andriantriputra.medium.com/bidding-2-menghidupkan-realtime-implementasi-websocket-di-golang-78a3f85b516e)
- Menghubungkan Realtime dengan Redis: Pub/Sub di Golang -> [Chapter-3](https://andriantriputra.medium.com/bidding-3-menghubungkan-realtime-dengan-redis-pub-sub-di-golang-8d23db424878)
- Membangun Logic Bidding: Validasi dan State di Golang -> [Chapter-4](https://andriantriputra.medium.com/bidding-4-membangun-logic-bidding-validasi-dan-state-di-golang-f8f6bea6a56c)
- Menjaga Konsistensi: Atomic Bid dan Race Condition di Golang -> [Chapter-5](https://medium.com/@andriantriputra/bidding-5-menjaga-konsistensi-atomic-bid-dan-race-condition-di-golang-7d36fc742dad)
- Dari Event ke Persistence: Menjaga Reliability Sistem Realtime -> [Chapter-6](https://andriantriputra.medium.com/bidding-6-dari-event-ke-persistence-menjaga-reliability-sistem-realtime-09e659a13bd3)
- Scaling Realtime Bidding: Multi Room dan Horizontal Scaling di Golang -> [Chapter-7](https://medium.com/@andriantriputra/bidding-7-scaling-realtime-bidding-multi-room-dan-horizontal-scaling-di-golang-36823b9b772b)
- Distributed Worker dan Event Processing Pipeline di Golang -> [Chapter-8](https://medium.com/@andriantriputra/bidding-8-distributed-worker-dan-event-processing-pipeline-di-golang-9119bfb180b8)
- Consumer Group dan Worker Coordination di Redis Streams -> [Chapter-9](https://medium.com/@andriantriputra/bidding-9-consumer-group-dan-worker-coordination-di-redis-streams-7218ee58547d)
- Snapshot dan Fast Recove -> [Chapter-10](https://medium.com/@andriantriputra/bidding-10-snapshot-dan-fast-recovery-mengurangi-replay-di-distributed-system-0509e18f9d35)
- Sharding dan Partition Ownership -> [Chapter-11]()
- Fault Tolerance dan Failover -> [Chapter-12]()
- Event Sourcing dan CQRS -> [Chapter-13]()
- Observability di Distributed System -> [Chapter-14]()
- Dari Realtime Server ke Event-Driven Platform -> [Chapter-15]()



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
