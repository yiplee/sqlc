package sqlc

import "context"

type builderContextKey struct{}

func WithBuilder(ctx context.Context, b *Builder) context.Context {
	return context.WithValue(ctx, builderContextKey{}, b)
}

func BuilderFrom(ctx context.Context) (*Builder, bool) {
	b, ok := ctx.Value(builderContextKey{}).(*Builder)
	return b, ok
}

func Build(ctx context.Context, f func(builder *Builder)) context.Context {
	b, ok := BuilderFrom(ctx)
	if !ok {
		b = &Builder{}
	} else {
		b = b.clone()
	}

	f(b)
	return WithBuilder(ctx, b)
}
