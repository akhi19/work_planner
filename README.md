# work_planner

## Steps to run
1. Install docker by following steps mentioned here: https://docs.docker.com/engine/install/
2. cd to the repository and run the following command

```
docker-compose up --build
```

## Steps to import postman collection
1. Download postman app: https://www.postman.com/downloads/
2. Open app and click on Import.
3. Select Link and copy the following url in the box: https://www.getpostman.com/collections/d0fb724681441edf46d8
4. Press Continue
5. You should now see collection name: Work Planner
6. Click Import
7. You can now start hitting API's from the imported collection

## Sample test case file
### ./pkg/worker/internal/services/worker_command_service_test.go

## Constraints
1. Only soft delete support.
2. Worker email should be unique.
3. For flexibilty purpose, have not hard coded shifts but provided an API to add/delete/get shifts

