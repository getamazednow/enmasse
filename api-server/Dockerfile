FROM enmasseproject/java-base:8-11

ARG version
ARG maven_version
ARG commit
ENV VERSION=${version} COMMIT=${commit} MAVEN_VERSION=${maven_version}

ADD target/api-server-${MAVEN_VERSION}.jar /api-server.jar

CMD ["/opt/run-java/launch_java.sh", "/api-server.jar"]
