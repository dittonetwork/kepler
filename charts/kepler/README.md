```bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 ignite chain build -o output_dir/

cp go/bin/app_binary output_dir/app_service

docker buildx build --platform linux/amd64 -t my-docker-repo/my-app:latest .

docker tag my-app:latest my-docker-repo/my-app:latest
docker push my-docker-repo/my-app:latest

helm upgrade --install my-app ./my-app-chart --namespace my-namespace
```