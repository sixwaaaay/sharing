package handler

import "go.uber.org/fx"

var HandlerMoudle = fx.Module("Public",
	fx.Provide(
		WrapPublicAnnotation(NewLogin),
		WrapPublicAnnotation(NewRegister),
		WrapOptionalAnnotation(NewCommentListHandler),
		WrapOptionalAnnotation(NewFeed),
		WrapOptionalAnnotation(NewFollowedListHandler),
		WrapOptionalAnnotation(NewFollowerListHandler),
		WrapOptionalAnnotation(NewUserInfoHandler),
		WrapPrivateAnnotation(NewCommentActionHandler),
		WrapPrivateAnnotation(NewFavoriteActionHandler),
		WrapPrivateAnnotation(NewFavoriteListHandler),
		WrapPrivateAnnotation(NewFollowActionHandler),
		WrapPrivateAnnotation(NewPublishListHandler),
		WrapPrivateAnnotation(NewUploadHandler)),
)

func WrapPublicAnnotation[T any](t T) any {
	return fx.Annotate(t, fx.ResultTags(`group:"public"`))
}

func WrapOptionalAnnotation[T any](t T) any {
	return fx.Annotate(t, fx.ResultTags(`group:"option"`))
}

func WrapPrivateAnnotation[T any](t T) any {
	return fx.Annotate(t, fx.ResultTags(`group:"private"`))
}
