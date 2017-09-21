### webHook for Kubernetes authenticate by keystone token

config for Kubernetes apiserver

    --authorization-mode=RBAC
    --authentication-token-webhook-config-file=/root/webhook.kubeconfig



webhook.kubeconfig

    apiVersion: v1
    clusters:
    - cluster:
        insecure-skip-tls-verify: true
        server: https://10.1.2.73:8443/webhook
      name: webhook
    contexts:
    - context:
        cluster: webhook
        user: webhook
      name: webhook
    current-context: webhook

generate certificate

    openssl genrsa -out cakey.pem 2048
    openssl req -new -x509 -key cakey.pem -out cacert.pem
    openssl genrsa -out nginx.key 2048
    openssl req -new -key nginx.key -out nginx.csr
    echo subjectAltName = IP:10.1.2.73 > extfile.cnf    # 10.1.2.7 listen ip
    openssl x509 -req -in nginx.csr -CA cacert.pem -CAkey cakey.pem -CAcreateserial -out server.crt -extfile extfile.cnf
