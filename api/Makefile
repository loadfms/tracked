.PHONY: build clean deploy

GO_BUILD := env GOARCH=arm64 GOOS=linux go build -ldflags="-s -w"

build:
	export GO111MODULE=on
	export CGO_ENABLED=1

	${GO_BUILD} -o bin/customer/create-customer/bootstrap cmd/customer/create/create-customer.go
	chmod +x bin/customer/create-customer/bootstrap
	cd bin/customer/create-customer && pwd && zip -r create-customer.zip .

	${GO_BUILD} -o bin/customer/login-customer/bootstrap cmd/customer/login/login-customer.go
	chmod +x bin/customer/login-customer/bootstrap
	cd bin/customer/login-customer && pwd && zip -r login-customer.zip .

	${GO_BUILD} -o bin/workspace/create-workspace/bootstrap cmd/workspace/create/create-workspace.go
	chmod +x bin/workspace/create-workspace/bootstrap
	cd bin/workspace/create-workspace && pwd && zip -r create-workspace.zip .

	${GO_BUILD} -o bin/workspace/query-workspace/bootstrap cmd/workspace/query/query-workspace.go
	chmod +x bin/workspace/query-workspace/bootstrap
	cd bin/workspace/query-workspace && pwd && zip -r query-workspace.zip .

	${GO_BUILD} -o bin/authorizer/bootstrap cmd/authorizer/authorizer.go
	chmod +x bin/authorizer/bootstrap
	cd bin/authorizer && pwd && zip -r authorizer.zip .

	${GO_BUILD} -o bin/site/create-site/bootstrap cmd/site/create/create-site.go
	chmod +x bin/site/create-site/bootstrap
	cd bin/site/create-site && pwd && zip -r create-site.zip .

	${GO_BUILD} -o bin/site/query-site/bootstrap cmd/site/query/query-site.go
	chmod +x bin/site/query-site/bootstrap
	cd bin/site/query-site && pwd && zip -r query-site.zip .

	${GO_BUILD} -o bin/privacypolicy/create-privacypolicy/bootstrap cmd/privacypolicy/create/create-privacypolicy.go
	chmod +x bin/privacypolicy/create-privacypolicy/bootstrap
	cd bin/privacypolicy/create-privacypolicy && pwd && zip -r create-privacypolicy.zip .

	${GO_BUILD} -o bin/privacypolicy/query-privacypolicy/bootstrap cmd/privacypolicy/query/query-privacypolicy.go
	chmod +x bin/privacypolicy/query-privacypolicy/bootstrap
	cd bin/privacypolicy/query-privacypolicy && pwd && zip -r query-privacypolicy.zip .

	${GO_BUILD} -o bin/cookie/create-cookie/bootstrap cmd/cookie/create/create-cookie.go
	chmod +x bin/cookie/create-cookie/bootstrap
	cd bin/cookie/create-cookie && pwd && zip -r create-cookie.zip .

	${GO_BUILD} -o bin/cookie/query-cookie/bootstrap cmd/cookie/query/query-cookie.go
	chmod +x bin/cookie/query-cookie/bootstrap
	cd bin/cookie/query-cookie && pwd && zip -r query-cookie.zip .

	${GO_BUILD} -o bin/consent/create-consent/bootstrap cmd/consent/create/create-consent.go
	chmod +x bin/consent/create-consent/bootstrap
	cd bin/consent/create-consent && pwd && zip -r create-consent.zip .

	${GO_BUILD} -o bin/consent/query-consent/bootstrap cmd/consent/query/query-consent.go
	chmod +x bin/consent/query-consent/bootstrap
	cd bin/consent/query-consent && pwd && zip -r query-consent.zip .

clean:
	rm -rf ./bin

deploy: clean build
	cdk bootstrap
	cdk deploy
