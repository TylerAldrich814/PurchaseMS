# Go Microservices Project

## Overview
This project is strictly for learning and understanding about microservices in Go. Demonstrating how multiple services can work together using modern tools and design patterns to create a scalable and maintainable architecture. The project is designed to showcase key concepts such as service communication, data management, and telemetry in a distributed system.

## Services
The project currently includes the following services:

1. **Gateway**: Handles all HTTP communication between the client and backend. Routes requests to the appropriate microservice through gRPC channels.
2. **Orders**: Records new orders, interacts with the Payments service, and manages order status.
3. **Payments**: Implements a hexagonal architecture for payment processing. Integrates with external payment providers (e.g., Stripe).
4. **Kitchen** (Planned): Will handle order preparation and management.
5. **Stock** (Planned): Will track inventory and validate item availability.

Additionally, a shared package **`common`** contains reusable logic and utilities used across the services.

## Technology Stack
The project leverages the following technologies:

- **gRPC with Protobuf**: For efficient and structured service-to-service communication.
- **RabbitMQ**: For inter-service messaging and event-driven communication.
- **Consul**: For service discovery and network mesh management.
- **Jaeger**: For distributed tracing and telemetry.
- **NoSQL Databases**: Each service owns and manages its own NoSQL database.

### Key Features

- **Hexagonal Architecture in Payments**: The Payments service follows the hexagonal (ports and adapters) design pattern to interact with external payment systems. Stripe is currently implemented as the payment handler.
- **Distributed Design**: Each service is isolated, owns its database, and communicates via gRPC or message queues.
- **Telemetry**: Jaeger provides end-to-end tracing for request flows across the microservices.

## How It Works

### General Flow
1. A user sends an HTTP request to the Gateway.
2. The Gateway translates the HTTP request into a gRPC message and routes it to the appropriate service.
3. Services communicate and perform their respective responsibilities via gRPC and RabbitMQ.
4. Responses are returned through the Gateway back to the user.

### Example Workflow: Placing an Order
1. **Gateway**: Receives an HTTP request to create an order.
2. **Orders**: Records the new order and sends a gRPC request to the Payments service.
3. **Payments**: 
    - Initiates a request to the payment handler (e.g., Stripe).
    - On success, returns the payment link to the Orders service.
4. **Orders**: Updates the order status and sends the response back through the Gateway to the user.


## Project Structure
```
├── common         # Shared logic and utilities
├── gateway          # HTTP server for client communication
├── orders            # Order management service
├── payments       # Payment processing service with hexagonal design
├── kitchen          # Planned: Order preparation service
├── stock             # Planned: Inventory management service
```

## Requirements
- Go (1.20+)
- Docker (optional, for running dependencies)
- RabbitMQ
- Consul
- Jaeger
- Stripe account (for Payments service)

## Getting Started
1. Clone the repository:
   ```bash
   git clone https://github.com/TylerAldrich814/PurchaseMS.git
   cd ./PurchaseMS
   ```

2. Install dependencies:
   ```bash
dirs=("gateway" "kitchen" "orders" "payments" "stock")
for dir in "${dirs[@]}"; do
       pushd ./${dir}} && go mod tidy && popd
done
   
   ```

3. Run services:
   Each service can be started individually by navigating to its directory and running:
   ```bash
   go run . 
// or
air init && air
   ```

4. Spin up Docker Images:
   ```bash
   docker-compose up
   ```

5. Access services through the Gateway:
   ```bash
   curl -X POST localhost:8080/api/customers/2/orders -d \
   '{"ID": 123, "Quantity": 2}'
   ```

## Flowchart
```flow
st=>start: User Request
op1=>operation: Gateway Receives Request
op2=>operation: Orders Service Processes Request
op3=>operation: Payments Service Processes Payment
op4=>operation: Stripe Payment Handler
cond1=>condition: Stock Validated?
cond2=>condition: Payment Successful?
e=>end: Response to User

st->op1->op2->cond1
cond1(yes)->op3->op4->cond2
cond2(yes)->e
cond1(no)->e
cond2(no)->e
```

## Sequence Diagram
```seq
User->Gateway: Sends HTTP Request
Gateway->Orders: gRPC Request for New Order
Orders->Stock: Validates Stock
Stock-->Orders: Returns Stock Status
Orders->Payments: Initiates Payment
Payments->Stripe: Processes Payment
Stripe-->Payments: Returns Payment Link
Payments-->Orders: Confirms Payment
Orders-->Gateway: Sends Response
Gateway-->User: Returns Payment Link
```

## Future Plans
- Implement Kitchen and Stock services.
- Add robust error handling and logging.
- Extend the Payments service with additional providers (e.g., PayPal).
- Integrate a CI/CD pipeline.

