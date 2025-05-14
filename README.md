# 🧩 Order Management System – Golang Microservices Architecture

A scalable, event-driven backend system for processing customer orders and payments in real time. Built with Go microservices, gRPC for communication, RabbitMQ for asynchronous messaging, and Consul for service discovery. Designed to mimic real-world e-commerce platforms and support production-ready features like observability, retries, and modularity.

---

## 🖼️ Architecture Diagram
![diagram](https://github.com/user-attachments/assets/62342f5a-3117-444a-9579-282a6b7681c4)

> _Note: Replace the above path with the actual location of your architecture image file._

---

## 🧠 Key Features

- **Microservices Design**: Modular services for orders, payments, and gateway using Go
- **Dynamic Service Discovery**: Consul registers and discovers services with built-in health checks
- **gRPC + Protobuf**: Efficient binary communication with well-defined API contracts
- **RabbitMQ Integration**: 
  - Publishes `order.created` events from OrderService
  - PaymentService consumes events and emits `order.paid`
- **Stripe Integration (Planned)**: Generate real payment links and handle webhook responses
- **Real-Time Fulfillment**: Event-driven design enables instant processing and downstream propagation
- **Message Reliability**: Dead Letter Queue (DLQ) and retries for failed delivery (Planned)
- **Observability**: Distributed tracing with OpenTelemetry (Planned)
- **Persistence**: MongoDB for storing orders and payments (Planned)

---

## 🧰 Tech Stack

| Category         | Technologies                          |
|------------------|----------------------------------------|
| Language         | Golang                                 |
| Communication    | gRPC, Protocol Buffers                 |
| Message Broker   | RabbitMQ (amqp091-go)                  |
| Discovery        | Consul                                 |
| Database         | MongoDB *(planned)*                    |
| Payments         | Stripe *(planned)*                     |
| Observability    | OpenTelemetry *(planned)*              |
| Development      | Air (hot reload), Docker, Postman      |

---

## 📁 Folder Structure

```
OrderManagementService/
├── common/              # Shared components: broker, discovery, proto
├── gateway/             # HTTP server acting as entrypoint
├── orders/              # OrderService (gRPC server + publisher)
├── payments/            # PaymentService (AMQP consumer + processor)
├── docs/                # Architecture images, design documents
└── README.md
```

---

## 🚀 Getting Started

### Prerequisites

- [Go 1.21+](https://go.dev/dl/)
- [RabbitMQ](https://www.rabbitmq.com/download.html)
- [Consul](https://www.consul.io/downloads)
- [Docker](https://www.docker.com/) (for containerized deployments)

### Setup

```bash
# Clone the project
$ git clone https://github.com/manush2312/Order-Management-System.git
$ cd Order-Management-System

# Start services (in separate terminals)
$ cd orders && air
$ cd payments && air
$ cd gateway && air
```

### Local RabbitMQ (Docker)
```bash
docker run -d --name rabbitmq \
  -p 5672:5672 -p 15672:15672 \
  rabbitmq:3-management
```

### Local Consul (Docker)
```bash
docker run -d --name=consul \
  -p 8500:8500 -p 8600:8600/udp \
  consul agent -dev -client=0.0.0.0
```

---

## ✅ Roadmap / In Progress

- [x] Order creation and event publishing
- [x] Payment microservice with AMQP consumer
- [x] Real-time fulfillment via RabbitMQ
- [x] Stripe payment gateway integration
- [x] `order.paid` fanout to multiple services
- [x] MongoDB integration for data persistence
- [x] Retry mechanisms and DLQ
- [x] OpenTelemetry for observability
- [x] Unit & integration testing

---

## 🤝 Contributing

Feel free to fork, submit PRs, or open issues. Contributions welcome!

---
