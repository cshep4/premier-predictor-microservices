package com.cshep4.premierpredictor.notification.config;

import com.cshep4.premierpredictor.auth.AuthServiceGrpc;
import com.mongodb.ConnectionString;
import com.mongodb.MongoClientSettings;
import com.mongodb.client.MongoClient;
import com.mongodb.client.MongoClients;
import io.grpc.ManagedChannelBuilder;
import lombok.val;
import lombok.var;
import org.bson.codecs.pojo.PojoCodecProvider;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

import static org.bson.codecs.configuration.CodecRegistries.fromProviders;
import static org.bson.codecs.configuration.CodecRegistries.fromRegistries;
import static org.springframework.util.StringUtils.isEmpty;

@Configuration
public class Config {
    @Value("${MONGO_SCHEME}")
    private String mongoScheme;

    @Value("${MONGO_HOST}")
    private String mongoHost;

    @Value("${MONGO_USERNAME}")
    private String mongoUsername;

    @Value("${MONGO_PASSWORD}")
    private String mongoPassword;

    @Value("${MONGO_PORT}")
    private String mongoPort;

    @Value("${AUTH_ADDR}")
    private String authServiceAddress;

    @Bean
    public MongoClient client() {
        var mongoUri = mongoScheme + "://" + mongoUsername + ":" + mongoPassword + "@" + mongoHost;

        if (!isEmpty(mongoPort)) {
            mongoUri += ":" + mongoPort;
        }

        mongoUri += "/?retryWrites=true";

        val pojoCodecRegistry = fromRegistries(
                MongoClientSettings.getDefaultCodecRegistry(),
                fromProviders(
                        PojoCodecProvider.builder()
                                .automatic(true)
                                .build()
                )
        );

        val settings = MongoClientSettings.builder()
                .codecRegistry(pojoCodecRegistry)
                .applyConnectionString(new ConnectionString(mongoUri))
                .build();

        return MongoClients.create(settings);
    }

    @Bean
    public AuthServiceGrpc.AuthServiceBlockingStub stub() {
        val channel = ManagedChannelBuilder
                .forTarget(authServiceAddress)
                .usePlaintext()
                .build();

        return AuthServiceGrpc.newBlockingStub(channel);
    }
}
