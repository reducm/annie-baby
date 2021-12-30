package controller

import "annie-baby/ent"

type ApplicationController struct {
	DB   *ent.Client
	Name string
}
