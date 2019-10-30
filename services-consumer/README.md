- to start the service in diferent environments, use te command bellow:   
> ./gradlew clean build   
> java -Dspring.profiles.active=dev -Dspring.cloud.config.uri=http://localhost:8888 -jar build/libs/licensing-service-0.0.1-SNAPSHOT.jar   

- test endpoints:   
> http://localhost:9000/v1/organizations/e254f8c-c442-4ebe-a82a-e2fc1d1ff78a/licenses   
> http://localhost:9000/v1/organizations/e254f8c-c442-4ebe-a82a-e2fc1d1ff78a/licenses/f3831f8c-c338-4ebe-a82a-e2fc1d1ff78a   


- Refreshing your properties using Spring Cloud configuration server:   
> localhost:9000/refresh - not working    


- to run local, put this in the environment variables:   
> ENCRYPT_KEY=IMSYMMETRIC;spring.profiles.active=dev


- exemplos com [redis](https://docs.spring.io/spring-data/data-redis/docs/current/reference/html/) 