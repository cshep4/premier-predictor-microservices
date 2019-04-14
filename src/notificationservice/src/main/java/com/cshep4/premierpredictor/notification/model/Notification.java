package com.cshep4.premierpredictor.notification.model;

import lombok.Builder;
import lombok.Data;

@Data
@Builder
public class Notification {
    private String title;
    private String message;

    public static Notification fromGrpc(com.cshep4.premierpredictor.notification.Notification notification) {
        return Notification.builder()
                .title(notification.getTitle())
                .message(notification.getMessage())
                .build();
    }
}
