package api

// payload of request received from a user through
// facebook messenger, including all details like
// the message content and the sender id to be
// used for answering
type PostRequestData struct {
	Entry []struct {
		ID int64 										`json:"id,string"`
		Messaging []struct {
			Message struct {
				Mid  string 						`json:"mid"`
				Seq  int64  						`json:"seq"`
				Text string 						`json:"text"`
			} 												`json:"message"`
			Recipient struct {
				ID int64 								`json:"id,string"`
			} 												`json:"recipient"`
			Sender struct {
				ID int64 								`json:"id,string"`
			} 												`json:"sender"`
			Timestamp int64 					`json:"timestamp"`
		} 													`json:"messaging"`
		Time int64 									`json:"time"`
	} 														`json:"entry"`
	Object string 								`json:"object"`
}

// payload of request to be sent to facebook endpoint
// https://graph.facebook.com/v2.6/me/messages for answering
// to a user through facebook messenger
type PostResponseData struct {
	AccessToken string 						`json:"access_token"`
	Recipient struct {
		ID int64 										`json:"id"`
	} 														`json:"recipient"`
	Message struct {
		Text string 								`json:"text"`
	} 														`json:"message"`
}
