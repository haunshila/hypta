# Run this from your project root directory

mkdir -p network/organizations/peerOrganizations/org1.example.com/{ca,msp,peers,users}
mkdir -p network/organizations/peerOrganizations/org2.example.com
mkdir -p network/organizations/ordererOrganizations/example.com/{ca,msp,orderers}
mkdir -p network/scripts
mkdir -p network/configtx
mkdir -p network/connection-profiles
touch network/connection-profiles/org1-connection.json
touch network/connection-profiles/org2-connection.json
mkdir -p network/docker-compose
touch network/docker-compose/docker-compose-ca.yaml
touch network/docker-compose/docker-compose-cli.yaml
touch network/docker-compose/docker-compose-test-net.yaml

mkdir -p chaincode/did-registry/lib
touch chaincode/did-registry/lib/did_contract.go
mkdir -p chaincode/did-registry/META-INF
touch chaincode/did-registry/package.json

mkdir -p applications/api-gateway/cmd
touch applications/api-gateway/cmd/main.go
mkdir -p applications/api-gateway/internal/handlers
touch applications/api-gateway/internal/handlers/handlers_test.go
mkdir -p applications/api-gateway/internal/routes
touch applications/api-gateway/go.mod
touch applications/api-gateway/go.sum
touch applications/api-gateway/Dockerfile

mkdir -p applications/credential-verifier/cmd
touch applications/credential-verifier/cmd/main.go
mkdir -p applications/credential-verifier/internal/services
touch applications/credential-verifier/internal/services/verifier_test.go
mkdir -p applications/credential-verifier/internal/models
touch applications/credential-verifier/go.mod
touch applications/credential-verifier/go.sum
touch applications/credential-verifier/Dockerfile

mkdir -p applications/policy-decision-point/cmd
touch applications/policy-decision-point/cmd/main.go
mkdir -p applications/policy-decision-point/internal/policies
touch applications/policy-decision-point/internal/policies/policies_test.go
mkdir -p applications/policy-decision-point/internal/rules
touch applications/policy-decision-point/go.mod
touch applications/policy-decision-point/go.sum
touch applications/policy-decision-point/Dockerfile

mkdir -p applications/credential-issuance/cmd
touch applications/credential-issuance/cmd/main.go
mkdir -p applications/credential-issuance/internal/services
touch applications/credential-issuance/internal/services/issuer_test.go
mkdir -p applications/credential-issuance/internal/data
touch applications/credential-issuance/go.mod
touch applications/credential-issuance/go.sum
touch applications/credential-issuance/Dockerfile

mkdir -p applications/idp-bridge/cmd
touch applications/idp-bridge/cmd/main.go
mkdir -p applications/idp-bridge/internal/bridge
touch applications/idp-bridge/internal/bridge/bridge_test.go
mkdir -p applications/idp-bridge/internal/connectors
touch applications/idp-bridge/go.mod
touch applications/idp-bridge/go.sum
touch applications/idp-bridge/Dockerfile

mkdir -p applications/digital-wallet-app/public
mkdir -p applications/digital-wallet-app/src/components
mkdir -p applications/digital-wallet-app/src/pages
touch applications/digital-wallet-app/src/App.js
touch applications/digital-wallet-app/package.json
touch applications/digital-wallet-app/Dockerfile

mkdir -p applications/admin-ui/public
mkdir -p applications/admin-ui/src/components
mkdir -p applications/admin-ui/src/pages
touch applications/admin-ui/src/App.js
touch applications/admin-ui/package.json
touch applications/admin-ui/Dockerfile

mkdir -p terraform/modules/network
touch terraform/modules/network/main.tf
touch terraform/modules/network/variables.tf
touch terraform/modules/network/outputs.tf

mkdir -p terraform/modules/eks_cluster
touch terraform/modules/eks_cluster/main.tf
touch terraform/modules/eks_cluster/variables.tf
touch terraform/modules/eks_cluster/outputs.tf

mkdir -p terraform/modules/amb_fabric
touch terraform/modules/amb_fabric/main.tf
touch terraform/modules/amb_fabric/variables.tf
touch terraform/modules/amb_fabric/outputs.tf

mkdir -p terraform/modules/api_gateway
touch terraform/modules/api_gateway/main.tf
touch terraform/modules/api_gateway/variables.tf
touch terraform/modules/api_gateway/outputs.tf

mkdir -p terraform/modules/database
touch terraform/modules/database/main.tf
touch terraform/modules/database/variables.tf
touch terraform/modules/database/outputs.tf

mkdir -p terraform/environments/dev
touch terraform/environments/dev/main.tf
touch terraform/environments/dev/variables.tf

mkdir -p terraform/environments/prod
touch terraform/environments/prod/main.tf
touch terraform/environments/prod/variables.tf

touch terraform/README.md
touch terraform/main.tf
touch terraform/variables.tf
touch terraform/outputs.tf

mkdir -p scripts
touch scripts/deploy-chaincode.sh
touch scripts/deploy-services.sh

mkdir -p docs
touch docs/architecture.md
touch docs/api-specs.md
touch docs/smart-contract-specs.md

mkdir -p tests/chaincode-tests
mkdir -p tests/service-tests
mkdir -p tests/e2e-tests
touch tests/README.md

touch README.md