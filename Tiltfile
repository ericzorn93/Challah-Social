# Primary Docker Compose Imports
docker_compose(["./docker-compose.yml"])

# Accounts GraphQL
docker_build('challah-social/accounts-graphql', '.',
    dockerfile="./apps/services/accounts-graphql/Dockerfile",
    live_update = [
        sync('./apps/services/accounts-graphql', '/app'),
        run('air --build.cmd "go build -o /bin/accounts-graphql ./apps/services/accounts-graphql/cmd/server/main.go" --build.bin "/bin/accounts-graphql"'),
        restart_container()
    ]
)

# Accounts API
docker_build('challah-social/accounts-api', '.',
    dockerfile="./apps/services/accounts-api/Dockerfile",
    live_update = [
        sync('./apps/services/accounts-api', '/app'),
        run('air --build.cmd "go build -o /bin/accounts-api ./apps/services/accounts-api/cmd/server/main.go" --build.bin "/bin/accounts-api"'),
        restart_container()
    ]
)

# Accounts Worker
docker_build('challah-social/accounts-worker', '.',
    dockerfile="./apps/services/accounts-worker/Dockerfile",
    live_update = [
        sync('./apps/services/accounts-worker', '/app'),
        run('air --build.cmd "go build -o /bin/accounts-worker ./apps/services/accounts-worker/cmd/server/main.go" --build.bin "/bin/accounts-worker"'),
        restart_container()
    ]
)

# Inbound Webhooks API
docker_build('challah-social/inbound-webhooks-api', '.',
    dockerfile="./apps/services/inbound-webhooks-api/Dockerfile",
    live_update = [
        sync('./apps/services/inbound-webhooks-api', '/app'),
        run('air --build.cmd "go build -o /bin/inbound-webhooks-api ./apps/services/inbound-webhooks-api/cmd/server/main.go" --build.bin "/bin/inbound-webhooks-api"'),
        restart_container()
    ]
)

# TODO: Adding Helm Chart Build
# load('ext://helm_resource', 'helm_resource', 'helm_repo')
# yaml = helm(
#     "./k8s/helm/backend-service",
#     name='accounts-api-helm',
#     namespace='default',
#     values=['./apps/services/accounts-api/provisions/k8s/values-dev.yml'],
# )
# k8s_yaml(yaml)
