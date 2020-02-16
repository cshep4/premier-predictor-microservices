FROM openjdk:8-jdk

RUN mkdir -p /opt/cprof && \
  wget -q -O- https://storage.googleapis.com/cloud-profiler/java/latest/profiler_java_agent.tar.gz \
  | tar xzv -C /opt/cprof

VOLUME /tmp
RUN mkdir /work
COPY . /work
WORKDIR /work
RUN /work/gradlew build -x test
RUN mv /work/build/libs/*.jar /work/app.jar

ENV JAVA_OPTS="-XX:+UseCompressedOops -XX:+UseContainerSupport -XX:+UseConcMarkSweepGC -XX:+CMSParallelRemarkEnabled -XX:+UseCMSInitiatingOccupancyOnly -XX:CMSInitiatingOccupancyFraction=70 -XX:+ScavengeBeforeFullGC -XX:+CMSScavengeBeforeRemark"

ENTRYPOINT ["java", \
            "-agentpath:/opt/cprof/profiler_java_agent.so=-cprof_service=core,-cprof_service_version=1.0.0, -logtostderr,-minloglevel=2", \
            "-Djava.security.egd=file:/dev/./urandom","-jar","/work/app.jar"]