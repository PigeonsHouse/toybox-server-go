package main

import "github.com/gin-gonic/gin"

func setRouter(r *gin.Engine) {
	api := r.Group("/api")
	// ユーザー系
	// ユーザー一覧取得
	api.GET("/v1/users")
	// 単一のユーザー取得
	api.GET("/v1/users/:id")
	// 自身のユーザー情報取得
	api.GET("/v1/users/@me")
	// 自身のユーザー情報更新
	api.PUT("/v1/users/@me")
	// 自身のアバター更新
	api.PUT("/v1/users/@me/avatar")
	// 自身のアバター削除
	api.DELETE("/v1/users/@me/avatar")

	// 認証系
	// Discord認証
	api.GET("/v1/auth/discord")
	// Discordコールバック
	api.GET("/v1/auth/discord/callback")
	// リフレッシュトークン取得
	api.POST("/v1/auth/token")

	// 作品系
	// 作品一覧取得
	api.GET("/v1/works")
	// 作品投稿
	api.POST("/v1/works")
	// 単一の作品取得
	api.GET("/v1/works/:id")
	// 作品更新
	api.PUT("/v1/works/:id")
	// 作品削除
	api.DELETE("/v1/works/:id")
	// ページネーションでの作品一覧取得
	api.GET("/v2/works")
	// ユーザーの作品一覧取得
	api.GET("/v1/users/:id/works")
	// 自身の作品一覧取得
	api.GET("/v1/users/@me/works")

	// アセット系
	// アセット投稿
	api.POST("/v1/assets")
	// アセットの削除
	api.DELETE("/v1/assets/:id")
	// ブログのアセット投稿
	api.POST("/v1/blogs/assets")

	// コメント系
	// 作品のコメント一覧取得
	api.GET("/v1/works/:id/comments")
	// 作品へのコメント投稿
	api.POST("/v1/works/:id/comments")
	// コメントへのリプライ取得
	api.GET("/v1/works/:id/comments/:id")
	// コメント削除
	api.DELETE("/v1/works/:id/comments/:id")

	// いいね系
	// 作品のいいね一覧取得
	api.GET("/v1/works/:id/favorite")
	// 作品へのいいね
	api.POST("/v1/works/:id/favorite")
	// 作品のいいね削除
	api.DELETE("/v1/works/:id/favorite")
	// ユーザーのいいねを送った作品一覧取得
	api.GET("/v1/users/:id/favorite")
	// 自身のいいねを送った作品一覧取得
	api.GET("/v1/users/@me/favorite")

	// タグ系
	// タグ一覧取得
	api.GET("/v1/tags")
	// タグ作成
	api.POST("/v1/tags")
	// 単一のタグ取得
	api.GET("/v1/tags/:id")
	// タグ更新
	api.PUT("/v1/tags/:id")
	// タグ削除
	api.DELETE("/v1/tags/:id")

	// ブログ系
	// ブログ一覧取得
	api.GET("/v1/blogs")
	// ブログ投稿
	api.POST("/v1/blogs")
	// 単一のブログ取得
	api.GET("/v1/blogs/:id")
	// ブログ更新
	api.PUT("/v1/blogs/:id")
	// ブログ削除
	api.DELETE("/v1/blogs/:id")
	// 自身のブログ一覧取得
	api.GET("/v1/users/@me/blogs")
}
