# Project Overview

This project sets up a cloud-based Identity and Access Management (IAM) service over Hyperledger Fabric, with Go microservices and Terraform for infrastructure provisioning.

## Project Structure

```

.
├── README.md                          // Project overview and setup instructions
├── network                            // Hyperledger Fabric network configuration
│   ├── organizations                  // Certificate Authorities (CAs) and MSPs for orgs
│   │   ├── peerOrganizations
│   │   │   ├── https://www.google.com/search?q=org1.example.com
│   │   │   │   ├── ca
│   │   │   │   ├── msp
│   │   │   │   ├── peers
│   │   │   │   └── users
│   │   │   └── https://www.google.com/search?q=org2.example.com
│   │   └── ordererOrganizations
│   │       └── example.com
│   │           ├── ca
│   │           ├── msp
│   │           └── orderers
│   ├── scripts                        // Network setup and teardown scripts (e.g., start.sh, stop.sh)
│   ├── configtx                       // Configuration transaction definitions (e.g., configtx.yaml)
│   ├── connection-profiles            // Connection profiles for applications to connect to the network
│   │   ├── org1-connection.json
│   │   └── org2-connection.json
│   └── docker-compose                 // Docker Compose files for local development network
│       ├── docker-compose-ca.yaml
│       ├── docker-compose-cli.yaml
│       └── docker-compose-test-net.yaml
│
├── chaincode                          // Smart Contracts (Chaincode)
│   ├── did-registry                   // DID Registry Smart Contract
│   │   ├── lib                        // Chaincode source code (e.g., Go, Node.js, Java)
│   │   │   └── did\_contract.go
│   │   ├── META-INF                   // Chaincode metadata (optional)
│   │   └── package.json               // For Node.js chaincode (if using Node.js chaincode)
│
├── applications                       // Cloud Service Layer Microservices and User Applications (Go-based)
│   ├── api-gateway                    // IAM Service API Gateway (e.g., Go/Gin, Echo, Fiber)
│   │   ├── cmd
│   │   │   └── main.go
│   │   ├── internal
│   │   │   ├── handlers
│   │   │   │   └── handlers\_test.go   // Unit tests for API handlers
│   │   │   └── routes
│   │   ├── go.mod
│   │   ├── go.sum
│   │   └── Dockerfile
│   ├── credential-verifier            // Credential Verifier Service
│   │   ├── cmd
│   │   │   └── main.go
│   │   ├── internal
│   │   │   ├── services
│   │   │   │   └── verifier\_test.go   // Unit tests for verifier logic
│   │   │   └── models                 // Example: for data structures
│   │   ├── go.mod
│   │   ├── go.sum
│   │   └── Dockerfile
│   ├── policy-decision-point          // Policy Decision Point (PDP) Service
│   │   ├── cmd
│   │   │   └── main.go
│   │   ├── internal
│   │   │   ├── policies
│   │   │   │   └── policies\_test.go   // Unit tests for policy evaluation
│   │   │   └── rules                  // Example: for policy rules
│   │   ├── go.mod
│   │   ├── go.sum
│   │   └── Dockerfile
│   ├── credential-issuance            // Credential Issuance Service
│   │   ├── cmd
│   │   │   └── main.go
│   │   ├── internal
│   │   │   ├── services
│   │   │   │   └── issuer\_test.go     // Unit tests for issuer logic
│   │   │   └── data                   // Example: for data access
│   │   ├── go.mod
│   │   ├── go.sum
│   │   └── Dockerfile
│   ├── idp-bridge                     // Identity Provider (IdP) Bridge Service
│   │   ├── cmd
│   │   │   └── main.go
│   │   ├── internal
│   │   │   ├── bridge
│   │   │   │   └── bridge\_test.go     // Unit tests for bridge logic
│   │   │   └── connectors             // Example: for external IdP connectors
│   │   ├── go.mod
│   │   ├── go.sum
│   │   └── Dockerfile
│   ├── digital-wallet-app             // User-facing Digital Wallet Application (e.g., React Native, Web App)
│   │   ├── public
│   │   ├── src
│   │   │   ├── components
│   │   │   ├── pages
│   │   │   └── App.js
│   │   ├── package.json               // Assuming this remains a web/mobile app (e.g., React)
│   │   └── Dockerfile
│   └── admin-ui                       // Web-based Admin UI
│       ├── public
│       ├── src
│       │   ├── components
│       │   ├── pages
│       │   └── App.js
│       ├── package.json               // Assuming this remains a web app (e.g., React)
│       └── Dockerfile
│
├── terraform                          // Terraform configurations for infrastructure provisioning
│   ├── README.md                      // Terraform project overview and setup
│   ├── main.tf                        // Main Terraform configuration file (calls modules)
│   ├── variables.tf                   // Common variables for the entire Terraform project
│   ├── outputs.tf                     // Outputs from the root Terraform module
│   ├── modules                        // Reusable Terraform modules
│   │   ├── network                  // VPC, subnets, security groups
│   │   │   ├── main.tf
│   │   │   ├── variables.tf
│   │   │   └── outputs.tf
│   │   ├── eks\_cluster            // Kubernetes cluster for Go services
│   │   │   ├── main.tf
│   │   │   ├── variables.tf
│   │   │   └── outputs.tf
│   │   ├── amb\_fabric             // Amazon Managed Blockchain for Hyperledger Fabric
│   │   │   ├── main.tf
│   │   │   ├── variables.tf
│   │   │   └── outputs.tf
│   │   ├── api\_gateway            // API Gateway for IAM services
│   │   │   ├── main.tf
│   │   │   ├── variables.tf
│   │   │   └── outputs.tf
│   │   └── rds\_database           // Optional: Managed database for non-blockchain data
│   │       ├── main.tf
│   │       ├── variables.tf
│   │       └── outputs.tf
│   └── environments               // Environment-specific configurations (dev, staging, prod)
│       ├── dev
│       │   ├── main.tf
│       │   └── variables.tf
│       └── prod
│           ├── main.tf
│           └── variables.tf
├── scripts                            // General project scripts (e.g., CI/CD, deployment)
│   ├── deploy-chaincode.sh
│   └── deploy-services.sh
│
├── docs                               // Project documentation
│   ├── architecture.md
│   ├── api-specs.md
│   └── smart-contract-specs.md
│
└── tests                              // Unit and integration tests
├── chaincode-tests                // Tests for Hyperledger Fabric chaincode
├── service-tests                  // Includes Go unit tests for application services
├── e2e-tests                      // End-to-end tests for the entire system
└── README.md                      // Explanation of testing strategy

```