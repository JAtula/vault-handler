#!/bin/bash
export VAULT_ADDR=http://localhost:8200


cat << EOF | vault policy write kube-admin -
path "secret/*" {
  capabilities = [ "read", "list" ]
}

path "sys/leases/renew" {
  capabilities = ["create"]
}
path "sys/leases/revoke" {
  capabilities = ["update"]
}
EOF

VAULT_SA_NAME=$(kubectl get sa vault-sa -o jsonpath="{.secrets[*]['name']}")
JWT_TOKEN=$(kubectl get secret $VAULT_SA_NAME -o jsonpath="{.data.token}" | base64 --decode; echo)
CA_CERT=$(kubectl get secret $VAULT_SA_NAME -o jsonpath="{.data['ca\.crt']}" | base64 --decode; echo)

vault auth enable kubernetes
vault write auth/kubernetes/config \
    token_reviewer_jwt="$JWT_TOKEN" \
    kubernetes_host="https://192.168.99.100:8443" \
    kubernetes_ca_cert="$CA_CERT"
vault write auth/kubernetes/role/kube-admin \
    bound_service_account_names=vault-sa \
    bound_service_account_namespaces="*" \
    policies=kube-admin \
    ttl=3h