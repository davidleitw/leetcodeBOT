package bot

import "errors"

var (
	FORMAT_ERROR = errors.New("指令格式錯誤， 請使用help指令確認指令使用方式。")

	SEARCH_NUMBER_ERROR    = errors.New("請輸入數字用以查詢題目。")
	SEARCH_NOT_FOUND_ERROR = errors.New("資料庫內查無資料，請確認problem ID是否正確。")

	ADD_REPORT_NUMBER_ERROR = errors.New("添加題目的時候請輸入數字。")
	ADD_REPORT_REPEAT_ERROR = errors.New("請勿重複添加題目。")
	ADD_TOO_MANY_REPORTS    = errors.New("請勿一次性添加超過五個題目。")

	REMOVE_TOO_MANY_REPORTS = errors.New("一次性刪除多個題目請使用clear指令。")
	REMOVE_REPROT_ERROR     = errors.New("請再次確認要刪除的題號。")

	SET_TIME_FORMAT_ERROR = errors.New("時間格式錯誤， 請參考help的範例格式。")
	SET_TIME_EARLY_ERROR  = errors.New("請輸入未來的一個時間點。")

	STUDYGROUP_NOT_FOUND_ERROR = errors.New("這個server還沒有舉辦過讀書會， 請先至少add一個題目讓bot建立新的讀書會。")

	LIST_NO_DATA_ERROR = errors.New("找不到資料，請再次確認是否有預定要在這次讀書會報告。")

	KIRITO_ERROR = errors.New("https://imgur.com/gallery/egb5S")
)
