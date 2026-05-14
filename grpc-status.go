package erw

import (
	"google.golang.org/grpc/codes"
)

const (
	Text_gRPC_OK                 = "OK"
	Text_gRPC_Canceled           = "CANCELLED"
	Text_gRPC_Unknown            = "UNKNOWN"
	Text_gRPC_InvalidArgument    = "INVALID_ARGUMENT"
	Text_gRPC_DeadlineExceeded   = "DEADLINE_EXCEEDED"
	Text_gRPC_NotFound           = "NOT_FOUND"
	Text_gRPC_AlreadyExists      = "ALREADY_EXISTS"
	Text_gRPC_PermissionDenied   = "PERMISSION_DENIED"
	Text_gRPC_ResourceExhausted  = "RESOURCE_EXHAUSTED"
	Text_gRPC_FailedPrecondition = "FAILED_PRECONDITION"
	Text_gRPC_Aborted            = "ABORTED"
	Text_gRPC_OutOfRange         = "OUT_OF_RANGE"
	Text_gRPC_Unimplemented      = "UNIMPLEMENTED"
	Text_gRPC_Internal           = "INTERNAL"
	Text_gRPC_Unavailable        = "UNAVAILABLE"
	Text_gRPC_DataLoss           = "DATA_LOSS"
	Text_gRPC_Unauthenticated    = "UNAUTHENTICATED"
)

// gRPCStatusText текстовое описание gRPC статуса
func gRPCStatusText(code codes.Code) string {
	switch code {
	case codes.OK:
		return Text_gRPC_OK
	case codes.Canceled:
		return Text_gRPC_Canceled
	case codes.Unknown:
		return Text_gRPC_Unknown
	case codes.InvalidArgument:
		return Text_gRPC_InvalidArgument
	case codes.DeadlineExceeded:
		return Text_gRPC_DeadlineExceeded
	case codes.NotFound:
		return Text_gRPC_NotFound
	case codes.AlreadyExists:
		return Text_gRPC_AlreadyExists
	case codes.PermissionDenied:
		return Text_gRPC_PermissionDenied
	case codes.ResourceExhausted:
		return Text_gRPC_ResourceExhausted
	case codes.FailedPrecondition:
		return Text_gRPC_FailedPrecondition
	case codes.Aborted:
		return Text_gRPC_Aborted
	case codes.OutOfRange:
		return Text_gRPC_OutOfRange
	case codes.Unimplemented:
		return Text_gRPC_Unimplemented
	case codes.Internal:
		return Text_gRPC_Internal
	case codes.Unavailable:
		return Text_gRPC_Unavailable
	case codes.DataLoss:
		return Text_gRPC_DataLoss
	case codes.Unauthenticated:
		return Text_gRPC_Unauthenticated
	default:
		return ""
	}
}
