image_name = random_video

run:
	docker run -v $(pwd)/videos:/app/videos -p 84:3000 -d ${image_name}
build:
	docker build -t ${image_name} .