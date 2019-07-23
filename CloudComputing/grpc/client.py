import grpc

import common.convertor_pb2_grpc as convertor_pb2_grpc
import common.convertor_pb2 as convertor_pb2

with grpc.insecure_channel('localhost:50051') as channel:
    stub = convertor_pb2_grpc.ConvertorStub(channel)
    response = stub.ConvertLength(convertor_pb2.LengthConvertorRequest(value=10))
    print(response)
