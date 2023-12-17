rabbit-up:
	docker run -d --name rhythmify_rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3-management

add-user:
	docker exec rhythmify_rabbitmq rabbitmqctl add_user fady fady

set-admin:
	docker exec rhythmify_rabbitmq rabbitmqctl set_user_tags fady administrator

set-auth:
	docker exec rhythmify_rabbitmq rabbitmqctl delete_user guest

set-vhost:
	docker exec rhythmify_rabbitmq rabbitmqctl add_vhost video_converter

set-permissions:
	docker exec rhythmify_rabbitmq rabbitmqctl set_permissions -p video_converter fady ".*" ".*" ".*"

set-exchange-perms: 
	docker exec rhythmify_rabbitmq rabbitmqctl set_topic_permissions -p customers percy customer_events "^customers.*" "^customers.*"

mongo-up:
	docker run -d -p 27017:27017 --name mongo_video mongo:latest
