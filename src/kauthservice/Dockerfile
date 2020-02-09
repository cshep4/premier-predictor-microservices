# GraalVM docker image used for AoT compilation
FROM oracle/graalvm-ce:19.1.1 AS build-aot
# Add maven wrapper
ADD mvnw app/
ADD .mvn app/.mvn/
# Add pom
ADD pom.xml app/
# Add sources
ADD src app/src/
# Set working dir
WORKDIR /app
# Build (java side)
RUN ./mvnw package -DskipTests
# Create new image from debian-slim (20Mb)
FROM debian:stable-slim
# Copy generated native executable from build-aot
COPY --from=build-aot /app/target/kauthservice /kauthservice
COPY --from=build-aot /opt/graalvm-ce-19.1.1/jre/lib/amd64/libsunec.so /libsunec.so

COPY --from=build-aot /opt/graalvm-ce-19.1.1/jre/lib/amd64 amd64
COPY --from=build-aot /opt/graalvm-ce-19.1.1/jre/lib/security/cacerts cacerts

#ENV MONGO_URI="mongodb://PremPredUser:oSdr4Q5nzEY4Q3sM@premierpredictor-shard-00-00-ty1lj.mongodb.net:27017,premierpredictor-shard-00-01-ty1lj.mongodb.net:27017,premierpredictor-shard-00-02-ty1lj.mongodb.net:27017/test?ssl=true&replicaSet=PremierPredictor-shard-0&authSource=admin&retryWrites=true&w=majority"
ENV MONGO_URI="mongodb://localhost:27017"
ENV JWT_SECRET="o62?G)Xep<[n:4Y%]a)Y0(!Wa|TM5|cFcCf<2Ve7%LI,ah)B.M`\;@!\\=-nDLXg"
ENV PORT=3000
ENV HEALTH_PORT=3001
ENV EMAIL_URL="https://jbemuyb0o3.execute-api.us-east-1.amazonaws.com/dev/email"

# Set the entrypoint
EXPOSE 3000
ENTRYPOINT [ "./kauthservice", "-Djava.library.path=./amd64", "-Djavax.net.ssl.trustStore=./cacerts" ]
