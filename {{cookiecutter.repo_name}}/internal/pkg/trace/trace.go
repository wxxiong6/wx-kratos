package trace

import (
	"context"
	"fmt"
	mysqlparse "github.com/go-sql-driver/mysql"
	"github.com/samber/lo"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"runtime"
	"strings"
)

const (
	gormSpanKey        = "gorm"
	callBackBeforeName = "tracing:before"
	callBackAfterName  = "tracing:after"
)

func before(db *gorm.DB) {
	db.Statement.Context = startTrace(
		db.Statement.Context,
		db,
	)
}

func after(scope *gorm.DB) {
	var fieldStrings []string
	if scope.Statement != nil {
		fieldStrings = lo.Map(scope.Statement.Vars, func(v any, i int) string {
			return fmt.Sprintf("($%v = %v)", i+1, v)
		})
	}
	// then add the vars to the span metadata
	span := trace.SpanFromContext(scope.Statement.Context)
	if span != nil && span.IsRecording() {
		span.SetAttributes(
			attribute.String("gorm.query.vars", strings.Join(fieldStrings, ", ")),
		)
	}
	endTrace(scope)
}

func startTrace(
	ctx context.Context,
	db *gorm.DB,
) context.Context {
	// Don't trace queries if they don't have a parent span.
	if span := trace.SpanFromContext(ctx); span == nil {
		return ctx
	}

	ctx, span := otel.Tracer(gormSpanKey).Start(ctx, ""+db.Statement.Table, trace.WithSpanKind(trace.SpanKindInternal))
	var (
		file string
		line int
	)
	for n := 5; n < 20; n++ {
		_, file, line, _ = runtime.Caller(n)
		if strings.Contains(file, "/gorm.io/") {
			// skip any helper code and go further up the call stack
			continue
		}
		break
	}

	db.InstanceSet(gormSpanKey, span)

	span.SetAttributes(attribute.String("caller", fmt.Sprintf("%s:%v", file, line)))
	span.SetAttributes(attribute.String("gorm.table", db.Statement.Table))
	return ctx
}

func endTrace(db *gorm.DB) {
	span := trace.SpanFromContext(db.Statement.Context)
	defer span.End()

	if span == nil || !span.IsRecording() {
		return
	}

	if db.Error != nil && db.Error != gorm.ErrRecordNotFound {
		err := db.Error
		span.SetStatus(codes.Error, db.Error.Error())
		span.RecordError(err)
	} else {
		span.SetStatus(codes.Ok, "OK")
	}

	dsn, _ := mysqlparse.ParseDSN(db.Dialector.(*mysql.Dialector).DSN)

	span.SetAttributes(attribute.String("db.type", "mysql"))
	span.SetAttributes(attribute.String("db.name", dsn.DBName))
	span.SetAttributes(attribute.String("db.user", dsn.User))

	adds := strings.Split(dsn.Addr, ":")
	if len(adds) == 2 {
		span.SetAttributes(attribute.String("db.ip", adds[0]))
		span.SetAttributes(attribute.String("db.port", adds[1]))
	}
	span.SetAttributes(attribute.String("db.instance", dsn.DBName))

	attrs := make([]attribute.KeyValue, 0)
	if db.Statement.Table != "" {
		attrs = append(attrs, semconv.DBSQLTableKey.String(db.Statement.Table))
	}
	if db.Statement.RowsAffected != -1 {
		attrs = append(attrs, attribute.Int64("db.rows_affected", db.Statement.RowsAffected))
	}
	attrs = append(attrs, attribute.String("db.statement", db.Dialector.Explain(db.Statement.SQL.String(), db.Statement.Vars...)))
	span.SetName(fmt.Sprintf("gorm.query.%s", db.Statement.Table))
	span.SetAttributes(attrs...)
}

type GormTrace struct {
}

func (op *GormTrace) Name() string {
	return "GormTrace"
}

func (op *GormTrace) Initialize(db *gorm.DB) (err error) {
	// 开始前
	db.Callback().Create().Before("gorm_plugin:before_create").Register(callBackBeforeName, before)
	db.Callback().Query().Before("gorm_plugin:query").Register(callBackBeforeName, before)
	db.Callback().Delete().Before("gorm_plugin:before_delete").Register(callBackBeforeName, before)
	db.Callback().Update().Before("gorm_plugin:setup_reflect_value").Register(callBackBeforeName, before)
	db.Callback().Row().Before("gorm_plugin:row").Register(callBackBeforeName, before)
	db.Callback().Raw().Before("gorm_plugin:raw").Register(callBackBeforeName, before)

	// 结束后
	db.Callback().Create().After("gorm_plugin:after_create").Register(callBackAfterName, after)
	db.Callback().Query().After("gorm_plugin:after_query").Register(callBackAfterName, after)
	db.Callback().Delete().After("gorm_plugin:after_delete").Register(callBackAfterName, after)
	db.Callback().Update().After("gorm_plugin:after_update").Register(callBackAfterName, after)
	db.Callback().Row().After("gorm_plugin:row").Register(callBackAfterName, after)
	db.Callback().Raw().After("gorm_plugin:raw").Register(callBackAfterName, after)
	return
}