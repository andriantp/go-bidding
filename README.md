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

- Fondasi Realtime Bidding: Menyiapkan Redis untuk Sistem WebSocket di Golang -> [chapter-1](https://andriantriputra.medium.com/bidding-1-fondasi-realtime-bidding-menyiapkan-redis-untuk-sistem-websocket-di-golang-14186134ebd0)
- Menghidupkan Realtime: Implementasi WebSocket di Golang -> [chapter-2](https://andriantriputra.medium.com/bidding-2-menghidupkan-realtime-implementasi-websocket-di-golang-78a3f85b516e)
- Menghubungkan Realtime dengan Redis: Pub/Sub di Golang
- Membangun Logic Bidding: Validasi dan State di Golang
- Menjaga Konsistensi: Atomic Bid dan Race Condition di Golang
- Dari Event ke Persistence: Menjaga Reliability Sistem Realtime
- Scaling Realtime Bidding: Multi Room dan Horizontal Scaling di Golang
- Distributed Worker dan Event Processing Pipeline di Golang
- Consumer Group dan Worker Coordination di Redis Streams
- Snapshot dan Fast Recove
- Sharding dan Partition Ownership
- Fault Tolerance dan Failover
- Event Sourcing dan CQRS
- Observability di Distributed System
- Dari Realtime Server ke Event-Driven Platform



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
