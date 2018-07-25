dev:
	~/.google-cloud-sdk/bin/dev_appserver.py \
		--skip_sdk_update_check=false \
		--log_level=debug \
		--port=8080 --admin_port=8000 \
		--storage_path=$(GOPATH)/.data \
		$(GOPATH)/src/go-app/.gae/app.yaml
