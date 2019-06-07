package models

import (
	"github.com/asaskevich/govalidator"
	"github.com/sirupsen/logrus"
)

func ValidatePaginationParams(fctx *FlowCtx, limit string, offset string) ([]string, bool) {
	logrus.WithField("TrackingID", fctx.TrackingID).Info("validating pagination params")

	errors := []string{}

	if !govalidator.InRangeInt(limit, 1, 5000) {
		errors = append(errors, "limit should be between 1 to 5000")
	}

	if !govalidator.InRangeInt(offset, 0, 5000) {
		errors = append(errors, "offset should be between 0 to 5000")
	}

	return errors, len(errors) <= 0
}

func ValidateID(fctx *FlowCtx, id string) ([]string, bool) {
	logrus.WithField("TrackingID", fctx.TrackingID).Info("validating ID for valid UUID")

	errors := []string{}

	if !govalidator.IsUUID(id) {
		errors = append(errors, "ID should be UUID")
	}

	return errors, len(errors) <= 0
}
