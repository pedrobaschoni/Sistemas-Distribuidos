# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc
import warnings

import validation_pb2 as validation__pb2

GRPC_GENERATED_VERSION = '1.66.2'
GRPC_VERSION = grpc.__version__
_version_not_supported = False

try:
    from grpc._utilities import first_version_is_lower
    _version_not_supported = first_version_is_lower(GRPC_VERSION, GRPC_GENERATED_VERSION)
except ImportError:
    _version_not_supported = True

if _version_not_supported:
    raise RuntimeError(
        f'The grpc package installed is at version {GRPC_VERSION},'
        + f' but the generated code in validation_pb2_grpc.py depends on'
        + f' grpcio>={GRPC_GENERATED_VERSION}.'
        + f' Please upgrade your grpc module to grpcio>={GRPC_GENERATED_VERSION}'
        + f' or downgrade your generated code using grpcio-tools<={GRPC_VERSION}.'
    )


class ValidationServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.ValidateCPF = channel.unary_unary(
                '/validation.ValidationService/ValidateCPF',
                request_serializer=validation__pb2.ValidationRequest.SerializeToString,
                response_deserializer=validation__pb2.ValidationResponse.FromString,
                _registered_method=True)
        self.ValidateCNPJ = channel.unary_unary(
                '/validation.ValidationService/ValidateCNPJ',
                request_serializer=validation__pb2.ValidationRequest.SerializeToString,
                response_deserializer=validation__pb2.ValidationResponse.FromString,
                _registered_method=True)


class ValidationServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def ValidateCPF(self, request, context):
        """Corrigido para ValidationRequest
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ValidateCNPJ(self, request, context):
        """Movido para dentro do serviço
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_ValidationServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'ValidateCPF': grpc.unary_unary_rpc_method_handler(
                    servicer.ValidateCPF,
                    request_deserializer=validation__pb2.ValidationRequest.FromString,
                    response_serializer=validation__pb2.ValidationResponse.SerializeToString,
            ),
            'ValidateCNPJ': grpc.unary_unary_rpc_method_handler(
                    servicer.ValidateCNPJ,
                    request_deserializer=validation__pb2.ValidationRequest.FromString,
                    response_serializer=validation__pb2.ValidationResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'validation.ValidationService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))
    server.add_registered_method_handlers('validation.ValidationService', rpc_method_handlers)


 # This class is part of an EXPERIMENTAL API.
class ValidationService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def ValidateCPF(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/validation.ValidationService/ValidateCPF',
            validation__pb2.ValidationRequest.SerializeToString,
            validation__pb2.ValidationResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def ValidateCNPJ(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/validation.ValidationService/ValidateCNPJ',
            validation__pb2.ValidationRequest.SerializeToString,
            validation__pb2.ValidationResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)
