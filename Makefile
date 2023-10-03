version=1.0.0

generate:
	$(MAKE) -C ./modules/go-client-generator/ generate-go-client service=product_configurator version=${version}
