# Eboolkiq
Asynchronous task queue service using gRPC

**Note: This is not an official Daangn project**

# About
Eboolkiq is async task queue service. It manages tasks using gRPC protocol so
that any language that support gRPC can publish and consume their async task.

# Motivation
There are good projects that process asynchronous tasks such as sidekiq, salary,
and bull, but these projects are language specific libraries. There are good
projects that manage messages such as kafka, RabbitMQ, and ActiveMQ, but these
projects need well managed libraries because they use their own protocol.

This project start from small idea: What if there is an async task queue service
using gRPC protocol? The gRPC protocol is easy to use, quite familiar, faster
than REST api, and managed by Google. It is not language specific so that any
language can publish and consume async tasks.

# Roadmap
- Feature
   - [x] publish and consume task using memory
   - [x] support task finished confirmation
   - [ ] improve server command
   - [ ] support admin dashboard
   - [ ] support dead letter queue to handle dead tasks
   - [ ] support streaming publish and consume
   - [ ] support delayed task scheduling
- HA
   - [ ] sync queue to file
   - [ ] support for cluster
- Plugin Support
  - [ ] opentracing
  - [ ] grafana metric
  - [ ] custom plugins
