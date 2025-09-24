# Create a service (e.g., albums service running at http://albums:8080)
curl -sS -X POST http://localhost:8001/services \
  -d name=albums \
  -d url=http://albums:8080

# Create a route to expose it on /albums
curl -sS -X POST http://localhost:8001/services/albums/routes \
  -d name=albums-api \
  -d paths=/albums \
  -d strip_path=true


curl -sS -X POST http://localhost:8001/services/albums/plugins \
  -d name=cors \
  -d config.origins='*' \
  -d config.methods=GET \
  -d config.methods=POST \
  -d config.methods=DELETE \
  -d config.methods=PUT \
  -d config.methods=PATCH \
  -d config.methods=OPTIONS \
  -d config.methods=HEAD \
  -d config.headers=Authorization \
  -d config.headers=Content-Type \
  -d config.credentials=true


# protect the service with key-auth
curl -sS -X POST http://localhost:8001/services/albums/plugins \
  -d name=key-auth \
  -d config.key_names=apikey

# # create a consumer
# curl -sS -X POST http://localhost:8001/consumers -d username=demo

# # create a key for the consumer
# curl -sS -X POST http://localhost:8001/consumers/demo/key-auth \
#   -d key=DEMO-KEY-123456


curl -sS -X POST http://localhost:8001/services \
  -d name=user-service \
  -d url=http://user-service:8080

# Create a route to expose it on /user-service
curl -sS -X POST http://localhost:8001/services/user-service/routes \
  -d name=user-service-api \
  -d paths=/user-service \
  -d strip_path=true


curl -sS -X POST http://localhost:8001/services/user-service/plugins \
  -d name=cors \
  -d config.origins='*' \
  -d config.methods=GET \
  -d config.methods=POST \
  -d config.methods=DELETE \
  -d config.methods=PUT \
  -d config.methods=PATCH \
  -d config.methods=OPTIONS \
  -d config.methods=HEAD \
  -d config.headers=Authorization \
  -d config.headers=Content-Type \
  -d config.credentials=true


# protect the service with key-auth
curl -sS -X POST http://localhost:8001/services/user-service/plugins \
  -d name=key-auth \
  -d config.key_names=apikey