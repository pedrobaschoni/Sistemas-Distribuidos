# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: validation.proto
# Protobuf Python Version: 5.27.2
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import runtime_version as _runtime_version
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_runtime_version.ValidateProtobufRuntimeVersion(
    _runtime_version.Domain.PUBLIC,
    5,
    27,
    2,
    '',
    'validation.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x10validation.proto\x12\nvalidation\"\"\n\x11ValidationRequest\x12\r\n\x05value\x18\x01 \x01(\t\"$\n\x12ValidationResponse\x12\x0e\n\x06result\x18\x01 \x01(\t2\xb0\x01\n\x11ValidationService\x12L\n\x0bValidateCPF\x12\x1d.validation.ValidationRequest\x1a\x1e.validation.ValidationResponse\x12M\n\x0cValidateCNPJ\x12\x1d.validation.ValidationRequest\x1a\x1e.validation.ValidationResponseb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'validation_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  DESCRIPTOR._loaded_options = None
  _globals['_VALIDATIONREQUEST']._serialized_start=32
  _globals['_VALIDATIONREQUEST']._serialized_end=66
  _globals['_VALIDATIONRESPONSE']._serialized_start=68
  _globals['_VALIDATIONRESPONSE']._serialized_end=104
  _globals['_VALIDATIONSERVICE']._serialized_start=107
  _globals['_VALIDATIONSERVICE']._serialized_end=283
# @@protoc_insertion_point(module_scope)
