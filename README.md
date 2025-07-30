# HYPTA - Decentralized Identity and Access Management (IAM)

## Overview

HYPTA is a cloud-native, enterprise-grade IAM solution leveraging Hyperledger Fabric, Go microservices, and Terraform for infrastructure provisioning. It implements decentralized identity using DIDs and Verifiable Credentials, enhancing privacy, security, and auditability.

## Features

- Decentralized User Management (DID-based)
- OAuth 2.0 and OIDC support
- Multi-Factor Authentication (MFA)
- Credential Issuance and Verification
- Policy Decision Point (PDP) for fine-grained access control
- Admin UI and Digital Wallet App
- Integration with external Identity Providers (IdP Bridge)
- Audit Logging

## Project Structure

See [`docs/structure.md`](docs/structure.md) for a detailed directory and component overview.

## Getting Started

1. **Clone the repository**
   ```bash
   git clone https://github.com/your-org/hypta.git
   cd hypta
   ```

2. **Provision infrastructure**
   ```bash
   cd terraform
   terraform init
   terraform apply
   ```

3. **Start Fabric network**
   ```bash
   cd network/scripts
   ./start.sh
   ```

4. **Deploy chaincode and services**
   ```bash
   ./scripts/deploy-chaincode.sh
   ./scripts/deploy-services.sh
   ```

## Documentation

- [Architecture](docs/architecture.md)
- [API Specs](docs/api-specs.md)
- [Smart Contract Specs](docs/smart-contract-specs.md)
