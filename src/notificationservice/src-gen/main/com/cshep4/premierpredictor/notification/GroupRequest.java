// Generated by the protocol buffer compiler.  DO NOT EDIT!
// source: notification.proto

package com.cshep4.premierpredictor.notification;

/**
 * Protobuf type {@code model.GroupRequest}
 */
public  final class GroupRequest extends
    com.google.protobuf.GeneratedMessageV3 implements
    // @@protoc_insertion_point(message_implements:model.GroupRequest)
    GroupRequestOrBuilder {
  // Use GroupRequest.newBuilder() to construct.
  private GroupRequest(com.google.protobuf.GeneratedMessageV3.Builder<?> builder) {
    super(builder);
  }
  private GroupRequest() {
    userIds_ = com.google.protobuf.LazyStringArrayList.EMPTY;
  }

  @java.lang.Override
  public final com.google.protobuf.UnknownFieldSet
  getUnknownFields() {
    return com.google.protobuf.UnknownFieldSet.getDefaultInstance();
  }
  private GroupRequest(
      com.google.protobuf.CodedInputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    this();
    int mutable_bitField0_ = 0;
    try {
      boolean done = false;
      while (!done) {
        int tag = input.readTag();
        switch (tag) {
          case 0:
            done = true;
            break;
          default: {
            if (!input.skipField(tag)) {
              done = true;
            }
            break;
          }
          case 10: {
            java.lang.String s = input.readStringRequireUtf8();
            if (!((mutable_bitField0_ & 0x00000001) == 0x00000001)) {
              userIds_ = new com.google.protobuf.LazyStringArrayList();
              mutable_bitField0_ |= 0x00000001;
            }
            userIds_.add(s);
            break;
          }
          case 18: {
            com.cshep4.premierpredictor.notification.Notification.Builder subBuilder = null;
            if (notification_ != null) {
              subBuilder = notification_.toBuilder();
            }
            notification_ = input.readMessage(com.cshep4.premierpredictor.notification.Notification.parser(), extensionRegistry);
            if (subBuilder != null) {
              subBuilder.mergeFrom(notification_);
              notification_ = subBuilder.buildPartial();
            }

            break;
          }
        }
      }
    } catch (com.google.protobuf.InvalidProtocolBufferException e) {
      throw e.setUnfinishedMessage(this);
    } catch (java.io.IOException e) {
      throw new com.google.protobuf.InvalidProtocolBufferException(
          e).setUnfinishedMessage(this);
    } finally {
      if (((mutable_bitField0_ & 0x00000001) == 0x00000001)) {
        userIds_ = userIds_.getUnmodifiableView();
      }
      makeExtensionsImmutable();
    }
  }
  public static final com.google.protobuf.Descriptors.Descriptor
      getDescriptor() {
    return com.cshep4.premierpredictor.notification.NotificationOuterClass.internal_static_model_GroupRequest_descriptor;
  }

  protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
      internalGetFieldAccessorTable() {
    return com.cshep4.premierpredictor.notification.NotificationOuterClass.internal_static_model_GroupRequest_fieldAccessorTable
        .ensureFieldAccessorsInitialized(
            com.cshep4.premierpredictor.notification.GroupRequest.class, com.cshep4.premierpredictor.notification.GroupRequest.Builder.class);
  }

  private int bitField0_;
  public static final int USERIDS_FIELD_NUMBER = 1;
  private com.google.protobuf.LazyStringList userIds_;
  /**
   * <code>repeated string userIds = 1;</code>
   */
  public com.google.protobuf.ProtocolStringList
      getUserIdsList() {
    return userIds_;
  }
  /**
   * <code>repeated string userIds = 1;</code>
   */
  public int getUserIdsCount() {
    return userIds_.size();
  }
  /**
   * <code>repeated string userIds = 1;</code>
   */
  public java.lang.String getUserIds(int index) {
    return userIds_.get(index);
  }
  /**
   * <code>repeated string userIds = 1;</code>
   */
  public com.google.protobuf.ByteString
      getUserIdsBytes(int index) {
    return userIds_.getByteString(index);
  }

  public static final int NOTIFICATION_FIELD_NUMBER = 2;
  private com.cshep4.premierpredictor.notification.Notification notification_;
  /**
   * <code>.model.Notification notification = 2;</code>
   */
  public boolean hasNotification() {
    return notification_ != null;
  }
  /**
   * <code>.model.Notification notification = 2;</code>
   */
  public com.cshep4.premierpredictor.notification.Notification getNotification() {
    return notification_ == null ? com.cshep4.premierpredictor.notification.Notification.getDefaultInstance() : notification_;
  }
  /**
   * <code>.model.Notification notification = 2;</code>
   */
  public com.cshep4.premierpredictor.notification.NotificationOrBuilder getNotificationOrBuilder() {
    return getNotification();
  }

  private byte memoizedIsInitialized = -1;
  public final boolean isInitialized() {
    byte isInitialized = memoizedIsInitialized;
    if (isInitialized == 1) return true;
    if (isInitialized == 0) return false;

    memoizedIsInitialized = 1;
    return true;
  }

  public void writeTo(com.google.protobuf.CodedOutputStream output)
                      throws java.io.IOException {
    for (int i = 0; i < userIds_.size(); i++) {
      com.google.protobuf.GeneratedMessageV3.writeString(output, 1, userIds_.getRaw(i));
    }
    if (notification_ != null) {
      output.writeMessage(2, getNotification());
    }
  }

  public int getSerializedSize() {
    int size = memoizedSize;
    if (size != -1) return size;

    size = 0;
    {
      int dataSize = 0;
      for (int i = 0; i < userIds_.size(); i++) {
        dataSize += computeStringSizeNoTag(userIds_.getRaw(i));
      }
      size += dataSize;
      size += 1 * getUserIdsList().size();
    }
    if (notification_ != null) {
      size += com.google.protobuf.CodedOutputStream
        .computeMessageSize(2, getNotification());
    }
    memoizedSize = size;
    return size;
  }

  private static final long serialVersionUID = 0L;
  @java.lang.Override
  public boolean equals(final java.lang.Object obj) {
    if (obj == this) {
     return true;
    }
    if (!(obj instanceof com.cshep4.premierpredictor.notification.GroupRequest)) {
      return super.equals(obj);
    }
    com.cshep4.premierpredictor.notification.GroupRequest other = (com.cshep4.premierpredictor.notification.GroupRequest) obj;

    boolean result = true;
    result = result && getUserIdsList()
        .equals(other.getUserIdsList());
    result = result && (hasNotification() == other.hasNotification());
    if (hasNotification()) {
      result = result && getNotification()
          .equals(other.getNotification());
    }
    return result;
  }

  @java.lang.Override
  public int hashCode() {
    if (memoizedHashCode != 0) {
      return memoizedHashCode;
    }
    int hash = 41;
    hash = (19 * hash) + getDescriptor().hashCode();
    if (getUserIdsCount() > 0) {
      hash = (37 * hash) + USERIDS_FIELD_NUMBER;
      hash = (53 * hash) + getUserIdsList().hashCode();
    }
    if (hasNotification()) {
      hash = (37 * hash) + NOTIFICATION_FIELD_NUMBER;
      hash = (53 * hash) + getNotification().hashCode();
    }
    hash = (29 * hash) + unknownFields.hashCode();
    memoizedHashCode = hash;
    return hash;
  }

  public static com.cshep4.premierpredictor.notification.GroupRequest parseFrom(
      java.nio.ByteBuffer data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static com.cshep4.premierpredictor.notification.GroupRequest parseFrom(
      java.nio.ByteBuffer data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static com.cshep4.premierpredictor.notification.GroupRequest parseFrom(
      com.google.protobuf.ByteString data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static com.cshep4.premierpredictor.notification.GroupRequest parseFrom(
      com.google.protobuf.ByteString data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static com.cshep4.premierpredictor.notification.GroupRequest parseFrom(byte[] data)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data);
  }
  public static com.cshep4.premierpredictor.notification.GroupRequest parseFrom(
      byte[] data,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws com.google.protobuf.InvalidProtocolBufferException {
    return PARSER.parseFrom(data, extensionRegistry);
  }
  public static com.cshep4.premierpredictor.notification.GroupRequest parseFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static com.cshep4.premierpredictor.notification.GroupRequest parseFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }
  public static com.cshep4.premierpredictor.notification.GroupRequest parseDelimitedFrom(java.io.InputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input);
  }
  public static com.cshep4.premierpredictor.notification.GroupRequest parseDelimitedFrom(
      java.io.InputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
  }
  public static com.cshep4.premierpredictor.notification.GroupRequest parseFrom(
      com.google.protobuf.CodedInputStream input)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input);
  }
  public static com.cshep4.premierpredictor.notification.GroupRequest parseFrom(
      com.google.protobuf.CodedInputStream input,
      com.google.protobuf.ExtensionRegistryLite extensionRegistry)
      throws java.io.IOException {
    return com.google.protobuf.GeneratedMessageV3
        .parseWithIOException(PARSER, input, extensionRegistry);
  }

  public Builder newBuilderForType() { return newBuilder(); }
  public static Builder newBuilder() {
    return DEFAULT_INSTANCE.toBuilder();
  }
  public static Builder newBuilder(com.cshep4.premierpredictor.notification.GroupRequest prototype) {
    return DEFAULT_INSTANCE.toBuilder().mergeFrom(prototype);
  }
  public Builder toBuilder() {
    return this == DEFAULT_INSTANCE
        ? new Builder() : new Builder().mergeFrom(this);
  }

  @java.lang.Override
  protected Builder newBuilderForType(
      com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
    Builder builder = new Builder(parent);
    return builder;
  }
  /**
   * Protobuf type {@code model.GroupRequest}
   */
  public static final class Builder extends
      com.google.protobuf.GeneratedMessageV3.Builder<Builder> implements
      // @@protoc_insertion_point(builder_implements:model.GroupRequest)
      com.cshep4.premierpredictor.notification.GroupRequestOrBuilder {
    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return com.cshep4.premierpredictor.notification.NotificationOuterClass.internal_static_model_GroupRequest_descriptor;
    }

    protected com.google.protobuf.GeneratedMessageV3.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return com.cshep4.premierpredictor.notification.NotificationOuterClass.internal_static_model_GroupRequest_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              com.cshep4.premierpredictor.notification.GroupRequest.class, com.cshep4.premierpredictor.notification.GroupRequest.Builder.class);
    }

    // Construct using com.cshep4.premierpredictor.notification.GroupRequest.newBuilder()
    private Builder() {
      maybeForceBuilderInitialization();
    }

    private Builder(
        com.google.protobuf.GeneratedMessageV3.BuilderParent parent) {
      super(parent);
      maybeForceBuilderInitialization();
    }
    private void maybeForceBuilderInitialization() {
      if (com.google.protobuf.GeneratedMessageV3
              .alwaysUseFieldBuilders) {
      }
    }
    public Builder clear() {
      super.clear();
      userIds_ = com.google.protobuf.LazyStringArrayList.EMPTY;
      bitField0_ = (bitField0_ & ~0x00000001);
      if (notificationBuilder_ == null) {
        notification_ = null;
      } else {
        notification_ = null;
        notificationBuilder_ = null;
      }
      return this;
    }

    public com.google.protobuf.Descriptors.Descriptor
        getDescriptorForType() {
      return com.cshep4.premierpredictor.notification.NotificationOuterClass.internal_static_model_GroupRequest_descriptor;
    }

    public com.cshep4.premierpredictor.notification.GroupRequest getDefaultInstanceForType() {
      return com.cshep4.premierpredictor.notification.GroupRequest.getDefaultInstance();
    }

    public com.cshep4.premierpredictor.notification.GroupRequest build() {
      com.cshep4.premierpredictor.notification.GroupRequest result = buildPartial();
      if (!result.isInitialized()) {
        throw newUninitializedMessageException(result);
      }
      return result;
    }

    public com.cshep4.premierpredictor.notification.GroupRequest buildPartial() {
      com.cshep4.premierpredictor.notification.GroupRequest result = new com.cshep4.premierpredictor.notification.GroupRequest(this);
      int from_bitField0_ = bitField0_;
      int to_bitField0_ = 0;
      if (((bitField0_ & 0x00000001) == 0x00000001)) {
        userIds_ = userIds_.getUnmodifiableView();
        bitField0_ = (bitField0_ & ~0x00000001);
      }
      result.userIds_ = userIds_;
      if (notificationBuilder_ == null) {
        result.notification_ = notification_;
      } else {
        result.notification_ = notificationBuilder_.build();
      }
      result.bitField0_ = to_bitField0_;
      onBuilt();
      return result;
    }

    public Builder clone() {
      return (Builder) super.clone();
    }
    public Builder setField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        Object value) {
      return (Builder) super.setField(field, value);
    }
    public Builder clearField(
        com.google.protobuf.Descriptors.FieldDescriptor field) {
      return (Builder) super.clearField(field);
    }
    public Builder clearOneof(
        com.google.protobuf.Descriptors.OneofDescriptor oneof) {
      return (Builder) super.clearOneof(oneof);
    }
    public Builder setRepeatedField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        int index, Object value) {
      return (Builder) super.setRepeatedField(field, index, value);
    }
    public Builder addRepeatedField(
        com.google.protobuf.Descriptors.FieldDescriptor field,
        Object value) {
      return (Builder) super.addRepeatedField(field, value);
    }
    public Builder mergeFrom(com.google.protobuf.Message other) {
      if (other instanceof com.cshep4.premierpredictor.notification.GroupRequest) {
        return mergeFrom((com.cshep4.premierpredictor.notification.GroupRequest)other);
      } else {
        super.mergeFrom(other);
        return this;
      }
    }

    public Builder mergeFrom(com.cshep4.premierpredictor.notification.GroupRequest other) {
      if (other == com.cshep4.premierpredictor.notification.GroupRequest.getDefaultInstance()) return this;
      if (!other.userIds_.isEmpty()) {
        if (userIds_.isEmpty()) {
          userIds_ = other.userIds_;
          bitField0_ = (bitField0_ & ~0x00000001);
        } else {
          ensureUserIdsIsMutable();
          userIds_.addAll(other.userIds_);
        }
        onChanged();
      }
      if (other.hasNotification()) {
        mergeNotification(other.getNotification());
      }
      onChanged();
      return this;
    }

    public final boolean isInitialized() {
      return true;
    }

    public Builder mergeFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      com.cshep4.premierpredictor.notification.GroupRequest parsedMessage = null;
      try {
        parsedMessage = PARSER.parsePartialFrom(input, extensionRegistry);
      } catch (com.google.protobuf.InvalidProtocolBufferException e) {
        parsedMessage = (com.cshep4.premierpredictor.notification.GroupRequest) e.getUnfinishedMessage();
        throw e.unwrapIOException();
      } finally {
        if (parsedMessage != null) {
          mergeFrom(parsedMessage);
        }
      }
      return this;
    }
    private int bitField0_;

    private com.google.protobuf.LazyStringList userIds_ = com.google.protobuf.LazyStringArrayList.EMPTY;
    private void ensureUserIdsIsMutable() {
      if (!((bitField0_ & 0x00000001) == 0x00000001)) {
        userIds_ = new com.google.protobuf.LazyStringArrayList(userIds_);
        bitField0_ |= 0x00000001;
       }
    }
    /**
     * <code>repeated string userIds = 1;</code>
     */
    public com.google.protobuf.ProtocolStringList
        getUserIdsList() {
      return userIds_.getUnmodifiableView();
    }
    /**
     * <code>repeated string userIds = 1;</code>
     */
    public int getUserIdsCount() {
      return userIds_.size();
    }
    /**
     * <code>repeated string userIds = 1;</code>
     */
    public java.lang.String getUserIds(int index) {
      return userIds_.get(index);
    }
    /**
     * <code>repeated string userIds = 1;</code>
     */
    public com.google.protobuf.ByteString
        getUserIdsBytes(int index) {
      return userIds_.getByteString(index);
    }
    /**
     * <code>repeated string userIds = 1;</code>
     */
    public Builder setUserIds(
        int index, java.lang.String value) {
      if (value == null) {
    throw new NullPointerException();
  }
  ensureUserIdsIsMutable();
      userIds_.set(index, value);
      onChanged();
      return this;
    }
    /**
     * <code>repeated string userIds = 1;</code>
     */
    public Builder addUserIds(
        java.lang.String value) {
      if (value == null) {
    throw new NullPointerException();
  }
  ensureUserIdsIsMutable();
      userIds_.add(value);
      onChanged();
      return this;
    }
    /**
     * <code>repeated string userIds = 1;</code>
     */
    public Builder addAllUserIds(
        java.lang.Iterable<java.lang.String> values) {
      ensureUserIdsIsMutable();
      com.google.protobuf.AbstractMessageLite.Builder.addAll(
          values, userIds_);
      onChanged();
      return this;
    }
    /**
     * <code>repeated string userIds = 1;</code>
     */
    public Builder clearUserIds() {
      userIds_ = com.google.protobuf.LazyStringArrayList.EMPTY;
      bitField0_ = (bitField0_ & ~0x00000001);
      onChanged();
      return this;
    }
    /**
     * <code>repeated string userIds = 1;</code>
     */
    public Builder addUserIdsBytes(
        com.google.protobuf.ByteString value) {
      if (value == null) {
    throw new NullPointerException();
  }
  checkByteStringIsUtf8(value);
      ensureUserIdsIsMutable();
      userIds_.add(value);
      onChanged();
      return this;
    }

    private com.cshep4.premierpredictor.notification.Notification notification_ = null;
    private com.google.protobuf.SingleFieldBuilderV3<
        com.cshep4.premierpredictor.notification.Notification, com.cshep4.premierpredictor.notification.Notification.Builder, com.cshep4.premierpredictor.notification.NotificationOrBuilder> notificationBuilder_;
    /**
     * <code>.model.Notification notification = 2;</code>
     */
    public boolean hasNotification() {
      return notificationBuilder_ != null || notification_ != null;
    }
    /**
     * <code>.model.Notification notification = 2;</code>
     */
    public com.cshep4.premierpredictor.notification.Notification getNotification() {
      if (notificationBuilder_ == null) {
        return notification_ == null ? com.cshep4.premierpredictor.notification.Notification.getDefaultInstance() : notification_;
      } else {
        return notificationBuilder_.getMessage();
      }
    }
    /**
     * <code>.model.Notification notification = 2;</code>
     */
    public Builder setNotification(com.cshep4.premierpredictor.notification.Notification value) {
      if (notificationBuilder_ == null) {
        if (value == null) {
          throw new NullPointerException();
        }
        notification_ = value;
        onChanged();
      } else {
        notificationBuilder_.setMessage(value);
      }

      return this;
    }
    /**
     * <code>.model.Notification notification = 2;</code>
     */
    public Builder setNotification(
        com.cshep4.premierpredictor.notification.Notification.Builder builderForValue) {
      if (notificationBuilder_ == null) {
        notification_ = builderForValue.build();
        onChanged();
      } else {
        notificationBuilder_.setMessage(builderForValue.build());
      }

      return this;
    }
    /**
     * <code>.model.Notification notification = 2;</code>
     */
    public Builder mergeNotification(com.cshep4.premierpredictor.notification.Notification value) {
      if (notificationBuilder_ == null) {
        if (notification_ != null) {
          notification_ =
            com.cshep4.premierpredictor.notification.Notification.newBuilder(notification_).mergeFrom(value).buildPartial();
        } else {
          notification_ = value;
        }
        onChanged();
      } else {
        notificationBuilder_.mergeFrom(value);
      }

      return this;
    }
    /**
     * <code>.model.Notification notification = 2;</code>
     */
    public Builder clearNotification() {
      if (notificationBuilder_ == null) {
        notification_ = null;
        onChanged();
      } else {
        notification_ = null;
        notificationBuilder_ = null;
      }

      return this;
    }
    /**
     * <code>.model.Notification notification = 2;</code>
     */
    public com.cshep4.premierpredictor.notification.Notification.Builder getNotificationBuilder() {
      
      onChanged();
      return getNotificationFieldBuilder().getBuilder();
    }
    /**
     * <code>.model.Notification notification = 2;</code>
     */
    public com.cshep4.premierpredictor.notification.NotificationOrBuilder getNotificationOrBuilder() {
      if (notificationBuilder_ != null) {
        return notificationBuilder_.getMessageOrBuilder();
      } else {
        return notification_ == null ?
            com.cshep4.premierpredictor.notification.Notification.getDefaultInstance() : notification_;
      }
    }
    /**
     * <code>.model.Notification notification = 2;</code>
     */
    private com.google.protobuf.SingleFieldBuilderV3<
        com.cshep4.premierpredictor.notification.Notification, com.cshep4.premierpredictor.notification.Notification.Builder, com.cshep4.premierpredictor.notification.NotificationOrBuilder> 
        getNotificationFieldBuilder() {
      if (notificationBuilder_ == null) {
        notificationBuilder_ = new com.google.protobuf.SingleFieldBuilderV3<
            com.cshep4.premierpredictor.notification.Notification, com.cshep4.premierpredictor.notification.Notification.Builder, com.cshep4.premierpredictor.notification.NotificationOrBuilder>(
                getNotification(),
                getParentForChildren(),
                isClean());
        notification_ = null;
      }
      return notificationBuilder_;
    }
    public final Builder setUnknownFields(
        final com.google.protobuf.UnknownFieldSet unknownFields) {
      return this;
    }

    public final Builder mergeUnknownFields(
        final com.google.protobuf.UnknownFieldSet unknownFields) {
      return this;
    }


    // @@protoc_insertion_point(builder_scope:model.GroupRequest)
  }

  // @@protoc_insertion_point(class_scope:model.GroupRequest)
  private static final com.cshep4.premierpredictor.notification.GroupRequest DEFAULT_INSTANCE;
  static {
    DEFAULT_INSTANCE = new com.cshep4.premierpredictor.notification.GroupRequest();
  }

  public static com.cshep4.premierpredictor.notification.GroupRequest getDefaultInstance() {
    return DEFAULT_INSTANCE;
  }

  private static final com.google.protobuf.Parser<GroupRequest>
      PARSER = new com.google.protobuf.AbstractParser<GroupRequest>() {
    public GroupRequest parsePartialFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
        return new GroupRequest(input, extensionRegistry);
    }
  };

  public static com.google.protobuf.Parser<GroupRequest> parser() {
    return PARSER;
  }

  @java.lang.Override
  public com.google.protobuf.Parser<GroupRequest> getParserForType() {
    return PARSER;
  }

  public com.cshep4.premierpredictor.notification.GroupRequest getDefaultInstanceForType() {
    return DEFAULT_INSTANCE;
  }

}

