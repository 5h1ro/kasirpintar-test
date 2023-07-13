package utils

type MerchatTokenInterface interface {
	MerchatTokenGenerator(timeStamp, iMid, referenceNo, amt, merchantKey string) (string, error)
}

func MerchatTokenGenerator(timeStamp, iMid, referenceNo, amt, merchantKey string) (string, error) {
	return string(timeStamp + iMid + referenceNo + amt + merchantKey), nil
}
