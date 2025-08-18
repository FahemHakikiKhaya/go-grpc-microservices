# Go gRPC Microservices

A collection of microservices built in **Go** using **gRPC** for inter-service communication.

---

## 📂 Project Structure
- **common/** → Shared proto files and utilities  
- **gateway/** → API Gateway that routes requests to services  
- **kitchen/** → Kitchen service (food preparation logic)  
- **menu/** → Menu service (manages menu items)  
- **orders/** → Orders service (handles customer orders)  

---

## 🚀 Features
- gRPC-based communication  
- API Gateway entry point  
- Modular service design  
- Go workspace (`go.work`) setup  

---

## 🛠️ Prerequisites
- Go 1.21+  
- Protocol Buffers (`protoc`)  
- Plugins:  
  - `protoc-gen-go`  
  - `protoc-gen-go-grpc`  

---

## ⚙️ Setup
```bash
# Clone repo
git clone https://github.com/FahemHakikiKhaya/go-grpc-microservices.git
cd go-grpc-microservices

# Install dependencies
go mod tidy
