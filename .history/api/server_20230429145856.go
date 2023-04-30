package api

import db "payzone/db/sqlc"

type Server struct {
	store *db.Store
	router 
}
