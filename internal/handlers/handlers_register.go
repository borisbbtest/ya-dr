package handlers

import (
	"net/http"
)

func (hook *WrapperHandler) PostJSONRegisterHandler(w http.ResponseWriter, r *http.Request) {

	// var reader io.Reader

	// if r.Header.Get(`Content-Encoding`) == `gzip` {
	// 	gz, err := gzip.NewReader(r.Body)
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}
	// 	reader = gz
	// 	defer gz.Close()
	// } else {
	// 	reader = r.Body
	// }

	// bytes, err := ioutil.ReadAll(reader)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// log.Info("PostJSONHandler")
	// defer r.Body.Close()

	// var m model.RequestAddDBURL
	// if err := json.Unmarshal(bytes, &m); err != nil {
	// 	log.Errorf("body error: %v", string(bytes))
	// 	log.Errorf("error decoding message: %v", err)
	// 	http.Error(w, "request body is not valid json", 400)
	// 	return
	// }

	// hashcode, err := storage.ParserDataURL(m.ReqNewURL)
	// if err != nil {
	// 	http.Error(w, "request body is not valid URL", 400)
	// 	return
	// }

	// hashcode.UserID = hook.UserID

	// w.Header().Set("Content-Type", "application/json; charset=utf-8")

	// ShortPath, err := hook.Storage.Put(hashcode.ShortPath, hashcode)
	// if err != nil {
	// 	log.Error("Put error ", err)
	// }

	// // проверяем что получили хеш сокращенного url
	// if ShortPath != "" {
	// 	hashcode.ShortPath = ShortPath
	// 	w.WriteHeader(http.StatusConflict)
	// } else {
	// 	w.WriteHeader(http.StatusCreated)
	// }

	// resp := model.ResponseURLShort{
	// 	ResNewURL: fmt.Sprintf("%s/%s", hook.ServerConf.BaseURL, hashcode.ShortPath),
	// }

	// json.NewEncoder(w).Encode(resp)

	log.Println("Post handler")
}
