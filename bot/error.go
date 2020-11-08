package bot

import "errors"

var (
	FORMAT_ERROR = errors.New("指令格式錯誤， 請使用help指令確認指令使用方式。")

	SEARCH_NUMBER_ERROR    = errors.New("請輸入數字用以查詢題目。")
	SEARCH_NOT_FOUND_ERROR = errors.New("資料庫內查無資料，請確認problem ID是否正確。")

	ADD_REPORT_NUMBER_ERROR = errors.New("添加題目的時候請輸入數字。")
	ADD_REPORT_REPEAT_ERROR = errors.New("請勿重複添加題目。")

	SET_TIME_FORMAT_ERROR = errors.New("時間格式錯誤， 請參考help的範例格式。")
	SET_TIME_EARLY_ERROR  = errors.New("請輸入未來的一個時間點")

	STUDYGROUP_NOT_FOUND_ERROR = errors.New("")
)
