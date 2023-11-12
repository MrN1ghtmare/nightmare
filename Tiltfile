docker_build(
    'nightmare-cli-alpha',
    '.',
    dockerfile='./build/alpha.Dockerfile'
)

docker_compose('./deployments/docker-compose.yml')

dc_resource('nightmare-cli', labels=['aplha'])