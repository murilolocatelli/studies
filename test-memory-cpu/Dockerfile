FROM openjdk:alpine
COPY springboot_swagger_example/target/spring-boot-web-0.0.1-SNAPSHOT.jar /
ENTRYPOINT ["java", "-Xms1200m", "-Xmx1200m", "-XX:MaxRAM=5000m", "-jar", "/spring-boot-web-0.0.1-SNAPSHOT.jar"]