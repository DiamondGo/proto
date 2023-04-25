# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: chat.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='chat.proto',
  package='service',
  syntax='proto3',
  serialized_options=b'Z\nservice/go',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\nchat.proto\x12\x07service\"\x80\x01\n\x07Message\x12\n\n\x02id\x18\x01 \x01(\t\x12\x1b\n\x04\x63hat\x18\x02 \x01(\x0b\x32\r.service.Chat\x12\x1d\n\x06sender\x18\x03 \x01(\x0b\x32\r.service.User\x12\x0c\n\x04text\x18\x04 \x01(\t\x12\x0c\n\x04\x64\x61ta\x18\x05 \x01(\x0c\x12\x11\n\treplyToId\x18\x06 \x01(\t\"2\n\x04\x43hat\x12\n\n\x02id\x18\x01 \x01(\t\x12\x1e\n\x07members\x18\x02 \x03(\x0b\x32\r.service.User\"9\n\x04User\x12\n\n\x02id\x18\x01 \x01(\t\x12\x12\n\nfirst_name\x18\x02 \x01(\t\x12\x11\n\tuser_name\x18\x03 \x01(\t2B\n\x08\x43hatting\x12\x36\n\nChatStream\x12\x10.service.Message\x1a\x10.service.Message\"\x00(\x01\x30\x01\x42\x0cZ\nservice/gob\x06proto3'
)




_MESSAGE = _descriptor.Descriptor(
  name='Message',
  full_name='service.Message',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='service.Message.id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='chat', full_name='service.Message.chat', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='sender', full_name='service.Message.sender', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='text', full_name='service.Message.text', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='data', full_name='service.Message.data', index=4,
      number=5, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=b"",
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='replyToId', full_name='service.Message.replyToId', index=5,
      number=6, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=24,
  serialized_end=152,
)


_CHAT = _descriptor.Descriptor(
  name='Chat',
  full_name='service.Chat',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='service.Chat.id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='members', full_name='service.Chat.members', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=154,
  serialized_end=204,
)


_USER = _descriptor.Descriptor(
  name='User',
  full_name='service.User',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='service.User.id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='first_name', full_name='service.User.first_name', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='user_name', full_name='service.User.user_name', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=206,
  serialized_end=263,
)

_MESSAGE.fields_by_name['chat'].message_type = _CHAT
_MESSAGE.fields_by_name['sender'].message_type = _USER
_CHAT.fields_by_name['members'].message_type = _USER
DESCRIPTOR.message_types_by_name['Message'] = _MESSAGE
DESCRIPTOR.message_types_by_name['Chat'] = _CHAT
DESCRIPTOR.message_types_by_name['User'] = _USER
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Message = _reflection.GeneratedProtocolMessageType('Message', (_message.Message,), {
  'DESCRIPTOR' : _MESSAGE,
  '__module__' : 'chat_pb2'
  # @@protoc_insertion_point(class_scope:service.Message)
  })
_sym_db.RegisterMessage(Message)

Chat = _reflection.GeneratedProtocolMessageType('Chat', (_message.Message,), {
  'DESCRIPTOR' : _CHAT,
  '__module__' : 'chat_pb2'
  # @@protoc_insertion_point(class_scope:service.Chat)
  })
_sym_db.RegisterMessage(Chat)

User = _reflection.GeneratedProtocolMessageType('User', (_message.Message,), {
  'DESCRIPTOR' : _USER,
  '__module__' : 'chat_pb2'
  # @@protoc_insertion_point(class_scope:service.User)
  })
_sym_db.RegisterMessage(User)


DESCRIPTOR._options = None

_CHATTING = _descriptor.ServiceDescriptor(
  name='Chatting',
  full_name='service.Chatting',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_start=265,
  serialized_end=331,
  methods=[
  _descriptor.MethodDescriptor(
    name='ChatStream',
    full_name='service.Chatting.ChatStream',
    index=0,
    containing_service=None,
    input_type=_MESSAGE,
    output_type=_MESSAGE,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
])
_sym_db.RegisterServiceDescriptor(_CHATTING)

DESCRIPTOR.services_by_name['Chatting'] = _CHATTING

# @@protoc_insertion_point(module_scope)
