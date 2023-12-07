api_image:
	docker build -t auth_img .
api_container:
	docker run -d --name auth_api -e PORT=5003 -p 5003:5003 auth_img