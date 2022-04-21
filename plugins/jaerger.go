package plugins

import (
	"context"
	"github.com/smallnest/rpcx/protocol"
)

type Jaeger struct {
}

func (j Jaeger) PreHandleRequest(ctx context.Context, r *protocol.Message) error {
	//wireContext, err := share.GetSpanContextFromContext(ctx)
	//
	//if err != nil || wireContext == nil {
	//	return err
	//}
	//span := opentracing.StartSpan(
	//	"执行:"+r.ServicePath+"."+r.ServiceMethod,
	//	ext.RPCServerOption(wireContext))
	//
	//clientConn := ctx.Value(server.RemoteConnContextKey).(net.Conn)
	//span.LogFields(log.String("remote_addr", clientConn.RemoteAddr().Network()))
	//
	//span.SetTag("开始执行时间", time.Now())
	//codec := share.Codecs[r.SerializeType()]
	//if codec != nil {
	//	var argv interface{}
	//	err = codec.Decode(r.Payload, &argv)
	//	if err == nil {
	//		span.LogKV("参数", argv)
	//	}
	//}
	//
	//if ctx, ok := ctx.(*share.Context); ok {
	//	ctx.SetValue(share.OpentracingSpanServerKey, span)
	//
	//	tracer := opentracing.GlobalTracer()
	//	metadata := opentracing.TextMapCarrier(make(map[string]string))
	//	err = tracer.Inject(span.Context(), opentracing.TextMap, metadata)
	//	if err == nil {
	//		ctx.SetValue(share.ReqMetaDataKey, (map[string]string)(metadata))
	//	}
	//}
	return nil

}

func (j Jaeger) PostWriteResponse(ctx context.Context, req *protocol.Message, res *protocol.Message, err error) error {

	return nil
}
