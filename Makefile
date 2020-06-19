IMAGE := udamliyanage/rest-api-mongo

test:
	true

image:
	docker build -t $(IMAGE) .

push-image:
	docker push $(IMAGE)


.PHONY: image push-image test