package app_errors

var (
	TenderNotFound            = newErr("Tender not found", 404)
	TenderNotClosed           = newErr("Tender is not closed", 400)
	BidNotFound               = newErr("Bid not found", 404)
	BidNotFoundOrAccessDenied = newErr("Bid not found or access denied", 404)
	BidNotPending             = newErr("Bid is not pending", 400)
	BidInvalidData            = newErr("Invalid bid data", 400)
	BidTenderIsNotOpen        = newErr("Tender is not open for bids", 400)
)
