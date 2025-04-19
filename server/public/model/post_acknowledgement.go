package model

import "net/http"

type PostAcknowledgement struct {
	UserId         string `json:"user_id"`
	PostId         string `json:"post_id"`
	AcknowledgedAt int64  `json:"acknowledged_at"`
}

func (o *PostAcknowledgement) IsValid() *AppError {
	if !IsValidId(o.UserId) {
		return NewAppError("PostAcknowledgement.IsValid", "model.acknowledgement.is_valid.user_id.app_error", nil, "user_id="+o.UserId, http.StatusBadRequest)
	}

	if !IsValidId(o.PostId) {
		return NewAppError("PostAcknowledgement.IsValid", "model.acknowledgement.is_valid.post_id.app_error", nil, "post_id="+o.PostId, http.StatusBadRequest)
	}

	return nil
}
