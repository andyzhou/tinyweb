package form

/*
 * base request form
 */

type BaseForm struct {
	Act string `form:"act"`
	Page int32 `form:"page"`
}

