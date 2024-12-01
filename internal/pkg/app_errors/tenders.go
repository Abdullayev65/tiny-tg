package app_errors

var (
	TenderInvalidInput           = newErr("Invalid input", 400)
	TenderInvalidData            = newErr("Invalid tender data", 400)
	TenderInvalidStatus          = newErr("Invalid tender status", 400)
	TenderNotFoundOrAccessDenied = newErr("Tender not found or access denied", 404)
)
