mongo:
	# run mongo in cli
	docker run -d  --name mongodb -e MONGO_INITDB_ROOT_USERNAME=chj -e MONGO_INITDB_ROOT_PASSWORD=hongchenjin -p 27017:27017 -v mongodb_data:/data/db mongo

app:
	docker build --tag datacenter:1.0 .\
	docker run -d --name datacenter -p 8080:8080 -v datacenter_data:/datacenter/store datacenter:1.0

	