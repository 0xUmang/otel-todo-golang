# otel-todo-golang


A basic Todo App in golang which is implementing Open Telemetry and Visualzing via grafana Tempo 


To get Started 


Clone the Repo in your local 

Ensure you have docker running 

Run 
docker compose -f docker-compose.yaml up -d


Thats it. 

Use postman to GET/POST/PUT/DELETE data to localhost:8888 

Example Api 

POST localhost:8888/v1/todo JSON data = {"Title":"Get Food","Description": "Get Breakfast","Status": "Started"}

GET localhost:8888/v1/todo

Grafana is running on localhost:3000 
After opening Grafana, the Explore button on the left pane would have Tempo as a data source listed, Selecting the Service Name as "Todo" from the drop down would show the API call traces 

