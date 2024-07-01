package app

import (
	"chap1/internal/admin"
	"chap1/internal/blogs"
	"net/http"
)

func SetupUserRoutes() *http.ServeMux {

	userRoutes := http.NewServeMux()

	userRoutes.HandleFunc("GET /get_blog", blogs.Get_blog_data)
	userRoutes.HandleFunc("POST /post_blog", blogs.Post_blog)
	userRoutes.HandleFunc("PUT /put_blog", blogs.Put_blog)
	userRoutes.HandleFunc("PATCH /patch_blog", blogs.Patch_blog)
	userRoutes.HandleFunc("DELETE /del_blog", blogs.Del_blog)

	return userRoutes
}

func SetupAdminRoutes() *http.ServeMux {

	adminRoutes := http.NewServeMux()

	adminRoutes.HandleFunc("/get_user_data/{user_id}", admin.Get_user_data)
	adminRoutes.HandleFunc("/post_data", blogs.Post_blog)

	return adminRoutes
}
